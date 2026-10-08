// Package thinking provides the AgentThinking indicator of BoardUI for
// the MyGo native toolkit: a row of a working indicator, a label and an
// elapsed timer, drawn while an agent works on a response. The indicator
// animates by itself, repainting without rebuilding the view.
//
//	thinking.Thinking(c, thinking.Props{Label: "Thinking"})
package thinking

import (
	"fmt"
	"math"
	"time"

	"github.com/egoist/mygo/ui"
)

// Variant is the pattern of the working indicator.
type Variant string

// The variants of the indicator, from BoardUI's AgentThinking.
const (
	// Wave sweeps a diagonal wavefront across a dot grid.
	Wave Variant = "wave"
	// Spin orbits a bright head around a dot grid.
	Spin Variant = "spin"
	// Stars twinkles a few sparkles.
	Stars Variant = "stars"
	// Infinity sweeps a comet along a figure-eight.
	Infinity Variant = "infinity"
)

// Tone is the color of the indicator and the label.
type Tone string

// The tones of the indicator and the label.
const (
	Subtle  Tone = "subtle"
	Default Tone = "default"
	Primary Tone = "primary"
	Accent  Tone = "accent"
)

// Props describes the indicator to draw.
type Props struct {
	// Variant is the pattern of the indicator; the zero value is Wave.
	Variant Variant
	// Label is the status text, such as "Thinking" or "Searching the docs".
	Label string
	// Tone colors the indicator and the label. The zero value picks the
	// default tone of the variant, which is Subtle for Stars and Default
	// for the rest.
	Tone Tone
	// ShowTimer draws the time since Started after the label.
	ShowTimer bool
	// Started is when the work began, for the timer. The zero value
	// draws 0.0s.
	Started time.Time
}

// Thinking draws the indicator and returns it.
func Thinking(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	color := toneColor(t, p.Tone)
	size := t.FontSize
	return ui.Row(c).AlignItems(ui.Center).Gap(10).Children(func() {
		indicator(c, p.Variant, color)
		ui.Text(c, p.Label).FontSize(size).TextColor(color)
		if p.ShowTimer {
			ui.Box(c).Shrink(0).Width(48).Draw(func(pt *ui.Painter, r ui.Rect) {
				base := p.Started
				if base.IsZero() {
					base = pt.Now()
				}
				elapsed := pt.Now().Sub(base)
				s := fmt.Sprintf("%.1fs", elapsed.Seconds())
				pt.After(100 * time.Millisecond)
				pt.Text(r.X, r.Y+r.H/2+size*0.35, s, size*0.85, t.TextMuted)
			})
		}
	})
}

// toneColor maps a tone to a theme color; the two muted tones share the
// muted text color, as the theme has no secondary text color.
func toneColor(t *ui.Theme, tone Tone) ui.Color {
	switch tone {
	case Primary:
		return t.Text
	case Accent:
		return t.Accent
	}
	return t.TextMuted
}

// indicator draws the animated mark of the variant.
func indicator(c *ui.Context, v Variant, color ui.Color) ui.Element {
	switch v {
	case Spin:
		return dotGrid(c, true, color)
	case Stars:
		return stars(c, color)
	case Infinity:
		return infinity(c, color)
	default: // Wave
		return dotGrid(c, false, color)
	}
}

// dotGrid draws the wave or spin pattern: a 3x3 grid of dots whose
// opacities follow a phase front across the grid, in a wave along the
// diagonals or a head orbiting the centre.
func dotGrid(c *ui.Context, spin bool, color ui.Color) ui.Element {
	const size, gap = 4, 2
	e := ui.Box(c).Size(20, 20).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		phase := float64(p.Now().UnixMilli()%80) / 80
		// The 16px grid sits centred in the 20px box, with a 2px margin.
		cx, cy := r.X+2, r.Y+2
		for row := 0; row < 3; row++ {
			for col := 0; col < 3; col++ {
				s := waveScalar(col, row)
				if spin {
					s = spinScalar(col, row)
				}
				behind := math.Mod(phase-s+1, 1)
				lit := math.Pow(math.Max(0, 1-behind/0.3), 1.5)
				alpha := float32(0.12 + 0.88*lit)
				x := cx + float32(col)*(size+gap)
				y := cy + float32(row)*(size+gap)
				var dot ui.Path
				dot.Circle(x+size/2, y+size/2, size/2)
				p.FillPath(&dot, color.Alpha(alpha))
			}
		}
		p.After(80 * time.Millisecond)
	})
	return e
}

// waveScalar is how far along the wave's travel direction a cell sits:
// the front walks the diagonals from top-left to bottom-right. The scalar
// is compressed below 1 so the wrap reads as the front leaving and
// re-entering the grid.
func waveScalar(col, row int) float64 {
	return float64(col+row) / (2 * 2) * (3.0 / 4.0)
}

// spinScalar is the angle of a cell around the grid centre, as a phase
// in [0, 1), so the bright head orbits clockwise.
func spinScalar(col, row int) float64 {
	return math.Mod(math.Atan2(float64(row-1), float64(col-1))/(2*math.Pi)+1, 1)
}

// stars twinkles five four-point sparkles, each on its own phase, so the
// sky does not pulse in unison.
func stars(c *ui.Context, color ui.Color) ui.Element {
	// Positions as percentages of the box, like BoardUI's STAR_LAYOUT.
	layout := []struct{ x, y, scale float64 }{
		{50, 46, 1}, {18, 22, 0.55}, {82, 26, 0.45}, {78, 76, 0.55}, {22, 78, 0.4},
	}
	const period = 1400.0 // ms per twinkle
	e := ui.Box(c).Size(21, 21).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		ms := float64(p.Now().UnixMilli())
		for i, s := range layout {
			phase := math.Mod(ms/period+float64(i)*0.7/5, 1)
			alpha := float32(0.25 + 0.75*math.Pow(math.Sin(2*math.Pi*phase), 2))
			arm := float32(3.5 * s.scale)
			x := r.X + float32(s.x/100)*r.W
			y := r.Y + float32(s.y/100)*r.H
			var spark ui.Path
			spark.MoveTo(x-arm, y).LineTo(x+arm, y)
			spark.MoveTo(x, y-arm).LineTo(x, y+arm)
			p.StrokePath(&spark, 1.2, color.Alpha(alpha))
			var core ui.Path
			core.Circle(x, y, 0.8)
			p.FillPath(&core, color.Alpha(alpha))
		}
		p.After(100 * time.Millisecond)
	})
	return e
}

// infinity draws a figure-eight faintly and sweeps a comet along it: the
// head and the fading dots of its trail, on the curve x = A sin t,
// y = B sin 2t, so the loop reads as a horizontal eight.
func infinity(c *ui.Context, color ui.Color) ui.Element {
	const lap = 1200.0 // ms per lap
	e := ui.Box(c).Size(24, 16).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		a, b := r.W/2-2, r.H/2-2
		at := func(t float64) (float32, float32) {
			return cx + float32(float64(a)*math.Sin(t)), cy + float32(float64(b)*math.Sin(2*t))
		}
		// The faint whole, one sampled path.
		var eight ui.Path
		first := true
		for i := 0; i <= 128; i++ {
			x, y := at(2 * math.Pi * float64(i) / 128)
			if first {
				eight.MoveTo(x, y)
				first = false
			} else {
				eight.LineTo(x, y)
			}
		}
		p.StrokePath(&eight, 2, color.Alpha(0.15))
		// The comet: a bright head and fading dots behind it along the path.
		ph := math.Mod(float64(p.Now().UnixMilli())/lap, 1)
		for i := 3; i >= 0; i-- {
			t := 2*math.Pi*ph - float64(i)*0.18
			x, y := at(t)
			var dot ui.Path
			dot.Circle(x, y, float32(2-0.45*float64(i)))
			p.FillPath(&dot, color.Alpha(float32(0.55-0.14*float64(i))))
		}
		p.After(40 * time.Millisecond)
	})
	return e
}
