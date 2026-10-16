package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/calendar"
	"github.com/smallnest/uix/components/checkbox"
	"github.com/smallnest/uix/components/checkboxgroup"
	"github.com/smallnest/uix/components/colorpicker"
	"github.com/smallnest/uix/components/colorwell"
	"github.com/smallnest/uix/components/datepicker"
	"github.com/smallnest/uix/components/datetimepicker"
	"github.com/smallnest/uix/components/meter"
	"github.com/smallnest/uix/components/radio"
	"github.com/smallnest/uix/components/rangeslider"
	"github.com/smallnest/uix/components/rating"
	"github.com/smallnest/uix/components/slider"
	sw "github.com/smallnest/uix/components/switch"
	"github.com/smallnest/uix/components/timeinput"
)

// demoCheckbox shows a check box with a label.
func demoCheckbox(c *ui.Context, t *ui.Theme) {
	checkbox.Checkbox(c, checkbox.Props{Checked: &Demo.Check, Label: "记住我"})
}

// demoCheckboxGroup shows a group of check boxes under a label.
func demoCheckboxGroup(c *ui.Context, t *ui.Theme) {
	checkboxgroup.CheckboxGroup(c, checkboxgroup.Props{
		Label: "语言",
		Children: func(c *ui.Context) {
			checkbox.Checkbox(c, checkbox.Props{Checked: &Demo.Check, Label: "Go"})
			checkbox.Checkbox(c, checkbox.Props{Checked: &Demo.Tg2, Label: "Rust"})
			checkbox.Checkbox(c, checkbox.Props{Checked: &Demo.Tg3, Label: "Zig"})
		},
	})
}

// demoSwitch shows a switch with a label.
func demoSwitch(c *ui.Context, t *ui.Theme) {
	sw.Switch(c, sw.Props{On: &Demo.Switch, Label: "自动保存"})
}

// demoRadio shows a radio group with one chosen.
func demoRadio(c *ui.Context, t *ui.Theme) {
	radio.Group(c, radio.Props[string]{
		Selected: &Demo.Radio,
		Options: []radio.Option[string]{
			{Value: "小", Label: "小"},
			{Value: "中", Label: "中"},
			{Value: "大", Label: "大"},
		},
	})
}

// demoSlider shows sliders: one continuous and one that snaps to a step,
// each with a label and a readout.
func demoSlider(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(340).Gap(t.Space(3)).Children(func() {
		slider.Slider(c, slider.Props{
			Value: &Demo.Slider,
			Min:   0, Max: 100, Step: 1,
			Label: "音量",
		})
		ui.Text(c, fmt.Sprintf("音量: %.0f", Demo.Slider)).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		slider.Slider(c, slider.Props{
			Value: &Demo.StepperV,
			Min:   0, Max: 10, Step: 1,
			Label: "步进",
		})
		ui.Text(c, fmt.Sprintf("步进: %.0f", Demo.StepperV)).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoRangeSlider shows a track with two knobs for a range.
func demoRangeSlider(c *ui.Context, t *ui.Theme) {
	rangeslider.RangeSlider(c, rangeslider.Props{
		Low: &Demo.RangeLo, High: &Demo.RangeHi,
		Min: 0, Max: 100, Step: 1,
		Label: "价格区间",
	})
}

// demoRating shows stars that set a score, and the same score read-only,
// as a review shows it.
func demoRating(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Gap(t.Space(3)).Children(func() {
		rating.Rating(c, rating.Props{Value: &Demo.Rating, Max: 5})
		ui.Text(c, fmt.Sprintf("可点击的评分: %d 星", Demo.Rating)).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		rating.Rating(c, rating.Props{Value: &Demo.Rating, Max: 5, ReadOnly: true})
		ui.Text(c, "只读的评分。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// demoMeter shows a progress meter that colors by its value.
func demoMeter(c *ui.Context, t *ui.Theme) {
	meter.Meter(c, meter.Props{Value: 72, Min: 0, Max: 100, Label: "磁盘占用"})
}

// demoDatePicker shows a calendar button that picks a date.
func demoDatePicker(c *ui.Context, t *ui.Theme) {
	datepicker.DatePicker(c, datepicker.Props{Value: &Demo.Date})
}

// demoTimeInput shows a time field with stepping and typing.
func demoTimeInput(c *ui.Context, t *ui.Theme) {
	timeinput.TimeInput(c, timeinput.Props{Value: &Demo.Time})
}

// demoDateTimePicker shows a date field and a time field together.
func demoDateTimePicker(c *ui.Context, t *ui.Theme) {
	datetimepicker.DateTimePicker(c, datetimepicker.Props{Value: &Demo.Both})
}

// demoCalendar shows a month's calendar.
func demoCalendar(c *ui.Context, t *ui.Theme) {
	calendar.Calendar(c, calendar.Props{Date: &Demo.Cal})
}

// demoColorPicker shows a swatch that opens a picker with channels.
func demoColorPicker(c *ui.Context, t *ui.Theme) {
	colorpicker.ColorPicker(c, colorpicker.Props{Value: &Demo.Color})
}

// demoColorWell shows a swatch that opens a picker below it.
func demoColorWell(c *ui.Context, t *ui.Theme) {
	colorwell.ColorWell(c, colorwell.Props{Color: &Demo.Well, Label: "色调"})
}

// demoStepSlider shows mygo's slider that snaps to ticks, as the arrows
// and the track step it.
func demoStepSlider(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(320).Gap(t.Space(2)).Children(func() {
		ui.StepSlider(c, &Demo.Step, 0, 100, 25)
		ui.Text(c, fmt.Sprintf("步进值: %.0f", Demo.Step)).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}
