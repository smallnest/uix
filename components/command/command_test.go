package command

import (
	"fmt"
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

// TestEnterRuns runs the command the query leaves first for Enter in the
// field, which keeps the focus while the palette is open.
func TestEnterRuns(t *testing.T) {
	ran := ""
	open := true
	items := []Item{
		{Label: "New file", OnSelect: func() { ran = "New file" }},
		{Label: "Open recent", OnSelect: func() { ran = "Open recent" }},
	}
	tt := ui.NewTester(palette(&open, &State{}, items), 500, 320)
	tt.Type("recent")
	tt.Key(0, ui.KeyEnter)
	if ran != "Open recent" {
		t.Fatalf("Enter ran %q, want Open recent", ran)
	}
	if open {
		t.Fatal("the palette stayed open")
	}
}

// TestArrowsChoose moves the choice with the arrows from the field: Down
// to the next command, Up back, neither past an end of the list.
func TestArrowsChoose(t *testing.T) {
	for _, c := range []struct {
		keys []ui.Key
		want string
	}{
		{nil, "New file"},
		{[]ui.Key{ui.KeyDown}, "Open recent"},
		{[]ui.Key{ui.KeyDown, ui.KeyDown, ui.KeyDown}, "Open recent"},
		{[]ui.Key{ui.KeyDown, ui.KeyUp}, "New file"},
		{[]ui.Key{ui.KeyUp}, "New file"},
	} {
		ran := ""
		open := true
		items := []Item{
			{Label: "New file", OnSelect: func() { ran = "New file" }},
			{Label: "Open recent", OnSelect: func() { ran = "Open recent" }},
		}
		tt := ui.NewTester(palette(&open, &State{}, items), 500, 320)
		for _, k := range c.keys {
			tt.Key(0, k)
		}
		tt.Key(0, ui.KeyEnter)
		if ran != c.want {
			t.Errorf("%v then Enter ran %q, want %q", c.keys, ran, c.want)
		}
	}
}

// TestArrowsScroll keeps the choice in view as the arrows take it past
// the commands the list shows.
func TestArrowsScroll(t *testing.T) {
	open := true
	var items []Item
	for i := range 30 {
		items = append(items, Item{Label: fmt.Sprintf("Command %d", i)})
	}
	tt := ui.NewTester(palette(&open, &State{}, items), 500, 320)
	if tt.HasText("Command 20") {
		t.Fatal("the list shows every command: nothing to scroll")
	}
	for range 20 {
		tt.Key(0, ui.KeyDown)
	}
	if !tt.HasText("Command 20") {
		t.Fatalf("the choice went out of view: %q", tt.Texts())
	}
}

// TestFocusReturns gives the focus back to what had it as the palette
// closes, for Escape and for a command run.
func TestFocusReturns(t *testing.T) {
	for _, key := range []ui.Key{ui.KeyEscape, ui.KeyEnter} {
		open := false
		name := ""
		tt := ui.NewTester(func(c *ui.Context) {
			ui.Column(c).Fill().Padding(16).Children(func() {
				ui.TextInput(c, &name).Label("Name")
				Command(c, Props{Open: &open, State: &State{}, Items: commands})
			})
		}, 500, 320)
		if err := tt.Click("Name"); err != nil {
			t.Fatal(err)
		}
		open = true
		tt.Frame()
		if tt.Focused("Name") {
			t.Fatal("the palette did not take the focus")
		}
		tt.Key(0, key)
		tt.Frame()
		if open {
			t.Fatalf("the palette stayed open after %v", key)
		}
		if !tt.Focused("Name") {
			t.Errorf("the focus did not return after %v", key)
		}
	}
}
