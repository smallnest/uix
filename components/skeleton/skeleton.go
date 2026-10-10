// Package skeleton provides a placeholder of content still loading, for
// the Skeleton of shadcn/ui: a rounded bar in the muted color that pulses
// softly while it stands in for the content.
//
//	skeleton.Skeleton(c, skeleton.Props{Width: 240, Height: 20})
package skeleton

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the placeholder to draw.
type Props struct {
	// Width and Height of the bar, which fill the container when zero.
	Width, Height float32
	// Radius rounds the corners; the theme's radius when zero.
	Radius float32
}

// Skeleton draws the bar and returns it, so a view can chain more calls
// on it. The bar pulses between the muted tones as time loops, a frame
// for each change.
func Skeleton(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	radius := p.Radius
	if radius <= 0 {
		radius = t.Radius - 2
	}
	e := ui.Box(c).Background(t.SurfaceHover).Radius(radius)
	if p.Width > 0 {
		e.Width(p.Width)
	} else {
		e.FillWidth()
	}
	if p.Height > 0 {
		e.Height(p.Height)
	} else {
		e.Height(t.FontSize)
	}
	phase := e.Loop("pulse", 1400*time.Millisecond, ui.Bounce(ui.EaseInOut))
	e.Opacity(0.45 + 0.55*phase)
	return e
}
