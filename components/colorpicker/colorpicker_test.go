package colorpicker

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			ColorPicker(c, p).Label("Color")
		})
	}
}

// TestRenders draws the swatch of the color.
func TestRenders(t *testing.T) {
	color := ui.Hex("#7c3aed")
	tt := ui.NewTester(frame(Props{Value: &color}), 400, 500)
	if _, ok := tt.Find("Color"); !ok {
		t.Fatal("swatch not found")
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestOpens shows the picker when the swatch is clicked.
func TestOpens(t *testing.T) {
	color := ui.Hex("#7c3aed")
	tt := ui.NewTester(frame(Props{Value: &color}), 400, 500)
	if err := tt.Click("Color"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"Red", "Green", "Black"} {
		if !tt.HasText(s) {
			t.Fatalf("missing the %q swatch of the open picker", s)
		}
	}
}

// TestSwatches picks a color from the swatches of the open picker.
func TestSwatches(t *testing.T) {
	color := ui.Hex("#7c3aed")
	tt := ui.NewTester(frame(Props{Value: &color}), 400, 500)
	tt.Click("Color")
	if err := tt.Click("Green"); err != nil {
		t.Fatal(err)
	}
	if color != ui.Hex("#22c55e") {
		t.Fatalf("chose %v, want green", color)
	}
}

// TestDisabled opens no picker when the swatch is clicked.
func TestDisabled(t *testing.T) {
	color := ui.Hex("#7c3aed")
	tt := ui.NewTester(frame(Props{Value: &color, Disabled: true}), 400, 500)
	if err := tt.Click("Color"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("Red") {
		t.Fatal("a disabled swatch opens its picker")
	}
}

// TestDarkMode draws the swatch and its picker under the dark
// appearance too.
func TestDarkMode(t *testing.T) {
	color := ui.Hex("#7c3aed")
	tt := ui.NewTester(frame(Props{Value: &color}), 400, 500)
	tt.SetDark(true)
	tt.Click("Color")
	if !tt.HasText("Red") {
		t.Fatal("picker missing under dark mode")
	}
}
