package main

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/i18n"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/fields/view"
)

// newTester draws the example headless in its initial state, in a
// 24-hour region: the format the fields of the Chinese locale show, and
// the headless tester otherwise defaults to 12-hour English.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	tt := ui.NewTester(view.FieldsView, view.Width, view.Height)
	tt.SetRegion(i18n.Region{HourCycle: 24})
	return tt
}

// TestRenders draws the heading, the fields and the date of the picker.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{"New event", "Invite guests", "10/18/2026", "15"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTypes types a title into the first floating label field.
func TestTypes(t *testing.T) {
	tt := newTester(t)
	// The first field is under the heading, near the top of the window.
	tt.ClickAt(60, 70)
	tt.Type("Team lunch")
	if view.State.Title != "Team lunch" {
		t.Fatalf("title %q, want Team lunch", view.State.Title)
	}
}

// TestPicksGuest chooses a guest in the multi-select's popup.
func TestPicksGuest(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Invite guests"); err != nil {
		t.Fatal(err)
	}
	r, ok := tt.Find("Ada")
	if !ok {
		t.Fatalf("the popup shows %q", tt.Texts())
	}
	tt.ClickAt(r.X+r.W/2, r.Y+r.H/2)
	if !slices.Equal(view.State.Guests, []string{"Ada"}) {
		t.Fatalf("guests %q, want [Ada]", view.State.Guests)
	}
}

// TestStepsTime raises the hour of the picker.
func TestStepsTime(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("15"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyUp)
	if view.State.When.Hour() != 16 {
		t.Fatalf("hour %d, want 16", view.State.When.Hour())
	}
}

// TestSaves counts a click of Save.
func TestSaves(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	if view.State.Saved != 1 {
		t.Fatalf("saved %d, want 1", view.State.Saved)
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("New event") {
		t.Fatal("fields missing under dark mode")
	}
}
