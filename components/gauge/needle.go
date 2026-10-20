package gauge

import (
	"github.com/egoist/mygo/ui"
)

// End names which end of a range a needle or dot follows: the far end,
// which is the value itself on a single-value gauge, or the start.
type End int

const (
	// EndTo follows the far end of the reading.
	EndTo End = iota
	// EndFrom follows the start.
	EndFrom
)

// NeedleStyle is the shape of a needle.
type NeedleStyle int

const (
	// NeedleLine is a stroke with round ends.
	NeedleLine NeedleStyle = iota
	// NeedlePointer is a tapered blade, widest at the center, ending in a
	// rounded tip, with a blunt stub behind.
	NeedlePointer
	// NeedleCompass is a lozenge whose two halves can take different
	// colors.
	NeedleCompass
	// NeedleArrow is a shaft ending in a triangular arrowhead, for wind
	// or flow direction.
	NeedleArrow
)

// NeedleOptions draws the needle pointing at the value.
type NeedleOptions struct {
	// Style is the shape of the needle; NeedleLine when zero.
	Style NeedleStyle
	// Length is the length as a fraction of the reference radius; 0.85
	// when zero.
	Length float32
	// Width is the stroke width for NeedleLine, and the width at the
	// center for the filled styles; 6 when zero.
	Width float32
	// Color is the needle; the theme's text when zero.
	Color ui.Color
	// Tail is how far the needle extends behind the center, in design
	// units; 30 when zero. A negative Tail asks for none, which is how a
	// needle that only reaches forward is drawn — zero itself is the
	// default, not "no tail".
	Tail float32
	// TailColor is the color of the tail half of a pointer, compass or
	// arrow; the Color when zero.
	TailColor ui.Color
	// TailDot is the radius of a disc drawn at the end of the tail in
	// TailColor, like the counterweight on a wind vane. 0 leaves the
	// tail plain, and so does a tail of none: the dot rides the end of
	// the tail, and a needle with no tail has none.
	TailDot float32
	// Turns sweeps of the arc the needle makes across the domain: a
	// clock's minute hand turns 12 times while the hour hand turns once.
	Turns int
	// Gap is the radius of a clear circle around the center, in design
	// units, that splits the needle so a readout can sit between the
	// halves. 0 draws it whole.
	Gap float32
	// End is the end of a range the needle points at.
	End End
}

// Needle draws a needle pointing at the current value, as a gauge-ui
// GaugeNeedle. Add a Hub for the pivot, or a readout at the center.
func (g *Canvas) Needle(o NeedleOptions) {
	length := o.Length
	if length == 0 {
		length = 0.85
	}
	width := o.Width
	if width == 0 {
		width = 6
	}
	tail := o.Tail
	switch {
	case tail == 0:
		tail = 30
	case tail < 0:
		tail = 0
	}
	turns := o.Turns
	if turns == 0 {
		turns = 1
	}
	color := g.color(o.Color)
	tailColor := o.TailColor
	if tailColor == (ui.Color{}) {
		tailColor = color
	}
	// A gap is a length like any other, so a negative one is none rather
	// than a circle struck outside the needle.
	gap := o.Gap
	if gap < 0 {
		gap = 0
	}

	value := g.value
	if o.End == EndFrom {
		value = g.from
	}
	angle := valueToTurnedAngle(value, g.min, g.max, g.startAngle, g.endAngle, turns)
	reach := g.radius * length
	// A tail dot rides the end of the tail, so with no tail it has no end
	// to ride: the pivot disc is the Hub's, not the needle's.
	tailDot := o.TailDot
	if tail == 0 {
		tailDot = 0
	}

	if o.Style == NeedleLine {
		// The gap is cut out of the stroke by starting it that far along
		// the needle: a stroke has no shape to hollow, so the two halves
		// are two lengths of line rather than one line and a hole.
		inner := gap
		if inner < 0 {
			inner = 0
		}
		if inner < reach {
			tx, ty := g.xy(polar(0, 0, reach, angle))
			sx, sy := g.xy(polar(0, 0, inner, angle))
			var p ui.Path
			p.MoveTo(sx, sy).LineTo(tx, ty)
			g.pt.StrokePath(&p, g.width(width), color)
		}
		backEnd := -tail
		if backEnd < -inner {
			bx, by := g.xy(polar(0, 0, backEnd, angle))
			sx, sy := g.xy(polar(0, 0, -inner, angle))
			var p ui.Path
			p.MoveTo(bx, by).LineTo(sx, sy)
			g.pt.StrokePath(&p, g.width(width), tailColor)
		}
		if tailDot > 0 {
			bx, by := g.xy(polar(0, 0, -tail, angle))
			g.fillCircle(bx, by, tailDot, tailColor)
		}
		return
	}

	// The filled styles are two polygons meeting at the center, so the
	// tail can take its own color. Their shared base runs across the
	// pivot, and each end keeps a fifth of the width and is rounded off
	// with a disc of that size, as a round cap would, so a thin needle
	// never vanishes to a point.
	//
	// A gap moves the base they meet at out to the circle, each half
	// closed along that circle's chord: the chord is perpendicular to the
	// needle, so it is the same across-the-needle pair as the base, at
	// the gap's distance instead of zero.
	left := pt(width/2, angle-90)
	right := pt(width/2, angle+90)
	endRadius := width * 0.2
	endLeft := pt(endRadius, angle-90)
	endRight := pt(endRadius, angle+90)
	tip := pt(reach, angle)
	back := pt(-tail, angle)
	blade := func(center [2]float32) [][2]float32 {
		return [][2]float32{
			pin(center, endLeft),
			pin(center, endRight),
			right,
			left,
		}
	}

	if o.Style == NeedleArrow {
		// A shaft of Width capped by a triangular head that scales with
		// it, never longer than the needle itself. The shaft stops where
		// the head begins so the two never overlap, and the tail is the
		// shaft carried on behind the center in its own color, rounded
		// off at the end.
		headLength := minf(width*4, reach)
		base := pt(reach-headLength, angle)
		shaftLeft := pt(width/2, angle-90)
		shaftRight := pt(width/2, angle+90)
		headLeft := pt(width*2, angle-90)
		headRight := pt(width*2, angle+90)
		if tail > 0 {
			// The tail runs back from the gap, not from the center, so a
			// readout between the halves is not covered by either. Its
			// corners are the axis point carried sideways, which is what
			// pin does; a polar point at a negative radius would swing
			// round the center instead of staying on the axis.
			rearLeft := pin(back, shaftLeft)
			rearRight := pin(back, shaftRight)
			backInnerLeft := pin(pt(-gap, angle), shaftLeft)
			backInnerRight := pin(pt(-gap, angle), shaftRight)
			g.fillPoly(g.dev([][2]float32{
				backInnerLeft, backInnerRight, rearRight, rearLeft,
			}), tailColor)
			bx, by := g.xy(back[0], back[1])
			g.fillCircle(bx, by, width/2, tailColor)
		}
		frontLeft, frontRight := shaftLeft, shaftRight
		if gap > 0 {
			// The shaft starts at the circle rather than the pivot, so
			// the hole is left clear between the head and the tail.
			atGap := pt(gap, angle)
			frontLeft, frontRight = pin(atGap, shaftLeft), pin(atGap, shaftRight)
		}
		g.fillPoly(g.dev([][2]float32{
			frontLeft,
			pin(base, shaftLeft),
			pin(base, headLeft),
			tip,
			pin(base, headRight),
			pin(base, shaftRight),
			frontRight,
		}), color)
		if tailDot > 0 {
			bx, by := g.xy(back[0], back[1])
			g.fillCircle(bx, by, tailDot, tailColor)
		}
		return
	}

	// A pointer's tail is a blunt stub; a compass tapers the same both
	// ways.
	var rear [][2]float32
	if o.Style == NeedleCompass {
		rear = blade(back)
	} else {
		rear = [][2]float32{
			pin(back, [2]float32{left[0] * 0.6, left[1] * 0.6}),
			pin(back, [2]float32{right[0] * 0.6, right[1] * 0.6}),
			right,
			left,
		}
	}
	if tail > 0 {
		switch {
		case gap <= 0:
			g.fillPoly(g.dev(rear), tailColor)
		case o.Style == NeedleCompass:
			// A compass tapers the same both ways, so its tail is the
			// same blade, cut at the circle: the taper runs from the
			// half width across the needle to the fine end behind.
			outer := pt(-tail, angle)
			atGap := pt(-gap, angle)
			g.fillPoly(g.dev([][2]float32{
				pin(outer, endLeft), pin(outer, endRight),
				pin(atGap, endRight), pin(atGap, endLeft),
			}), tailColor)
		default:
			// A pointer's stub keeps its blunt end but starts at the
			// circle, so neither half covers the readout between them.
			outer := pt(-tail, angle)
			atGap := pt(-gap, angle)
			stub := func(p [2]float32) [2]float32 {
				return pin(p, [2]float32{left[0] * 0.6, left[1] * 0.6})
			}
			wide := func(p [2]float32) [2]float32 {
				return pin(p, [2]float32{right[0] * 0.6, right[1] * 0.6})
			}
			g.fillPoly(g.dev([][2]float32{
				stub(outer), wide(outer), wide(atGap), stub(atGap),
			}), tailColor)
		}
		if o.Style == NeedleCompass {
			bx, by := g.xy(back[0], back[1])
			g.fillCircle(bx, by, endRadius, tailColor)
		}
	}
	if gap <= 0 {
		g.fillPoly(g.dev(blade(tip)), color)
	} else {
		// The blade's wide end is the gap's chord rather than the pivot,
		// so the tip half starts outside the hole; the fine end stays at
		// the tip, where the disc below rounds it off.
		g.fillPoly(g.dev([][2]float32{
			pin(tip, endLeft), pin(tip, endRight), right, left,
		}), color)
	}
	tx, ty := g.xy(tip[0], tip[1])
	g.fillCircle(tx, ty, endRadius, color)
	if tailDot > 0 {
		bx, by := g.xy(back[0], back[1])
		g.fillCircle(bx, by, tailDot, tailColor)
	}
}

// pt is a design-unit point of the gauge, relative to the center.
func pt(r float32, a float64) [2]float32 {
	x, y := polar(0, 0, r, a)
	return [2]float32{x, y}
}

// pin is the sum of two design-unit points and a design-unit vector.
func pin(center, d [2]float32) [2]float32 {
	return [2]float32{center[0] + d[0], center[1] + d[1]}
}

// dev converts design-unit points, relative to the center, to device
// coordinates on the box.
func (g *Canvas) dev(pts [][2]float32) [][2]float32 {
	out := make([][2]float32, len(pts))
	for i, p := range pts {
		out[i][0], out[i][1] = g.xy(p[0], p[1])
	}
	return out
}

// fillPoly fills the closed polygon of the device points in c.
func (g *Canvas) fillPoly(pts [][2]float32, c ui.Color) {
	if len(pts) < 3 {
		return
	}
	var p ui.Path
	p.MoveTo(pts[0][0], pts[0][1])
	for _, pnt := range pts[1:] {
		p.LineTo(pnt[0], pnt[1])
	}
	p.Close()
	g.pt.FillPath(&p, c)
}

// HubOptions draws the pivot disc at the center.
type HubOptions struct {
	// Radius of the disc, in design units; 12 when zero, and a negative
	// one draws no disc.
	Radius float32
	// Color of the disc; the theme's text when zero.
	Color ui.Color
}

// Hub draws the pivot disc at the center, drawn once however many
// needles there are, as a gauge-ui GaugeHub. A negative Radius draws
// none, as gauge-ui's zero does.
func (g *Canvas) Hub(o HubOptions) {
	radius := o.Radius
	if radius == 0 {
		radius = 12
	}
	if radius < 0 {
		return
	}
	g.fillCircle(g.cx+g.dx, g.cy+g.dy, radius, g.color(o.Color))
}

// DotOptions draws a disc riding the arc.
type DotOptions struct {
	// Radius of the disc, in design units; 10 when zero, and a negative
	// one draws no disc.
	Radius float32
	// Color of the disc; the theme's text when zero.
	Color ui.Color
	// Opacity of the disc; 1 when zero.
	Opacity float32
	// Offset moves the disc's center from the reference radius.
	Offset float32
	// Turns sweeps of the arc the disc makes across the domain.
	Turns int
	// Halo is the width of a ring drawn behind the disc in HaloColor, so
	// it stands off whatever band it rides. 0 leaves the disc bare.
	Halo float32
	// HaloColor is the color of the halo: the surface under the gauge,
	// usually; the theme's background when zero.
	HaloColor ui.Color
	// End is the end of a range the disc rides; two dots, one at each
	// end, make the handles of a range.
	End End
}

// Dot draws a disc riding the arc at the current value, as a gauge-ui
// GaugeDot: the head of a progress ring, a body on its orbit, the handle
// on a dial. It follows the value the way a needle does.
func (g *Canvas) Dot(o DotOptions) {
	radius := o.Radius
	if radius == 0 {
		radius = 10
	}
	if radius < 0 {
		return
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}
	turns := o.Turns
	if turns == 0 {
		turns = 1
	}
	value := g.value
	if o.End == EndFrom {
		value = g.from
	}
	angle := valueToTurnedAngle(value, g.min, g.max, g.startAngle, g.endAngle, turns)
	x, y := g.point(g.radius+o.Offset, angle)
	haloColor := o.HaloColor
	if haloColor == (ui.Color{}) {
		haloColor = g.t.Background
	}
	if o.Halo > 0 {
		g.fillCircle(x, y, radius+o.Halo, haloColor.Alpha(opacity))
	}
	g.fillCircle(x, y, radius, g.color(o.Color).Alpha(opacity))
}
