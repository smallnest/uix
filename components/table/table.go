// Package table provides a shadcn/ui-style data table for the MyGo native
// toolkit, on top of the Table of MyGo: a bordered box with a header
// band, rows that highlight under the pointer, and a chosen row in the
// accent color. The sort, the choice and the column layout live in the
// ListState the caller owns, as the state of any control does.
//
//	table.Table(c, table.Props{
//		State:   list, // &ui.ListState, owned by the view
//		Columns: []ui.TableColumn{{Title: "Name", Sortable: true}, {Title: "Size", Align: ui.End, Sortable: true}},
//		Rows:    len(rows),
//		Cell: func(row, col int) {
//			ui.Text(c, rows[row][col]).SingleLine()
//		},
//	}).Grow(1)
package table

import "github.com/egoist/mygo/ui"

// Props describes the table to draw.
type Props struct {
	// State is the ListState the table keeps its choice, its sort and
	// its column layout in; the caller owns it and keeps it across
	// frames. A nil State works, but then rows cannot be chosen, sorted
	// or rearranged.
	State *ui.ListState
	// Columns are the columns of the table. The user resizes them by
	// dragging the edges of their headers, and moves them by dragging
	// the headers.
	Columns []ui.TableColumn
	// Rows is the number of rows the table draws. It builds only those
	// in view, so a table of thousands is as cheap as one of ten.
	Rows int
	// Cell builds the content of the cell of a row and a column, usually
	// a Text.
	Cell func(row, col int)
}

// Table draws a table and returns it, so a view can chain more calls on
// it, such as Grow and Submitted.
func Table(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	e := ui.Table(c, p.State, p.Columns, p.Rows, p.Cell)
	e.Background(t.Background).Border(1, t.Border).Radius(t.Radius)
	// The header band: MyGo's table draws no face for its header row,
	// which the band gives it, as the muted band of shadcn/ui.
	//
	// Why not set a background on the header itself: it is built inside
	// MyGo's Table, out of reach, so the band is drawn behind it, the
	// height of the header the table sets.
	e.Draw(func(pp *ui.Painter, r ui.Rect) {
		pp.Fill(ui.Rect{X: r.X, Y: r.Y, W: r.W, H: t.Space(8)}, t.Surface, 0)
	})
	return e
}
