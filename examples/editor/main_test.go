package main

import (
	"slices"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/editor/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	tt := ui.NewTester(view.EditorView, view.Width, view.Height)
	tt.SetPreferences(ui.Preferences{ReduceMotion: true, TextScale: 1})
	return tt
}

// TestRenders draws the editor headless: the sidebar, the toolbar and the
// note of the file chosen.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Editor", "Files", "Notes", "Todo", "Journal", "File",
		"B", "I", "U", "S", "Save",
		"Editing notes.md", "Buy milk", "Style: bold",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSidebar opens the note of another file when its item is clicked.
func TestSidebar(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Todo"); err != nil {
		t.Fatal(err)
	}
	if view.State.Place != "todo" {
		t.Fatalf("place %q, want todo", view.State.Place)
	}
	if !tt.HasText("Editing todo.md") || !tt.HasText("- finish uix batch 5") {
		t.Fatal("the note did not follow the choice")
	}
	if tt.HasText("Buy milk") {
		t.Fatal("the old note still shows")
	}
}

// TestToggles restyles the note when a toggle is clicked.
func TestToggles(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("B"); err != nil {
		t.Fatal(err)
	}
	if view.State.Bold || !tt.HasText("Style: none") {
		t.Fatal("Bold did not go off")
	}
	if err := tt.Click("I"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Italic || !tt.HasText("Style: italic") {
		t.Fatal("Italic did not go on")
	}
}

// TestMenu opens the file menu and chooses an item, which opens its file.
func TestMenu(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("File")
	if !ok {
		t.Fatal("no File button")
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
	if got := tt.Menu(); !slices.Equal(got, []string{"New note", "Open todo.md", "-", "Save"}) {
		t.Fatalf("menu %q, want the file actions", got)
	}
	if err := tt.ChooseMenuItem("Open todo.md"); err != nil {
		t.Fatal(err)
	}
	if view.State.Place != "todo" || !tt.HasText("Opened todo.md.") {
		t.Fatalf("the menu did not open the file: place %q", view.State.Place)
	}
}

// TestTooltip shows the tip of a control as the pointer rests on it.
func TestTooltip(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("Save")
	if !ok {
		t.Fatal("no Save button")
	}
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	time.Sleep(650 * time.Millisecond)
	tt.Frame()
	if !tt.HasText("Save the note") {
		t.Fatal("the tip did not show")
	}
}

// TestDarkMode draws the editor under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Editor") {
		t.Fatal("dark editor missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
