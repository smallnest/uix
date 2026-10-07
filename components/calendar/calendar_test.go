package calendar

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Calendar(c, p)
		})
	}
}

func day() time.Time { return time.Date(2026, 10, 15, 9, 41, 0, 0, time.UTC) }

// TestRenders shows the month and the day chosen.
func TestRenders(t *testing.T) {
	d := day()
	tt := ui.NewTester(frame(Props{Date: &d}), 400, 300)
	for _, s := range []string{"October 2026", "15"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestChooses picks a day with a click, which keeps the time of day.
func TestChooses(t *testing.T) {
	d := day()
	changes := 0
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			if Calendar(c, Props{Date: &d}).Changed() {
				changes++
			}
		})
	}, 400, 300)
	if err := tt.Click("October 20, 2026"); err != nil {
		t.Fatal(err)
	}
	if d.Day() != 20 || d.Month() != 10 || d.Hour() != 9 || changes != 1 {
		t.Fatalf("the click chose %v (%d changes), want Oct 20 09:00 (1)", d, changes)
	}
}

// TestKeys moves the choice with the keys, as a click does: Right the
// next day, End the last of the month, Page Down the first of the next.
func TestKeys(t *testing.T) {
	d := day()
	tt := ui.NewTester(frame(Props{Date: &d}), 400, 300)
	if err := tt.Click("October 20, 2026"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyRight)
	if got := d.Format("2006-01-02"); got != "2026-10-21" {
		t.Fatalf("Right chose %s, want 2026-10-21", got)
	}
	tt.Key(0, ui.KeyEnd)
	if got := d.Format("2006-01-02"); got != "2026-10-31" {
		t.Fatalf("End chose %s, want 2026-10-31", got)
	}
	tt.Key(0, ui.KeyHome)
	if got := d.Format("2006-01-02"); got != "2026-10-01" {
		t.Fatalf("Home chose %s, want 2026-10-01", got)
	}
	tt.Key(0, ui.KeyPageDown)
	if got := d.Format("2006-01-02"); got != "2026-11-01" {
		t.Fatalf("Page Down chose %s, want 2026-11-01", got)
	}
}
