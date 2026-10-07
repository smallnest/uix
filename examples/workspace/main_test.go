package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/workspace/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.WorkspaceView, view.Width, view.Height)
}

// TestRenders draws the workspace headless: the title, the links, the
// form fields, the days and the board.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Workspace", "Open the guide", "Report an issue",
		"Project", "Project name", "Owner", "Sprint", "Sprint 1",
		"Days", "Day 1", "Board", "r0c0",
		"Pick a day of the sprint.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTypes edits the project name in the form, which the field edits in
// place.
func TestTypes(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Project name"); err != nil {
		t.Fatal(err)
	}
	tt.Type("my-project")
	if view.State.Name != "my-project" {
		t.Fatalf("the form did not take the name: %q", view.State.Name)
	}
}

// TestLinkOpens opens the guide in the browser when the link is clicked,
// which the tester records.
func TestLinkOpens(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Open the guide"); err != nil {
		t.Fatal(err)
	}
	urls := tt.OpenedURLs()
	if len(urls) != 1 || urls[0] != "https://example.com/guide" {
		t.Fatalf("the link opened %q, want the guide", urls)
	}
}

// TestScrollsDays moves the day chips sideways as the pointer wheels.
func TestScrollsDays(t *testing.T) {
	tt := newTester(t)
	r, _ := tt.Find("Day 2")
	x0 := r.X
	tt.Scroll(r.X, r.Y, 30, 0)
	r, _ = tt.Find("Day 2")
	if r.W == 0 || r.X >= x0 {
		t.Fatalf("the days did not scroll left: x %g from %g", r.X, x0)
	}
}

// TestChoosesDay picks a day with a click, which the status names. Day 2
// sits in the first view of the days, so the click reaches it.
func TestChoosesDay(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Day 2"); err != nil {
		t.Fatal(err)
	}
	if view.State.Day != 2 {
		t.Fatalf("the click chose day %d, want 2", view.State.Day)
	}
	if !tt.HasText("Day 2 of the sprint.") {
		t.Fatalf("the status does not name the day: %q", tt.Texts())
	}
}

// TestChoosesCell picks a cell of the board with a click, which the
// status names. r0c0 sits in the corner of the board, so the click
// reaches it.
func TestChoosesCell(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("r0c0"); err != nil {
		t.Fatal(err)
	}
	if view.State.Cell != "r0c0" {
		t.Fatalf("the click chose cell %q, want r0c0", view.State.Cell)
	}
	if !tt.HasText("Cell r0c0 chosen.") {
		t.Fatalf("the status does not name the cell: %q", tt.Texts())
	}
}

// TestDarkMode draws the workspace under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Workspace") {
		t.Fatal("dark workspace missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
