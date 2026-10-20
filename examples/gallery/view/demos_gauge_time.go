package view

import (
	"fmt"
	"math"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/gauge"
)

// The time page is a wall of clocks: a clock face whose three hands turn
// at their own rates, a stopwatch with a minute register inside it, and
// the day's alarms on a twenty-four hour dial. The readings are the ones
// gauge-ui's time dashboard ships, frozen at the hour it opens on.

// The time readings, as upstream seeds them: a morning alarm and a night
// one, and a stopwatch a few minutes in.
const (
	tmHour      = 10.35 // the hour the page shows: 10:21
	tmStopwatch = 229.9 // seconds on the stopwatch
	tmAlarmWoke = 6.5   // 06:30
	tmAlarmLamp = 22.5  // 22:30
	tmTimerTea  = 3.0   // the tea timer's length, in minutes
	tmTeaLeft   = 1.2   // and what is left of it
)

// tmCities are the offsets of the world clocks, in hours from the page's
// own hour.
var tmCities = []struct {
	name   string
	offset float64
}{
	{"旧金山", -8}, {"纽约", -3}, {"雷克雅未克", 2}, {"伦敦", 2}, {"东京", 11}, {"悉尼", 13},
}

// demoGaugeTime is the time page: the turning primitives, which is what a
// clock face needs and nothing else in the library does.
func demoGaugeTimePage(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(4)).Children(func() {
		gaugeCaption(c, t, "世界时钟")
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			tmClock(c, t, "本地", tmHour)
			for _, city := range tmCities {
				tmClock(c, t, city.name, tmHour+city.offset)
			}
		})
		ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
			tmStopwatchDial(c, t)
			tmAlarms(c, t)
			tmTimer(c, t)
		})
	})
}

// tmClock is a clock face: the hours, minutes and seconds as three hands
// of one reading, each turning its own number of times across the twelve
// hour domain, so the minute hand laps the hour hand twelve times over.
func tmClock(c *ui.Context, t *ui.Theme, name string, hours float64) ui.Element {
	turned := math.Mod(math.Mod(hours, 12)+12, 12)
	label := gauge.ClockTime(hours)
	return gauge.Gauge(c, gauge.Options{
		Value:       turned,
		Min:         0,
		Max:         12,
		StartAngle:  180,
		EndAngle:    540,
		Size:        150,
		Label:       name,
		Description: fmt.Sprintf("%s的时间：%s。", name, label),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Ticks(gauge.TicksOptions{Count: 60, Length: 5, Width: 2, Offset: -4, Opacity: 0.5})
			g.Ticks(gauge.TicksOptions{Count: 12, Length: 12, Width: 3, Offset: -8})
			g.TickLabels(gauge.TickLabelsOptions{Count: 12, Offset: -30, FontSize: 15,
				Format: gauge.ClockLabel})
			// The hour hand turns once and the minute hand twelve times
			// across the same domain, which is the whole of Turns.
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Length: 0.44, Width: 11, Tail: -1, Color: t.Text})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Length: 0.68, Width: 8, Tail: -1, Turns: 12, Color: t.Text})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedleLine, Length: 0.78, Width: 2, Tail: -1, Turns: 720, Color: t.Accent})
			g.Hub(gauge.HubOptions{Radius: 7, Color: t.Text})
			g.Text(gauge.TextOptions{Y: 46, FontSize: 15, Color: t.TextMuted, Text: label})
		},
	})
}

// tmStopwatchDial is the chronograph: the seconds round the dial at a
// hundredth's resolution, with the minutes on a register inside it.
func tmStopwatchDial(c *ui.Context, t *ui.Theme) ui.Element {
	seconds := tmStopwatch
	return gauge.Gauge(c, gauge.Options{
		Value:       seconds,
		Min:         0,
		Max:         60,
		StartAngle:  180,
		EndAngle:    540,
		Size:        200,
		Label:       "秒表",
		Description: fmt.Sprintf("秒表计时 %.1f 秒。", seconds),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Ticks(gauge.TicksOptions{Count: 240, Length: 5, Width: 1.5, Offset: -3, Opacity: 0.4})
			g.Ticks(gauge.TicksOptions{Count: 60, Length: 10, Width: 2, Offset: -6})
			g.Ticks(gauge.TicksOptions{Count: 12, Length: 16, Width: 3, Offset: -9})
			g.TickLabels(gauge.TickLabelsOptions{Count: 12, Offset: -30, FontSize: 16,
				Format: func(v float64) string {
					if v == 0 {
						return "60"
					}
					return fmt.Sprintf("%.0f", v)
				}})
			// The minute register is a dial inside the dial, at a third
			// of the scale and the same sweep: half an hour round once.
			g.Inset(gauge.InsetOptions{
				Value: math.Mod(seconds/60, 30), Min: 0, Max: 30,
				Y: 46, Scale: 0.28,
				StartAngle: 180, EndAngle: 540,
				Draw: func(in *gauge.Canvas) {
					in.Track(gauge.TrackOptions{Width: 5, Offset: -3, Opacity: 0.25})
					in.Ticks(gauge.TicksOptions{Count: 30, Length: 12, Width: 4, Offset: -12, Opacity: 0.5})
					in.Ticks(gauge.TicksOptions{Count: 6, Length: 22, Width: 7, Offset: -18})
					in.Needle(gauge.NeedleOptions{Style: gauge.NeedlePointer, Length: 0.78, Width: 26, Tail: -1, Color: t.Text})
					in.Hub(gauge.HubOptions{Radius: 16, Color: t.Text})
				}})
			g.Needle(gauge.NeedleOptions{Style: gauge.NeedleLine, Length: 0.9, Width: 3, Tail: 30,
				TailDot: 7, Color: t.Accent})
			g.Hub(gauge.HubOptions{Radius: 9, Color: t.Text})
			g.Text(gauge.TextOptions{Y: -60, FontSize: 22, Font: gauge.FontMono, Color: t.Text,
				Text: fmt.Sprintf("03:%05.2f", seconds)})
		},
	})
}

// tmAlarms is the day's alarms on a twenty-four hour dial: each alarm as
// a dot of its own inset, the next one read out in the middle.
func tmAlarms(c *ui.Context, t *ui.Theme) ui.Element {
	alarms := []struct {
		name string
		at   float64
		on   bool
	}{
		{"起床", tmAlarmWoke, true},
		{"灯光", tmAlarmLamp, true},
	}
	return gauge.Gauge(c, gauge.Options{
		Value:       tmAlarmWoke,
		Min:         0,
		Max:         24,
		StartAngle:  180,
		EndAngle:    540,
		Size:        200,
		Label:       "闹钟",
		Description: "两个闹钟：起床 06:30，灯光 22:30。",
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 8, Opacity: 0.12})
			// The night's sleep is the arc from the lamp going out to
			// the alarm going off, which crosses midnight: Wrap carries
			// it round the seam in one piece rather than clamping it.
			g.Arc(gauge.ArcOptions{Width: 8, From: ptr(tmAlarmLamp), To: ptr(tmAlarmWoke + 24),
				Wrap: true, Color: t.Accent.Alpha(0.5)})
			g.Ticks(gauge.TicksOptions{Count: 96, Length: 5, Width: 1.5, Offset: -18, Opacity: 0.4})
			g.Ticks(gauge.TicksOptions{Count: 24, Length: 9, Width: 2, Offset: -20, Opacity: 0.8})
			g.TickLabels(gauge.TickLabelsOptions{Count: 4, Offset: -42, FontSize: 15,
				Format: func(v float64) string { return gauge.ClockTime(v) }})
			for _, a := range alarms {
				radius := float32(7)
				color := t.TextMuted
				opacity := float32(0.6)
				if a.on {
					radius, color, opacity = 10, t.Text, 1
				}
				g.Inset(gauge.InsetOptions{
					Value: a.at, Min: 0, Max: 24,
					StartAngle: 180, EndAngle: 540,
					Draw: func(in *gauge.Canvas) {
						in.Dot(gauge.DotOptions{Radius: radius, Color: color, Opacity: opacity})
					}})
			}
			g.Text(gauge.TextOptions{Y: -6, FontSize: 26, Font: gauge.FontMono, Weight: 600,
				Text: gauge.ClockTime(tmAlarmWoke)})
			g.Text(gauge.TextOptions{Y: 24, FontSize: 14, Color: t.TextMuted, Text: "起床"})
		},
	})
}

// tmTimer is the kitchen timer running down: the tea on the outer ring,
// the focus timer inside it, each counting to its own length.
func tmTimer(c *ui.Context, t *ui.Theme) ui.Element {
	return gauge.Gauge(c, gauge.Options{
		Value:       tmTimerTea - tmTeaLeft,
		Min:         0,
		Max:         tmTimerTea,
		StartAngle:  180,
		EndAngle:    540,
		Size:        200,
		Label:       "计时",
		Description: fmt.Sprintf("泡茶还剩 %.1f 分钟，共 %.0f 分钟。", tmTeaLeft, tmTimerTea),
		Transition:  gauge.Transition{Enabled: true},
		Draw: func(g *gauge.Canvas) {
			g.Track(gauge.TrackOptions{Width: 7, Opacity: 0.12})
			g.Arc(gauge.ArcOptions{Width: 7, Color: t.Accent})
			g.Ticks(gauge.TicksOptions{Count: 12, Length: 7, Width: 2, Offset: 16})
			g.Ticks(gauge.TicksOptions{Count: 60, Length: 4, Width: 1.5, Offset: 14, Opacity: 0.5})
			g.Inset(gauge.InsetOptions{
				Value: 18.667, Min: 0, Max: 25,
				Scale:      0.72,
				StartAngle: 180, EndAngle: 540,
				Draw: func(in *gauge.Canvas) {
					in.Track(gauge.TrackOptions{Width: 7, Opacity: 0.12})
					in.Arc(gauge.ArcOptions{Width: 7, Color: t.TextMuted})
				}})
			g.Text(gauge.TextOptions{Y: -34, FontSize: 14, Color: t.TextMuted, Text: "泡茶"})
			g.Text(gauge.TextOptions{Y: 0, FontSize: 30, Font: gauge.FontMono, Weight: 600,
				Text: gauge.DurationLabel(tmTeaLeft / 60)})
			g.Text(gauge.TextOptions{Y: 28, FontSize: 13, Color: t.TextMuted,
				Text: fmt.Sprintf("of %.0f:00", tmTimerTea)})
		},
	})
}
