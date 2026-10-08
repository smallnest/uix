// Package scroll provides a container that scrolls its content, as a
// shadcn/ui component for the Scroll of MyGo.
//
//	scroll.Scroll(c, scroll.Props{
//		Children: func(c *ui.Context) {
//			for _, line := range lines {
//				ui.Text(c, line)
//			}
//		},
//	}).Grow(1)
package scroll

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the scroll container to draw.
type Props struct {
	// Children are the content, which the scroll moves when it goes
	// past the edge.
	Children func(c *ui.Context)
}

// Scroll draws a container that scrolls its children vertically, and
// returns it. Size it, or Grow it within its parent; the content sizes
// to itself, so the scroll can move it.
func Scroll(c *ui.Context, p Props) ui.Element {
	s := ui.Scroll(c)
	if p.Children != nil {
		s.Children(func() { p.Children(c) })
	}
	return s
}
