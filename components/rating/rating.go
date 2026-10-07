// Package rating provides a shadcn/ui-style Rating: a row of stars that
// show and set a score, by a click or the arrows.
//
//	rating.Rating(c, rating.Props{Value: &stars, Max: 5})
package rating

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// Props describes the rating to draw.
type Props struct {
	// Value is the score, from 0 to Max.
	Value *int
	// Max is the number of stars.
	Max int
	// ReadOnly draws the stars without letting the user change the score.
	ReadOnly bool
}

// Rating draws Max stars showing *Value, and returns the row, so a view
// can chain more calls on it. Label names it for assistive technology. A
// click on a star sets the score to it, and a click on the set star
// clears it; the arrows and Home and End move the score while focused.
// ReadOnly draws the stars as they are, without the interaction.
func Rating(c *ui.Context, p Props) *ui.Element {
	if p.ReadOnly {
		return readOnly(c, p)
	}
	return ui.Rating(c, p.Value, p.Max)
}

// readOnly draws the stars for ReadOnly. Why not MyGo's rating: it takes
// the pointer and the arrows, and no flag turns that off from outside.
func readOnly(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	e := ui.Row(c).Gap(t.Space(0.5)).Shrink(0).Role(ui.RoleSlider).
		Range(0, float64(p.Max), float64(*p.Value))
	star := t.FontSize * 1.25
	e.Children(func() {
		for i := range p.Max {
			b := ui.Box(c).Size(star, star).Shrink(0).Role(ui.RoleNone)
			b.Draw(func(pp *ui.Painter, r ui.Rect) {
				path := starPath(r)
				if i < *p.Value {
					pp.FillPath(path, t.Warning)
				} else {
					pp.StrokePath(path, 1.2, t.TextMuted)
				}
			})
		}
	})
	return e
}

// starPath returns a five-pointed star in r. Why not MyGo's helper: it
// is unexported, and the ten points are easy to place.
func starPath(r ui.Rect) *ui.Path {
	var path ui.Path
	cx, cy := r.X+r.W/2, r.Y+r.H/2+r.H*0.04
	outer, inner := r.W*0.48, r.W*0.2
	for i := range 10 {
		rad := outer
		if i%2 == 1 {
			rad = inner
		}
		a := float64(i)*math.Pi/5 - math.Pi/2
		x, y := cx+rad*float32(math.Cos(a)), cy+rad*float32(math.Sin(a))
		if i == 0 {
			path.MoveTo(x, y)
		} else {
			path.LineTo(x, y)
		}
	}
	path.Close()
	return &path
}
