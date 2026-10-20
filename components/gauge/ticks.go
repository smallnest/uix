package gauge

import (
	"math"
	"strconv"

	"github.com/egoist/mygo/ui"
)

// ValueColor colors the marks: a single Color for every mark, or a
// PerValue color of each mark worked out from its value, as the ticks
// inside a range or each tick in the color of its zone. The zero value
// uses the theme's text.
type ValueColor struct {
	// Color is the one color of every mark.
	Color ui.Color
	// PerValue is a color of each mark's value; it wins over Color.
	PerValue func(value float64) ui.Color
}

// paint is the color of the mark at the value.
func (g *Canvas) paint(vc ValueColor) func(float64) ui.Color {
	if vc.PerValue != nil {
		return vc.PerValue
	}
	if vc.Color != (ui.Color{}) {
		return func(float64) ui.Color { return vc.Color }
	}
	return func(float64) ui.Color { return g.t.Text }
}

// tickValues are the evenly spaced domain values of the ticks; on a
// closed ring the last step lands on the first, so it is dropped rather
// than drawn twice.
func (g *Canvas) tickValues(count int) []float64 {
	values := stepValues(g.min, g.max, count)
	if g.endAngle-g.startAngle >= 360 {
		values = values[:len(values)-1]
	}
	return values
}

// radialLines draws one line per value, straddling the reference radius
// by the value's half length, in the value's color. Every line of a call
// goes into one stroked path, so a stroke is one op whatever the tick
// count; the painter's own Line only reaches the axis-aligned.
func (g *Canvas) radialLines(values []float64, length float32, lengthOf func(float64) float32, width float32, color ValueColor, opacity float32, offset float32) {
	paint := g.paint(color)
	w := g.width(width)
	type seg struct {
		c      ui.Color
		x0, y0 float32
		x1, y1 float32
	}
	segs := make([]seg, 0, len(values))
	for _, v := range values {
		half := length / 2
		if lengthOf != nil {
			half = lengthOf(v) / 2
		}
		a := g.angleOf(v)
		x0, y0 := g.point(g.radius+offset-half, a)
		x1, y1 := g.point(g.radius+offset+half, a)
		segs = append(segs, seg{paint(v).Alpha(opacity), x0, y0, x1, y1})
	}
	for _, s := range segs {
		var p ui.Path
		p.MoveTo(s.x0, s.y0).LineTo(s.x1, s.y1)
		g.pt.StrokePath(&p, w, s.c)
	}
}

// TicksOptions draws the evenly spaced tick marks of the whole sweep.
type TicksOptions struct {
	// Count is the number of intervals across the domain; it draws
	// Count+1 ticks. 10 when zero.
	Count int
	// Length of each line, in design units; 12 when zero.
	Length float32
	// Width of each line, in design units; 2 when zero.
	Width float32
	// Color colors the ticks, one or per value; the theme's text when
	// zero.
	Color ValueColor
	// Opacity of the ticks; 1 when zero.
	Opacity float32
	// Offset moves the line's midpoint from the reference radius;
	// positive moves outward.
	Offset float32
}

// Ticks draws the evenly spaced tick marks across the whole arc, as a
// gauge-ui GaugeTicks. A count below one draws no marks at all, which
// lets the count come from configuration; only the zero value asks for
// the default.
func (g *Canvas) Ticks(o TicksOptions) {
	count := o.Count
	if count == 0 {
		count = 10
	}
	if count < 0 {
		return
	}
	length := o.Length
	if length == 0 {
		length = 12
	}
	width := o.Width
	if width == 0 {
		width = 2
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}
	g.radialLines(g.tickValues(count), length, nil, width, o.Color, opacity, o.Offset)
}

// FillMode is how a TickFill lights its ticks.
type FillMode int

const (
	// FillModeFill lights every tick between the ends, as an arc would
	// fill.
	FillModeFill FillMode = iota
	// FillModeNearest lights only the tick nearest each end: on a single
	// value a pointer made of ticks, on a range one at either end.
	FillModeNearest
)

// TickFillOptions draws the fill made of ticks.
type TickFillOptions struct {
	// Count is the number of intervals across the domain; 40 when zero.
	Count int
	// From starts the lit stretch, defaulting to the start of the
	// reading, and To ends it, defaulting to the value.
	From, To *float64
	// Mode lights the fill or only the tick nearest each end.
	Mode FillMode
	// Color colors the lit ticks, one or per value; the theme's text
	// when zero.
	Color ValueColor
	// DimColor colors the ticks left unlit; a faint text when zero.
	DimColor ValueColor
	// Length is the length of a tick and LitLength that of a lit one,
	// for meters whose live segments stand out. 12 and Length when zero.
	Length, LitLength float32
	// Width of each line; 2 when zero.
	Width float32
	// Opacity of the ticks; 1 when zero.
	Opacity float32
	// Offset moves the line's midpoint from the reference radius.
	Offset float32
}

// TickFill lights the ticks between the ends of the value while the rest
// stay dim, as a gauge-ui GaugeTickFill. It follows the value frame by
// frame as an Arc does, fills a range between its ends, and on a wrapped
// gauge runs forward from the start across the seam. With ModeNearest it
// lights a single tick at each end instead, for a pointer drawn in
// ticks.
func (g *Canvas) TickFill(o TickFillOptions) {
	count := o.Count
	if count == 0 {
		count = 40
	}
	if count < 0 {
		return
	}
	length := o.Length
	if length == 0 {
		length = 12
	}
	litLength := o.LitLength
	if litLength == 0 {
		litLength = length
	}
	width := o.Width
	if width == 0 {
		width = 2
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}

	values := g.tickValues(count)
	tail := g.from
	if o.From != nil {
		tail = *o.From
	}
	head := g.value
	if o.To != nil {
		head = *o.To
	}
	span := math.Abs(g.max - g.min)
	// Slack for ticks that land on an end, whose step values carry
	// floating point error.
	eps := span * 1e-6

	// ahead is how far round from the start a value sits, on a closed
	// ring; between whether it lies between the ends.
	ahead := func(v float64) float64 { return Fold(v-tail+eps, 0, span) }
	between := func(v float64) bool {
		if g.wrap {
			return ahead(v) <= ahead(head)+eps
		}
		return v >= math.Min(tail, head)-eps && v <= math.Max(tail, head)+eps
	}
	// nearestTo is the tick nearest a value, measured round the seam on a
	// ring.
	nearestTo := func(target float64) float64 {
		gap := func(v float64) float64 {
			d := math.Abs(v - target)
			if g.wrap {
				return math.Min(d, span-d)
			}
			return d
		}
		best := values[0]
		for _, v := range values[1:] {
			if gap(v) < gap(best) {
				best = v
			}
		}
		return best
	}

	lit := make(map[float64]bool, len(values))
	switch o.Mode {
	case FillModeNearest:
		ends := []float64{nearestTo(head)}
		if g.rangeMode || o.From != nil {
			ends = append(ends, nearestTo(tail))
		}
		for _, v := range values {
			for _, e := range ends {
				if v == e {
					lit[v] = true
				}
			}
		}
	default:
		for _, v := range values {
			lit[v] = between(v)
		}
	}

	paint := g.paint(o.Color)
	dim := o.DimColor
	if dim.PerValue == nil && dim.Color == (ui.Color{}) {
		dim = ValueColor{Color: g.t.Text.Alpha(0.15)}
	}
	g.radialLines(values, length, func(v float64) float32 {
		if lit[v] {
			return litLength
		}
		return length
	}, width, ValueColor{PerValue: func(v float64) ui.Color {
		if lit[v] {
			return paint(v)
		}
		return g.paint(dim)(v)
	}}, opacity, o.Offset)
}

// MarksOptions draws the radial marks at the cutoffs.
type MarksOptions struct {
	// Values are the domain values to mark, for example the cutoffs
	// between zones. Nil draws the equal cuts of Count instead.
	Values []float64
	// Count cuts the domain into that many equal segments; the marks
	// fall between the segments, none at the two ends. On a closed ring
	// the seam is cut too.
	Count int
	// Length, Width, Color, Opacity and Offset are as the Ticks'.
	Length, Width float32
	Color         ValueColor
	Opacity       float32
	Offset        float32
}

// Marks draws radial marks at arbitrary domain values, or at equal cuts
// of the domain, as a gauge-ui GaugeMarks. Drawn in the surface's color
// over a band, they slice it the way a water ring is cut into glasses.
func (g *Canvas) Marks(o MarksOptions) {
	length := o.Length
	if length == 0 {
		length = 12
	}
	width := o.Width
	if width == 0 {
		width = 2
	}
	opacity := o.Opacity
	if opacity == 0 {
		opacity = 1
	}
	marks := o.Values
	if marks == nil && o.Count > 1 {
		marks = cutValues(g.min, g.max, o.Count)
		if g.endAngle-g.startAngle >= 360 {
			// The seam is a cut like any other on a closed ring, since the
			// last segment meets the first there; it is the end of the
			// domain, which the equal cuts stop short of.
			marks = append([]float64{g.max}, marks...)
		}
	}
	if len(marks) == 0 {
		return
	}
	g.radialLines(marks, length, nil, width, o.Color, opacity, o.Offset)
}

// TickLabelsOptions draws the numeric labels at the evenly spaced
// values.
type TickLabelsOptions struct {
	// Count is the number of intervals across the domain; 10 when zero.
	Count int
	// Offset moves the label center from the reference radius; -48 when
	// zero, which sits the labels just inside the dial.
	Offset float32
	// FontSize is the label size in design units; 16 when zero.
	FontSize float32
	// Font is the family; the theme's when zero.
	Font FontFamily
	// Color colors every label, or each by its value; the theme's text
	// when zero.
	Color ValueColor
	// Weight is the font weight of the labels; 500 when zero.
	Weight int
	// Decimals rounds the label; 0 when zero.
	Decimals int
	// Format renders a label in place of the rounded value.
	Format func(value float64) string
}

// TickLabels draws the numeric labels at evenly spaced values across the
// arc, as a gauge-ui GaugeTickLabels. A count below one draws no labels,
// as the Ticks' does.
func (g *Canvas) TickLabels(o TickLabelsOptions) {
	count := o.Count
	if count == 0 {
		count = 10
	}
	if count < 0 {
		return
	}
	offset := o.Offset
	if offset == 0 {
		offset = -48
	}
	size := o.FontSize
	if size == 0 {
		size = 16
	}
	weight := o.Weight
	if weight == 0 {
		weight = 500
	}
	paint := g.paint(o.Color)
	for _, v := range g.tickValues(count) {
		label := strconv.FormatFloat(v, 'f', o.Decimals, 64)
		if o.Format != nil {
			label = o.Format(v)
		}
		x, y := g.point(g.radius+offset, g.angleOf(v))
		g.textAt(label, x, y, size, weight, TextMiddle, paint(v), o.Font)
	}
}
