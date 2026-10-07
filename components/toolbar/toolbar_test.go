package toolbar

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/menu"
)

// frame draws the toolbar of p at the top of a window.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Children(func() {
			Toolbar(c, p)
		})
	}
}

// TestRenders draws the controls of the toolbar.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: func(c *ui.Context) {
		ui.Button(c, "New")
	}}), 300, 80)
	if !tt.HasText("New") {
		t.Fatalf("texts %q, want the New button", tt.Texts())
	}
}

// TestClicks runs the action of a button of the toolbar.
func TestClicks(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(frame(Props{Children: func(c *ui.Context) {
		if ui.Button(c, "New").Clicked() {
			clicks++
		}
	}}), 300, 80)
	if err := tt.Click("New"); err != nil {
		t.Fatal(err)
	}
	if clicks != 1 {
		t.Fatalf("clicks %d, want 1", clicks)
	}
}

// TestMenu opens the menu button of the toolbar.
func TestMenu(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: func(c *ui.Context) {
		menu.Menu(c, menu.Props{Label: "File", Items: []menu.Item{
			{Label: "New"},
			{Label: "Open"},
		}})
	}}), 400, 120)
	r, ok := tt.Find("File")
	if !ok {
		t.Fatal("no File button")
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
	if got := tt.Menu(); !slices.Equal(got, []string{"New", "Open"}) {
		t.Fatalf("menu %q, want the items", got)
	}
}

// TestDisabled keeps the controls of the toolbar from acting.
func TestDisabled(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(frame(Props{Disabled: true, Children: func(c *ui.Context) {
		if ui.Button(c, "New").Clicked() {
			clicks++
		}
	}}), 300, 80)
	if err := tt.Click("New"); err != nil {
		t.Fatal(err)
	}
	if clicks != 0 {
		t.Fatalf("clicks %d, want the disabled toolbar to ignore it", clicks)
	}
}

// TestDarkMode draws the toolbar under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: func(c *ui.Context) {
		ui.Button(c, "New")
	}}), 300, 80)
	tt.SetDark(true)
	if !tt.HasText("New") {
		t.Fatal("toolbar missing under dark mode")
	}
}
