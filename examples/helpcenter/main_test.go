package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/helpcenter/view"
)

// newTester resets the example and returns a tester of its page with the
// animation off, so the content is gone the frame a section closes, as
// the tests assert.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	tt := ui.NewTester(view.HelpView, view.Width, view.Height)
	tt.SetPreferences(ui.Preferences{ReduceMotion: true, TextScale: 1})
	return tt
}

// TestRenders draws the page headless: the trail, the sections, the
// collapsible and the help button, with the first answer and the more
// options in view.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Help Center", "Home", "Docs", "FAQ",
		"You are at Home › Docs › FAQ",
		"What is uix?", "How do I add components?", "How do I build a theme?",
		"uix is a component registry for MyGo apps.",
		"More help options",
		"Email support at help@uix.dev, or open an issue on GitHub.",
		"Help",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestBreadcrumbs moves the trail with a click on a link, which the path
// line shows at once.
func TestBreadcrumbs(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if view.State.Chosen != 1 {
		t.Fatalf("chosen %d, want the Docs item", view.State.Chosen)
	}
	if !tt.HasText("You are at Home › Docs") || tt.HasText("You are at Home › Docs › FAQ") {
		t.Fatal("the path line did not follow the trail")
	}
}

// TestFAQ opens the answer to a question when its header is clicked.
func TestFAQ(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("How do I add components?"); err != nil {
		t.Fatal(err)
	}
	if !view.State.FAQ2 {
		t.Fatal("the click did not open the section")
	}
	if !tt.HasText("Run uix add button in your Go module") {
		t.Fatal("the answer did not show")
	}
}

// TestCollapsible hides the more options when its label is clicked.
func TestCollapsible(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("More help options"); err != nil {
		t.Fatal(err)
	}
	if view.State.More || tt.HasText("Email support at help@uix.dev") {
		t.Fatal("a click did not close the options")
	}
}

// TestPopover opens the help popover with the button, and jumps to the
// Docs section with a click in it, which closes the popover.
func TestPopover(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Help"); err != nil {
		t.Fatal(err)
	}
	if !view.State.HelpOpen || !tt.HasText("Go to Docs") {
		t.Fatal("the help button did not open the popover")
	}
	if err := tt.Click("Go to Docs"); err != nil {
		t.Fatal(err)
	}
	if view.State.HelpOpen || view.State.Chosen != 1 {
		t.Fatalf("the link did not jump and close: open %v, chosen %d", view.State.HelpOpen, view.State.Chosen)
	}
	if tt.HasText("Go to Docs") {
		t.Fatal("the popover still shows after the jump")
	}
}

// TestDarkMode draws the page under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Help Center") {
		t.Fatal("dark page missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
