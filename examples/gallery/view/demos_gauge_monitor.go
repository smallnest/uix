package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/gauge"
)

// The monitor page is a server's load: a CPU dial split into user and
// kernel, a ring per core, memory and swap on one dial, and the network
// on a twin half dial. The readings are the ones gauge-ui's own monitor
// dashboard ships, so the shapes are the shapes they were drawn for.

// The monitor's readings, seeded as upstream's are.
const (
	monCores   = 8
	monRAM     = 32
	monSwap    = 8
	monLink    = 1000
	monKernel  = 0.28
	monWired   = 3.4
	monApp     = 12.6
	monCached  = 6.2
	monSwapped = 1.1
)

// monCoreLoad is the load of each core, as upstream seeds them.
var monCoreLoad = [monCores]float64{46, 38, 61, 29, 52, 34, 44, 27}

// demoGaugeMonitor is the monitor page: the layered primitives at their
// clearest, where each dial needs two or three of them over one another.
func demoGaugeMonitorPage(c *ui.Context, t *ui.Theme) {
	busy := 0.0
	for _, v := range monCoreLoad {
		busy += v
	}
	busy /= monCores
	used := monWired + monApp + monCached
	ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			monCPU(c, t, busy)
			monMemory(c, t, used)
			monNetwork(c, t)
		})
		gaugeCaption(c, t, "每核占用")
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
			for i := range monCoreLoad {
				monCore(c, t, i)
			}
		})
		gaugeCaption(c, t, "卷占用")
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			monDisk(c, t, "/", 64)
			monDisk(c, t, "/var/log", 91)
			monDisk(c, t, "/data", 47)
		})
	})
}

// gaugeCaption is the small muted line a page puts above a row of dials.
func gaugeCaption(c *ui.Context, t *ui.Theme, text string) {
	ui.Text(c, text).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
}

// monCPU is the whole processor: the busy share drawn as a stack of the
// user half and the kernel half, over a fine scale.
//
// The stack is the weighted kind, which is what a partition of one
// reading needs: the two parts keep their share of the fill frame by
// frame rather than sitting at fixed cutoffs.
func monCPU(c *ui.Context, t *ui.Theme, busy float64) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       busy,
		Max:         100,
		Size:        200,
		Label:       "CPU 占用",
		Description: fmt.Sprintf("处理器占用 %.0f%%，其中内核 %.0f%%。", busy, monKernel*busy),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 26, Opacity: 0.1})
			g.Stack(gauge.StackOptions{Width: 26, Parts: []gauge.StackPart{
				{Weight: 1 - monKernel, Color: t.Text},
				{Weight: monKernel, Color: t.TextMuted, Opacity: 0.6},
			}})
			g.Ticks(gauge.TicksOptions{Count: 50, Length: 6, Width: 2, Offset: -30, Opacity: 0.4})
			g.Ticks(gauge.TicksOptions{Count: 10, Length: 14, Width: 4, Offset: -34})
			g.TickLabels(gauge.TickLabelsOptions{Count: 5, Offset: -70, FontSize: 15})
			g.Value(gauge.ValueOptions{Y: -10, FontSize: 40, Font: gauge.FontMono})
			g.Text(gauge.TextOptions{Y: 30, FontSize: 15, Color: t.TextMuted, Text: "% busy"})
		},
	})
}

// monCore is one core's ring: a full circle filled to its load, with the
// last tenth called out by a danger zone.
func monCore(c *ui.Context, t *ui.Theme, index int) ui.Element {
	load := monCoreLoad[index]
	return gauge.Gauge(c, gauge.Options{
		Value:       load,
		Max:         100,
		StartAngle:  180,
		EndAngle:    540,
		Size:        108,
		Label:       fmt.Sprintf("核心 %d", index),
		Description: fmt.Sprintf("第 %d 个核心占用 %.0f%%。", index, load),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 16, Opacity: 0.1})
			g.Zones(gauge.ZonesOptions{Width: 5, Offset: -20, Zones: []gauge.Zone{
				{To: 90, Color: t.Success},
				{Color: t.Danger},
			}})
			g.Arc(gauge.ArcOptions{Width: 16, Color: t.Accent})
			g.Value(gauge.ValueOptions{Y: -4, FontSize: 26, Font: gauge.FontMono})
		},
	})
}

// monMemory is the memory in its three parts on one dial, with swap on a
// ring inside it. The inset takes its own domain and its own sweep, so
// the inner ring is a different scale of the same drawing.
func monMemory(c *ui.Context, t *ui.Theme, used float64) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       used,
		Max:         monRAM,
		Size:        200,
		Label:       "内存",
		Description: fmt.Sprintf("内存已用 %.1f GB，共 %d GB，交换 %.1f GB。", used, monRAM, monSwapped),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 26, Opacity: 0.1})
			g.Stack(gauge.StackOptions{Width: 26, Parts: []gauge.StackPart{
				{Weight: monWired, Color: t.Text},
				{Weight: monApp, Color: t.Text, Opacity: 0.55},
				{Weight: monCached, Color: t.TextMuted, Opacity: 0.35},
			}})
			// The cuts are drawn in the background color over the band, so
			// they slice it the way a water ring is cut into glasses.
			g.Marks(gauge.MarksOptions{Count: 4, Length: 34, Width: 5,
				Color: gauge.ValueColor{Color: t.Background}})
			// Swap is a smaller ring inside, on the top half alone: an
			// inset with its own sweep and its own domain.
			g.Inset(gauge.InsetOptions{
				Value: monSwapped, Min: 0, Max: monSwap,
				StartAngle: 90, EndAngle: 270,
				Draw: func(i *gauge.Canvas) {
					i.Track(gauge.TrackOptions{Width: 9, Offset: -40, Opacity: 0.12})
					i.Arc(gauge.ArcOptions{Width: 9, Offset: -40, Color: t.TextMuted})
				}})
			g.Value(gauge.ValueOptions{Y: -10, FontSize: 36, Font: gauge.FontMono, Decimals: 1})
			g.Text(gauge.TextOptions{Y: 30, FontSize: 15, Color: t.TextMuted,
				Text: fmt.Sprintf("of %d GB", monRAM)})
		},
	})
}

// monNetwork is the link: the download on the outer half and the upload
// on a ring inside it, each growing from its own end of the sweep.
func monNetwork(c *ui.Context, t *ui.Theme) ui.Element {
	const rx, tx = 412.0, 88.0
	return gauge.Gauge(c, gauge.Options{
		Value:       rx,
		Max:         monLink,
		StartAngle:  90,
		EndAngle:    270,
		Padding:     56,
		Size:        200,
		Label:       "网络",
		Description: fmt.Sprintf("下载 %.0f Mbps，上传 %.0f Mbps，链路 %d Mbps。", rx, tx, monLink),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 14, Opacity: 0.1})
			g.Arc(gauge.ArcOptions{Width: 14, Opacity: 0.35, Color: t.Accent})
			g.Inset(gauge.InsetOptions{
				Value: tx, Max: monLink,
				StartAngle: 90, EndAngle: 270,
				Draw: func(i *gauge.Canvas) {
					i.Track(gauge.TrackOptions{Width: 14, Offset: -26, Opacity: 0.1})
					i.Arc(gauge.ArcOptions{Width: 14, Offset: -26, Color: t.TextMuted})
				}})
			g.Ticks(gauge.TicksOptions{Count: 20, Length: 8, Width: 2, Offset: 16, Opacity: 0.4})
			g.Ticks(gauge.TicksOptions{Count: 4, Length: 14, Width: 4, Offset: 19})
			g.TickLabels(gauge.TickLabelsOptions{Count: 4, Offset: 48, FontSize: 15,
				Format: func(v float64) string {
					if v >= monLink {
						return "1G"
					}
					return fmt.Sprintf("%.0f", v)
				}})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Length: 0.86, Width: 14, Tail: -1, Color: t.Text})
			g.Hub(gauge.HubOptions{Radius: 12, Color: t.Text})
			g.Value(gauge.ValueOptions{Y: 42, FontSize: 20, Font: gauge.FontMono,
				Format: func(v float64) string { return fmt.Sprintf("%.0f / %.0f", v, tx) }})
		},
	})
}

// monDisk is one volume: the used share of it, with the ninety per cent
// line marked so a nearly full disk is obvious.
func monDisk(c *ui.Context, t *ui.Theme, name string, used float64) ui.Element {
	full := used >= 90
	return gauge.Gauge(c, gauge.Options{
		Value:       used,
		Max:         100,
		StartAngle:  90,
		EndAngle:    270,
		Padding:     40,
		Size:        150,
		Label:       name,
		Description: fmt.Sprintf("卷 %s 已用 %.0f%%。", name, used),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 22, Opacity: 0.1})
			opacity := float32(0.4)
			if full {
				opacity = 1
			}
			g.Arc(gauge.ArcOptions{Width: 22, Opacity: opacity, Color: t.Accent})
			g.Marks(gauge.MarksOptions{Values: []float64{90}, Length: 30, Width: 4,
				Color: gauge.ValueColor{Color: t.TextMuted}})
			g.Value(gauge.ValueOptions{Y: -22, FontSize: 30,
				Format: func(v float64) string { return fmt.Sprintf("%.0f%%", v) }})
			g.Text(gauge.TextOptions{Y: 20, FontSize: 14, Color: t.TextMuted, Text: name})
		},
	})
}
