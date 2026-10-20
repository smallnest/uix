package view

import (
	"fmt"
	"math"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/gauge"
)

// The cockpit page is the six-pack of a light aircraft: the airspeed, the
// attitude, the altitude, the turn coordinator, the heading and the
// vertical speed. The readings are the ones gauge-ui's cockpit ships, at
// cruise.

// The cockpit readings, as upstream seeds them at cruise.
const (
	ckAirspeed = 112
	ckAltitude = 5500
	ckHeading  = 270
	ckVSpeed   = 0
	ckBank     = 12
)

// demoGaugeCockpit is the cockpit page: the turning card, the marks that
// make the case marks, and a pair of hubs on one needle.
func demoGaugeCockpitPage(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		gaugeCaption(c, t, "六件套")
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			ckAirspeedDial(c, t)
			ckAttitudeDial(c, t)
			ckAltimeter(c, t)
		})
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			ckTurn(c, t)
			ckHeadingDial(c, t)
			ckVSpeedDial(c, t)
		})
	})
}

// ckScale is the two-tier scale the cockpit dials share: fine minor ticks
// under heavier majors, inside a ring of numerals.
func ckScale(g *gauge.Canvas, minor, major int, labelSize float32, format func(float64) string) {
	g.Ticks(gauge.TicksOptions{Count: minor, Length: 10, Width: 1.5, Offset: -5, Opacity: 0.5})
	g.Ticks(gauge.TicksOptions{Count: major, Length: 20, Width: 3, Offset: -10})
	g.TickLabels(gauge.TickLabelsOptions{Count: major, Offset: -46, FontSize: labelSize, Format: format})
}

// ckPointer is the cockpit's needle: a slim one on a pivot capped in the
// face color, which is why there are two hubs.
func ckPointer(g *gauge.Canvas, t *ui.Theme, gap float32) {
	g.Needle(gauge.NeedleOptions{Style: gauge.NeedleLine, Length: 0.88, Width: 5, Tail: 28, Gap: gap, Color: t.Text})
	g.Hub(gauge.HubOptions{Radius: 14, Color: t.Text})
	g.Hub(gauge.HubOptions{Radius: 5, Color: t.Background})
}

// ckAirspeedDial is the airspeed: the white arc of the flap range, the
// green and yellow arcs of the normal ones, and the never-exceed line.
func ckAirspeedDial(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       ckAirspeed,
		Min:         40,
		Max:         200,
		StartAngle:  200,
		EndAngle:    520,
		Size:        190,
		Label:       "空速",
		Description: fmt.Sprintf("空速 %d 节。", ckAirspeed),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Arc(gauge.ArcOptions{Width: 4, Offset: -15, Opacity: 0.5, From: ptr(40.0), To: ptr(85.0)})
			g.Arc(gauge.ArcOptions{Width: 10, Offset: -5, Opacity: 0.15, From: ptr(50.0), To: ptr(130.0)})
			g.Arc(gauge.ArcOptions{Width: 10, Offset: -5, Opacity: 0.35, From: ptr(130.0), To: ptr(165.0)})
			g.Marks(gauge.MarksOptions{Values: []float64{165}, Length: 24, Width: 4, Offset: -12,
				Color: gauge.ValueColor{Color: t.Danger}})
			ckScale(g, 32, 8, 15, func(v float64) string { return fmt.Sprintf("%.0f", v) })
			ckPointer(g, t, 0)
			g.Text(gauge.TextOptions{Y: -44, FontSize: 14, Color: t.TextMuted, Text: "KNOTS"})
		},
	})
}

// ckAttitudeDial is the attitude indicator: the sky and the ground turned
// by the bank, with the pitch ladder over them.
//
// It rotates the whole dial by the bank and reads the pitch on the value,
// so the horizon tilts under a fixed aircraft rather than the other way.
func ckAttitudeDial(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       -ckBank,
		Min:         -90,
		Max:         90,
		StartAngle:  90,
		EndAngle:    270,
		Rotate:      ckBank,
		Size:        190,
		Label:       "姿态",
		Description: fmt.Sprintf("坡度 %d 度，俯仰 2 度。", ckBank),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			// The horizon is the band across the top of the half dial:
			// the sky above it in the accent and the ground below in the
			// warning tone, with the whole dial turned by the bank, so
			// the horizon stays level with the world while the aircraft
			// symbol does not.
			//
			// The two arcs are the sky and the ground halves of the
			// pitch scale, so From and To are domain values, not angles.
			// An instrument face is a filled field rather than a ring,
			// which the band primitives do not draw; the horizon is
			// therefore the thick arc a real one is read from, not the
			// painted disc an artificial horizon has.
			g.Arc(gauge.ArcOptions{Width: 40, Offset: -46, Color: t.Accent.Alpha(0.45),
				From: ptr(-90.0), To: ptr(0.0)})
			g.Arc(gauge.ArcOptions{Width: 40, Offset: -46, Color: t.Warning.Alpha(0.45),
				From: ptr(0.0), To: ptr(90.0)})
			// The pitch ladder: the fifteen- and thirty-degree lines
			// either side of the horizon, longer as they leave it. The
			// marks of one call share a length, so the longer thirty-
			// degree lines take their own call and the horizon its own,
			// heaviest of all. They all turn with the dial, as the lines
			// of a real one do.
			g.Marks(gauge.MarksOptions{Values: []float64{-30, 30}, Length: 20, Width: 3, Offset: -8,
				Color: gauge.ValueColor{Color: t.Text}})
			g.Marks(gauge.MarksOptions{Values: []float64{-15, 15}, Length: 12, Width: 3, Offset: -8,
				Color: gauge.ValueColor{Color: t.Text}})
			g.Marks(gauge.MarksOptions{Values: []float64{0}, Length: 30, Width: 5, Offset: -14,
				Color: gauge.ValueColor{Color: t.Text}})
			// The aircraft symbol is fixed to the instrument rather than
			// to the world, so an inset draws it on the unrotated canvas:
			// its own narrow sweep keeps it level while the ladder turns
			// under it, which is what the wings of a real one do.
			//
			// A mark is drawn across its own radius, so the sweep of the
			// inset runs through three o'clock to lay the wings out
			// horizontally, as the aircraft they stand for sits.
			g.Inset(gauge.InsetOptions{
				Value: 0, Min: -1, Max: 1,
				StartAngle: 90, EndAngle: 90,
				// The symbol is drawn at the dial's own scale rather than
				// a fraction of it, so its lengths are design units of the
				// host and a length of 150 reaches most of the way across.
				Scale: 1,
				Draw: func(in *gauge.Canvas) {
					// The wings straddle the pivot, so the line is pulled
					// back by the whole reference radius and half its
					// length falls each side of the centre.
					in.Marks(gauge.MarksOptions{Values: []float64{0}, Length: 150, Width: 7,
						Offset: -160, Color: gauge.ValueColor{Color: t.Text}})
					// The dot is the aircraft itself: a disc on the pivot,
					// so it takes the same pull back and no offset of its
					// own beyond that.
					in.Dot(gauge.DotOptions{Radius: 8, Offset: -160, Color: t.Text})
				}})
			// The hub is the aircraft, not the horizon: it stays on the
			// pivot while the field turns around it.
			g.Hub(gauge.HubOptions{Radius: 5, Color: t.TextMuted})
			g.Text(gauge.TextOptions{X: -62, Y: -16, FontSize: 13, Color: t.TextMuted, Text: "SKY"})
			g.Text(gauge.TextOptions{X: 62, Y: -16, FontSize: 13, Color: t.TextMuted, Text: "GND"})
		},
	})
}

// ckAltimeter is the altitude: the hundreds on a short blunt hand, the
// thousands on a long fine one turning ten times over, and the barometer
// setting beside the dial.
func ckAltimeter(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       ckAltitude,
		Min:         0,
		Max:         10000,
		StartAngle:  180,
		EndAngle:    540,
		Size:        190,
		Label:       "高度",
		Description: fmt.Sprintf("高度 %d 英尺。", ckAltitude),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			ckScale(g, 50, 10, 15, func(v float64) string { return fmt.Sprintf("%.0f", v/1000) })
			// Two hands on one reading: the hundreds turn once round the
			// dial and the thousands ten times, which is what a real
			// altimeter's hands do.
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Length: 0.55, Width: 16, Tail: -1, Color: t.Text})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedleLine, Length: 0.88, Width: 5, Tail: 28, Turns: 10, Color: t.Text})
			g.Hub(gauge.HubOptions{Radius: 14, Color: t.Text})
			g.Hub(gauge.HubOptions{Radius: 5, Color: t.Background})
			g.Text(gauge.TextOptions{Y: -48, FontSize: 14, Color: t.TextMuted, Text: "ALT"})
			g.Text(gauge.TextOptions{X: 62, Y: -30, FontSize: 14, Font: gauge.FontMono, Color: t.TextMuted, Text: "29.92"})
		},
	})
}

// ckTurn is the turn coordinator: the rate of turn on the wings and the
// slip on the ball below, the two dials a coordinated turn is read from.
func ckTurn(c *ui.Context, t *ui.Theme) ui.Element {
	const rate = 0.6
	return gauge.Gauge(c, gauge.Options{
		Value:       rate,
		Min:         -1.5,
		Max:         1.5,
		StartAngle:  60,
		EndAngle:    300,
		Size:        190,
		Label:       "转弯侧滑",
		Description: "转弯侧滑仪：转弯率 0.6，侧滑球居中。",
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Arc(gauge.ArcOptions{Width: 10, Offset: -5, Opacity: 0.15, From: ptr(0.0), To: ptr(rate)})
			g.Ticks(gauge.TicksOptions{Count: 6, Length: 10, Width: 1.5, Offset: -5, Opacity: 0.5})
			g.Marks(gauge.MarksOptions{Values: []float64{0, rate}, Length: 22, Width: 4, Offset: -12,
				Color: gauge.ValueColor{Color: t.Text}})
			g.Text(gauge.TextOptions{X: -44, Y: -32, FontSize: 16, Text: "L"})
			g.Text(gauge.TextOptions{X: 44, Y: -32, FontSize: 16, Text: "R"})
			g.Text(gauge.TextOptions{Y: -66, FontSize: 13, Color: t.TextMuted, Text: "TURN"})
			// The slip ball is a track across the bottom with a disc on
			// it, centred when the turn is coordinated.
			g.Inset(gauge.InsetOptions{
				Value: 0, Min: -1, Max: 1,
				Radius:     120,
				StartAngle: 338, EndAngle: 382,
				Draw: func(in *gauge.Canvas) {
					in.Track(gauge.TrackOptions{Width: 26, Opacity: 0.12})
					in.Track(gauge.TrackOptions{Width: 1.5, Offset: -13, Opacity: 0.4})
					in.Track(gauge.TrackOptions{Width: 1.5, Offset: 13, Opacity: 0.4})
					in.Marks(gauge.MarksOptions{Values: []float64{-0.3, 0.3}, Length: 26, Width: 2, Opacity: 0.6})
					in.Dot(gauge.DotOptions{Radius: 10, Color: t.Text})
				}})
		},
	})
}

// ckHeadingDial is the heading: the compass card turning under a fixed
// lubber line, so the dial reads like the instrument it copies.
func ckHeadingDial(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       0,
		Min:         0,
		Max:         360,
		StartAngle:  180,
		EndAngle:    540,
		Rotate:      -Demo.GaugeHeading,
		Padding:     32,
		Size:        190,
		Label:       "航向",
		Description: fmt.Sprintf("航向 %s。", gauge.CompassLabel(Demo.GaugeHeading)),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Marks(gauge.MarksOptions{Values: []float64{gauge.Fold(Demo.GaugeHeading, 0, 360)},
				Length: 12, Width: 12, Offset: -6, Color: gauge.ValueColor{Color: t.Accent}})
			g.Ticks(gauge.TicksOptions{Count: 72, Length: 6, Width: 2, Offset: -3, Opacity: 0.5})
			g.Ticks(gauge.TicksOptions{Count: 36, Length: 10, Width: 3, Offset: -5})
			// A compass card names the cardinals and numbers every
			// thirty degrees between them, as "3" for 30 and "30" for
			// 300: the last digit of the heading.
			g.TickLabels(gauge.TickLabelsOptions{Count: 12, Offset: -30, FontSize: 14,
				Format: func(v float64) string {
					switch gauge.Fold(v, 0, 360) {
					case 0:
						return "N"
					case 90:
						return "E"
					case 180:
						return "S"
					case 270:
						return "W"
					}
					return fmt.Sprintf("%.0f", gauge.Fold(v, 0, 360)/10)
				}})
			// The lubber line is fixed to the aircraft, so it is drawn
			// by an inset that does not turn with the card.
			g.Inset(gauge.InsetOptions{
				Value: 0, Min: 0, Max: 360,
				StartAngle: 180, EndAngle: 540,
				Draw: func(in *gauge.Canvas) {
					in.Marks(gauge.MarksOptions{Values: []float64{0}, Length: 12, Width: 5, Offset: 14,
						Color: gauge.ValueColor{Color: t.Text}})
				}})
			g.Text(gauge.TextOptions{Y: 2, FontSize: 24, Font: gauge.FontMono, Weight: 600,
				Text: fmt.Sprintf("%03.0f", Demo.GaugeHeading)})
			g.Text(gauge.TextOptions{Y: -28, FontSize: 13, Color: t.TextMuted, Text: "HDG"})
		},
	})
}

// ckVSpeedDial is the vertical speed: the climb above the zero mark and
// the descent below it, read in thousands of feet a minute.
func ckVSpeedDial(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       ckVSpeed / 1000,
		Min:         -2,
		Max:         2,
		StartAngle:  -90,
		EndAngle:    270,
		Size:        190,
		Label:       "升降速度",
		Description: "升降速度零，平飞。",
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Ticks(gauge.TicksOptions{Count: 40, Length: 10, Width: 1.5, Offset: -5, Opacity: 0.5})
			g.Ticks(gauge.TicksOptions{Count: 8, Length: 20, Width: 3, Offset: -10})
			g.TickLabels(gauge.TickLabelsOptions{Count: 8, Offset: -34, FontSize: 14,
				Format: func(v float64) string {
					if v == 0 {
						return ""
					}
					return fmt.Sprintf("%.0f", math.Abs(v)*10)
				}})
			g.Marks(gauge.MarksOptions{Values: []float64{0}, Length: 34, Width: 5, Offset: -17,
				Color: gauge.ValueColor{Color: t.Text}})
			g.Text(gauge.TextOptions{X: -58, Y: -22, FontSize: 14, Color: t.TextMuted, Text: "UP"})
			g.Text(gauge.TextOptions{X: -58, Y: 22, FontSize: 14, Color: t.TextMuted, Text: "DN"})
			ckPointer(g, t, 0)
			g.Text(gauge.TextOptions{Y: 58, FontSize: 14, Color: t.TextMuted, Text: "FT/MIN"})
		},
	})
}
