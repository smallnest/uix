package command

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// palette draws the command palette open, with the app's state, the
// commands and a child to see.
func palette(open *bool, st *State, items []Item) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Command(c, Props{Open: open, State: st, Items: items})
		})
	}
}

// commands are the two of the tests.
var commands = []Item{
	{Label: "New file", Keywords: []string{"create", "blank"}},
	{Label: "Open recent"},
}

// TestShows draws the palette with the commands.
func TestShows(t *testing.T) {
	open := true
	tt := ui.NewTester(palette(&open, &State{}, commands), 500, 320)
	for _, want := range []string{"New file", "Open recent"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestFilters keeps only the commands that match the query, which the
// field takes the focus of as the palette opens.
func TestFilters(t *testing.T) {
	open := true
	tt := ui.NewTester(palette(&open, &State{}, commands), 500, 320)
	tt.Type("new")
	if !tt.HasText("New file") {
		t.Fatalf("the match is missing from %q", tt.Texts())
	}
	if tt.HasText("Open recent") {
		t.Fatalf("a non-match stayed in %q", tt.Texts())
	}
}

// TestKeywords match the query too: "blank" finds New file.
func TestKeywords(t *testing.T) {
	open := true
	tt := ui.NewTester(palette(&open, &State{}, commands), 500, 320)
	tt.Type("blank")
	if !tt.HasText("New file") {
		t.Fatalf("a keyword did not match: %q", tt.Texts())
	}
}

// TestRuns closes the palette and runs the item for a click on it.
func TestRuns(t *testing.T) {
	ran := false
	open := true
	items := []Item{{Label: "New file", OnSelect: func() { ran = true }}}
	tt := ui.NewTester(palette(&open, &State{}, items), 500, 320)
	if err := tt.Click("New file"); err != nil {
		t.Fatal(err)
	}
	if !ran {
		t.Fatal("the command did not run")
	}
	if open {
		t.Fatal("the palette stayed open")
	}
}

// TestEscapeCloses closes the palette for Escape.
func TestEscapeCloses(t *testing.T) {
	open := true
	tt := ui.NewTester(palette(&open, &State{}, commands), 500, 320)
	tt.Key(0, ui.KeyEscape)
	if open {
		t.Fatal("Escape did not close the palette")
	}
}
