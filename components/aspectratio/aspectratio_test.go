package aspectratio

import (
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"
)

// red is the color the tests fill the ratio box with.
var red = color.RGBA{R: 239, G: 68, B: 68, A: 255}

// TestKeepsRatio draws a red box at a 2:1 ratio in a column: the box is
// as wide as the column and half as tall, so a point below it is the
// background again.
func TestKeepsRatio(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			AspectRatio(c, Props{
				Ratio: 2,
				Children: func() {
					ui.Box(c).Fill().Background(ui.Hex("#ef4444"))
				},
			})
		})
	}, 240, 200)
	img := tt.Image()
	if img.RGBAAt(100, 100) != red {
		t.Fatalf("the box did not fill the ratio: %v at 100,100", img.RGBAAt(100, 100))
	}
	// The column is 208 wide, so the box is 104 tall: below it the
	// background shows.
	if img.RGBAAt(100, 150) == red {
		t.Fatalf("the box is taller than its ratio: %v at 100,150", img.RGBAAt(100, 150))
	}
}

// TestDefaultRatio draws a square when the ratio is zero.
func TestDefaultRatio(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			AspectRatio(c, Props{
				Children: func() { ui.Box(c).Fill().Background(ui.Hex("#ef4444")) },
			})
		})
	}, 240, 200)
	img := tt.Image()
	// The column is 208 wide, so the square is 208 tall and out of the
	// window: a point far down the middle is still the box.
	if img.RGBAAt(100, 150) != red {
		t.Fatalf("a zero ratio did not draw a square: %v at 100,150", img.RGBAAt(100, 150))
	}
}
