package datetimepicker

import (
	"testing"
	"time"

	"github.com/egoist/mygo/i18n"
	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			DateTimePicker(c, p).Label("When")
		})
	}
}

// tester draws the picker in a 24-hour region, the format the Chinese
// locale of the gallery shows; the headless tester otherwise defaults to
// 12-hour English.
func tester(p Props, w, h int) *ui.Tester {
	tt := ui.NewTester(frame(p), w, h)
	tt.SetRegion(i18n.Region{HourCycle: 24})
	return tt
}

// TestRenders draws the date and the time in the row.
func TestRenders(t *testing.T) {
	when := time.Date(2026, 10, 11, 14, 30, 0, 0, time.UTC)
	tt := tester(Props{Value: &when}, 420, 100)
	if _, ok := tt.Find("When"); !ok {
		t.Fatal("picker not found")
	}
	for _, s := range []string{"10/11/2026", "14", ":", "30"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestEditsTime steps the hours without touching the date.
func TestEditsTime(t *testing.T) {
	when := time.Date(2026, 10, 11, 14, 30, 0, 0, time.UTC)
	tt := tester(Props{Value: &when}, 420, 100)
	if err := tt.Click("14"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyUp)
	if when.Hour() != 15 || when.Day() != 11 {
		t.Fatalf("stepping the hours changed the time to %v", when)
	}
}

// TestDisabled keeps the fields from being changed.
func TestDisabled(t *testing.T) {
	when := time.Date(2026, 10, 11, 14, 30, 0, 0, time.UTC)
	tt := tester(Props{Value: &when, Disabled: true}, 420, 100)
	if err := tt.Click("14"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyUp)
	if when.Hour() != 14 {
		t.Fatalf("a disabled picker changed to %v", when)
	}
}

// TestDarkMode draws the picker under the dark appearance too.
func TestDarkMode(t *testing.T) {
	when := time.Date(2026, 10, 11, 14, 30, 0, 0, time.UTC)
	tt := tester(Props{Value: &when}, 420, 100)
	tt.SetDark(true)
	if _, ok := tt.Find("When"); !ok {
		t.Fatal("picker missing under dark mode")
	}
}
