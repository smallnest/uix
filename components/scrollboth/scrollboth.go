// Package scrollboth provides a container that scrolls its content both
// ways, as a canvas or a wide table does, for the ScrollBoth of MyGo.
//
//	scrollboth.ScrollBoth(c, scrollboth.Props{
//		Children: func(c *ui.Context) {
//			for _, row := range grid {
//				ui.Text(c, row)
//			}
//		},
//	})
package scrollboth

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the scroll container to draw.
type Props struct {
	// Children are the content, which the scroll moves up and sideways
	// when it goes past the edge.
	Children func(c *ui.Context)
}

// ScrollBoth draws a container that scrolls its children up and sideways,
// and returns it. Size it, or Grow it within its parent; the content
// sizes to itself, so the scroll can move it.
func ScrollBoth(c *ui.Context, p Props) *ui.Element {
	s := ui.ScrollBoth(c)
	if p.Children != nil {
		s.Children(func() { p.Children(c) })
	}
	return s
}
