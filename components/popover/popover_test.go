package popover

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Popover(c, p)
		})
	}
}

// help turns the props into one whose trigger is a Help button and whose
// panel holds two lines.
func help(p Props) func(c *ui.Context) {
	p.Trigger = func(c *ui.Context) *ui.Element { return ui.Button(c, "Help") }
	p.Content = func(c *ui.Context) {
		ui.Column(c).Gap(4).Children(func() {
			ui.Text(c, "View docs")
			ui.Text(c, "Report an issue")
		})
	}
	return frame(p)
}

// TestRenders draws the trigger, and no panel while the popover is
// closed.
func TestRenders(t *testing.T) {
	open := false
	tt := ui.NewTester(help(Props{Open: &open}), 300, 220)
	if !tt.HasText("Help") {
		t.Fatalf("texts %q", tt.Texts())
	}
	if tt.HasText("View docs") {
		t.Fatal("closed, it shows its panel")
	}
}

// TestOpens shows the panel when the trigger is clicked.
func TestOpens(t *testing.T) {
	open := false
	tt := ui.NewTester(help(Props{Open: &open}), 300, 220)
	if err := tt.Click("Help"); err != nil {
		t.Fatal(err)
	}
	if !open || !tt.HasText("View docs") {
		t.Fatal("a click did not open the panel")
	}
}

// TestCloses hides the panel when the trigger is clicked again, and when
// Escape is pressed.
func TestCloses(t *testing.T) {
	open := false
	tt := ui.NewTester(help(Props{Open: &open}), 300, 220)
	tt.Click("Help")
	if err := tt.Click("Help"); err != nil {
		t.Fatal(err)
	}
	if open || tt.HasText("View docs") {
		t.Fatal("a click did not close the panel")
	}
	tt.Click("Help")
	tt.Key(0, ui.KeyEscape)
	if open || tt.HasText("View docs") {
		t.Fatal("Escape did not close the panel")
	}
}

// TestDisabled keeps the trigger from opening the panel.
func TestDisabled(t *testing.T) {
	open := false
	tt := ui.NewTester(help(Props{Open: &open, Disabled: true}), 300, 220)
	if err := tt.Click("Help"); err != nil {
		t.Fatal(err)
	}
	if open || tt.HasText("View docs") {
		t.Fatal("a disabled popover opened")
	}
}

// TestDarkMode draws the trigger and the panel under the dark appearance
// too.
func TestDarkMode(t *testing.T) {
	open := false
	tt := ui.NewTester(help(Props{Open: &open}), 300, 220)
	tt.SetDark(true)
	tt.Click("Help")
	if !tt.HasText("View docs") {
		t.Fatal("panel missing under dark mode")
	}
}
