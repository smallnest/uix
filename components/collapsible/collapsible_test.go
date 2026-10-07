package collapsible

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Collapsible(c, p)
		})
	}
}

// calm turns the animation of the tester off, so the content is gone the
// frame a collapsible closes, as the tests assert.
func calm(tt *ui.Tester) {
	tt.SetPreferences(ui.Preferences{ReduceMotion: true, TextScale: 1})
}

// TestRenders draws the label, and no content while the collapsible is
// closed.
func TestRenders(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{
		Label:    "Advanced",
		Open:     &open,
		Children: func(c *ui.Context) { ui.Text(c, "Inside") },
	}), 300, 160)
	calm(tt)
	if !tt.HasText("Advanced") {
		t.Fatalf("texts %q", tt.Texts())
	}
	if tt.HasText("Inside") {
		t.Fatal("closed, it shows its content")
	}
}

// TestOpens shows the content when the label is clicked.
func TestOpens(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{
		Label:    "Advanced",
		Open:     &open,
		Children: func(c *ui.Context) { ui.Text(c, "Inside") },
	}), 300, 160)
	calm(tt)
	if err := tt.Click("Advanced"); err != nil {
		t.Fatal(err)
	}
	if !open || !tt.HasText("Inside") {
		t.Fatal("a click did not open it")
	}
}

// TestKeyboard opens it with a click, and closes it with Space once the
// label has the focus.
func TestKeyboard(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{
		Label:    "Advanced",
		Open:     &open,
		Children: func(c *ui.Context) { ui.Text(c, "Inside") },
	}), 300, 160)
	calm(tt)
	if err := tt.Click("Advanced"); err != nil {
		t.Fatal(err)
	}
	if !open || !tt.HasText("Inside") {
		t.Fatal("the click did not open it")
	}
	tt.Key(0, ui.KeySpace)
	if open || tt.HasText("Inside") {
		t.Fatal("Space did not close it")
	}
}

// TestDisabled keeps it from opening and closing.
func TestDisabled(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{
		Label:    "Advanced",
		Open:     &open,
		Children: func(c *ui.Context) { ui.Text(c, "Inside") },
		Disabled: true,
	}), 300, 160)
	calm(tt)
	if err := tt.Click("Advanced"); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Fatal("a disabled collapsible opened")
	}
}

// TestDarkMode draws it under the dark appearance too.
func TestDarkMode(t *testing.T) {
	open := true
	tt := ui.NewTester(frame(Props{
		Label:    "Advanced",
		Open:     &open,
		Children: func(c *ui.Context) { ui.Text(c, "Inside") },
	}), 300, 160)
	calm(tt)
	tt.SetDark(true)
	if !tt.HasText("Advanced") || !tt.HasText("Inside") {
		t.Fatal("collapsible missing under dark mode")
	}
}
