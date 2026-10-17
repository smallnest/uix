package internal

import "github.com/egoist/mygo/ui"

// Remeasure rebuilds the view when layout gives e another size than the w×h it was built for.
func Remeasure(c *ui.Context, e ui.Element, w, h float32) ui.Element {
	// MyGo v0.5.5 builders return a value, not a pointer.
	out := e.DrawOver(func(_ *ui.Painter, r ui.Rect) {
		if abs(r.W-w) > 0.5 || abs(r.H-h) > 0.5 {
			c.Invalidate()
		}
	})
	return out
}

func abs(v float32) float32 { return max(v, -v) }
