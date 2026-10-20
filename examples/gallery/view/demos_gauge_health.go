package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/gauge"
)

// The health page is a day in a health app: the activity rings with the
// week beside them, the heart rate over its training zones, last night's
// sleep round a twenty-four hour dial, and the day's steps and water.
// The readings are the ones gauge-ui's health dashboard ships.

// The health readings, as upstream seeds them.
const (
	hlMove     = 412
	hlMoveGoal = 600
	hlExer     = 24
	hlExerGoal = 30
	hlStand    = 9
	hlStandGoal
	hlBPM      = 72
	hlResting  = 58
	hlSteps    = 8042
	hlStepGoal = 10000
	hlGlasses  = 7
	hlGlassMax = 10
)

// hlWeek is the move ring of each of the last seven days.
var hlWeek = [7]float64{520, 610, 380, 690, 470, 530, 412}

// hlStages is last night in the order it happened: the stage, and the
// hour it ended at. A fixed cutoff each, so the stack draws the night on
// in order rather than splitting one reading between parts.
var hlStages = []struct {
	name string
	to   float64
}{
	{"awake", -7.2},
	{"light", -5.4},
	{"deep", -3.1},
	{"light", -1.4},
	{"rem", -0.5},
}

// demoGaugeHealth is the health page: the stacked and ranged primitives,
// which is what a day of readings needs.
func demoGaugeHealthPage(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		gaugeCaption(c, t, "活动圆环")
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			hlActivity(c, t, "今天", hlMove, 132)
			for i, v := range hlWeek {
				hlActivity(c, t, fmt.Sprintf("星期%d", i+1), v, 96)
			}
		})
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			hlHeart(c, t)
			hlSleep(c, t)
		})
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			hlStepsDial(c, t)
			hlWater(c, t)
		})
	})
}

// hlActivity is the three activity rings: the move, exercise and stand
// goals as three arcs of one dial, each inside the last. Each ring is its
// own inset, so each takes its own goal as its domain; the radial offset
// is what keeps them from drawing over one another.
func hlActivity(c *ui.Context, t *ui.Theme, name string, move float64, size float32) ui.Element {
	goals := [3]struct {
		goal   float64
		value  float64
		offset float32
	}{
		{hlMoveGoal, move, 0},
		{hlExerGoal, move / hlMoveGoal * hlExerGoal, -(hlRingWidth + hlRingGap)},
		{hlStandGoal, move / hlMoveGoal * hlStandGoal, -2 * (hlRingWidth + hlRingGap)},
	}
	return gauge.Gauge(c, gauge.Options{
		Value:       0,
		StartAngle:  180,
		EndAngle:    540,
		Size:        size,
		Padding:     3*hlRingWidth/2 + 2*hlRingGap + 8,
		Label:       name,
		Description: fmt.Sprintf("%s的活动圆环：活动 %.0f/%.0f 千卡。", name, move, float64(hlMoveGoal)),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			for i, r := range goals {
				alpha := float32(1 - 0.34*float64(i))
				g.Inset(gauge.InsetOptions{
					Value: r.value, Min: 0, Max: r.goal,
					StartAngle: 180, EndAngle: 540,
					Draw: func(in *gauge.Canvas) {
						in.Track(gauge.TrackOptions{Width: hlRingWidth, Offset: r.offset, Opacity: 0.1})
						in.Arc(gauge.ArcOptions{Width: hlRingWidth, Offset: r.offset, Color: t.Text.Alpha(alpha)})
					}})
			}
		},
	})
}

// The activity rings: the width of a ring and the clear gap between two.
const (
	hlRingWidth = 13
	hlRingGap   = 3
)

// hlHeart is the heart rate over its training zones: the resting floor
// as an arc under the reading, the zones as a band around the rim, and
// the reading itself as a dot riding the sweep.
func hlHeart(c *ui.Context, t *ui.Theme) ui.Element {
	const low, high = 52.0, 148.0
	return gauge.Gauge(c, gauge.Options{
		Value:       hlBPM,
		Min:         40,
		Max:         200,
		StartAngle:  60,
		EndAngle:    300,
		Size:        200,
		Label:       "心率",
		Description: fmt.Sprintf("心率 %d 次每分，静息 %d。", hlBPM, hlResting),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 14, Opacity: 0.14})
			g.Arc(gauge.ArcOptions{Width: 14, Opacity: 0.4, From: ptr(low), To: ptr(high)})
			g.Zones(gauge.ZonesOptions{Width: 4, Offset: -18, Gap: 1.5, EndCap: ptr(gauge.CapRound),
				Zones: []gauge.Zone{
					{To: 100, Color: t.Text.Alpha(0.25)},
					{To: 160, Color: t.Text.Alpha(0.55)},
					{Color: t.Text},
				}})
			g.Dot(gauge.DotOptions{Radius: 11, Halo: 5, HaloColor: t.Background, Color: t.Text})
			g.Value(gauge.ValueOptions{Y: -8, FontSize: 44, Font: gauge.FontRounded, Weight: 700})
			g.Text(gauge.TextOptions{Y: 30, FontSize: 15, Color: t.TextMuted, Text: "BPM"})
		},
	})
}

// hlSleep is last night round a twenty-four hour dial: each stage of the
// night as a part of the stack, the whole night as a range, and the time
// asleep read out of the range's far end.
func hlSleep(c *ui.Context, t *ui.Theme) ui.Element {
	bed := -7.8
	wake := -0.5
	asleep := wake - bed - 0.4
	parts := make([]gauge.StackPart, 0, len(hlStages))
	for _, s := range hlStages {
		parts = append(parts, gauge.StackPart{HasTo: true, To: s.to, Color: t.Accent.Alpha(float32(0.3 + 0.7*float64(len(parts)+1)/float64(len(hlStages))))})
	}
	return gauge.Gauge(c, gauge.Options{
		Value:       wake,
		Range:       &[2]float64{bed, wake},
		Min:         -12,
		Max:         12,
		StartAngle:  0,
		EndAngle:    360,
		Size:        200,
		Label:       "睡眠",
		Description: fmt.Sprintf("昨夜睡了 %.1f 小时。", asleep),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 26, Opacity: 0.12})
			// The stack's parts each end at their own cutoff, so the
			// night draws on in the order the stages happened.
			g.Stack(gauge.StackOptions{From: ptr(bed), Width: 26, Parts: parts})
			g.Ticks(gauge.TicksOptions{Count: 24, Length: 6, Width: 2, Offset: -30, Opacity: 0.5})
			// The clock labels ride just inside the band, so they clear
			// the readout in the middle of the dial rather than running
			// into it at the hours either side.
			g.TickLabels(gauge.TickLabelsOptions{Count: 4, Offset: 22, FontSize: 15,
				Format: func(v float64) string { return gauge.ClockTime(v) }})
			// The two ends of the range are the bedtime and the hour
			// woken, which is what the two dots mark.
			g.Dot(gauge.DotOptions{Radius: 9, End: gauge.EndFrom, Color: t.TextMuted})
			g.Dot(gauge.DotOptions{Radius: 9, Color: t.Text})
			// The readout is the long night, so it is given the whole
			// middle of the dial, the caption a line under it.
			g.Text(gauge.TextOptions{Y: -12, FontSize: 24, Font: gauge.FontRounded, Weight: 700,
				Text: gauge.DurationLabel(asleep)})
			g.Text(gauge.TextOptions{Y: 20, FontSize: 15, Color: t.TextMuted, Text: "asleep"})
		},
	})
}

// hlStepsDial is the day's steps filling towards their goal, on the top
// half of a dial.
func hlStepsDial(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       hlSteps,
		Max:         hlStepGoal,
		StartAngle:  90,
		EndAngle:    270,
		Padding:     46,
		Size:        200,
		Label:       "步数",
		Description: fmt.Sprintf("今天 %d 步，目标 %d 步。", hlSteps, hlStepGoal),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 24, Opacity: 0.16})
			g.Arc(gauge.ArcOptions{Width: 24, Color: t.Accent})
			g.Ticks(gauge.TicksOptions{Count: 10, Length: 8, Width: 2, Offset: -34, Opacity: 0.4})
			g.Value(gauge.ValueOptions{Y: -30, FontSize: 34, Font: gauge.FontRounded, Weight: 700,
				Format: thousands})
			g.Text(gauge.TextOptions{Y: 4, FontSize: 15, Color: t.TextMuted,
				Text: fmt.Sprintf("of %s steps", thousands(hlStepGoal))})
		},
	})
}

// hlWater is the day's water cut into glasses: a full ring, sliced into
// as many marks as the goal has glasses, with the glasses drunk in a
// brighter tone than the rest.
func hlWater(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       hlGlasses,
		Max:         hlGlassMax,
		StartAngle:  180,
		EndAngle:    540,
		Size:        200,
		Label:       "饮水",
		Description: fmt.Sprintf("已喝 %d 杯，共 %d 杯。", hlGlasses, hlGlassMax),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 28, Opacity: 0.16})
			// The cut of each glass is drawn in the background color, so
			// the band reads as ten separate glasses.
			g.Marks(gauge.MarksOptions{Count: hlGlassMax, Length: 32, Width: 5,
				Color: gauge.ValueColor{Color: t.Background}})
			g.Arc(gauge.ArcOptions{Width: 28, Color: t.Accent})
			g.Text(gauge.TextOptions{Y: -6, FontSize: 28, Font: gauge.FontRounded, Weight: 700,
				Text: fmt.Sprintf("%.2f L", float64(hlGlasses)*0.25)})
			g.Text(gauge.TextOptions{Y: 26, FontSize: 15, Color: t.TextMuted,
				Text: fmt.Sprintf("of %.1f L", float64(hlGlassMax)*0.25)})
		},
	})
}

// thousands groups a count in threes, as "8,042".
func thousands(v float64) string {
	n := int(v)
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	out := ""
	for i, r := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			out += ","
		}
		out += string(r)
	}
	return out
}

// ptr takes the address of a value, for the option fields that ask for
// one: a nil pointer means the default, so a literal cannot be used.
func ptr[T any](v T) *T { return &v }
