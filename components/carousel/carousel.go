// Package carousel provides the gallery carousel of BoardUI for the MyGo
// native toolkit: a horizontal run of slides you slide through with the
// previous and next buttons or the position dots, or with a trackpad.
//
//	carousel.Carousel(c, carousel.Props{
//		Count:  len(shots),
//		Slide: func(c *ui.Context, i int) { shots[i].Show(c) },
//		Scroll: &state.Scroll,
//		State:  &state.Carousel,
//	})
//
// The track is a horizontal scroll container, so the user scrolls it as
// a native app scrolls. Each slide is as wide as the track, one slide a
// view. The buttons sit above the track on the right and disable at the
// ends; the dots below tell which slide is showing and jump to the
// slide they name. The app keeps the ScrollState of the track and the
// State of the carousel, as it keeps the scroll of a chat.
package carousel

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Props describes the carousel to draw.
type Props struct {
	// Count is the number of slides.
	Count int
	// Slide builds slide i into the track; nil draws empty slides.
	Slide func(c *ui.Context, i int)
	// Scroll is the scroll state of the track, for the app to keep.
	// Without it the carousel draws but does not move.
	Scroll *ui.ScrollState
	// State is the carousel's own state, for the app to keep: which
	// slide is showing and whether either end is reached.
	State *State
	// HideArrows hides the previous and next buttons above the track.
	HideArrows bool
	// HideDots hides the position dots below the track.
	HideDots bool
	// Gap is the space between slides; 0 draws the theme's spacing.
	Gap float32
	// Height is the height of the track; the slides size it when 0.
	Height float32
}

// State is where the carousel is, kept by the app as the slides move.
type State struct {
	// Active is the slide showing, from 0.
	Active int
	// AtStart and AtEnd say whether either end of the track is reached.
	AtStart, AtEnd bool
	// viewport is how wide the track is, measured from its layout so the
	// slides can be as wide as the track a frame later.
	viewport float32
	// settling is true for one frame after the slide width changed, so
	// the width is not measured against the layout it has not had yet.
	settling bool
}

// Carousel draws the carousel and returns it, so a view can chain more
// calls on it.
func Carousel(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	n := max(1, p.Count)
	gap := p.Gap
	if gap == 0 {
		gap = t.Space(4)
	}
	// The width of a slide is the width of the track. A scroll row
	// measures its children with no bound on their width, so the width
	// must be in points, not percents: the track's width of the last
	// frame is read back from how far it can scroll, with the width of
	// the window for the very first frame.
	slideW := float32(0)
	measuring := p.Scroll != nil && p.State != nil && n > 1
	if measuring {
		st := p.State
		if st.viewport > 0 {
			slideW = st.viewport
		} else {
			w, _ := c.Size()
			slideW = w
			// A second frame, to measure the track once it is laid out.
			c.Invalidate()
		}
		if st.settling {
			// The layout has not yet seen the new slide width; measure
			// it a frame from now, when it has.
			st.settling = false
		} else if p.Scroll.MaxX > 0 {
			// The track shows slides and gaps, and can scroll the rest:
			// viewport = content - MaxX.
			v := float32(n)*slideW + float32(n-1)*gap - p.Scroll.MaxX
			if v != st.viewport {
				st.viewport = v
				st.settling = true
				c.Invalidate()
			}
		}
	}
	if p.Scroll != nil && p.State != nil {
		st := p.State
		if n > 1 && p.Scroll.MaxX > 0 {
			pitch := p.Scroll.MaxX / float32(n-1)
			st.Active = clamp(int((p.Scroll.X+pitch/2)/pitch), 0, n-1)
		} else {
			st.Active = 0
		}
		st.AtStart = p.Scroll.X <= 1
		st.AtEnd = p.Scroll.MaxX > 0 && p.Scroll.X >= p.Scroll.MaxX-1
	}
	return ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		if !p.HideArrows && n > 1 {
			arrowRow(c, t, p, n)
		}
		scroll := ui.ScrollHorizontal(c).FillWidth().Gap(gap).TrackScroll(p.Scroll)
		if p.Height > 0 {
			scroll.Height(p.Height)
		}
		scroll.Children(func() {
			for i := 0; i < n; i++ {
				// The slide is as wide as the track and does not shrink,
				// so exactly one shows at a time.
				ui.Box(c).Width(slideW).Shrink(0).Children(func() {
					if p.Slide != nil {
						p.Slide(c, i)
					}
				})
			}
		})
		if !p.HideDots && n > 1 {
			dots(c, t, p, n)
		}
	})
}

// arrowRow is the previous and next buttons above the track on the
// right, as BoardUI draws them.
func arrowRow(c *ui.Context, t *ui.Theme, p Props, n int) {
	ui.Row(c).FillWidth().Justify(ui.End).Gap(t.Space(2)).Children(func() {
		arrow(c, t, p, "Previous slide", false, p.State.AtStart, n, p.State.Active-1)
		arrow(c, t, p, "Next slide", true, p.State.AtEnd, n, p.State.Active+1)
	})
}

// arrow is one of the previous and next buttons: a chevron in a round
// box, disabled at its end of the track.
func arrow(c *ui.Context, t *ui.Theme, p Props, label string, next bool, disabled bool, n, to int) {
	fg := t.Text
	if disabled {
		fg = t.TextMuted
	}
	b := ui.ButtonBase(c).Size(28, 28).Radius(t.Radius).Center()
	b.Label(label)
	b.Children(func() {
		ui.Box(c).Size(16, 16).Draw(func(pt *ui.Painter, r ui.Rect) {
			cx, cy := r.X+r.W/2, r.Y+r.H/2
			var p ui.Path
			if next {
				p.MoveTo(cx-3, cy-4).LineTo(cx+3, cy).LineTo(cx-3, cy+4)
			} else {
				p.MoveTo(cx+3, cy-4).LineTo(cx-3, cy).LineTo(cx+3, cy+4)
			}
			pt.StrokePath(&p, 1.8, fg)
		})
	})
	if disabled {
		b.Disabled(true)
	}
	if b.Clicked() {
		goTo(c, p, n, to)
	}
}

// dots is the position indicator below the track: a wide pill for the
// slide showing, small dots for the others, each jumping to its slide.
func dots(c *ui.Context, t *ui.Theme, p Props, n int) {
	ui.Row(c).FillWidth().Justify(ui.Center).Gap(6).Children(func() {
		for i := 0; i < n; i++ {
			active := p.State != nil && p.State.Active == i
			w := float32(6)
			bg := t.SurfaceHover
			if active {
				w, bg = 16, t.Text
			}
			dot := ui.ButtonBase(c).Width(w).Height(6).Radius(3).Background(bg)
			dot.Label(fmt.Sprintf("Go to slide %d", i+1))
			if dot.Clicked() {
				goTo(c, p, n, i)
			}
		}
	})
}

// goTo scrolls the track to slide i, keeping it within the slides.
func goTo(c *ui.Context, p Props, n, i int) {
	if p.Scroll == nil {
		return
	}
	i = clamp(i, 0, n-1)
	if n > 1 && p.Scroll.MaxX > 0 {
		p.Scroll.X = float32(i) * (p.Scroll.MaxX / float32(n-1))
	}
}

// clamp keeps i within [lo, hi].
func clamp(i, lo, hi int) int {
	if i < lo {
		return lo
	}
	if i > hi {
		return hi
	}
	return i
}
