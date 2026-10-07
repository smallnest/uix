package form

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws a form of two fields inside a padded column.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Form(c, p)
		})
	}
}

// fields builds a form with two text fields, the second wider than the
// first, so their labels must line up at the widest.
func fields() func(c *ui.Context) {
	return func(c *ui.Context) {
		var name, email string
		ui.Field(c, "Name", func() { ui.TextInput(c, &name) })
		ui.Field(c, "Email address", func() { ui.TextInput(c, &email) })
	}
}

// TestRenders shows the labels of the fields and the controls beside them.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: fields()}), 320, 160)
	for _, s := range []string{"Name", "Email address"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestAlignsLinesUp lines the labels up at the right: the labels end at
// one column, so their controls start together.
func TestAlignsLinesUp(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: fields()}), 320, 160)
	short, ok := tt.Find("Name")
	if !ok {
		t.Fatal("no Name label")
	}
	long, ok := tt.Find("Email address")
	if !ok {
		t.Fatal("no Email address label")
	}
	// Both labels right-align to the widest, so the right edges match.
	if d := (short.X + short.W) - (long.X + long.W); d > 0.5 || d < -0.5 {
		t.Fatalf("labels do not line up: %v ends at %g, %v at %g", short, short.X+short.W, long, long.X+long.W)
	}
}

// TestClickFocuses focuses the control of a field when its label is
// clicked, which the label names for the field.
func TestClickFocuses(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: fields()}), 320, 160)
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	if !tt.Focused("Name") {
		t.Fatal("the label did not focus its field")
	}
}
