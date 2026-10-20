package gauge

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// Cap is how the end of a band is shaped, as a stroke line cap: rounded
// or square.
type Cap int

const (
	// CapButt ends a band squarely, as the default of a stroke.
	CapButt Cap = iota
	// CapRound ends a band with a half-disc, as a round stroke cap.
	CapRound
)

// TrackOptions draws the background arc of the whole sweep.
type TrackOptions struct {
	// Width of the band, in design units; 24 when zero.
	Width float32
	// Color of the band; the theme's text when zero.
	Color ui.Color
	// Opacity of the band, the quiet of a track; 0.15 when zero.
	Opacity float32
	// Offset moves the band radially from the reference radius; positive
	// moves outward.
	Offset float32
}

// Track draws the full background arc from the start to the end of the
// sweep, as a gauge-ui GaugeTrack.
func (g *Canvas) Track(o TrackOptions) {
	width := o.Width
	if width == 0 {
		width = 24
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 0.15
	}
	g.fillBand(g.radius+o.Offset, width, g.startAngle, g.endAngle, g.color(o.Color).Alpha(opacity), true)
}

// ArcOptions draws the filled arc of the value.
type ArcOptions struct {
	// Width of the band, in design units; 24 when zero.
	Width float32
	// Color of the band; the theme's text when zero.
	Color ui.Color
	// Opacity of the band; 1 when zero.
	Opacity float32
	// Offset moves the band radially from the reference radius.
	Offset float32
	// From anchors the arc at a domain value instead of the start of the
	// reading, and To ends it at one instead of the value.
	From, To *float64
	// Reverse grows the arc back from the end of the sweep instead of
	// forward from its start: the same length, anchored at the far end,
	// which is what a dial filling towards its start needs.
	Reverse bool
	// Wrap carries the arc on round a closed ring instead of clamping it
	// at the domain: a span from eleven at night to two in the morning
	// crosses the seam in one piece.
	Wrap bool
}

// Arc draws the filled arc from the start of the reading (or From) to
// the value (or To), as a gauge-ui GaugeArc. The ends may arrive either
// way round, and on a wrapped gauge the arc runs forward from the start
// across the seam where the reading does.
func (g *Canvas) Arc(o ArcOptions) {
	width := o.Width
	if width == 0 {
		width = 24
	}
	tail := g.from
	if o.From != nil {
		tail = *o.From
	}
	head := g.value
	if o.To != nil {
		head = *o.To
	}
	// A range on a closed ring is read forward from its start, so the
	// far end is carried round past the seam when it sits below the
	// start, as gauge-ui's GaugeArc does.
	forward := g.rangeMode && g.wrap && o.From == nil && o.To == nil
	if forward {
		head = tail + foldValue(head-tail, 0, math.Abs(g.max-g.min))
	}
	lo, hi := math.Min(tail, head), math.Max(tail, head)

	var a0, a1 float64
	if o.Wrap || forward {
		span := g.max - g.min
		if span <= 0 {
			return
		}
		perUnit := (g.endAngle - g.startAngle) / span
		a0 = g.startAngle + foldValue(lo-g.min, 0, span)*perUnit
		a1 = a0 + math.Min((hi-lo)*perUnit, 360)
	} else if o.Reverse {
		a0 = g.angleOf(g.max - (hi - lo))
		a1 = g.angleOf(g.max)
	} else {
		a0 = g.angleOf(lo)
		a1 = g.angleOf(hi)
	}
	if a1 <= a0 {
		return
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}
	g.fillBand(g.radius+o.Offset, width, a0, a1, g.color(o.Color).Alpha(opacity), true)
}

// StackPart is one part of a GaugeStack.
type StackPart struct {
	// Weight is the part's share of the fill. Weighted parts split the
	// value between them in proportion.
	Weight float64
	// To is a fixed upper cutoff in domain units, for a part that stands
	// still while the value sweeps across it; it runs from the end of the
	// part before. Parts with a To ignore Weight.
	To float64
	// HasTo tells To is set rather than zero, so a part can sit at zero.
	HasTo bool
	// Color fills the part.
	Color ui.Color
	// Opacity of the part; 1 when zero.
	Opacity float32
}

// StackOptions draws the value arc split into parts that sweep as one.
type StackOptions struct {
	// Parts are the pieces of the fill, in order. Use one kind or the
	// other in a single stack: weighted parts, or parts with fixed To.
	Parts []StackPart
	// Width of the band, in design units; 24 when zero.
	Width float32
	// Opacity of the whole stack; 1 when zero.
	Opacity float32
	// Offset moves the band radially from the reference radius.
	Offset float32
	// From starts the stack at a domain value instead of the start of
	// the reading.
	From *float64
}

// Stack draws the value arc split into parts: a weighted part takes its
// share of the fill frame by frame, so a memory ring keeps wired, app
// and cache in proportion; a part with a fixed To runs from the end of
// the one before and is clipped to the value, so a night of sleep stages
// draws on in order as the dial sweeps in.
func (g *Canvas) Stack(o StackOptions) {
	width := o.Width
	if width == 0 {
		width = 24
	}
	if len(o.Parts) == 0 {
		return
	}
	total := 0.0
	for _, p := range o.Parts {
		if !p.HasTo {
			total += p.Weight
		}
	}
	start := g.from
	if o.From != nil {
		start = *o.From
	}
	fill := g.value - start

	// Each part ends at its own cutoff, or its share of the fill after
	// the part before.
	ends := make([]float64, len(o.Parts))
	for i, p := range o.Parts {
		if p.HasTo {
			ends[i] = p.To
		} else {
			prev := start
			if i > 0 {
				prev = ends[i-1]
			}
			share := 0.0
			if total > 0 {
				share = fill * p.Weight / total
			}
			ends[i] = prev + share
		}
	}

	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}
	for i, p := range o.Parts {
		from := start
		if i > 0 {
			from = ends[i-1]
		}
		a0 := g.angleOf(from)
		a1 := g.angleOf(math.Min(ends[i], g.value))
		if a1 <= a0 {
			continue
		}
		// A part's own opacity, then the stack's, multiply as an SVG
		// group's opacity and a stroke's would.
		alpha := p.Opacity
		if alpha == 0 {
			alpha = 1
		}
		g.fillBand(g.radius+o.Offset, width, a0, a1, p.Color.Alpha(alpha).Alpha(opacity), false)
	}
}

// ZonesOptions draws the colored bands of the cutoffs.
type ZonesOptions struct {
	// Zones are the ordered bands: each runs from the previous zone's To
	// (or Min) up to its own To; the last always extends to Max.
	Zones []Zone
	// Width of the band, in design units; 8 when zero.
	Width float32
	// Opacity of the bands; 1 when zero.
	Opacity float32
	// Offset moves the band radially from the reference radius.
	Offset float32
	// Gap is the gap between neighboring zones, in gauge degrees.
	Gap float64
	// Cap shapes the joints between zones; CapButt when zero.
	Cap Cap
	// EndCap shapes the two outer ends of the whole band; the Cap when
	// nil, so a single zone still rounds or squares every end.
	EndCap *Cap
}

// Zones draws the colored bands marking the cutoffs of the domain, each
// zone from the previous cutoff to its own, as a gauge-ui GaugeZones.
// The joints and the outer ends may take their own caps, so a single
// value still rounds or squares every end.
func (g *Canvas) Zones(o ZonesOptions) {
	width := o.Width
	if width == 0 {
		width = 8
	}
	if len(o.Zones) == 0 {
		return
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}
	joint := o.Cap
	endCap := o.Cap
	if o.EndCap != nil {
		endCap = *o.EndCap
	}

	// Each zone runs from the previous cutoff to its own, clamped to the
	// domain.
	segments := make([]struct{ from, to float64 }, len(o.Zones))
	for i := range o.Zones {
		from := g.min
		if i > 0 {
			from = segments[i-1].to
		}
		to := g.max
		if i < len(o.Zones)-1 {
			to = clamp(o.Zones[i].To, from, g.max)
		}
		segments[i] = struct{ from, to float64 }{from, to}
	}

	for i, z := range o.Zones {
		seg := segments[i]
		a0 := g.angleOf(seg.from)
		a1 := g.angleOf(seg.to)
		if i > 0 {
			a0 += o.Gap / 2
		}
		if i < len(o.Zones)-1 {
			a1 -= o.Gap / 2
		}
		if a1 <= a0 {
			continue
		}
		g.fillBand(g.radius+o.Offset, width, a0, a1, z.Color.Alpha(opacity), false)
		// A round cap at an end is a disc of the stroke width on the
		// centerline, which is exactly the shape a round cap adds; drawn
		// separately, the joints and the outer ends can differ.
		first, last := i == 0, i == len(o.Zones)-1
		if first && endCap == CapRound {
			g.capDot(g.radius+o.Offset, width, a0, z.Color.Alpha(opacity))
		}
		if last && endCap == CapRound {
			g.capDot(g.radius+o.Offset, width, a1, z.Color.Alpha(opacity))
		}
		if !first && joint == CapRound {
			g.capDot(g.radius+o.Offset, width, a0, z.Color.Alpha(opacity))
		}
		if !last && joint == CapRound {
			g.capDot(g.radius+o.Offset, width, a1, z.Color.Alpha(opacity))
		}
	}
}

// capDot fills the disc of a round cap at the gauge angle a on the band.
func (g *Canvas) capDot(r, w float32, a float64, c ui.Color) {
	x, y := g.point(r, a)
	g.fillCircle(x, y, w/2, c)
}

// color is the primitive's color, the theme's text when the zero color
// is passed, as gauge-ui's currentColor resolves to the text color.
func (g *Canvas) color(c ui.Color) ui.Color {
	if c == (ui.Color{}) {
		return g.t.Text
	}
	return c
}
