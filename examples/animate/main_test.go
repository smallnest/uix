package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/animate/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() { view.Reset() }

// TestRenders draws the animate view headless and checks that every part
// came through: the header, the cards, and the slide-in list. The
// entrances start at their first values, but the text is in the tree.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ExampleView, view.Width, view.Height)
	for _, s := range []string{"Animate", "Recording", "Replay", "Fade", "Rise", "Slide", "Scale", "First row", "Second row", "Third row"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestReplay clicks the replay button, which bumps the Run that replays
// the entrances.
func TestReplay(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ExampleView, view.Width, view.Height)
	if err := tt.Click("Replay"); err != nil {
		t.Fatal(err)
	}
	if view.State.Run == 0 {
		t.Fatal("the replay button did not bump the run")
	}
}

// TestDarkMode draws the animate view under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ExampleView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Animate") {
		t.Fatal("dark view missing the title")
	}
}
