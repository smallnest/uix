package colorwell

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			ColorWell(c, p)
		})
	}
}

// TestRenders shows the well, found by its label.
func TestRenders(t *testing.T) {
	color := ui.Hex("#3b82f6")
	tt := ui.NewTester(frame(Props{Color: &color, Label: "Tint"}), 240, 120)
	if !tt.HasText("Tint") {
		t.Fatalf("missing the well in %q", tt.Texts())
	}
}

// TestOpensPicker opens the picker on a click, and the swatches choose a
// new color, which Changed reports.
func TestOpensPicker(t *testing.T) {
	color := ui.Hex("#3b82f6")
	changes := 0
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			w := ColorWell(c, Props{Color: &color, Label: "Tint"})
			if w.Changed() {
				changes++
			}
		})
	}, 400, 500)
	if err := tt.Click("Tint"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Red") {
		t.Fatal("the picker did not open")
	}
	tt.Click("Red")
	if color != ui.Hex("#ef4444") || changes != 1 {
		t.Fatalf("the swatch chose %v, want red", color)
	}
}

// TestEscCloses closes the picker with Escape, leaving the color.
func TestEscCloses(t *testing.T) {
	color := ui.Hex("#3b82f6")
	tt := ui.NewTester(frame(Props{Color: &color, Label: "Tint"}), 400, 500)
	if err := tt.Click("Tint"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyEscape)
	if tt.HasText("Red") {
		t.Fatal("Escape left the picker open")
	}
	if color != ui.Hex("#3b82f6") {
		t.Fatalf("Escape changed the color to %v", color)
	}
}
