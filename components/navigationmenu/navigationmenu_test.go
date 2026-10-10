package navigationmenu

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws the bar across the window, with the app's state.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			NavigationMenu(c, p)
		})
	}
}

// props returns the props of a bar with a fresh state for the app to
// keep.
func props() Props { return Props{State: &State{}} }

// TestRenders draws the labels of the links.
func TestRenders(t *testing.T) {
	p := props()
	p.Items = []Item{
		{Label: "Docs"},
		{Label: "Source"},
	}
	tt := ui.NewTester(frame(p), 400, 120)
	for _, want := range []string{"Docs", "Source"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestOpens shows the panel when the link with children is clicked, and
// hides it for a click outside.
func TestOpens(t *testing.T) {
	p := props()
	p.Items = []Item{
		{Label: "Docs", Children: []Item{
			{Label: "Getting started"},
			{Label: "Guides"},
		}},
	}
	tt := ui.NewTester(frame(p), 400, 160)
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Getting started") || !tt.HasText("Guides") {
		t.Fatalf("missing the panel in %q", tt.Texts())
	}
	tt.ClickAt(200, 140)
	if tt.HasText("Getting started") {
		t.Fatal("a click outside kept the panel open")
	}
}

// TestCloses toggles the panel for a click on the same link.
func TestCloses(t *testing.T) {
	p := props()
	p.Items = []Item{
		{Label: "Docs", Children: []Item{{Label: "Guides"}}},
	}
	tt := ui.NewTester(frame(p), 400, 160)
	tt.Click("Docs")
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("Guides") {
		t.Fatal("a second click kept the panel open")
	}
}

// TestChooses runs the OnClick of a plain link, and of a link of the
// panel.
func TestChooses(t *testing.T) {
	picked := ""
	p := props()
	p.Items = []Item{
		{Label: "Docs", Children: []Item{
			{Label: "Guides", OnClick: func() { picked = "guides" }},
		}},
		{Label: "Source", OnClick: func() { picked = "source" }},
	}
	tt := ui.NewTester(frame(p), 400, 160)
	if err := tt.Click("Source"); err != nil {
		t.Fatal(err)
	}
	if picked != "source" {
		t.Fatalf("picked %q, want source", picked)
	}
	tt.Click("Docs")
	if err := tt.Click("Guides"); err != nil {
		t.Fatal(err)
	}
	if picked != "guides" {
		t.Fatalf("picked %q, want guides", picked)
	}
}

// TestDarkMode draws the bar under the dark appearance too.
func TestDarkMode(t *testing.T) {
	p := props()
	p.Items = []Item{{Label: "Docs"}}
	tt := ui.NewTester(frame(p), 400, 120)
	tt.SetDark(true)
	if !tt.HasText("Docs") {
		t.Fatal("navigation menu missing under dark mode")
	}
}
