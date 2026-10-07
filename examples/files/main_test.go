package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/files/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() {
	view.State = view.TableState{Selected: -1, List: &ui.ListState{}}
	view.State.List.Selected = &view.State.Selected
	view.State.List.Sort = &view.State.Sort
}

// TestRenders draws the file browser headless and checks that the
// headers and the rows came through: MyGo renders frames in memory, no
// window needed.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FilesView, view.Width, view.Height)
	for _, s := range []string{"Files", "Name", "Kind", "Size", "README.md", "go.mod", "No file selected"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestSorts clicks a header and checks both the sort state and the order
// of the rows the view draws, then reverses the order with a second
// click.
func TestSorts(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FilesView, view.Width, view.Height)
	// The initial order is the order of the files.
	if view.Sorted()[0].Name != "README.md" {
		t.Fatalf("first row %q, want README.md", view.Sorted()[0].Name)
	}
	if err := tt.Click("Size"); err != nil {
		t.Fatal(err)
	}
	if view.State.Sort.Column != "Size" || view.State.Sort.Descending {
		t.Fatalf("sort %+v, want Size ascending", view.State.Sort)
	}
	// Ascending by size puts the empty directory first.
	if view.Sorted()[0].Name != "internal/ui/components/table" {
		t.Fatalf("ascending: first row %q, want the directory", view.Sorted()[0].Name)
	}
	if err := tt.Click("Size"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Sort.Descending {
		t.Fatal("second click did not reverse the sort")
	}
	// Descending by size puts the largest file first.
	if view.Sorted()[0].Name != "go.sum" {
		t.Fatalf("descending: first row %q, want go.sum", view.Sorted()[0].Name)
	}
}

// TestSelects chooses a row with a click, then opens it with Enter.
func TestSelects(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FilesView, view.Width, view.Height)
	if err := tt.Click("go.mod"); err != nil {
		t.Fatal(err)
	}
	if view.State.Selected != 2 {
		t.Fatalf("chose %d, want 2 (go.mod)", view.State.Selected)
	}
	if !tt.HasText("Selected: go.mod") {
		t.Fatal("the detail line does not name the row")
	}
	tt.Key(0, ui.KeyEnter)
	if view.State.Opened != "go.mod" {
		t.Fatalf("opened %q, want go.mod", view.State.Opened)
	}
	if !tt.HasText("Opened go.mod") {
		t.Fatal("the open line is missing")
	}
}

// TestDarkMode draws the file browser under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FilesView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Files") {
		t.Fatal("dark file browser missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
