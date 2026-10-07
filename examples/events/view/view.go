// Package view builds the events example: a window of MyGo native UI
// that shows the datepicker, timeinput, colorpicker and toggle components
// together, as an event the fields edit and a summary line reflects.
// The main package shows it in a window; the tests and the snapshot
// command draw it headless.
package view

import (
	"fmt"
	"time"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/colorpicker"
	"github.com/smallnest/uix/components/datepicker"
	"github.com/smallnest/uix/components/separator"
	"github.com/smallnest/uix/components/timeinput"
	"github.com/smallnest/uix/components/toggle"
)

// Width and Height are the size of the example window.
const Width, Height = 420, 460

// State is the state of the example; the controls edit it in place.
var State = EventState{}

// EventState holds the event, as the controls edit it.
type EventState struct {
	When   time.Time
	Color  ui.Color
	Remind bool
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = EventState{
		When:   time.Date(2026, 10, 15, 14, 30, 0, 0, time.Local),
		Color:  ui.Hex("#22c55e"),
		Remind: false,
	}
}

// hexOf returns the color as a web hex string.
func hexOf(c ui.Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// EventView draws the example: the fields of an event above a summary
// line that reflects them, so every change shows at once.
func EventView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Event").FontSize(24).Bold()
		// The fields name their controls as the tests find them: a field's
		// visible label shares its text, so the control carries a name of
		// its own to click.
		ui.Field(c, "Date", func() {
			datepicker.DatePicker(c, datepicker.Props{Value: &State.When}).Label("Event date")
		})
		ui.Field(c, "Time", func() {
			timeinput.TimeInput(c, timeinput.Props{Value: &State.When}).Label("Event time")
		})
		ui.Field(c, "Color", func() {
			colorpicker.ColorPicker(c, colorpicker.Props{Value: &State.Color}).Label("Event color")
		})
		// The toggle names its state: it reads "Off" until it is pressed,
		// then "On".
		label := "Off"
		if State.Remind {
			label = "On"
		}
		ui.Field(c, "Remind me", func() {
			toggle.Toggle(c, toggle.Props{On: &State.Remind, Label: label})
		})
		separator.Separator(c, separator.Props{})
		ui.Text(c, fmt.Sprintf("The event is on %s at %s.",
			State.When.Format("January 2, 2006"), State.When.Format("15:04")))
		ui.Text(c, "Tag color "+hexOf(State.Color)).TextColor(t.TextMuted)
		if State.Remind {
			ui.Text(c, "A reminder is set.").TextColor(t.TextMuted)
		} else {
			ui.Text(c, "No reminder.").TextColor(t.TextMuted)
		}
	})
}
