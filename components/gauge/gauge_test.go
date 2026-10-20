package gauge

import (
	"bytes"
	"math"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// frame renders a gauge of o in a square box and returns the tester. The
// theme it drew with is kept for lit, which needs the background.
var lastBackground ui.Color

func frame(o Options) *ui.Tester {
	if o.Size == 0 {
		o.Size = 300
	}
	tt := ui.NewTester(func(c *ui.Context) {
		lastBackground = c.Theme().Background
		ui.Column(c).Fill().Padding(16).Children(func() { Gauge(c, o) })
	}, int(o.Size)+32, int(o.Size)+32)
	tt.Frame()
	return tt
}

// lit counts the pixels that are not the background, so a demo that drew
// nothing shows as an empty frame.
func lit(tt *ui.Tester) int {
	img := tt.Image()
	b := img.Bounds()
	n := 0
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			r, g, bl, _ := img.At(x, y).RGBA()
			if !near(ui.Color{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(bl >> 8)}, lastBackground) {
				n++
			}
		}
	}
	return n
}

// near reports whether c is within a small distance of o, so a pixel
// antialiased at a band's edge still reads as the color behind it.
func near(c, o ui.Color) bool {
	d := math.Abs(float64(c.R)-float64(o.R)) + math.Abs(float64(c.G)-float64(o.G)) + math.Abs(float64(c.B)-float64(o.B))
	return d < 12
}

// TestRendersAValue draws a full face and checks it puts ink on the
// canvas, so a gauge that silently drew nothing fails.
func TestRendersAValue(t *testing.T) {
	tt := frame(Options{Value: 72, Label: "Speed", Draw: func(g *Canvas) {
		g.Track(TrackOptions{Width: 20})
		g.Arc(ArcOptions{Width: 20})
		g.Needle(NeedleOptions{Style: NeedlePointer})
		g.Hub(HubOptions{})
		g.Value(ValueOptions{Y: 40})
	}})
	if n := lit(tt); n < 2000 {
		t.Fatalf("the gauge drew %d pixels, want a full face", n)
	}
	// The readout is painted, not an element text; the Label is what the
	// accessibility tree reads.
	if !tt.HasText("Speed") {
		t.Fatalf("the label is missing: %q", tt.Texts())
	}
}

// TestValueMovesTheFill draws two values and requires the frames to
// differ, so the value reaches the arc.
func TestValueMovesTheFill(t *testing.T) {
	draw := func(v float64) *ui.Tester {
		return frame(Options{Value: v, Draw: func(g *Canvas) {
			g.Arc(ArcOptions{Width: 20})
		}})
	}
	low, high := draw(20), draw(80)
	if bytes.Equal(low.Image().Pix, high.Image().Pix) {
		t.Fatal("the arc does not change with the value")
	}
}

// TestClampsAndDefaults drives values past both ends of the domain and
// the zero Min and Max; neither may panic, and the default domain is
// 0 to 100, so 150 fills as 100.
func TestClampsAndDefaults(t *testing.T) {
	tests := []Options{
		{Value: 150},
		{Value: -20},
		{Value: 50, Min: 100, Max: 0},
		{Value: 0},
	}
	for i, o := range tests {
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Fatalf("case %d panicked: %v", i, r)
				}
			}()
			frame(o)
		}()
	}
	full := frame(Options{Value: 100, Draw: func(g *Canvas) { g.Arc(ArcOptions{Width: 20}) }})
	over := frame(Options{Value: 150, Draw: func(g *Canvas) { g.Arc(ArcOptions{Width: 20}) }})
	if !bytes.Equal(full.Image().Pix, over.Image().Pix) {
		t.Fatal("a value past Max does not clamp to Max")
	}
}

// TestZonesAndStackColorTheBands draws the zones and the stack, and
// checks that each paints near its own color, so a band that took the
// wrong color fails.
func TestZonesAndStackColorTheBands(t *testing.T) {
	green := ui.RGB(20, 200, 90)
	red := ui.RGB(230, 40, 60)
	tt := frame(Options{Value: 90, Draw: func(g *Canvas) {
		g.Zones(ZonesOptions{Zones: []Zone{{To: 50, Color: green}, {Color: red}}, Width: 24})
	}})
	if !hasColor(tt, green) || !hasColor(tt, red) {
		t.Fatal("the zones did not draw both colors")
	}
	tt = frame(Options{Value: 100, Draw: func(g *Canvas) {
		g.Stack(StackOptions{Width: 24, Parts: []StackPart{
			{Weight: 1, Color: green}, {Weight: 1, Color: red},
		}})
	}})
	if !hasColor(tt, green) || !hasColor(tt, red) {
		t.Fatal("the stack did not draw both parts")
	}
}

// hasColor reports whether any pixel is within a small distance of c,
// which antialiasing of a nearby edge keeps intact at the core of a band.
func hasColor(tt *ui.Tester, c ui.Color) bool {
	img := tt.Image()
	b := img.Bounds()
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			r, g, bl, _ := img.At(x, y).RGBA()
			if near(ui.Color{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(bl >> 8)}, c) {
				return true
			}
		}
	}
	return false
}

// count is how many pixels of the frame take the color c, so a test can
// weigh a stretch of ink rather than only notice that it exists.
func count(tt *ui.Tester, c ui.Color) int {
	img := tt.Image()
	b := img.Bounds()
	n := 0
	for x := b.Min.X; x < b.Max.X; x++ {
		for y := b.Min.Y; y < b.Max.Y; y++ {
			r, g, bl, _ := img.At(x, y).RGBA()
			if near(ui.Color{R: uint8(r >> 8), G: uint8(g >> 8), B: uint8(bl >> 8)}, c) {
				n++
			}
		}
	}
	return n
}

// TestTickFillLightsTheValue draws the tick fill over the whole sweep
// and checks the lit stretch takes the lit color while the rest does
// not, so a fill that painted every tick fails. It also moves the value
// and requires the lit stretch to grow, so a fill that ignored the value
// and lit a fixed stretch fails too.
func TestTickFillLightsTheValue(t *testing.T) {
	lit_ := ui.RGB(255, 0, 0)
	dim := ui.RGB(0, 0, 255)
	draw := func(g *Canvas) {
		g.TickFill(TickFillOptions{
			Count: 20, Length: 24, Width: 5,
			Color:    ValueColor{Color: lit_},
			DimColor: ValueColor{Color: dim},
		})
	}
	tt := frame(Options{Value: 50, Draw: draw})
	if !hasColor(tt, lit_) {
		t.Fatal("no lit tick took the lit color")
	}
	if !hasColor(tt, dim) {
		t.Fatal("no dim tick took the dim color")
	}
	half := count(tt, lit_)
	full := count(frame(Options{Value: 100, Draw: draw}), lit_)
	if full <= half {
		t.Fatalf("a full value lit %d pixels of the lit color, no more than half's %d", full, half)
	}
	empty := count(frame(Options{Value: 0, Draw: draw}), lit_)
	// The tick under the start stays lit at zero: it is what a pointer
	// made of ticks looks like at rest. The dim stretch still takes every
	// other tick, so the two colors must not trade places.
	if empty >= half {
		t.Fatalf("an empty value lit %d pixels, as many as half's %d", empty, half)
	}
}

// TestRangeFillsBetweenTheEnds draws a range and its single value as the
// same far end, and requires the range's arc to start further along, so
// the near end reaches the fill.
func TestRangeFillsBetweenTheEnds(t *testing.T) {
	r := [2]float64{40, 80}
	rangeTT := frame(Options{Value: 80, Range: &r, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
	}})
	pointTT := frame(Options{Value: 80, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
	}})
	if bytes.Equal(rangeTT.Image().Pix, pointTT.Image().Pix) {
		t.Fatal("a range fills as much as a single value at its far end")
	}
}

// TestWrappedDomainFolds checks that a wrapped gauge folds a value past
// the end of the domain back to its start, as 370° on a compass reads as
// 10°: it draws as the value at the start, not pinned to the far end.
func TestWrappedDomainFolds(t *testing.T) {
	inside := frame(Options{Value: 250, Min: 0, Max: 300, Wrap: true, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
	}})
	// 550 is past the end of a 0-300 ring, so it folds back to 250.
	over := frame(Options{Value: 550, Min: 0, Max: 300, Wrap: true, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
	}})
	if !bytes.Equal(inside.Image().Pix, over.Image().Pix) {
		t.Fatal("a wrapped value past Max does not fold back inside the domain")
	}
	// An unwrapped gauge at the same value pins to Max instead, so the
	// fold is what made them agree.
	pinned := frame(Options{Value: 550, Min: 0, Max: 300, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
	}})
	if bytes.Equal(pinned.Image().Pix, over.Image().Pix) {
		t.Fatal("the wrapped and unwrapped gauges draw the same, so the test proves nothing")
	}
}

// TestInsetDrawsInside draws the same face with and without a nested
// gauge and requires the nested one to add ink, so an inset that drew
// nothing fails.
func TestInsetDrawsInside(t *testing.T) {
	plain := frame(Options{Value: 60, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
	}})
	nested := frame(Options{Value: 60, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 20})
		g.Inset(InsetOptions{Scale: 0.4, Value: 30, Draw: func(i *Canvas) {
			i.Arc(ArcOptions{Width: 20})
		}})
	}})
	if lit(nested) <= lit(plain) {
		t.Fatal("the inset drew nothing")
	}
}

// TestNeedleStylesRender draws every needle style, so a broken polygon
// fails rather than passing unnoticed.
func TestNeedleStylesRender(t *testing.T) {
	for _, s := range []NeedleStyle{NeedleLine, NeedlePointer, NeedleCompass, NeedleArrow} {
		tt := frame(Options{Value: 40, Draw: func(g *Canvas) {
			g.Needle(NeedleOptions{Style: s, TailDot: 5})
		}})
		if lit(tt) < 100 {
			t.Fatalf("needle style %d drew nothing", s)
		}
	}
}

// TestMarksAndLabelsDraw draws the marks at equal cuts and the tick
// labels, and requires each to put ink of its own on the canvas: the
// marks alone light pixels, and the labels alone light more, so neither
// line is dead weight. Counting only the union of the two would pass
// with either one of them missing.
func TestMarksAndLabelsDraw(t *testing.T) {
	marks := lit(frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.Marks(MarksOptions{Count: 4, Length: 16, Width: 3})
	}}))
	if marks < 100 {
		t.Fatalf("the marks drew %d pixels; four cuts of 16 by 3 should ink more", marks)
	}
	labels := lit(frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.TickLabels(TickLabelsOptions{Count: 5})
	}}))
	if labels < 100 {
		t.Fatalf("the labels drew %d pixels; six numerals should ink more", labels)
	}
	// Four cuts of the 0-100 domain fall on 20, 40, 60 and 80, which is
	// 100 wide in a 360 box at the reference radius, so the marks sit well
	// inside the reference radius rather than on it.
	on := lit(frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.Marks(MarksOptions{Count: 4, Length: 16, Width: 3, Offset: -80})
	}}))
	if on == 0 {
		t.Fatal("the marks ignored their radial offset")
	}
}

// TestTransitionAnimatesAndSettles moves the value with a transition and
// advances the clock, and requires the drawn frames to differ on the way
// and to agree once the motion settles, so the gauge animates and then
// stops rather than jittering forever.
func TestTransitionAnimatesAndSettles(t *testing.T) {
	value := 0.0
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Gauge(c, Options{Value: value, Size: 300, Transition: Transition{Enabled: true},
				Draw: func(g *Canvas) { g.Arc(ArcOptions{Width: 24}) }})
		})
	}, 332, 332)
	tt.Frame()
	first := append([]byte(nil), tt.Image().Pix...)
	value = 90
	// Advance the clock without forcing a frame at each step: the motion
	// itself is what asks for the next one, and the tester draws then.
	tt.Advance(50 * time.Millisecond)
	moving := append([]byte(nil), tt.Image().Pix...)
	if bytes.Equal(first, moving) {
		t.Fatal("the value jumped to its target instead of easing")
	}
	for range 40 {
		tt.Advance(50 * time.Millisecond)
	}
	settled := append([]byte(nil), tt.Image().Pix...)
	if !bytes.Equal(settled, frameValue(90).Image().Pix) {
		t.Fatal("the settled frame differs from the same value drawn without a transition")
	}
	tt.Advance(time.Second)
	if !bytes.Equal(settled, tt.Image().Pix) {
		t.Fatal("the motion never settles")
	}
}

// frameValue is the arc at v, drawn without a transition, for comparison
// with a settled animated one.
func frameValue(v float64) *ui.Tester {
	return frame(Options{Value: v, Size: 300, Draw: func(g *Canvas) {
		g.Arc(ArcOptions{Width: 24})
	}})
}

// TestEmptyOptionsRender draws the zero Options and every primitive with
// its zero value, so a nil map, a zero radius or an empty domain fails
// here rather than in a caller.
func TestEmptyOptionsRender(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("the zero options panicked: %v", r)
		}
	}()
	frame(Options{Size: 200})
	tt := frame(Options{Value: 0, Min: 0, Max: 0, Size: 200, Draw: func(g *Canvas) {
		g.Track(TrackOptions{})
		g.Arc(ArcOptions{})
		g.Stack(StackOptions{})
		g.Zones(ZonesOptions{})
		g.Ticks(TicksOptions{})
		g.TickFill(TickFillOptions{})
		g.Marks(MarksOptions{})
		g.TickLabels(TickLabelsOptions{})
		g.Needle(NeedleOptions{})
		g.Hub(HubOptions{})
		g.Dot(DotOptions{})
		g.Text(TextOptions{})
		g.Value(ValueOptions{})
		g.Inset(InsetOptions{})
	}})
	// A zero domain is not a division by zero: the gauge still answers a
	// frame, and the zero face still fits the box it was given. "Some
	// pixel is lit" would pass on an empty frame, since the two halves of
	// the disjunction each cover a different mistake.
	if lit(tt) == 0 {
		t.Fatal("the zero face drew nothing at all")
	}
	fill := float64(lit(tt)) / float64(200*200)
	if fill > 0.5 {
		t.Fatalf("the zero face covered %.0f%% of the box; a zero Sweep should not fill it", fill*100)
	}
}

// TestZeroCountsDrawNothing draws the primitives with a count that names
// no marks, and requires no panic: a negative count once reached
// make([]float64, 0, count) and crashed the whole window.
func TestZeroCountsDrawNothing(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("a count of -1 panicked: %v", r)
		}
	}()
	base := frame(Options{Value: 50, Draw: func(g *Canvas) {}})
	bare := lit(base)
	tt := frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.Ticks(TicksOptions{Count: -1})
		g.Marks(MarksOptions{Count: -1})
		g.Marks(MarksOptions{Count: 1})
		g.TickFill(TickFillOptions{Count: -1})
		g.TickLabels(TickLabelsOptions{Count: -1})
		g.Stack(StackOptions{Parts: nil})
	}})
	if lit(tt) != bare {
		t.Fatalf("a negative or absent count still drew ink: %d pixels over the bare %d", lit(tt), bare)
	}
}

// TestTurnedAngleHoldsAtTheEnds sweeps a needle four turns and requires
// it to stop at the end of the domain, as it does on a single turn, so a
// clock's hand does not carry on past 24:00. Without the clamp the extra
// turns keep going: 133.33° at the far end instead of 320.
func TestTurnedAngleHoldsAtTheEnds(t *testing.T) {
	end, over := 320.0, 0.0
	for _, v := range []float64{24, 25, 100} {
		got := valueToTurnedAngle(v, 0, 24, 40, end, 4)
		if v == 24 {
			over = got
		}
		if math.Abs(got-over) > 1e-9 {
			t.Fatalf("value %g turned to %g, past the domain's own end %g", v, got, over)
		}
	}
	start := valueToTurnedAngle(0, 0, 24, 40, end, 4)
	if math.Abs(start-40) > 1e-9 {
		t.Fatalf("the start of the domain turned to %g, not the start angle 40", start)
	}
}

// TestCompassLabelFoldsNegativeBearings requires a bearing below zero to
// read round the far side of the rose, so -10° is 350 and -45 is NW,
// rather than a negative degree or a reflected point.
func TestCompassLabelFoldsNegativeBearings(t *testing.T) {
	cases := []struct {
		degrees float64
		want    string
	}{
		{0, "N"}, {90, "E"}, {180, "S"}, {270, "W"}, {360, "N"},
		{-90, "W"}, {-10, "350°"}, {315, "NW"}, {-45, "NW"},
	}
	for _, c := range cases {
		if got := compassLabel(c.degrees); got != c.want {
			t.Fatalf("compassLabel(%g) = %q, want %q", c.degrees, got, c.want)
		}
	}
}

// TestDurationLabelKeepsOneSign requires the sign of a negative duration
// to stay on the hours, so a reading thirty minutes before midnight is
// "-0h 30m" and never the impossible "0h -30m".
func TestDurationLabelKeepsOneSign(t *testing.T) {
	cases := []struct {
		hours float64
		want  string
	}{
		{7.5, "7h 30m"}, {0, "0h 00m"}, {-0.5, "-0h 30m"},
		{-7.5, "-7h 30m"}, {23.999, "24h 00m"}, {1.99999, "2h 00m"},
		// A duration that rounds to zero minutes is zero, so it takes no
		// sign: "-0h 00m" would be a minus in front of nothing.
		{-0.001, "0h 00m"}, {-0.004, "0h 00m"},
	}
	for _, c := range cases {
		if got := durationLabel(c.hours); got != c.want {
			t.Fatalf("durationLabel(%g) = %q, want %q", c.hours, got, c.want)
		}
	}
}

// TestZoneColorFollowsTheDrawnCutoffs requires the color of a reading to
// agree with the band Zones draws it in: cutoffs out of order are clamped
// into the band they share, and a reading outside the domain takes the
// nearest zone rather than falling through a gap.
func TestZoneColorFollowsTheDrawnCutoffs(t *testing.T) {
	a, b, c := ui.RGB(1, 1, 1), ui.RGB(2, 2, 2), ui.RGB(3, 3, 3)
	ordered := []Zone{{To: 60, Color: a}, {To: 80, Color: b}, {Color: c}}
	for _, tc := range []struct {
		value float64
		want  ui.Color
	}{{0, a}, {59, a}, {60, b}, {79, b}, {80, c}, {100, c}, {-5, a}, {999, c}} {
		if got := ZoneColor(ordered, 0, 100, tc.value, a); got != tc.want {
			t.Fatalf("ZoneColor(ordered, %g) = %v, want %v", tc.value, got, tc.want)
		}
	}
	// The second cutoff falls below the first, so it is clamped up to it
	// and the middle zone has no width; the bands Zones draws agree, and
	// so must the color.
	crossed := []Zone{{To: 80, Color: a}, {To: 40, Color: b}, {Color: c}}
	for _, tc := range []struct {
		value float64
		want  ui.Color
	}{{10, a}, {79, a}, {80, c}, {100, c}} {
		if got := ZoneColor(crossed, 0, 100, tc.value, a); got != tc.want {
			t.Fatalf("ZoneColor(crossed, %g) = %v, want %v", tc.value, got, tc.want)
		}
	}
	if got := ZoneColor(nil, 0, 100, 50, b); got != b {
		t.Fatalf("ZoneColor(nil, 50) = %v, want the fallback %v", got, b)
	}
}

// TestMarksSeamOnAClosedRing draws the equal cuts of a closed ring and
// requires the seam to be cut as well as the inner cuts, and the seam
// value to be the end of the domain rather than a repeat of the start.
func TestMarksSeamOnAClosedRing(t *testing.T) {
	if got := cutValues(0, 100, 4); len(got) != 3 || got[0] != 25 || got[2] != 75 {
		t.Fatalf("cutValues(0, 100, 4) = %v, want the three inner cuts", got)
	}
	marks := append([]float64{100}, cutValues(0, 100, 4)...)
	if marks[0] != 100 || len(marks) != 4 {
		t.Fatalf("the closed ring's marks = %v, want the seam at the end plus three cuts", marks)
	}
	// The seam must not repeat a value the equal cuts already carry, or
	// the closed ring draws one mark twice and the seam stays uncut.
	for _, v := range marks[1:] {
		if v == marks[0] {
			t.Fatalf("the seam %g is drawn twice", v)
		}
	}
}

// TestNeedleTailNegativeDrawsNoTail draws the same needle with the
// default tail and with a negative one, and requires the negative to
// stop at the center, so a needle that only reaches forward has a way to
// say so — zero is the default, not "none".
func TestNeedleTailNegativeDrawsNoTail(t *testing.T) {
	draw := func(tail float32) *ui.Tester {
		return frame(Options{Value: 50, Draw: func(g *Canvas) {
			g.Needle(NeedleOptions{Style: NeedleLine, Width: 8, Length: 0.9, Tail: tail})
		}})
	}
	defaulted := lit(draw(0))
	tailed := lit(draw(40))
	tailless := lit(draw(-1))
	if tailless >= defaulted {
		t.Fatalf("a negative tail lit %d pixels, no fewer than the default's %d", tailless, defaulted)
	}
	if tailed <= defaulted {
		t.Fatalf("a tail of 40 lit %d pixels, no more than the default's %d", tailed, defaulted)
	}
	if got := lit(frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.Needle(NeedleOptions{Style: NeedleLine, Width: 8, Length: 0.9, Tail: -1, TailDot: 6})
	}})); got != tailless {
		t.Fatalf("a tail dot with no tail drew %d pixels, want the tailless %d", got, tailless)
	}
}

// TestHubAndDotNegativeRadiusDrawNothing draws a negative radius for the
// hub and the dot, and requires neither to ink the canvas, so a caller
// can switch a pivot or a body off from configuration.
func TestHubAndDotNegativeRadiusDrawNothing(t *testing.T) {
	bare := lit(frame(Options{Value: 50, Draw: func(g *Canvas) {}}))
	got := lit(frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.Hub(HubOptions{Radius: -1})
		g.Dot(DotOptions{Radius: -1})
	}}))
	if got != bare {
		t.Fatalf("a negative radius drew %d pixels over the bare %d", got, bare)
	}
}

// TestInsetZeroScaleStillDraws draws an inset with a zero scale and
// requires it to render: the zero value would collapse every width and
// radius of the inset's face, so it takes the documented third instead.
func TestInsetZeroScaleStillDraws(t *testing.T) {
	tt := frame(Options{Value: 50, Draw: func(g *Canvas) {
		g.Inset(InsetOptions{Value: 50, Draw: func(in *Canvas) { in.Arc(ArcOptions{Width: 20}) }})
	}})
	if lit(tt) == 0 {
		t.Fatal("an inset with a zero scale drew nothing")
	}
}
