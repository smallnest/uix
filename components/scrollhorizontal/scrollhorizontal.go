// Package scrollhorizontal provides a container that scrolls its content
// horizontally, as a shadcn/ui component for the ScrollHorizontal of
// MyGo.
//
//	scrollhorizontal.ScrollHorizontal(c, scrollhorizontal.Props{
//		Children: func(c *ui.Context) {
//			for _, t := range tabs {
//				ui.Text(c, t)
//			}
//		},
//	})
package scrollhorizontal

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the scroll container to draw.
type Props struct {
	// Children are the content, which the scroll moves sideways when it
	// goes past the edge.
	Children func(c *ui.Context)
}

// ScrollHorizontal draws a row that scrolls its children horizontally,
// and returns it. Size it, or Grow it within its parent; the content
// sizes to itself, so the scroll can move it.
func ScrollHorizontal(c *ui.Context, p Props) ui.Element {
	s := ui.ScrollHorizontal(c)
	if p.Children != nil {
		s.Children(func() { p.Children(c) })
	}
	return s
}
