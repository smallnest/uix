package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/docreader/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.DocView, view.Width, view.Height)
}

// TestRenders draws the reader headless: the title, the segments, the
// find button, the headings and bodies of the document, and the status
// line.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"uix notes", "Reader", "Outline", "Find",
		"Welcome", "Find", "Rename", "Outline",
		"component registry for the MyGo native UI toolkit",
		"Double click the title to rename it, or press Find to search.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestFind opens the find bar, types a query, and steps through the
// matches, which the status counts and names.
func TestFind(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Find"); err != nil {
		t.Fatal(err)
	}
	if !view.State.FindOpen {
		t.Fatal("Find did not open the bar")
	}
	tt.Type("the")
	if view.State.Query != "the" {
		t.Fatalf("the field has %q, want the", view.State.Query)
	}
	if !tt.HasText("Match 1 of ") {
		t.Fatalf("the status does not count the matches: %q", tt.Texts())
	}
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if view.State.Current != 1 {
		t.Fatalf("Next set the current %d, want 1", view.State.Current)
	}
	if !tt.HasText("Match 2 of ") {
		t.Fatalf("the status does not follow: %q", tt.Texts())
	}
}

// TestSegmented switches to the outline and jumps back to a paragraph.
func TestSegmented(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Outline"); err != nil {
		t.Fatal(err)
	}
	if view.State.Pane != 1 {
		t.Fatal("the segment did not switch to the outline")
	}
	if !tt.HasText("Switched to the outline.") {
		t.Fatal("the status does not name the switch")
	}
	if err := tt.Click("Welcome"); err != nil {
		t.Fatal(err)
	}
	if view.State.Pane != 0 || view.State.FocusPara != 0 {
		t.Fatalf("the outline did not jump back: pane %d, focus %d", view.State.Pane, view.State.FocusPara)
	}
	if !tt.HasText("Jumped to “Welcome”.") {
		t.Fatal("the status does not name the jump")
	}
}

// TestTitle renames the title in place with a double click.
func TestTitle(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("uix notes")
	if !ok {
		t.Fatal("no title")
	}
	tt.ClickAt(r.X+5, r.Y+5)
	tt.ClickAt(r.X+5, r.Y+5)
	tt.Type("Readme")
	tt.Key(0, ui.KeyTab)
	if view.State.Title != "Readme" {
		t.Fatalf("the title is %q, want Readme", view.State.Title)
	}
	if !tt.HasText("Renamed to “Readme”.") {
		t.Fatal("the status does not name the rename")
	}
}

// TestDarkMode draws the reader under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("uix notes") {
		t.Fatal("dark reader missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
