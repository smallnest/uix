package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	sl "github.com/smallnest/uix/components/slider"
)

// The gauge pages show the composable gauge of the gauge-ui port: the
// primitives compose in a draw callback rather than nesting as elements,
// and there is a page for each of the shapes they build.
//
// The set follows gauge-ui's own gallery, which shows whole instrument
// clusters rather than isolated dials, because a gauge is only as good as
// the cluster it sits in. Each page here is one of those clusters, at the
// readings it ships with.

// demoGauge is the dashboard: an animated percentage over its zones, a
// tick meter with its labels, and a speedometer with a fuel dial nested
// inside it. The slider under the row drives all three, so the motion,
// the geometry and the layers can be seen at once.
func demoGauge(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(3)).Children(func() {
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(4)).AlignItems(ui.Center).Children(func() {
			gaugeZones(c, t)
			gaugeTicks(c, t)
			gaugeSpeed(c, t)
		})
		ui.Box(c).MaxWidth(360).Children(func() {
			slSlider(c, t, &Demo.GaugeValue, 0, 100, "读数", "%.0f")
		})
	})
}

// slSlider is the slider a gauge page drives its reading with.
func slSlider(c *ui.Context, t *ui.Theme, value *float64, min, max float64, label, format string) {
	sl.Slider(c, sl.Props{
		Value: value,
		Min:   min,
		Max:   max,
		Label: label,
		Format: func(v float64) string {
			return fmt.Sprintf(format, v)
		},
	})
}
