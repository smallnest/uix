package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/placeholders/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.PlaceholdersView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading, the sections and
// the search hint.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Placeholders",
		"Skeleton", "Empty state", "Aspect ratio",
		"Type a query, then press Enter to search.",
		"The box keeps 16:9 as the window changes.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestEmpty shows the empty state for a search that found nothing, and
// hides it for the Clear action.
func TestEmpty(t *testing.T) {
	tt := newTester(t)
	// The search field is under the Empty state label; click it, type a
	// query and submit it.
	heading, _ := tt.Find("Empty state")
	tt.ClickAt(heading.X+100, heading.Y+heading.H+16)
	tt.Frame()
	tt.Type("nope")
	tt.Frame()
	tt.Key(0, ui.KeyEnter)
	tt.Frame()
	if !tt.HasText("No results") {
		t.Fatalf("missing the empty state in %q", tt.Texts())
	}
	if err := tt.Click("Clear search"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("No results") {
		t.Fatal("the empty state stayed after Clear search")
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	for _, s := range []string{"Placeholders", "Aspect ratio"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in the dark frame %q", s, tt.Texts())
		}
	}
}
