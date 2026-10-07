// Package grid provides a grid layout, as a shadcn/ui component for the
// Grid of MyGo: children placed in cells of columns, which they fill in
// order, or place themselves in with the placement of a cell.
//
//	grid.Grid(c, grid.Props{
//		Columns: 3,
//		Gap:     12,
//		Children: func(c *ui.Context) {
//			for _, p := range photos {
//				ui.Image(c, p).AspectRatio(1).Fit(ui.Cover)
//			}
//		},
//	})
package grid

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the grid to draw.
type Props struct {
	// Columns are the columns of equal width.
	Columns int
	// ColumnTracks set the columns instead, each a width: Fixed, Fr or
	// Auto.
	ColumnTracks []ui.Track
	// Gap spaces the cells.
	Gap float32
	// Children are the cells, which fill the columns in order or place
	// themselves with the placement of their elements, as a header
	// across a grid.
	Children func(c *ui.Context)
}

// Grid draws the grid and returns it. The children fill the columns in
// order; a child places itself with the ColumnStart, RowStart, ColumnSpan
// and RowSpan methods of its element.
func Grid(c *ui.Context, p Props) *ui.Element {
	g := ui.Grid(c)
	if p.Columns > 0 {
		g.Columns(p.Columns)
	}
	if p.ColumnTracks != nil {
		g.ColumnTracks(p.ColumnTracks...)
	}
	if p.Gap > 0 {
		g.Gap(p.Gap)
	}
	if p.Children != nil {
		g.Children(func() { p.Children(c) })
	}
	return g
}
