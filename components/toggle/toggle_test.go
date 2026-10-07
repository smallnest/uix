package toggle

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Toggle(c, p).Label("Bold")
		})
	}
}

// TestRenders draws the toggle with its label.
func TestRenders(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Bold"}), 200, 100)
	if !tt.HasText("Bold") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestToggles presses the toggle with a click, and lets it go with
// another.
func TestToggles(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Bold"}), 200, 100)
	if err := tt.Click("Bold"); err != nil {
		t.Fatal(err)
	}
	if !on {
		t.Fatal("the first click did not press the toggle")
	}
	if err := tt.Click("Bold"); err != nil {
		t.Fatal(err)
	}
	if on {
		t.Fatal("the second click did not let the toggle go")
	}
}

// TestKeyboard lets Space press the toggle once it has the focus.
func TestKeyboard(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Bold"}), 200, 100)
	if err := tt.Click("Bold"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeySpace)
	if on {
		t.Fatal("Space did not let the toggle go")
	}
}

// TestDisabled keeps the toggle from being pressed.
func TestDisabled(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Bold", Disabled: true}), 200, 100)
	if err := tt.Click("Bold"); err != nil {
		t.Fatal(err)
	}
	if on {
		t.Fatal("a disabled toggle was pressed")
	}
}

// TestDarkMode draws the toggle under the dark appearance too.
func TestDarkMode(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Bold"}), 200, 100)
	tt.SetDark(true)
	if !tt.HasText("Bold") {
		t.Fatal("toggle missing under dark mode")
	}
}
