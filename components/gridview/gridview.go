// Package gridview provides a scrolling grid of items, as a shadcn/ui
// component for the GridView of MyGo: as many columns as items of a
// minimum width fit, a click chooses an item, and the arrow keys move
// the choice around.
//
//	gridview.GridView(c, gridview.Props{
//		State:    &state,
//		Count:    len(photos),
//		MinWidth: 140,
//		Height:   120,
//		Item: func(c *ui.Context, i int) {
//			ui.Image(c, photos[i]).Fit(ui.Contain).Grow(1)
//		},
//	})
package gridview

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the grid view to draw.
type Props struct {
	// State keeps the place and the choice of the grid, which the app
	// owns. Set its Selected to let a click choose an item, as the
	// arrow keys do while the grid has the focus.
	State *ui.GridState
	// Count is the number of items.
	Count int
	// MinWidth is the least width of an item; as many columns as fit
	// share the width.
	MinWidth float32
	// Height is the height of every item.
	Height float32
	// Item draws item i, in a cell the grid styles when it is chosen.
	Item func(c *ui.Context, i int)
	// OnChoose runs when a click or the arrow keys choose an item.
	OnChoose func(i int)
}

// GridView draws the items in a scrolling grid, and returns it. A click
// or the arrow keys choose an item when State.Selected is set, and the
// chosen cell shows its choice.
func GridView(c *ui.Context, p Props) ui.Element {
	g := ui.GridView(c, p.State, p.Count, p.MinWidth, p.Height, func(i int) {
		if p.Item != nil {
			p.Item(c, i)
		}
	})
	if p.OnChoose != nil && g.Changed() && p.State != nil && p.State.Selected != nil {
		p.OnChoose(*p.State.Selected)
	}
	return g
}
