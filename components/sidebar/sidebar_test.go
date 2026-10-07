package sidebar

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
			Sidebar(c, p)
			ui.Spacer(c)
		})
	}
}

// calm turns the animation of the tester off, so the items of a section
// are gone the frame its title is clicked, as the tests assert.
func calm(tt *ui.Tester) {
	tt.SetPreferences(ui.Preferences{ReduceMotion: true, TextScale: 1})
}

// TestRenders draws the sections and their items.
func TestRenders(t *testing.T) {
	selected, open := "notes", true
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Sections: []Section{
			{Title: "Files", Open: &open, Items: []Item{{"notes", "Notes"}, {"todo", "Todo"}}},
		},
	}), 400, 200)
	calm(tt)
	for _, s := range []string{"Files", "Notes", "Todo"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestChooses selects the item clicked.
func TestChooses(t *testing.T) {
	selected := "notes"
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Sections: []Section{
			{Title: "Files", Items: []Item{{"notes", "Notes"}, {"todo", "Todo"}}},
		},
	}), 400, 200)
	calm(tt)
	if err := tt.Click("Todo"); err != nil {
		t.Fatal(err)
	}
	if selected != "todo" {
		t.Fatalf("selected %q, want todo", selected)
	}
}

// TestSection hides the items of a section when its title is clicked.
func TestSection(t *testing.T) {
	selected, open := "notes", true
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Sections: []Section{
			{Title: "Files", Open: &open, Items: []Item{{"notes", "Notes"}, {"todo", "Todo"}}},
		},
	}), 400, 200)
	calm(tt)
	if err := tt.Click("Files"); err != nil {
		t.Fatal(err)
	}
	if open || tt.HasText("Todo") {
		t.Fatal("a click on the title left the section open")
	}
}

// TestDisabled keeps the items from being chosen.
func TestDisabled(t *testing.T) {
	selected := "notes"
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Sections: []Section{
			{Title: "Files", Items: []Item{{"notes", "Notes"}, {"todo", "Todo"}}},
		},
		Disabled: true,
	}), 400, 200)
	calm(tt)
	if err := tt.Click("Todo"); err != nil {
		t.Fatal(err)
	}
	if selected != "notes" {
		t.Fatalf("a disabled sidebar chose %q", selected)
	}
}

// TestDarkMode draws the sidebar under the dark appearance too.
func TestDarkMode(t *testing.T) {
	selected := "notes"
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Sections: []Section{
			{Title: "Files", Items: []Item{{"notes", "Notes"}}},
		},
	}), 400, 200)
	calm(tt)
	tt.SetDark(true)
	if !tt.HasText("Notes") {
		t.Fatal("sidebar missing under dark mode")
	}
}
