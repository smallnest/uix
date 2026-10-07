package icon

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// square is a black square in currentColor, as icon sets draw them.
var square = ui.MustParseSVG([]byte(`<svg viewBox="0 0 10 10"><rect width="10" height="10"/></svg>`))

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Icon(c, p)
		})
	}
}

// TestRenders draws the icon at the height of the font, found by its
// label.
func TestRenders(t *testing.T) {
	var fontSize float32
	tt := ui.NewTester(func(c *ui.Context) {
		fontSize = c.Theme().FontSize
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Icon(c, Props{SVG: square, Label: "save"})
		})
	}, 200, 80)
	r, ok := tt.Find("save")
	if !ok {
		t.Fatal("no icon labeled save")
	}
	if r.H != fontSize {
		t.Fatalf("an icon is %v high, not the font size %v", r.H, fontSize)
	}
}

// TestTakesTextColor draws the icon in the color of the text, which a
// row can set on it.
func TestTakesTextColor(t *testing.T) {
	blue := ui.RGB(0, 0, 255)
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Row(c).Fill().Padding(16).TextColor(blue).Children(func() {
			Icon(c, Props{SVG: square, Label: "tint"})
		})
	}, 200, 80)
	img := tt.Image()
	r, _ := tt.Find("tint")
	c := img.RGBAAt(int(r.X+r.W/2), int(r.Y+r.H/2))
	if c.R != 0 || c.G != 0 || c.B != 255 {
		t.Fatalf("an icon drew %v, want blue", c)
	}
}

// TestSized fits the icon into a box of another size.
func TestSized(t *testing.T) {
	tt := ui.NewTester(frame(Props{SVG: square, Label: "big", Width: 40, Height: 20}), 200, 80)
	r, ok := tt.Find("big")
	if !ok {
		t.Fatal("no icon labeled big")
	}
	if r.W != 40 || r.H != 20 {
		t.Fatalf("a sized icon is %vx%v, want 40x20", r.W, r.H)
	}
}

// TestGrayscale draws the icon without color.
func TestGrayscale(t *testing.T) {
	red := ui.MustParseSVG([]byte(`<svg viewBox="0 0 10 10"><rect width="10" height="10" fill="#f00"/></svg>`))
	tt := ui.NewTester(frame(Props{SVG: red, Label: "gray", Grayscale: true}), 200, 80)
	img := tt.Image()
	r, _ := tt.Find("gray")
	c := img.RGBAAt(int(r.X+r.W/2), int(r.Y+r.H/2))
	if c.R != c.G || c.G != c.B {
		t.Fatalf("a grayscale icon drew %v, not gray", c)
	}
}
