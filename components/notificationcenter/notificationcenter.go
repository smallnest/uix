// Package notificationcenter provides the NotificationCenter of BoardUI
// for the MyGo native toolkit: a panel of notifications with a title, an
// unread count, a mark-all-read control, and tabs that filter the list.
//
//	notificationcenter.NotificationCenter(c, notificationcenter.Props{
//		Items: items,
//		Tab:   &tab,
//	})
//
// The panel is a rounded card on the theme's Background. The list sits
// on a recessed band; each notification is a card with a status icon,
// the title and timestamp, the description, and its actions. The app
// owns the list: it sets the unread flags, and Mark all read asks it to
// clear them.
package notificationcenter

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/segmented"
)

// Status is the tone of a notification's icon.
type Status int

const (
	// Neutral is a general notice, the bell.
	Neutral Status = iota
	// Info is an informational notice, the i.
	Info
	// Success is a notice of something that went well, the check.
	Success
	// Error is a notice of something that went wrong, the bang.
	Error
)

// Action is a button a notification carries, such as "View".
type Action struct {
	// ID names the action for OnAction.
	ID string
	// Label is the text of the button.
	Label string
	// OnClick runs when the button is clicked.
	OnClick func()
}

// Item is one notification.
type Item struct {
	// ID names the notification for OnAction.
	ID string
	// Category is "mentions" or "system"; anything else is general
	// activity, shown on the All tab.
	Category string
	// Title is the headline of the notification.
	Title string
	// Description is the detail under it.
	Description string
	// Timestamp is the time, as the app formats it, such as "2m".
	Timestamp string
	// Unread draws the accent dot and counts toward the unread total.
	Unread bool
	// Status tones the icon; the zero value is Neutral.
	Status Status
	// Actions are the buttons of the notification, in order.
	Actions []Action
}

// Props describes the panel to draw.
type Props struct {
	// Items are the notifications, in order.
	Items []Item
	// Tab is the chosen tab: 0 All, 1 Mentions, 2 System.
	Tab *int
	// OnTabChange runs when a tab is chosen.
	OnTabChange func(i int)
	// OnMarkAllRead runs when Mark all read is clicked; the app clears
	// the unread flags of its items.
	OnMarkAllRead func()
	// OnAction runs when a notification's action is clicked, with the
	// notification and the action, and then the action's own OnClick.
	OnAction func(itemID, actionID string)
	// Title is the heading of the panel; "Notifications" when empty.
	Title string
	// EmptyMessage is shown for an empty tab; "You're all caught up."
	// when empty.
	EmptyMessage string
	// Height, when set, caps the list and scrolls it; 0 grows the panel
	// to its items.
	Height float32
}

// NotificationCenter draws the panel and returns it, so a view can chain
// more calls on it.
func NotificationCenter(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	title := p.Title
	if title == "" {
		title = "Notifications"
	}
	empty := p.EmptyMessage
	if empty == "" {
		empty = "You're all caught up."
	}
	unread := 0
	for _, it := range p.Items {
		if it.Unread {
			unread++
		}
	}
	return ui.Column(c).FillWidth().Background(t.Background).Radius(t.Radius * 3).
		Border(1, t.Border).Children(func() {
		ui.Column(c).Padding(t.Space(4), t.Space(4), t.Space(1.5), t.Space(4)).
			Gap(t.Space(3)).Children(func() {
			header(c, t, p, title, unread)
			tabs(c, t, p)
		})
		list(c, t, p, empty)
	})
}

// header is the title, the unread count and the mark-all-read control.
func header(c *ui.Context, t *ui.Theme, p Props, title string, unread int) {
	ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Start).
		Gap(t.Space(4)).Children(func() {
		ui.Column(c).Gap(2).Children(func() {
			ui.Text(c, title).FontSize(t.FontSize * 1.1).FontWeight(600).TextColor(t.Text)
			sub := "No unread notifications"
			if unread > 0 {
				sub = fmt.Sprintf("%d unread", unread)
			}
			ui.Text(c, sub).TextColor(t.TextMuted)
		})
		markAll(c, t, p, unread)
	})
}

// markAll is the ghost control that clears the unread flags, disabled
// when there is nothing unread.
func markAll(c *ui.Context, t *ui.Theme, p Props, unread int) {
	b := ui.ButtonBase(c).Radius(t.Radius).Padding(t.Space(1.5), t.Space(2.5))
	b.Disabled(unread == 0)
	b.Label("Mark all read")
	if b.Clicked() && p.OnMarkAllRead != nil {
		p.OnMarkAllRead()
	}
	b.Children(func() {
		ui.Text(c, "Mark all read").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted).SingleLine()
	})
}

// tabs is the segmented control with the per-tab counts.
func tabs(c *ui.Context, t *ui.Theme, p Props) {
	var all, mentions, system int
	for _, it := range p.Items {
		all++
		switch it.Category {
		case "mentions":
			mentions++
		case "system":
			system++
		}
	}
	segmented.Segmented(c, segmented.Props{
		Selected: p.Tab,
		Labels: []string{
			fmt.Sprintf("All %d", all),
			fmt.Sprintf("Mentions %d", mentions),
			fmt.Sprintf("System %d", system),
		},
		OnChange: p.OnTabChange,
		Label:    "Notification category",
	})
}

// list is the recessed band of notifications, or the empty message.
func list(c *ui.Context, t *ui.Theme, p Props, empty string) {
	visible := p.Items
	tab := 0
	if p.Tab != nil {
		tab = *p.Tab
	}
	switch tab {
	case 1:
		visible = filterCategory(p.Items, "mentions")
	case 2:
		visible = filterCategory(p.Items, "system")
	}
	band := func() {
		ui.Column(c).Gap(t.Space(2)).Children(func() {
			if len(visible) == 0 {
				emptyState(c, t, empty)
				return
			}
			for _, it := range visible {
				itemRow(c, t, p, it)
			}
		})
	}
	if p.Height > 0 {
		ui.Scroll(c).Height(p.Height).Children(band).Padding(t.Space(1.5)).
			Background(t.SurfaceHover).Radius(t.Radius * 2)
	} else {
		ui.Column(c).Padding(t.Space(1.5)).Background(t.SurfaceHover).
			Radius(t.Radius * 2).Children(band)
	}
}

// filterCategory keeps the notifications of one category.
func filterCategory(items []Item, category string) []Item {
	out := make([]Item, 0, len(items))
	for _, it := range items {
		if it.Category == category {
			out = append(out, it)
		}
	}
	return out
}

// itemRow is one notification card.
func itemRow(c *ui.Context, t *ui.Theme, p Props, it Item) {
	ui.Row(c).FillWidth().Gap(t.Space(3)).Padding(t.Space(3)).
		Background(t.Background).Radius(t.Radius * 2).Children(func() {
		statusCircle(c, t, it.Status)
		ui.Column(c).Grow(1).Gap(2).Children(func() {
			ui.Row(c).FillWidth().Gap(t.Space(2)).Children(func() {
				ui.Text(c, it.Title).Grow(1).TextColor(t.Text).SingleLine()
				ui.Text(c, it.Timestamp).FontSize(t.FontSize * 0.8).
					TextColor(t.TextMuted).SingleLine()
				if it.Unread {
					ui.Box(c).Size(8, 8).Radius(4).Background(t.Accent).
						Label("Unread")
				}
			})
			ui.Text(c, it.Description).TextColor(t.TextMuted)
			if len(it.Actions) > 0 {
				ui.Row(c).Gap(t.Space(2)).Children(func() {
					for _, a := range it.Actions {
						actionButton(c, t, p, it, a)
					}
				})
			}
		})
	})
}

// actionButton is one action of a notification.
func actionButton(c *ui.Context, t *ui.Theme, p Props, it Item, a Action) {
	b := ui.ButtonBase(c).Radius(t.Radius).Background(t.Surface).
		Padding(t.Space(1.5), t.Space(2.5))
	if b.Clicked() {
		if p.OnAction != nil {
			p.OnAction(it.ID, a.ID)
		}
		if a.OnClick != nil {
			a.OnClick()
		}
	}
	b.Children(func() {
		ui.Text(c, a.Label).FontSize(t.FontSize * 0.85).TextColor(t.Text).SingleLine()
	})
}

// statusCircle is the round tile with the status icon.
func statusCircle(c *ui.Context, t *ui.Theme, status Status) {
	bg, fg := t.SurfaceHover, t.TextMuted
	switch status {
	case Info:
		bg, fg = t.Accent.Alpha(0.12), t.Accent
	case Success:
		bg, fg = t.Success.Alpha(0.15), t.Success
	case Error:
		bg, fg = t.Danger.Alpha(0.12), t.Danger
	}
	ui.Box(c).Size(40, 40).Radius(20).Background(bg).Shrink(0).Center().
		Children(func() {
			switch status {
			case Info:
				info(c, fg)
			case Success:
				check(c, fg)
			case Error:
				bang(c, fg)
			default:
				bell(c, fg)
			}
		})
}

// emptyState is the message of an empty tab.
func emptyState(c *ui.Context, t *ui.Theme, message string) {
	ui.Column(c).FillWidth().Center().Gap(t.Space(2)).
		Padding(0, t.Space(6)).Children(func() {
		ui.Box(c).Size(44, 44).Radius(22).Background(t.Surface).Center().
			Children(func() { bellOff(c, t.TextMuted) })
		ui.Text(c, message).TextColor(t.Text)
		ui.Text(c, "New activity will appear here when it arrives.").
			TextColor(t.TextMuted).TextAlign(ui.Center)
	})
}

// bell draws the neutral bell: a dome, a base, a clapper.
func bell(c *ui.Context, color ui.Color) {
	icon(c, func(pt *ui.Painter, cx, cy float32) {
		var dome ui.Path
		dome.MoveTo(cx-6, cy+2.5)
		dome.QuadTo(cx-6.5, cy-2.5, cx-3, cy-5.5)
		dome.QuadTo(cx-1, cy-6.5, cx, cy-6.5)
		dome.QuadTo(cx+1, cy-6.5, cx+3, cy-5.5)
		dome.QuadTo(cx+6.5, cy-2.5, cx+6, cy+2.5)
		pt.StrokePath(&dome, 1.6, color)
		var base ui.Path
		base.MoveTo(cx-5.5, cy+2.5).LineTo(cx+5.5, cy+2.5)
		pt.StrokePath(&base, 1.6, color)
		var clapper ui.Path
		clapper.Circle(cx, cy+4.5, 1.3)
		pt.FillPath(&clapper, color)
	})
}

// bellOff draws the bell with a slash over it, for an empty tab.
func bellOff(c *ui.Context, color ui.Color) {
	icon(c, func(pt *ui.Painter, cx, cy float32) {
		var dome ui.Path
		dome.MoveTo(cx-6, cy+2.5)
		dome.QuadTo(cx-6.5, cy-2.5, cx-3, cy-5.5)
		dome.QuadTo(cx-1, cy-6.5, cx, cy-6.5)
		dome.QuadTo(cx+1, cy-6.5, cx+3, cy-5.5)
		dome.QuadTo(cx+6.5, cy-2.5, cx+6, cy+2.5)
		pt.StrokePath(&dome, 1.6, color)
		var slash ui.Path
		slash.MoveTo(cx-7, cy-7).LineTo(cx+7, cy+7)
		pt.StrokePath(&slash, 1.6, color)
	})
}

// info draws the i of an informational notice.
func info(c *ui.Context, color ui.Color) {
	icon(c, func(pt *ui.Painter, cx, cy float32) {
		var ring ui.Path
		ring.Circle(cx, cy, 8)
		pt.StrokePath(&ring, 1.6, color)
		var dot ui.Path
		dot.Circle(cx, cy-3.5, 1.2)
		pt.FillPath(&dot, color)
		pt.Line(cx, cy-1, cx, cy+4.5, 1.6, color)
	})
}

// check draws the check of a success notice.
func check(c *ui.Context, color ui.Color) {
	icon(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.MoveTo(cx-5, cy).LineTo(cx-1.5, cy+3.5).LineTo(cx+5.5, cy-3.5)
		pt.StrokePath(&p, 2, color)
	})
}

// bang draws the bang of an error notice.
func bang(c *ui.Context, color ui.Color) {
	icon(c, func(pt *ui.Painter, cx, cy float32) {
		var ring ui.Path
		ring.Circle(cx, cy, 8)
		pt.StrokePath(&ring, 1.6, color)
		pt.Line(cx, cy-4.5, cx, cy+0.5, 1.8, color)
		var dot ui.Path
		dot.Circle(cx, cy+3.8, 1.2)
		pt.FillPath(&dot, color)
	})
}

// icon runs fn over a 20-point icon centered in the box.
func icon(c *ui.Context, fn func(pt *ui.Painter, cx, cy float32)) {
	ui.Box(c).Size(20, 20).Shrink(0).Draw(func(pt *ui.Painter, r ui.Rect) {
		fn(pt, r.X+r.W/2, r.Y+r.H/2)
	})
}
