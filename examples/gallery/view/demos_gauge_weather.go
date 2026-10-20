package view

import (
	"fmt"
	"math"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/gauge"
)

// The weather page is a Reykjavík afternoon: the temperature over the
// hours ahead, a wind vane, the air quality on its scale, and an
// aneroid barometer with its set hand. The readings are the ones
// gauge-ui's weather dashboard ships.

// The weather readings, as upstream seeds them.
const (
	wtTemp     = 8.4
	wtLow      = 4.8
	wtHigh     = 10.3
	wtPressure = 1003.6
	wtSet      = 1008.9
	wtAQI      = 18
)

// demoGaugeWeather is the weather page: the two needle styles the other
// pages do not use, the compass rose, and the mark primitives.
func demoGaugeWeatherPage(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			wtTemperature(c, t)
			wtWind(c, t)
			wtCompass(c, t)
		})
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			wtBarometer(c, t)
			wtAirQuality(c, t)
		})
		ui.Box(c).MaxWidth(360).Children(func() {
			slSlider(c, t, &Demo.GaugeWind, 0, 30, "风速", "%.1f m/s")
		})
		ui.Box(c).MaxWidth(360).Children(func() {
			slSlider(c, t, &Demo.GaugeHeading, 0, 360, "风向", "%.0f°")
		})
	})
}

// wtTemperature is the temperature over its range: the day's low and
// high marked on the band, the bands by warmth behind them.
func wtTemperature(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       wtTemp,
		Min:         -20,
		Max:         30,
		Size:        200,
		Label:       "气温",
		Description: fmt.Sprintf("气温 %.1f 摄氏度，今日 %.1f 到 %.1f 度。", wtTemp, wtLow, wtHigh),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 16, Opacity: 0.14})
			g.Arc(gauge.ArcOptions{Width: 16, Color: t.Accent})
			g.Zones(gauge.ZonesOptions{Width: 6, Offset: -20, Gap: 1, Zones: []gauge.Zone{
				{To: 0, Color: t.Accent.Alpha(0.22)},
				{To: 15, Color: t.Accent.Alpha(0.5)},
				{Color: t.Accent},
			}})
			// The day's own low and high, marked where they fall rather
			// than at an equal cut.
			g.Marks(gauge.MarksOptions{Values: []float64{wtLow, wtHigh}, Length: 14, Width: 3, Offset: -20,
				Color: gauge.ValueColor{Color: t.Text}})
			g.Ticks(gauge.TicksOptions{Count: 50, Length: 6, Width: 1.5, Offset: -34, Opacity: 0.5})
			g.Value(gauge.ValueOptions{Y: 6, FontSize: 44, Font: gauge.FontRounded, Weight: 700,
				Format: func(v float64) string { return fmt.Sprintf("%.0f°", v) }})
			g.Text(gauge.TextOptions{Y: -30, FontSize: 15, Color: t.TextMuted, Text: "°C"})
			g.Text(gauge.TextOptions{Y: 34, FontSize: 14, Color: t.TextMuted,
				Text: fmt.Sprintf("%.0f° · %.0f°", wtLow, wtHigh)})
		},
	})
}

// wtWind is the vane: the arrow pointing where the wind comes from, its
// counterweight behind the pivot, and the speed read out between the two
// halves of the needle.
//
// This is the one needle the gap exists for: the readout sits in the
// circle the needle is cut around, where a whole needle would run under
// the text.
func wtWind(c *ui.Context, t *ui.Theme) ui.Element {
	from := Demo.GaugeHeading
	speed := Demo.GaugeWind
	return gauge.Gauge(c, gauge.Options{
		Value:       from,
		Min:         0,
		Max:         360,
		StartAngle:  180,
		EndAngle:    540,
		Size:        200,
		Label:       "风向风速",
		Description: fmt.Sprintf("风来自 %s，风速 %.1f 米每秒。", gauge.CompassLabel(from), speed),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 24, Opacity: 0.12})
			g.Ticks(gauge.TicksOptions{Count: 4, Length: 8, Width: 3, Opacity: 0.7})
			g.Ticks(gauge.TicksOptions{Count: 36, Length: 4, Width: 3, Opacity: 0.5})
			g.TickLabels(gauge.TickLabelsOptions{Count: 4, Offset: -34, FontSize: 18, Weight: 700,
				Color: gauge.ValueColor{Color: t.TextMuted}, Format: gauge.CompassLabel})
			g.Needle(gauge.NeedleOptions{
				Style: gauge.NeedleArrow, Length: 1.05, Width: 7,
				Tail: 150, TailDot: 9, Gap: 60,
				Color: t.Text, TailColor: t.TextMuted,
			})
			g.Value(gauge.ValueOptions{Y: -6, FontSize: 36, Font: gauge.FontRounded, Weight: 700,
				Format: func(float64) string { return fmt.Sprintf("%.0f", speed) }})
			g.Text(gauge.TextOptions{Y: 28, FontSize: 14, Color: t.TextMuted, Text: "m/s"})
		},
	})
}

// wtCompass is the rose: the wind on a compass card, the eight points
// named, the needle a lozenge of two colors as a magnetic one is.
//
// No gauge-ui scene draws the compass needle, so this page is where it
// is shown: the two halves take the two poles' colors.
func wtCompass(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       Demo.GaugeHeading,
		Min:         0,
		Max:         360,
		StartAngle:  180,
		EndAngle:    540,
		Size:        200,
		Label:       "罗盘",
		Description: fmt.Sprintf("罗盘指向 %s。", gauge.CompassLabel(Demo.GaugeHeading)),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 20, Opacity: 0.12})
			g.Ticks(gauge.TicksOptions{Count: 72, Length: 5, Width: 2, Offset: -14, Opacity: 0.5})
			g.Ticks(gauge.TicksOptions{Count: 8, Length: 12, Width: 3, Offset: -18})
			g.TickLabels(gauge.TickLabelsOptions{Count: 8, Offset: -38, FontSize: 15,
				Format: gauge.CompassLabel})
			// A magnetic needle is red toward the north and dark toward
			// the south, which is what the two halves are for.
			g.Needle(gauge.NeedleOptions{
				Style: gauge.NeedleCompass, Length: 0.8, Width: 14, Tail: 66,
				Color: t.Danger, TailColor: t.Text,
			})
			g.Hub(gauge.HubOptions{Radius: 8, Color: t.Text})
			g.Value(gauge.ValueOptions{Y: 54, FontSize: 22, Font: gauge.FontMono, Weight: 600,
				Format: func(v float64) string { return fmt.Sprintf("%.0f°", v) }})
		},
	})
}

// wtBarometer is the aneroid: the reading against the pressure it was
// set to, with the set hand left where it was and the fall drawn between
// the two.
func wtBarometer(c *ui.Context, t *ui.Theme) ui.Element {
	words := []string{"风暴", "雨", "变化", "晴", "很晴"}
	return gauge.Gauge(c, gauge.Options{
		Value:       wtPressure,
		Min:         960,
		Max:         1040,
		StartAngle:  45,
		EndAngle:    315,
		Size:        200,
		Label:       "气压",
		Description: fmt.Sprintf("气压 %.1f 百帕，设定 %.1f。", wtPressure, wtSet),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 14, Opacity: 0.1})
			// The fall since the reading was set: an arc anchored at the
			// set hand rather than at the start of the sweep.
			g.Arc(gauge.ArcOptions{Width: 14, From: ptr(wtSet), Opacity: 0.35, Color: t.Accent})
			g.Ticks(gauge.TicksOptions{Count: 80, Length: 6, Width: 1.5, Offset: -18, Opacity: 0.5})
			g.Ticks(gauge.TicksOptions{Count: 8, Length: 14, Width: 3, Offset: -22})
			g.TickLabels(gauge.TickLabelsOptions{Count: 4, Offset: -50, FontSize: 14,
				Format: func(v float64) string {
					i := int(math.Round((v - 960) / 20))
					if i < 0 || i >= len(words) {
						return ""
					}
					return words[i]
				}})
			// The set hand stays where it was put, so the two readings
			// can be compared at a glance.
			g.Marks(gauge.MarksOptions{Values: []float64{wtSet}, Length: 30, Width: 5,
				Color: gauge.ValueColor{Color: t.Text}})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Length: 0.56, Width: 18, Tail: 26, Color: t.Text})
			g.Hub(gauge.HubOptions{Radius: 13, Color: t.Text})
			g.Hub(gauge.HubOptions{Radius: 4, Color: t.Background})
			g.Value(gauge.ValueOptions{Y: 44, FontSize: 20, Font: gauge.FontMono, Decimals: 1})
		},
	})
}

// wtAirQuality is the air: the index against its bands, the healthy
// range called out.
func wtAirQuality(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       wtAQI,
		Max:         500,
		StartAngle:  90,
		EndAngle:    270,
		Padding:     40,
		Size:        200,
		Label:       "空气质量",
		Description: fmt.Sprintf("空气质量指数 %d，优良。", wtAQI),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 20, Opacity: 0.14})
			g.Arc(gauge.ArcOptions{Width: 20, Color: t.Success})
			g.Zones(gauge.ZonesOptions{Width: 6, Offset: -22, Gap: 2, EndCap: ptr(gauge.CapRound),
				Zones: []gauge.Zone{
					{To: 50, Color: t.Success.Alpha(0.35)},
					{To: 100, Color: t.Warning.Alpha(0.5)},
					{Color: t.Danger.Alpha(0.6)},
				}})
			g.Marks(gauge.MarksOptions{Values: []float64{50, 100}, Length: 10, Width: 2, Offset: -22})
			g.Value(gauge.ValueOptions{Y: -30, FontSize: 36, Font: gauge.FontRounded, Weight: 700})
			g.Text(gauge.TextOptions{Y: 8, FontSize: 15, Color: t.TextMuted, Text: "AQI"})
		},
	})
}
