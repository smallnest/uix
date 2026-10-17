package mujica

import (
	"math"

	"github.com/egoist/mygo/ui"

	"git.zacharyzhang.com/ZacharyZhang-NY/MujicaUI/theme"
)

// Palette returns six series colors at 3:1 against Surface.
//
// This is a copy of MujicaUI's internal.Palette: the gallery's Charts II
// demos draw with it, and Go forbids importing another module's
// internal package. Why not call one of the chart packages instead:
// they keep these helpers unexported, so the port re-homes them.
func Palette(k theme.Tokens) []ui.Color {
	out := []ui.Color{k.AccentText, k.Info, k.Success, k.Warning, k.Ornament, k.TextMuted}
	for i, col := range out {
		for t := float32(0.1); theme.Contrast(col, k.Surface) < 3; t += 0.1 {
			col = out[i].Mix(k.Text, t)
		}
		out[i] = col
	}
	return out
}

// Marker draws series shape i centered at (cx, cy).
func Marker(p *ui.Painter, i int, cx, cy, size float32, col ui.Color) {
	cx, cy = Snap(p, cx), Snap(p, cy)
	h := size / 2
	var path ui.Path
	switch i % 5 {
	case 0:
		path.Circle(cx, cy, h)
	case 1:
		h *= 0.85
		path.MoveTo(cx-h, cy-h).LineTo(cx+h, cy-h).LineTo(cx+h, cy+h).LineTo(cx-h, cy+h).Close()
	case 2:
		p.FillPath(Diamond(cx, cy, size*1.2), col)
		return
	case 3:
		path.MoveTo(cx, cy-h).LineTo(cx+h, cy+h*0.8).LineTo(cx-h, cy+h*0.8).Close()
	default:
		path.MoveTo(cx-h, cy-h*0.8).LineTo(cx+h, cy-h*0.8).LineTo(cx, cy+h).Close()
	}
	p.FillPath(&path, col)
}

// Nearest returns the closest point within reach, or -1.
func Nearest(n int, at func(i int) (float32, float32), x, y, reach float32, byX bool) int {
	best, dist := -1, reach*reach
	for i := range n {
		px, py := at(i)
		d := (px - x) * (px - x)
		if !byX {
			d += (py - y) * (py - y)
		}
		if d <= dist {
			best, dist = i, d
		}
	}
	return best
}

// Diamond returns a rhombus of size s centered at (cx, cy).
func Diamond(cx, cy, s float32) *ui.Path {
	var p ui.Path
	h := s / 2
	p.MoveTo(cx, cy-h)
	p.LineTo(cx+h, cy)
	p.LineTo(cx, cy+h)
	p.LineTo(cx-h, cy)
	p.Close()
	return &p
}

// Snap rounds v to MyGo's quarter-device-pixel path cache grid.
func Snap(p *ui.Painter, v float32) float32 {
	g := 4 * p.Scale()
	return float32(math.Round(float64(v*g))) / g
}
