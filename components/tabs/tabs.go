// Package tabs provides a shadcn/ui-style tab list: a row of triggers
// that switch a pane of content, of which one is chosen. The chosen one
// shows in the text color, with a bar in the accent color under it; the
// others show muted, over a hairline. Draw the content of the chosen tab
// yourself, next to the list.
//
//	tabs.Tabs(c, tabs.Props{Selected: &tab, Labels: []string{"General", "About"}})
//	switch *tab {
//	case 0:
//		general(c)
//	case 1:
//		about(c)
//	}
package tabs

import "github.com/egoist/mygo/ui"

// Props describes the tab list to draw.
type Props struct {
	// Selected is the index of the chosen tab; the list edits it, and a
	// click or the arrow keys choose.
	Selected *int
	// Labels are the tabs, in order.
	Labels []string
}

// Tabs draws a list of tabs for props: the chosen one in the text color
// with an accent bar under it, the others muted, on a hairline, as the
// tabs of shadcn/ui. It returns the list, so a view can chain more calls
// on it.
func Tabs(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	tp := ui.TabsBase(c, p.Selected, len(p.Labels))
	list := tp.List.Gap(0)
	// The hairline sits behind the tabs, so the accent bar of the chosen
	// tab draws over it.
	list.Draw(func(pp *ui.Painter, r ui.Rect) {
		pp.Fill(ui.Rect{X: r.X, Y: r.Y + r.H - 1, W: r.W, H: 1}, t.Border, 0)
	})
	list.Children(func() {
		for i, label := range p.Labels {
			tab := tp.Tab(i).Padding(t.Space(2), t.Space(3)).FocusRing(false)
			on := i == *p.Selected
			tab.TextColor(t.TextMuted)
			if on {
				tab.TextColor(t.Text).FontWeight(600)
			}
			tab.DrawOver(func(pp *ui.Painter, r ui.Rect) {
				if on {
					in, h := t.Space(1.5), t.Space(0.5)
					pp.Fill(ui.Rect{X: r.X + in, Y: r.Y + r.H - h, W: r.W - 2*in, H: h}, t.Accent, h/2)
				}
			})
			tab.Children(func() { ui.Text(c, label).SingleLine() })
		}
	})
	return list
}
