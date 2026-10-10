// Package aspectratio provides a box that keeps an aspect ratio, for the
// Aspect Ratio of shadcn/ui: the width of its container, the height the
// ratio makes, and the children filling it.
//
//	aspectratio.AspectRatio(c, aspectratio.Props{
//		Ratio:    16.0 / 9,
//		Children: func(c *ui.Context) { frame(c) },
//	})
package aspectratio

import "github.com/egoist/mygo/ui"

// Props describes the box to draw.
type Props struct {
	// Ratio is the width divided by the height; 1 when zero.
	Ratio float32
	// Children draw inside the box, filling it.
	Children func()
}

// AspectRatio draws a box of the ratio and returns it, so a view can
// chain more calls on it. The box fills the width of its container.
func AspectRatio(c *ui.Context, p Props) ui.Element {
	r := p.Ratio
	if r <= 0 {
		r = 1
	}
	e := ui.Box(c).FillWidth().AspectRatio(r)
	if p.Children != nil {
		e.Children(p.Children)
	}
	return e
}
