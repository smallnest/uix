// Package dropdown provides a shadcn/ui-style dropdown menu: a button
// that opens a list of actions under it, to run one. The menu is a panel
// styled like a select's popup, over the view; it closes on a choice, on
// a press outside the menu or the button, or on Escape.
//
//	dropdown.Dropdown(c, dropdown.Props{
//		Trigger: "Account",
//		Open:    &open,
//		Items: []dropdown.Item{
//			{Label: "View profile", OnClick: viewProfile},
//			{Label: "Copy link", OnClick: copyLink},
//			{Separator: true},
//			{Label: "Sign out", OnClick: signOut},
//		},
//	})
package dropdown

import "github.com/egoist/mygo/ui"

// Item is one entry of a dropdown menu.
type Item struct {
	// Label is the text of the item.
	Label string
	// Disabled keeps the item from acting; it still shows, muted.
	Disabled bool
	// Separator draws a divider instead of a row; the other fields of a
	// separator are ignored.
	Separator bool
	// OnClick runs when the item is chosen; the menu closes either way.
	OnClick func()
}

// Props describes the dropdown to draw.
type Props struct {
	// Trigger is the label of the button that opens the menu.
	Trigger string
	// Open shows the menu while true; the button toggles it, and a
	// choice, a press outside the menu and the button, and Escape set
	// it to false.
	Open *bool
	// Items are the entries of the menu, in order.
	Items []Item
}

// Dropdown draws a button for props, with a menu of its items under it,
// and returns the button. The button opens the menu; choosing an item
// runs its OnClick and closes the menu. Why not a modal backdrop: a
// modal popover is not exposed by PopoverBase, and without one the press
// that closes the menu goes on to what is under it, as with the web's
// dropdown menus.
func Dropdown(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).
		Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	b.Children(func() {
		ui.Row(c).Gap(t.Space(1.5)).Children(func() {
			ui.Text(c, p.Trigger).SingleLine()
			chevron(c, t)
		})
	})
	// The button toggles the menu. Why not a popover that closes on a
	// press of the button: the popover leaves the presses on its anchor
	// to the anchor, so the toggle is the whole story of the button, and
	// the presses elsewhere are the popover's to close.
	if b.Clicked() {
		*p.Open = !*p.Open
	}
	ui.PopoverBase(c, b, p.Open, func(panel ui.Element) {
		panel.Margin(t.Space(1), 0, 0, 0).Padding(t.Space(1)).Radius(t.Radius+2).
			Background(t.Background).Border(1, t.Border).MinWidth(t.Space(36))
		panel.Shadow(0, 6, 20, 0, ui.RGBA(0, 0, 0, 0.18))
		panel.Children(func() {
			for _, it := range p.Items {
				item(c, t, it, p.Open)
			}
		})
	})
	return b
}

// item draws one entry of the menu: a divider for a separator, else a
// row that runs its OnClick and closes the menu.
func item(c *ui.Context, t *ui.Theme, it Item, open *bool) {
	if it.Separator {
		ui.Box(c).FillWidth().Height(1).Margin(t.Space(0.5), 0, t.Space(0.5), 0).Background(t.Border)
		return
	}
	row := ui.ButtonBase(c).FillWidth().Padding(t.Space(1.25), t.Space(2.5)).Radius(t.Radius)
	fg := t.Text
	if it.Disabled {
		fg = t.TextMuted
	}
	row.TextColor(fg)
	if !it.Disabled {
		row.Draw(func(p *ui.Painter, r ui.Rect) {
			if row.Hovered() {
				p.Fill(r, t.SurfaceHover, t.Radius)
			}
		})
	}
	row.Children(func() { ui.Text(c, it.Label).SingleLine() })
	if row.Clicked() && !it.Disabled {
		*open = false
		if it.OnClick != nil {
			it.OnClick()
		}
	}
}

// chevron draws the arrow of the button.
func chevron(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Size(t.Space(2.5), t.Space(2.5)).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		var path ui.Path
		path.MoveTo(r.X+r.W*0.1, r.Y+r.H*0.3).LineTo(r.X+r.W*0.5, r.Y+r.H*0.7).LineTo(r.X+r.W*0.9, r.Y+r.H*0.3)
		p.StrokePath(&path, 1.5, t.TextMuted)
	})
}
