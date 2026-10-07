package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/events/view"
)

// TestRenders draws the panel headless: every field and the summary line
// that reflects the initial state.
func TestRenders(t *testing.T) {
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	for _, s := range []string{
		"Event", "Date", "Time", "Color", "Remind me", "Off",
		"The event is on October 15, 2026 at 14:30.",
		"Tag color #22c55e",
		"No reminder.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestDate chooses another day in the calendar of the date field, which
// the summary line shows at once.
func TestDate(t *testing.T) {
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	// The field label "Date" shares its text with the field, so the field
	// itself is clicked by the name it carries.
	if err := tt.Click("Event date"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("October 2026") {
		t.Fatal("the calendar did not open")
	}
	if err := tt.Click("October 20, 2026"); err != nil {
		t.Fatal(err)
	}
	if view.State.When.Day() != 20 {
		t.Fatalf("chose the %dth, want the 20th", view.State.When.Day())
	}
	if !tt.HasText("The event is on October 20, 2026 at 14:30.") {
		t.Fatal("the summary did not follow the date")
	}
}

// TestTime raises the hours of the time field with Up, which the summary
// line shows at once.
func TestTime(t *testing.T) {
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	if err := tt.Click("14"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyUp)
	if view.State.When.Hour() != 15 {
		t.Fatalf("Up stepped the hours to %d, want 15", view.State.When.Hour())
	}
	if !tt.HasText("The event is on October 15, 2026 at 15:30.") {
		t.Fatal("the summary did not follow the time")
	}
}

// TestColor picks a color from the swatches of the color field, which the
// summary line shows at once.
func TestColor(t *testing.T) {
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	if err := tt.Click("Event color"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Blue"); err != nil {
		t.Fatal(err)
	}
	if view.State.Color != ui.Hex("#3b82f6") {
		t.Fatalf("chose %v, want blue", view.State.Color)
	}
	if !tt.HasText("Tag color #3b82f6") {
		t.Fatal("the summary did not follow the color")
	}
}

// TestToggle arms the reminder with a click and disarms it with another,
// which the summary line shows at once.
func TestToggle(t *testing.T) {
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	if err := tt.Click("Off"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Remind {
		t.Fatal("the first click did not arm the reminder")
	}
	if !tt.HasText("A reminder is set.") {
		t.Fatal("the summary did not follow the toggle")
	}
	if err := tt.Click("On"); err != nil {
		t.Fatal(err)
	}
	if view.State.Remind {
		t.Fatal("the second click did not disarm the reminder")
	}
	if !tt.HasText("No reminder.") {
		t.Fatal("the summary did not follow the toggle back")
	}
}

// TestDarkMode draws the panel under the dark appearance too.
func TestDarkMode(t *testing.T) {
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Event") {
		t.Fatal("dark panel missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
