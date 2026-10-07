package datepicker

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			DatePicker(c, p).Label("Date")
		})
	}
}

// TestRenders draws the field with the date in it.
func TestRenders(t *testing.T) {
	d := time.Date(2026, 10, 7, 9, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &d}), 300, 120)
	if _, ok := tt.Find("Date"); !ok {
		t.Fatal("field not found")
	}
	if !tt.HasText("2026-10-07") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestOpens shows the month's calendar when the field is clicked.
func TestOpens(t *testing.T) {
	d := time.Date(2026, 10, 7, 9, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &d}), 300, 160)
	if err := tt.Click("Date"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("October 2026") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestChooses picks a day in the calendar, which closes it.
func TestChooses(t *testing.T) {
	d := time.Date(2026, 10, 7, 9, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &d}), 300, 160)
	tt.Click("Date")
	if err := tt.Click("October 15, 2026"); err != nil {
		t.Fatal(err)
	}
	if want := time.Date(2026, 10, 15, 9, 30, 0, 0, time.UTC); !d.Equal(want) {
		t.Fatalf("chose %v, want the 15th with the time of day", d)
	}
	if tt.HasText("October 2026") {
		t.Fatal("the calendar still shows after the choice")
	}
}

// TestKeyboard moves with the arrows and chooses with Enter.
func TestKeyboard(t *testing.T) {
	d := time.Date(2026, 10, 7, 9, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &d}), 300, 160)
	tt.Click("Date")
	tt.Key(0, ui.KeyRight)
	tt.Key(0, ui.KeyDown)
	tt.Key(0, ui.KeyEnter)
	if want := time.Date(2026, 10, 15, 9, 30, 0, 0, time.UTC); !d.Equal(want) {
		t.Fatalf("Right and Down chose %v, want the 15th", d)
	}
}

// TestDisabled opens no calendar when the field is clicked.
func TestDisabled(t *testing.T) {
	d := time.Date(2026, 10, 7, 9, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &d, Disabled: true}), 300, 160)
	if err := tt.Click("Date"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("October 2026") {
		t.Fatal("a disabled field opens its calendar")
	}
}

// TestDarkMode draws the field and its calendar under the dark
// appearance too.
func TestDarkMode(t *testing.T) {
	d := time.Date(2026, 10, 7, 9, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &d}), 300, 160)
	tt.SetDark(true)
	tt.Click("Date")
	if !tt.HasText("October 2026") {
		t.Fatal("calendar missing under dark mode")
	}
}
