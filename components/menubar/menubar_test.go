package menubar

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws the bar across the window.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Menubar(c, p)
		})
	}
}

// TestRenders draws the labels of the menus.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Menus: []Menu{
		{Label: "File", Items: []Item{{Label: "New"}}},
		{Label: "Edit", Items: []Item{{Label: "Undo"}}},
	}}), 400, 100)
	for _, want := range []string{"File", "Edit"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// open opens the menu of the button at label, as the pointer goes down
// on it.
func open(t *testing.T, tt *ui.Tester, label string) {
	t.Helper()
	r, ok := tt.Find(label)
	if !ok {
		t.Fatalf("no %s button in %q", label, tt.Texts())
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
}

// TestOpens shows the items when the button goes down, with "-" for
// the separator.
func TestOpens(t *testing.T) {
	tt := ui.NewTester(frame(Props{Menus: []Menu{
		{Label: "File", Items: []Item{
			{Label: "New"},
			{Separator: true},
			{Label: "Save"},
		}},
	}}), 400, 160)
	open(t, tt, "File")
	if got := tt.Menu(); !slices.Equal(got, []string{"New", "-", "Save"}) {
		t.Fatalf("menu %q, want the items", got)
	}
}

// TestChooses runs the action of an item when it is chosen.
func TestChooses(t *testing.T) {
	picked := ""
	tt := ui.NewTester(frame(Props{Menus: []Menu{
		{Label: "File", Items: []Item{
			{Label: "New", Action: func() { picked = "new" }},
			{Label: "Save", Action: func() { picked = "save" }},
		}},
	}}), 400, 160)
	open(t, tt, "File")
	if err := tt.ChooseMenuItem("Save"); err != nil {
		t.Fatal(err)
	}
	if picked != "save" {
		t.Fatalf("picked %q, want save", picked)
	}
}

// TestDarkMode draws the bar under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Menus: []Menu{
		{Label: "File"},
	}}), 400, 100)
	tt.SetDark(true)
	if !tt.HasText("File") {
		t.Fatal("menubar missing under dark mode")
	}
}
