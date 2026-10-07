package main

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/explorer/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.ExplorerView, view.Width, view.Height)
}

// TestRenders draws the explorer headless: the toolbar, the tree, the
// list of the folder chosen and the status line.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"New", "View", "Open",
		"src", "lib", "docs", "go.mod", "Makefile",
		"main.go", "app.go", "1204 B",
		"Showing src.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTree changes the folder when a branch of the tree is clicked, and
// keeps a branch open that the Right key opened.
func TestTree(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("docs"); err != nil {
		t.Fatal(err)
	}
	if view.State.Folder != "docs" || !tt.HasText("guide.md") {
		t.Fatalf("the tree did not choose the folder %q", view.State.Folder)
	}
	// lib is closed; the Right key opens it. Then choosing src shows
	// the files of src, and lib's leaf shows only because the branch
	// stayed open.
	if err := tt.Click("lib"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyRight)
	if err := tt.Click("src"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("util.go") {
		t.Fatal("the Right key did not open lib")
	}
}

// TestList chooses a row of the list, which names it in the status.
func TestList(t *testing.T) {
	tt := newTester(t)
	// lib is closed, so its only file appears in the list alone, and a
	// click on it chooses the row.
	if err := tt.Click("lib"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("util.go"); err != nil {
		t.Fatal(err)
	}
	if view.State.Chosen != 0 || !tt.HasText("Chosen util.go.") {
		t.Fatalf("the list did not choose its row: %d", view.State.Chosen)
	}
}

// TestToolbar creates a file with New and hides the sizes with the View
// menu.
func TestToolbar(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("New"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("new1.txt") || !tt.HasText("Created new1.txt.") {
		t.Fatalf("New did not add a file: %q", tt.Texts())
	}
	// The menu of the View button opens on a press.
	r, ok := tt.Find("View")
	if !ok {
		t.Fatal("no View button")
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
	if got := tt.Menu(); !slices.Equal(got, []string{"Show sizes"}) {
		t.Fatalf("menu %q, want the view actions", got)
	}
	if err := tt.ChooseMenuItem("Show sizes"); err != nil {
		t.Fatal(err)
	}
	if view.State.ShowSizes || tt.HasText("1204 B") {
		t.Fatal("Show sizes did not hide the sizes")
	}
}

// TestSplit drags the divider, which widens the tree pane.
func TestSplit(t *testing.T) {
	tt := newTester(t)
	before := view.State.Divider
	tt.Press(before, 300)
	tt.Frame()
	tt.Move(before+60, 300)
	tt.Frame()
	tt.Release(before+60, 300)
	tt.Frame()
	if view.State.Divider <= before {
		t.Fatalf("divider %v, want the drag to move it", view.State.Divider)
	}
}

// TestDarkMode draws the explorer under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("src") {
		t.Fatal("dark explorer missing the tree")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
