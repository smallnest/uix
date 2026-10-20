package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/gauge"
)

// gaugeZonesGauge is the animated percentage: a track, the color zones
// of a load's levels, ticks, and an arc and needle over them, all
// easing to the value the slider sets.
func gaugeZones(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       Demo.GaugeValue,
		Size:        230,
		Label:       "负载",
		Description: "当前负载占满量的百分比。",
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 16})
			zones := []gauge.Zone{
				{To: 60, Color: t.Success},
				{To: 80, Color: t.Warning},
				{Color: t.Danger},
			}
			// The needle and the readout take the color of the zone the
			// reading falls in, so the dial says "still fine" or "too hot"
			// without anyone reading the number.
			tone := gauge.ZoneColor(zones, 0, 100, Demo.GaugeValue, t.Text)
			g.Zones(gauge.ZonesOptions{Width: 16, Gap: 3, Zones: zones})
			g.Ticks(gauge.TicksOptions{Count: 10, Length: 8, Width: 2, Offset: -30,
				Color: gauge.ValueColor{Color: t.TextMuted}})
			g.Arc(gauge.ArcOptions{Width: 5, Offset: 42, Color: tone})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Width: 7, Length: 0.92, Tail: 26, TailDot: 4, Color: tone})
			g.Hub(gauge.HubOptions{Radius: 9, Color: t.Text})
			g.Value(gauge.ValueOptions{Y: 44, FontSize: 34, Color: tone,
				Format: func(v float64) string { return fmt.Sprintf("%.0f%%", v) }})
		},
	})
}

// gaugeTicksGauge is the meter drawn from ticks: the whole sweep is
// ticked, those below the value take the accent while the rest stay
// faint, and the labels stand inside them on the radius.
func gaugeTicks(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       Demo.GaugeValue,
		Size:        230,
		Label:       "刻度表",
		Description: "以刻度表示的百分比读数。",
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.TickFill(gauge.TickFillOptions{Count: 44, Length: 22, Width: 3,
				Color:    gauge.ValueColor{Color: t.Accent},
				DimColor: gauge.ValueColor{Color: t.Text.Alpha(0.15)}})
			g.TickLabels(gauge.TickLabelsOptions{Count: 10, Offset: -54, FontSize: 13,
				Color: gauge.ValueColor{Color: t.TextMuted}, Format: func(v float64) string { return fmt.Sprintf("%.0f", v) }})
			g.Value(gauge.ValueOptions{Y: 26, FontSize: 32, Color: t.Text,
				Format: func(v float64) string { return fmt.Sprintf("%.0f%%", v) }})
		},
	})
}

// gaugeSpeedGauge is the speedometer: a wide track with a dot riding it
// and a needle inside, and a nested fuel dial at a third of the size, so
// the inset's own domain and scale show.
func gaugeSpeed(c *ui.Context, t *ui.Theme) ui.Element {
	const top = 240
	speed := Demo.GaugeValue / 100 * top
	return gauge.Gauge(c, gauge.Options{
		Value:       speed,
		Min:         0,
		Max:         top,
		Size:        230,
		Label:       "车速",
		Description: "车速，单位为公里每小时，内圈为油量。",
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 22, Opacity: 0.12})
			g.Arc(gauge.ArcOptions{Width: 22, Color: t.Accent})
			g.Ticks(gauge.TicksOptions{Count: 24, Length: 7, Width: 3, Offset: -34,
				Color: gauge.ValueColor{Color: t.TextMuted}})
			g.Dot(gauge.DotOptions{Radius: 11, Color: t.Accent, Halo: 4, HaloColor: t.Background})
			g.Inset(gauge.InsetOptions{X: 0, Y: -8, Scale: 0.40, Value: Demo.GaugeFuel, Min: 0, Max: 100,
				Draw: func(i *gauge.Canvas) {
					i.Track(gauge.TrackOptions{Width: 12, Opacity: 0.22})
					i.Arc(gauge.ArcOptions{Width: 12, Color: t.Success})
				}})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedleLine, Width: 4, Length: 0.58, Tail: -1, Color: t.Text})
			g.Hub(gauge.HubOptions{Radius: 7, Color: t.Text})
			g.Value(gauge.ValueOptions{Y: 62, FontSize: 24, Color: t.Text,
				Format: func(v float64) string { return fmt.Sprintf("%.0f km/h", v) }})
		},
	})
}
