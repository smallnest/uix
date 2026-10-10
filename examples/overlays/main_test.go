package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/overlays/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.OverlaysView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading, the three buttons
// and the command line.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Overlays",
		"Open filters", "Delete the project", "Run a command  ⌘K",
		"The last command: ",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSheet opens the filters sheet and closes it with Escape.
func TestSheet(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Open filters"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Minimum price") || !tt.HasText("In stock only") {
		t.Fatalf("missing the sheet in %q", tt.Texts())
	}
	tt.Key(0, ui.KeyEscape)
	if tt.HasText("Minimum price") {
		t.Fatal("Escape did not close the sheet")
	}
}

// TestDrawer opens the drawer and deletes the project with its button.
func TestDrawer(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Delete the project"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("This removes the project and its files. This cannot be undone.") {
		t.Fatalf("missing the drawer in %q", tt.Texts())
	}
	if err := tt.Click("Delete"); err != nil {
		t.Fatal(err)
	}
	if view.State.Last != "project deleted" {
		t.Fatalf("Last = %q, want project deleted", view.State.Last)
	}
	if !tt.HasText("The last command: project deleted") {
		t.Fatalf("missing the last command in %q", tt.Texts())
	}
}

// TestCommand runs a command of the palette, for a click on it.
func TestCommand(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Run a command  ⌘K"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Save"); err != nil {
		t.Fatal(err)
	}
	if view.State.Last != "saved" {
		t.Fatalf("Last = %q, want saved", view.State.Last)
	}
}

// TestCommandKey opens the palette for ⌘K, and runs a command of it.
func TestCommandKey(t *testing.T) {
	tt := newTester(t)
	tt.Key(ui.Cmd, ui.KeyK)
	if !tt.HasText("New file") || !tt.HasText("Save") {
		t.Fatalf("missing the palette in %q", tt.Texts())
	}
	if err := tt.Click("New file"); err != nil {
		t.Fatal(err)
	}
	if view.State.Last != "new file" {
		t.Fatalf("Last = %q, want new file", view.State.Last)
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	for _, s := range []string{"Overlays", "Open filters"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in the dark frame %q", s, tt.Texts())
		}
	}
}
