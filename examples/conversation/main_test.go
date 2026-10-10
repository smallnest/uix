package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/conversation/view"
)

// tickTime is the time the tests advance the simulated upload by.
var tickTime = time.Now()

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.ConversationView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading, the turns, the
// markers and the attachment.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Conversation",
		"AI", "Thanks for the report. I have looked at the numbers, and the growth is steady.",
		"You", "Here is the file you asked for.",
		"report.pdf", "1.2 MB · uploading",
		"Upload a file",
		"The upload is simulated, so the example runs anywhere.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestRemove removes the attachment for a click on its X, at the far
// right of the card.
func TestRemove(t *testing.T) {
	tt := newTester(t)
	tt.ClickAt(385, 205)
	if view.State.File != "" {
		t.Fatalf("File = %q, want removed", view.State.File)
	}
}

// TestUpload starts an upload with the button, and removes the old one.
func TestUpload(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	if view.State.File != "sales-report.xlsx" {
		t.Fatalf("File = %q, want sales-report.xlsx", view.State.File)
	}
}

// TestTick advances the simulated upload to its completion: the reset
// leaves it at 0.35, and each tick adds 0.05.
func TestTick(t *testing.T) {
	newTester(t)
	for i := 0; i < 14; i++ {
		view.Tick(tickTime)
	}
	if view.State.Progress != 1 || view.State.Desc != "1.2 MB · uploaded" {
		t.Fatalf("progress %v, desc %q; want a finished upload", view.State.Progress, view.State.Desc)
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	for _, s := range []string{"Conversation", "AI", "You"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in the dark frame %q", s, tt.Texts())
		}
	}
}
