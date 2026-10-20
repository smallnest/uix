// Package gauge draws composable gauges for the MyGo native toolkit, a
// port of the gauge-ui primitives: a track, arcs, zones, ticks, labels,
// a needle and its hub, a dot on the arc, and the value as text, all
// sharing one domain and one reference radius. Angles follow gauge-ui:
// "gauge degrees" where 0 points straight down and values grow clockwise,
// so a sweep from 40° to 320° leaves an opening at the bottom.
//
//	Gauge(c, Options{
//		Value: 72,
//		Transition: Transition{Enabled: true},
//		Draw: func(g *Canvas) {
//			g.Track(TrackOptions{})
//			g.Zones(ZonesOptions{Zones: []Zone{
//				{To: 60, Color: t.Success},
//				{To: 80, Color: t.Warning},
//				{Color: t.Danger},
//			}})
//			g.Ticks(TicksOptions{Count: 10})
//			g.Arc(ArcOptions{Color: t.Accent, Width: 24})
//			g.Needle(NeedleOptions{Style: NeedlePointer})
//			g.Hub(HubOptions{})
//			g.Value(ValueOptions{FontSize: 40, Y: 60})
//		},
//	})
//
// The gauge draws into a square box: the given Size, or the width it is
// given. Radius is the reference circle in design units, and the whole
// gauge scales to the box the way an SVG viewBox would, so a thicker
// stroke and a longer tick keep their proportion to the dial at any size.
// A Transition animates the value towards its target with MyGo's tween
// engine; the desktop's reduce-motion preference skips the motion.
package gauge

import (
	"math"
	"time"

	"github.com/egoist/mygo/ui"
)

// Options describes the gauge to draw: its domain, its reading, and the
// layers of the face, which the Draw function composes like gauge-ui's
// children.
type Options struct {
	// Value is the reading, or Range a span to fill between. On an open
	// domain a range is put in order; on a wrapped one it runs forward
	// from its first end, so [22, 6] on a 24-hour ring is the night.
	Value float64
	Range *[2]float64

	// Min and Max bound the domain; Min defaults to 0 and Max to 100.
	Min, Max float64
	// StartAngle and EndAngle are the sweep in gauge degrees, clockwise
	// from six o'clock. Both zero draws the default 40° to 320°.
	StartAngle, EndAngle float64
	// Wrap treats the domain as a closed ring, as a compass rose or clock
	// face is: a value folds back inside the domain instead of clamping,
	// so 370° on a compass reads as 10°.
	Wrap bool
	// Rotate turns the whole scale this many degrees clockwise, the way a
	// compass card turns under a fixed lubber line. It is applied without
	// animating.
	Rotate float64

	// Radius is the reference circle in design units; 160 when zero. The
	// whole gauge scales to its box, so Radius and every primitive's
	// lengths are proportions, not pixels.
	Radius float32
	// Padding is the room around the reference circle for thick strokes,
	// ticks and labels, in design units; 48 when zero.
	Padding float32
	// Size is the side of the square box, in pixels. Zero fills the width
	// it is given with a square.
	Size float32

	// Transition animates the value towards its target. Off by default;
	// true moves it over the house duration, and the fields tune it.
	Transition Transition
	// InitialValue and InitialRange are where the value starts on mount
	// when animating, for a sweep-in effect: the reading opens out from
	// there. A single number starts both ends of a range at it.
	InitialValue *float64
	InitialRange *[2]float64

	// Label names the gauge for assistive technology, and Description
	// describes it; the glyphs the primitives paint are not text the
	// accessibility tree can read, so a gauge that matters should carry
	// them. The Value primitive appends its own reading to the label, so
	// a gauge of a value needs no Label of its own.
	Label, Description string

	// Draw composes the face; nil draws only the empty box.
	Draw func(g *Canvas)
}

// Transition animates the value towards its target. Zero disables the
// motion; the value snaps to its target as gauge-ui without a transition
// does.
type Transition struct {
	// Enabled turns the motion on.
	Enabled bool
	// Duration is how long the value takes to reach its target; 600 ms
	// when zero.
	Duration time.Duration
	// Ease is how the value goes; EaseOut when nil.
	Ease ui.Easing
}

func (t Transition) duration() time.Duration {
	if t.Duration > 0 {
		return t.Duration
	}
	return 600 * time.Millisecond
}

func (t Transition) ease() ui.Easing {
	if t.Ease != nil {
		return t.Ease
	}
	return ui.EaseOut
}

// Gauge draws the gauge into a square box and returns it. The box is
// o.Size by o.Size when Size is set, and a square filling the width it is
// given otherwise, so a view can size it with Grow.
func Gauge(c *ui.Context, o Options) ui.Element {
	t := c.Theme()
	radius := o.Radius
	if radius <= 0 {
		radius = 160
	}
	padding := o.Padding
	if padding == 0 {
		padding = 48
	}
	design := 2 * (radius + padding)

	lo, hi := math.Min(o.Min, o.Max), math.Max(o.Min, o.Max)
	// Min and Max default to 0 and 100, so a gauge of just a value reads
	// as a percentage, as gauge-ui's do.
	if o.Min == 0 && o.Max == 0 {
		hi = 100
	}
	wrap := o.Wrap
	settle := func(v float64) float64 {
		if wrap {
			return foldValue(v, lo, hi)
		}
		return clamp(v, lo, hi)
	}
	rangeMode := o.Range != nil
	targetFrom, targetTo := lo, settle(o.Value)
	startFrom, startTo := targetFrom, targetTo
	if rangeMode {
		if wrap {
			targetFrom, targetTo = settle(o.Range[0]), settle(o.Range[1])
		} else {
			targetFrom, targetTo = settle(math.Min(o.Range[0], o.Range[1])), settle(math.Max(o.Range[0], o.Range[1]))
		}
		if o.InitialRange != nil {
			startFrom, startTo = settle(o.InitialRange[0]), settle(o.InitialRange[1])
		} else {
			startFrom, startTo = targetFrom, targetTo
		}
	} else if o.InitialValue != nil {
		startFrom, startTo = settle(*o.InitialValue), settle(*o.InitialValue)
	}

	e := ui.Box(c)
	if o.Size > 0 {
		e.Width(o.Size).Height(o.Size)
	} else {
		e.FillWidth().AspectRatio(1).MinWidth(0)
	}
	// The painted glyphs are not text the accessibility tree can read, so
	// the gauge carries the reading as its label; a Value primitive adds
	// the numbers it draws to it.
	if o.Label != "" {
		e.Label(o.Label)
	}
	if o.Description != "" {
		e.Description(o.Description)
	}
	m := ui.Local(e, "gauge.motion", func() motion { return motion{} })
	shownFrom, shownTo := animateValue(c, e, "gauge.progress", m, startFrom, startTo, targetFrom, targetTo, o.Transition)
	// A wrapped domain folds what a tween overshoots, so a bounce past
	// the seam reads as the value past it, as the spring of gauge-ui does.
	if wrap {
		shownFrom, shownTo = foldValue(shownFrom, lo, hi), foldValue(shownTo, lo, hi)
	}
	startAngle, endAngle := angles(o.StartAngle, o.EndAngle)
	startAngle += o.Rotate
	endAngle += o.Rotate

	e.Draw(func(pt *ui.Painter, r ui.Rect) {
		d := float32(math.Min(float64(r.W), float64(r.H)))
		if d <= 0 || design <= 0 {
			return
		}
		g := &Canvas{
			c:          c,
			t:          t,
			pt:         pt,
			sc:         d / design,
			cx:         r.X + r.W/2,
			cy:         r.Y + r.H/2,
			value:      shownTo,
			from:       shownFrom,
			rangeMode:  rangeMode,
			wrap:       wrap,
			min:        lo,
			max:        hi,
			startAngle: startAngle,
			endAngle:   endAngle,
			radius:     radius,
		}
		if o.Draw != nil {
			o.Draw(g)
		}
	})
	return e
}

// angles is the sweep: the default 40° to 320° when neither is set.
func angles(start, end float64) (float64, float64) {
	if start == 0 && end == 0 {
		return 40, 320
	}
	return start, end
}

// Canvas is the geometry and reading the primitives draw against. The
// Draw function of an Options receives one and composes the face with its
// methods; everything it draws is scaled from design units to the box.
type Canvas struct {
	c *ui.Context
	t *ui.Theme
	// pt paints the face; it is only valid during the Draw callback.
	pt *ui.Painter
	// sc scales design units to the box; cx and cy center it, and dx and
	// dy offset it further for an inset.
	sc, cx, cy, dx, dy float32
	// value is the reading being drawn this frame, from the start of a
	// range, and rangeMode whether it is a range rather than a value.
	value, from     float64
	rangeMode, wrap bool
	min, max        float64
	startAngle      float64
	endAngle        float64
	radius          float32
}

// point returns the position on the box of the design-unit radius r at
// the gauge angle a, centered on the box and scaled to it.
func (g *Canvas) point(r float32, a float64) (float32, float32) {
	x, y := polar(0, 0, r, a)
	return g.cx + g.dx + g.sc*x, g.cy + g.dy + g.sc*y
}

// xy converts a design-unit point relative to the center to device
// coordinates on the box.
func (g *Canvas) xy(x, y float32) (float32, float32) {
	return g.cx + g.dx + g.sc*x, g.cy + g.dy + g.sc*y
}

// width scales a design-unit stroke width to the box.
func (g *Canvas) width(w float32) float32 { return w * g.sc }

// fontSize scales a design-unit font size to the box.
func (g *Canvas) fontSize(s float32) float32 { return s * g.sc }

// angleOf maps a domain value onto the sweep.
func (g *Canvas) angleOf(v float64) float64 {
	return valueToAngle(v, g.min, g.max, g.startAngle, g.endAngle)
}

// fillBand draws the filled arc band of width w around the design-unit
// radius r between the gauge angles a0 and a1 in c, rounding the ends as
// round asks.
func (g *Canvas) fillBand(r, w float32, a0, a1 float64, c ui.Color, round bool) {
	wd := g.width(w)
	rd := r * g.sc
	var p ui.Path
	band(&p, g.cx+g.dx, g.cy+g.dy, rd, wd, a0, a1, round)
	g.pt.FillPath(&p, c)
}

// fillCircle fills a disc of the design-unit radius r centered at the
// position of the radius rp at the angle a, in c.
func (g *Canvas) fillCircle(cx, cy, r float32, c ui.Color) {
	var p ui.Path
	p.Circle(cx, cy, g.width(r))
	g.pt.FillPath(&p, c)
}

// Inset draws a whole gauge inside this one: a fuel dial under a
// speedometer, a second ring inside the first. It takes its own value
// and domain, and draws at the design-unit center (o.X, o.Y) at the
// fraction o.Scale of the host, so every primitive works inside it
// unchanged. The inset's reading draws at its value, without animating.
func (g *Canvas) Inset(o InsetOptions) {
	scale := o.Scale
	if scale == 0 {
		// A zero scale would collapse the inset: every width and radius of
		// its face multiplies by it, so nothing would reach the painter.
		// The documented default of a third is the useful reading.
		scale = 0.3
	}
	ig := *g
	ig.dx = g.dx + g.sc*o.X
	ig.dy = g.dy + g.sc*o.Y
	ig.sc = g.sc * scale
	lo, hi := math.Min(o.Min, o.Max), math.Max(o.Min, o.Max)
	if o.Min == 0 && o.Max == 0 {
		hi = 100
	}
	if o.Radius > 0 {
		ig.radius = o.Radius
	}
	ig.min, ig.max = lo, hi
	ig.startAngle, ig.endAngle = angles(o.StartAngle, o.EndAngle)
	ig.wrap = o.Wrap
	ig.rangeMode = o.Range != nil
	ig.from = lo
	ig.value = o.Value
	if o.Range != nil {
		if o.Wrap {
			ig.from, ig.value = foldValue(o.Range[0], lo, hi), foldValue(o.Range[1], lo, hi)
		} else {
			ig.from, ig.value = clamp(math.Min(o.Range[0], o.Range[1]), lo, hi), clamp(math.Max(o.Range[0], o.Range[1]), lo, hi)
		}
	} else if o.Wrap {
		ig.value = foldValue(o.Value, lo, hi)
	} else {
		ig.value = clamp(o.Value, lo, hi)
	}
	if o.Draw != nil {
		o.Draw(&ig)
	}
}

// InsetOptions places a nested gauge inside its host.
type InsetOptions struct {
	// X and Y center the inset in the host gauge's design units; positive
	// Y is down.
	X, Y float32
	// Scale is the size relative to the host: 0.3 draws it at a third.
	Scale float32
	// Radius is the inset's own reference radius; the host's when zero.
	Radius float32

	// Value, Range, Min, Max, StartAngle, EndAngle and Wrap set the
	// inset's own domain, as the Options fields do.
	Value                float64
	Range                *[2]float64
	Min, Max             float64
	StartAngle, EndAngle float64
	Wrap                 bool

	// Draw composes the inset's face.
	Draw func(g *Canvas)
}

// motion is the gauge's own animation state, kept on the element.
type motion struct {
	seen                 bool
	fromFrom, fromTo     float64 // where the reading was when the current move began
	targetFrom, targetTo float64 // where the reading is heading
}

// animateValue follows the reading towards its target along the
// transition, keeping frames coming while it moves. A change of target
// mid-move resamples the ongoing motion and sets off again, so a value
// dragged by a slider is followed rather than restarted. With the motion
// reduced, or no transition, it returns the target at once.
func animateValue(c *ui.Context, e ui.Element, key string, m *motion, startFrom, startTo, targetFrom, targetTo float64, o Transition) (float64, float64) {
	if o.Enabled && o.duration() > 0 && !c.Preferences().ReduceMotion && finite(targetFrom) && finite(targetTo) {
		if !m.seen || m.targetFrom != targetFrom || m.targetTo != targetTo {
			// A new target takes over from where the reading is now, so a
			// value dragged by a slider is followed rather than restarted.
			// Sample the frame that is on screen before adopting the new
			// one, as MyGo's own charts do, or the first frame of the move
			// would show the target already.
			fromFrom, fromTo := startFrom, startTo
			if m.seen {
				fromFrom, fromTo = shown(m.fromFrom, m.fromTo, m.targetFrom, m.targetTo, e.AnimateWith(key, 1, o.duration(), o.ease()))
			}
			m.seen = true
			m.fromFrom, m.fromTo = fromFrom, fromTo
			m.targetFrom, m.targetTo = targetFrom, targetTo
			e.Jump(key, 0)
		}
		prog := e.AnimateWith(key, 1, o.duration(), o.ease())
		if e.Settled(key) {
			return targetFrom, targetTo
		}
		return shown(m.fromFrom, m.fromTo, m.targetFrom, m.targetTo, prog)
	}
	m.seen = true
	m.targetFrom, m.targetTo = targetFrom, targetTo
	return targetFrom, targetTo
}

// shown is the reading partway along the move, at the progress prog.
func shown(fromFrom, fromTo, targetFrom, targetTo float64, prog float32) (float64, float64) {
	return fromFrom + (targetFrom-fromFrom)*float64(prog),
		fromTo + (targetTo-fromTo)*float64(prog)
}

// finite reports whether v is not NaN or infinite.
func finite(v float64) bool { return !math.IsNaN(v) && !math.IsInf(v, 0) }
