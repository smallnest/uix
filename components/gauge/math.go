package gauge

import (
	"math"

	"github.com/egoist/mygo/ui"
)

// The geometry of the gauge, in the coordinates gauge-ui defines: "gauge
// degrees" where 0 points straight down (six o'clock) and values grow
// clockwise, so a sweep from 40° to 320° leaves an opening at the bottom.
// A reference radius is the unit every primitive measures itself against,
// exactly as an SVG viewBox would.

// degToRad turns gauge degrees into radians.
func degToRad(d float64) float64 { return d * math.Pi / 180 }

// polar returns the point on a circle of radius r at the gauge angle a,
// around the center (cx, cy).
func polar(cx, cy, r float32, a float64) (x, y float32) {
	t := degToRad(a + 90)
	return cx + r*float32(math.Cos(t)), cy + r*float32(math.Sin(t))
}

// clamp bounds v to [lo, hi].
func clamp(v, lo, hi float64) float64 { return math.Max(lo, math.Min(hi, v)) }

// foldValue folds v back inside [min, max) modulo the span, as a closed
// ring: 370° on a compass reads as 10°, and 25 o'clock as 1.
func foldValue(v, min, max float64) float64 {
	lo := math.Min(min, max)
	span := math.Abs(max - min)
	if span == 0 {
		return lo
	}
	m := math.Mod(v-lo, span)
	if m < 0 {
		m += span
	}
	return lo + m
}

// valueToAngle maps a value in [min, max] onto the sweep [startAngle,
// endAngle], clamped so a value past an end holds there.
func valueToAngle(v, min, max, startAngle, endAngle float64) float64 {
	span := max - min
	var t float64
	if span != 0 {
		t = (clamp(v, min, max) - min) / span
	}
	return startAngle + t*(endAngle-startAngle)
}

// valueToTurnedAngle maps a value onto the sweep the way a clock's hand
// does: one sweep per turn across the domain, and a turn more for each
// extra turn, so a minute hand goes round twelve times while the hour
// hand goes round once.
func valueToTurnedAngle(v, min, max, startAngle, endAngle float64, turns int) float64 {
	if turns <= 1 {
		return valueToAngle(v, min, max, startAngle, endAngle)
	}
	// A value past an end holds there, as it does with one turn, so the
	// hand of a clock stops where the domain does rather than sweeping on.
	// Only the turning is floorless.
	v = clamp(v, min, max)
	span := max - min
	var t float64
	if span != 0 {
		t = (v - min) / span
	}
	return startAngle + math.Mod(t*float64(turns), 1)*(endAngle-startAngle)
}

// stepValues are count+1 evenly spaced values across [min, max],
// inclusive of both ends. A count below one is one step, so the two ends
// alone; a negative count never reaches make.
func stepValues(min, max float64, count int) []float64 {
	if count < 0 {
		count = 0
	}
	n := math.Max(1, float64(count))
	out := make([]float64, 0, count+1)
	for i := 0; i <= count; i++ {
		out = append(out, min+(max-min)*float64(i)/n)
	}
	return out
}

// cutValues are the values between count equal segments, none at the two
// ends: ten glasses of water make nine cuts. Fewer than two segments
// leave no cut between them, so the result is empty.
func cutValues(min, max float64, count int) []float64 {
	if count < 2 {
		return nil
	}
	steps := stepValues(min, max, count)
	return steps[1 : len(steps)-1]
}

// arcTo draws an arc of radius r around (cx, cy) from the gauge angle a0
// to a1 onto path, in cubic curves of up to a quarter turn, so the curve
// stays smooth at any stroke width. Its tangent comes from the derivative
// of polar: (-sin(a+90°), cos(a+90°)) per radian.
func arcTo(path *ui.Path, cx, cy, r float32, a0, a1 float64) {
	// The angles arrive in degrees, so a quarter turn is 90 of them, and
	// the tangent's magnitude comes from the segment's sweep in radians.
	n := int(math.Ceil(math.Abs(a1-a0) / 90))
	if n == 0 {
		return
	}
	step := (a1 - a0) / float64(n)
	k := float32(4.0 / 3 * math.Tan(degToRad(step)/4))
	for i := range n {
		b0, b1 := a0+float64(i)*step, a0+float64(i+1)*step
		x0, y0 := polar(cx, cy, r, b0)
		x1, y1 := polar(cx, cy, r, b1)
		// The tangent is the derivative of polar, whose own +90° offset
		// carries into the sine and cosine, so the control points stay on
		// the circle rather than bowing inside it.
		t0x, t0y := -math.Sin(degToRad(b0+90)), math.Cos(degToRad(b0+90))
		t1x, t1y := -math.Sin(degToRad(b1+90)), math.Cos(degToRad(b1+90))
		path.CubeTo(x0+k*r*float32(t0x), y0+k*r*float32(t0y),
			x1-k*r*float32(t1x), y1-k*r*float32(t1y), x1, y1)
	}
}

// band adds the filled shape of a stroke of width w along the arc of
// radius r from a0 to a1 to path: the ring between r-w/2 and r+w/2,
// closed at both ends. round adds a half-disc at each end, the shape a
// round stroke cap would, as a subpath like the band itself.
func band(path *ui.Path, cx, cy, r, w float32, a0, a1 float64, round bool) {
	ro, ri := r+w/2, r-w/2
	if a1 <= a0 || ri <= 0 {
		return
	}
	ox0, oy0 := polar(cx, cy, ro, a0)
	path.MoveTo(ox0, oy0)
	arcTo(path, cx, cy, ro, a0, a1)
	ix1, iy1 := polar(cx, cy, ri, a1)
	path.LineTo(ix1, iy1)
	// The inner arc runs back the other way, so the subpath winds once
	// round the ring and the fill leaves its middle empty; both running
	// the same way would wind the ring twice and fill the whole disc.
	arcTo(path, cx, cy, ri, a1, a0)
	path.Close()
	if round {
		// A full disc of the stroke width on the centerline at each end:
		// the band already covers the half that faces inward.
		x0, y0 := polar(cx, cy, r, a0)
		path.Circle(x0, y0, w/2)
		x1, y1 := polar(cx, cy, r, a1)
		path.Circle(x1, y1, w/2)
	}
}

// minf is the smaller of two design-unit lengths.
func minf(a, b float32) float32 { return min(a, b) }
