// Package separator provides a shadcn/ui-style Separator: a thin line
// that divides content, horizontal or vertical.
//
//	separator.Separator(c, separator.Props{})
package separator

import "github.com/egoist/mygo/ui"

// Orientation of the line.
type Orientation int

const (
	// Horizontal draws a line across the width. The default.
	Horizontal Orientation = iota
	// Vertical draws a line down the height.
	Vertical
)

// Props describes the separator to draw.
type Props struct {
	// Orientation of the line. Horizontal fills the width of the parent,
	// vertical fills the height.
	Orientation Orientation
}

// Separator draws a 1-point line in the border color, labeled for
// assistive technology, and returns it.
func Separator(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	e := ui.Box(c).Background(t.Border).Label("separator").Shrink(0)
	if p.Orientation == Vertical {
		return e.Width(1).FillHeight()
	}
	return e.Height(1).FillWidth()
}
