package buttongroup

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

// view draws the group in a padded column, with the app's choice.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			ButtonGroup(c, p)
		})
	}
}

// options are the three of the tests, as alignment of a paragraph.
var options = []Option{
	{Label: "Left", Value: "left"},
	{Label: "Center", Value: "center"},
	{Label: "Right", Value: "right"},
}

// TestSingle chooses one button alone: a click on another moves the
// choice, which stays on the one chosen.
func TestSingle(t *testing.T) {
	selected := []string{"left"}
	tt := ui.NewTester(view(Props{Options: options, Selected: &selected}), 300, 80)
	if err := tt.Click("Right"); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(selected, []string{"right"}) {
		t.Fatalf("Selected = %v, want [right]", selected)
	}
	// The chosen one stays chosen: clicking it again does not clear it.
	tt.Click("Right")
	if !slices.Equal(selected, []string{"right"}) {
		t.Fatalf("Selected = %v, want [right] still", selected)
	}
}

// TestOnChange runs for a click that moves the choice.
func TestOnChange(t *testing.T) {
	selected := []string{"left"}
	runs := 0
	tt := ui.NewTester(view(Props{Options: options, Selected: &selected, OnChange: func() { runs++ }}), 300, 80)
	tt.Click("Center")
	if runs != 1 {
		t.Fatalf("OnChange ran %d times, want 1", runs)
	}
}

// TestMultiple toggles each button on and off as its own.
func TestMultiple(t *testing.T) {
	selected := []string{}
	tt := ui.NewTester(view(Props{Options: options, Selected: &selected, Multiple: true}), 300, 80)
	tt.Click("Left")
	tt.Click("Right")
	if !slices.Equal(selected, []string{"left", "right"}) {
		t.Fatalf("Selected = %v, want [left right]", selected)
	}
	tt.Click("Left")
	if !slices.Equal(selected, []string{"right"}) {
		t.Fatalf("Selected = %v, want [right] after taking out Left", selected)
	}
}
