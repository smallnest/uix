package timeinput

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			TimeInput(c, p).Label("Time")
		})
	}
}

// TestRenders draws the field with the hours and the minutes in it.
func TestRenders(t *testing.T) {
	tm := time.Date(2026, 10, 7, 14, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &tm}), 260, 100)
	if _, ok := tt.Find("Time"); !ok {
		t.Fatal("field not found")
	}
	for _, s := range []string{"14", ":", "30"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestStepsUp raises the hours with Up while they have the focus.
func TestStepsUp(t *testing.T) {
	tm := time.Date(2026, 10, 7, 14, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &tm}), 260, 100)
	if err := tt.Click("14"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyUp)
	if tm.Hour() != 15 {
		t.Fatalf("Up stepped the hours to %d, want 15", tm.Hour())
	}
}

// TestStepsDown wraps the minutes around with Down.
func TestStepsDown(t *testing.T) {
	tm := time.Date(2026, 10, 7, 14, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &tm}), 260, 100)
	if err := tt.Click("30"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyDown)
	if tm.Minute() != 29 {
		t.Fatalf("Down stepped the minutes to %d, want 29", tm.Minute())
	}
}

// TestTypes sets the hours from the digits typed.
func TestTypes(t *testing.T) {
	tm := time.Date(2026, 10, 7, 14, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &tm}), 260, 100)
	if err := tt.Click("14"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.Key2)
	tt.Key(0, ui.Key1)
	if tm.Hour() != 21 {
		t.Fatalf("typed hours %d, want 21", tm.Hour())
	}
}

// TestDisabled keeps the field from being changed.
func TestDisabled(t *testing.T) {
	tm := time.Date(2026, 10, 7, 14, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &tm, Disabled: true}), 260, 100)
	if err := tt.Click("14"); err != nil {
		t.Fatal(err)
	}
	tt.Type("5")
	tt.Key(0, ui.KeyUp)
	if tm.Hour() != 14 || tm.Minute() != 30 {
		t.Fatalf("a disabled field changed to %02d:%02d", tm.Hour(), tm.Minute())
	}
}

// TestDarkMode draws the field under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tm := time.Date(2026, 10, 7, 14, 30, 0, 0, time.UTC)
	tt := ui.NewTester(frame(Props{Value: &tm}), 260, 100)
	tt.SetDark(true)
	if _, ok := tt.Find("Time"); !ok {
		t.Fatal("field missing under dark mode")
	}
}
