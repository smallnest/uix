// Package navigationmenu provides a shadcn/ui-style Navigation Menu:
// the bar of links across the top of a site, where a link with children
// opens a panel below it with its links. A click opens and closes the
// panel; a click outside it, or Escape, closes it, as the popover of
// MyGo does. Links without children act on their own.
//
//	navigationmenu.NavigationMenu(c, navigationmenu.Props{
//		Items: []navigationmenu.Item{
//			{Label: "Docs", Children: []navigationmenu.Item{
//				{Label: "Getting started", OnClick: app.showStart},
//				{Label: "Guides", OnClick: app.showGuides},
//			}},
//			{Label: "Source", OnClick: app.openSource},
//		},
//	})
package navigationmenu

import "github.com/egoist/mygo/ui"

// Item is one link of the bar, or of the panel of a link with children.
type Item struct {
	// Label is the text of the link.
	Label string
	// OnClick runs when the link is chosen.
	OnClick func()
	// Children are the links of the panel the link opens below it; nil
	// makes it a plain link of the bar.
	Children []Item
}

// State is where the bar is, kept by the app.
type State struct {
	// open holds whether the panel of each link with children shows.
	open []bool
}

// Props describes the bar to draw.
type Props struct {
	// State is where the bar is, for the app to keep.
	State *State
	// Items are the links of the bar, in order.
	Items []Item
}

// NavigationMenu draws the bar and returns it, so a view can chain more
// calls on it. The link with children shows an arrow, and opens its
// panel of links below it on a click.
func NavigationMenu(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	st := p.State
	if st == nil {
		st = &State{}
	}
	if len(st.open) != len(p.Items) {
		st.open = make([]bool, len(p.Items))
	}
	bar := ui.Row(c).FillWidth().Gap(t.Space(1))
	bar.Children(func() {
		for i := range p.Items {
			it := &p.Items[i]
			if len(it.Children) == 0 {
				link(c, t, it)
			} else {
				menuLink(c, t, st, i, it)
			}
		}
	})
	return bar
}

// link draws one link of the bar: a flat button that fills on hover and
// press, and whose OnClick runs for the click.
func link(c *ui.Context, t *ui.Theme, it *Item) {
	b := ui.ButtonBase(c).Padding(t.Space(1), t.Space(2.5)).Radius(t.Radius)
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if b.Hovered() || b.Pressed() {
			pp.Fill(r, t.SurfaceHover, t.Radius)
		}
	})
	b.Children(func() { ui.Text(c, it.Label).SingleLine() })
	if b.Clicked() && it.OnClick != nil {
		it.OnClick()
	}
}

// menuLink draws the link of the bar that opens a panel: an arrow says
// it does, and a click opens its panel of links below it, and closes
// the panels of the other links.
func menuLink(c *ui.Context, t *ui.Theme, st *State, i int, it *Item) {
	b := ui.ButtonBase(c).Padding(t.Space(1), t.Space(2.5)).Radius(t.Radius)
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if b.Hovered() || b.Pressed() {
			pp.Fill(r, t.SurfaceHover, t.Radius)
		}
	})
	b.Children(func() {
		ui.Row(c).Gap(t.Space(1.5)).Children(func() {
			ui.Text(c, it.Label).SingleLine()
			ui.Icon(c, chevronDown).Size(t.Space(3), t.Space(3)).TextColor(t.TextMuted)
		})
	})
	if b.Clicked() {
		st.open[i] = !st.open[i]
		for j := range st.open {
			if j != i {
				st.open[j] = false
			}
		}
	}
	if !st.open[i] {
		return
	}
	ui.Popover(c, b, &st.open[i], func() {
		ui.Column(c).Gap(t.Space(0.5)).Children(func() {
			for j := range it.Children {
				child := &it.Children[j]
				cb := ui.ButtonBase(c).FillWidth().Justify(ui.Start).
					Padding(t.Space(1.5), t.Space(2.5)).Radius(t.Radius)
				cb.Draw(func(pp *ui.Painter, r ui.Rect) {
					if cb.Hovered() || cb.Pressed() {
						pp.Fill(r, t.SurfaceHover, t.Radius)
					}
				})
				cb.Children(func() { ui.Text(c, child.Label).SingleLine() })
				if cb.Clicked() {
					st.open[i] = false
					if child.OnClick != nil {
						child.OnClick()
					}
				}
			}
		})
	})
}

var chevronDown = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>`))
