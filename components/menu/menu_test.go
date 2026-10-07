package menu

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Menu(c, p)
		})
	}
}

// TestRenders draws the button of the menu.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "File", Items: []Item{
		{Label: "New"},
		{Separator: true},
		{Label: "Open"},
	}}), 300, 100)
	if !tt.HasText("File") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// open opens the menu of the File button, as the pointer goes down on
// it.
func open(t *testing.T, tt *ui.Tester) {
	t.Helper()
	r, ok := tt.Find("File")
	if !ok {
		t.Fatalf("no File button in %q", tt.Texts())
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
}

// TestOpens shows the items when the button goes down.
func TestOpens(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "File", Items: []Item{
		{Label: "New"},
		{Label: "Open"},
	}}), 300, 160)
	open(t, tt)
	if got := tt.Menu(); !slices.Equal(got, []string{"New", "Open"}) {
		t.Fatalf("menu %q, want the items", got)
	}
}

// TestChooses runs the action of an item when it is chosen.
func TestChooses(t *testing.T) {
	picked := ""
	tt := ui.NewTester(frame(Props{Label: "File", Items: []Item{
		{Label: "New", Action: func() { picked = "new" }},
		{Label: "Open", Action: func() { picked = "open" }},
	}}), 300, 160)
	open(t, tt)
	if err := tt.ChooseMenuItem("New"); err != nil {
		t.Fatal(err)
	}
	if picked != "new" {
		t.Fatalf("picked %q, want new", picked)
	}
}

// TestDisabled keeps the menu from opening.
func TestDisabled(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "File", Items: []Item{
		{Label: "New"},
	}, Disabled: true}), 300, 160)
	if err := tt.Click("File"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("New") {
		t.Fatal("a disabled menu opened")
	}
}

// TestDarkMode draws the button under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "File"}), 300, 100)
	tt.SetDark(true)
	if !tt.HasText("File") {
		t.Fatal("menu missing under dark mode")
	}
}
