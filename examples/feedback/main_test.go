package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/feedback/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() { view.State = view.FeedbackState{} }

// TestRenders draws the feedback view headless and checks that every
// control came through: MyGo renders frames in memory, no window needed.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
	for _, s := range []string{
		"Feedback", "v1.2.0", "Beta", "Needs attention",
		"Disk usage", "64%", "Indexing", "Show toast", "Delete account",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestToast shows a toast and closes it with its close button.
func TestToast(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
	if err := tt.Click("Show toast"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"Settings saved", "They take effect at once."} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
	if err := tt.Click("✕"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("Settings saved") {
		t.Fatal("toast did not close")
	}
}

// TestDialog opens the delete dialog, cancels it, and checks the toast
// the cancel pushes.
func TestDialog(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
	if err := tt.Click("Delete account"); err != nil {
		t.Fatal(err)
	}
	if !view.State.DeleteOpen {
		t.Fatal("dialog did not open")
	}
	for _, s := range []string{"Delete your account?", "This removes your data and cannot be undone."} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
	if err := tt.Click("Keep account"); err != nil {
		t.Fatal(err)
	}
	if view.State.DeleteOpen {
		t.Fatal("cancel did not close the dialog")
	}
	if !tt.HasText("Account kept") {
		t.Fatal("cancel toast missing")
	}
}

// TestDialogConfirm confirms the delete and checks the error toast.
func TestDialogConfirm(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
	if err := tt.Click("Delete account"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Delete"); err != nil {
		t.Fatal(err)
	}
	if view.State.DeleteOpen {
		t.Fatal("confirm did not close the dialog")
	}
	if !tt.HasText("Account deleted") {
		t.Fatal("confirm toast missing")
	}
}

// TestDarkMode draws the feedback view under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Feedback") {
		t.Fatal("dark view missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
