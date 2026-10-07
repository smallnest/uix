// Package view builds the composer example: a window of MyGo native UI
// that shows the autocomplete, tokenfield, calendar and checkboxgroup
// components together in a form creating an appointment. The task field
// completes what is typed from the tasks known, the tag field keeps a
// set of tags, the calendar picks the due day, and the reminder group
// chooses how to be reminded. The main package shows it in a window; the
// tests and the snapshot command draw it headless.
package view

import (
	"fmt"
	"strings"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/autocomplete"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/calendar"
	"github.com/smallnest/uix/components/checkbox"
	"github.com/smallnest/uix/components/checkboxgroup"
	"github.com/smallnest/uix/components/separator"
	"github.com/smallnest/uix/components/tokenfield"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 600

// State is the state of the example; the controls edit it in place.
var State = ComposerState{}

// ComposerState holds the appointment, as the controls edit it.
type ComposerState struct {
	// Title is the task, edited by the field that completes it.
	Title string
	// Tags are the tags of the appointment, edited by the field of
	// tokens.
	Tags []string
	// Due is the day the appointment falls on, edited by the calendar.
	Due time.Time
	// Mail, Schedule and Bounce are the reminders, edited by the group.
	Mail, Schedule, Bounce bool
	// Status is the last thing the controls did, shown at the bottom.
	Status string
}

// lastRemind lets the view tell a change of the reminders from a plain
// rebuild: the group has no change of its own to report.
var lastRemind [3]bool

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = ComposerState{
		Due:    time.Date(2026, 10, 15, 9, 0, 0, 0, time.UTC),
		Status: "Type a title, or pick a due day.",
	}
	lastRemind = [3]bool{}
}

// tasks are the tasks the field knows, which complete what is typed.
var tasks = []string{
	"Water the plants",
	"Write the report",
	"Book the room",
	"Call the client",
	"Review the pull requests",
}

// tagSuggestions are the tags the field offers.
var tagSuggestions = []string{"urgent", "later", "work", "home", "call"}

// ComposerView draws the example: an appointment form whose controls all
// work.
func ComposerView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).PaddingX(16).PaddingY(10).Gap(12).AlignItems(ui.Center).Children(func() {
			ui.Text(c, "New appointment").Bold().Grow(1)
		})
		separator.Separator(c, separator.Props{})
		ui.Column(c).Padding(24).Gap(14).Children(func() {
			caption(c, "Task")
			title := autocomplete.Autocomplete(c, autocomplete.Props{
				Value:       &State.Title,
				Suggestions: tasks,
				Label:       "Task title",
			})
			if title.Submitted() && State.Title != "" {
				State.Status = "Added “" + State.Title + "” to the schedule."
			}
			caption(c, "Tags")
			tag := tokenfield.TokenField(c, tokenfield.Props{
				Tokens:      &State.Tags,
				Suggestions: tagSuggestions,
				Label:       "Add a tag",
			})
			if tag.Changed() {
				State.Status = "Tags: " + strings.Join(State.Tags, ", ") + "."
			}
			// The calendar and the reminders sit side by side, so the
			// form fits the window.
			ui.Row(c).Gap(24).AlignItems(ui.Start).Children(func() {
				ui.Column(c).Gap(14).Shrink(0).Children(func() {
					caption(c, "Due")
					due := calendar.Calendar(c, calendar.Props{Date: &State.Due})
					if due.Changed() {
						State.Status = "Due “" + State.Due.Format("Mon, Jan 2, 2006") + "”."
					}
				})
				ui.Column(c).Gap(14).Shrink(0).Children(func() {
					caption(c, "Remind")
					checkboxgroup.CheckboxGroup(c, checkboxgroup.Props{
						Label: "Remind me by",
						Children: func(c *ui.Context) {
							checkbox.Checkbox(c, checkbox.Props{Checked: &State.Mail, Label: "Mail"})
							checkbox.Checkbox(c, checkbox.Props{Checked: &State.Schedule, Label: "Calendar"})
							checkbox.Checkbox(c, checkbox.Props{Checked: &State.Bounce, Label: "Bounce"})
						},
					})
					if reminded() {
						State.Status = "Reminders updated."
						lastRemind = current()
					}
				})
			})
		})
		separator.Separator(c, separator.Props{})
		ui.Row(c).PaddingX(24).PaddingY(10).Gap(12).AlignItems(ui.Center).Children(func() {
			ui.Text(c, State.Status).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted).Grow(1)
			if ui.Button(c, "Create").Clicked() {
				State.Status = fmt.Sprintf("Appointment “%s” on %s created.", State.Title, State.Due.Format("Mon, Jan 2, 2006"))
			}
		})
	})
}

// caption draws the muted label above a field.
func caption(c *ui.Context, s string) {
	t := c.Theme()
	ui.Text(c, s).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
}

// current returns the reminders as an array.
func current() [3]bool { return [3]bool{State.Mail, State.Schedule, State.Bounce} }

// reminded reports whether the reminders changed since the last frame.
func reminded() bool { return current() != lastRemind }
