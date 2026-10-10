// Package drawer provides a shadcn/ui-style Drawer: a panel that rises
// over the bottom of the window, with a handle to grab, as a phone's
// sheet or a confirmation that needs the context of the page. The rest
// of the window dims and turns inert while it shows.
//
//	drawer.Drawer(c, drawer.Props{
//		Open:  &open,
//		Title: "Confirm",
//		Children: func() { confirmForm(c) },
//	})
package drawer

import "github.com/egoist/mygo/ui"

// Props describes the drawer to draw.
type Props struct {
	// Open is the bool that shows and hides the drawer; the backdrop
	// and Escape set it to false.
	Open *bool
	// Title is the heading of the drawer.
	Title string
	// Height is the height of the drawer; 380 when zero.
	Height float32
	// Children draw in the body of the drawer.
	Children func()
}

// Drawer draws the drawer for props, or nothing while it is closed, and
// returns the panel. The backdrop dims the window; clicking it or
// pressing Escape closes the drawer.
func Drawer(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	h := p.Height
	if h <= 0 {
		h = 380
	}
	return ui.DialogBase(c, p.Open, func(back, panel *ui.Element) {
		back.Background(ui.RGBA(0, 0, 0, 0.4))
		panel.Absolute().Left(0).Right(0).Bottom(0).Height(h).
			Padding(t.Space(4), t.Space(5)).Radius(t.Radius+2, t.Radius+2, 0, 0).
			Background(t.Background).Shadow(0, -10, 30, 0, ui.RGBA(0, 0, 0, 0.3))
		panel.Children(func() {
			ui.Column(c).AlignItems(ui.Center).Gap(t.Space(3)).Children(func() {
				// The handle to grab, as a phone's drawer has.
				ui.Box(c).Size(t.Space(10), 4).Radius(2).Background(t.Border)
				ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
					if p.Title != "" {
						ui.Text(c, p.Title).Bold()
					}
					if p.Children != nil {
						p.Children()
					}
				})
			})
		})
	})
}
