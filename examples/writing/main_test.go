package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/writing/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.WritingView, view.Width, view.Height)
}

// TestRenders draws the heading, the prose and the block quote.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Getting Started",
		"go get github.com/smallnest/uix/components/button",
		"Simplicity is the soul of efficiency.",
		"Austin Freeman",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTypes shows the typewriter line once the typing is done.
func TestTypes(t *testing.T) {
	tt := newTester(t)
	tt.Advance(1100 * time.Millisecond)
	if !tt.HasText("Build native apps with uix.") {
		t.Fatalf("the typewriter line is missing in %q", tt.Texts())
	}
}

// TestCopies copies the anchor of the heading to the clipboard.
func TestCopies(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("Getting Started")
	if !ok {
		t.Fatal("no heading")
	}
	tt.ClickAt(r.X+r.W+12, r.Y+r.H/2)
	if got := tt.Clipboard(); got != "uix#getting-started" {
		t.Fatalf("clipboard %q, want uix#getting-started", got)
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Getting Started") {
		t.Fatal("writing missing under dark mode")
	}
}
