package image

import (
	"image"
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"
)

// redBitmap is a 10×10 red bitmap, its top-left corner (2, 2) green so
// the tests can see the fit.
func redBitmap() *ui.Bitmap {
	src := image.NewRGBA(image.Rect(0, 0, 10, 10))
	for i := 0; i < len(src.Pix); i += 4 {
		src.Pix[i], src.Pix[i+3] = 255, 255 // red
	}
	src.SetRGBA(2, 2, color.RGBA{R: 0, G: 255, B: 0, A: 255}) // a green corner
	return ui.NewBitmap(src)
}

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Image(c, p)
		})
	}
}

// TestAtOwnSize shows the picture at the size of its source, in DIPs.
func TestAtOwnSize(t *testing.T) {
	tt := ui.NewTester(frame(Props{Src: redBitmap()}), 200, 80)
	// A picture has no text; find its box by the pixel it draws. The
	// frame has one picture, 10 DIPs at the top-left of the column.
	img := tt.Image()
	if c := img.RGBAAt(20, 20); c.R != 255 || c.G != 0 {
		t.Fatalf("the picture did not draw at its own size: %v at 20,20", c)
	}
}

// TestFitsInto sizes the picture into a box, keeping its aspect: a red
// bitmap in a wide box still draws red at its center.
func TestFitsInto(t *testing.T) {
	tt := ui.NewTester(frame(Props{Src: redBitmap(), Width: 100, Height: 50}), 200, 100)
	img := tt.Image()
	if c := img.RGBAAt(50, 41); c.R != 255 || c.G != 0 {
		t.Fatalf("the picture did not fit: %v at the center", c)
	}
}

// TestGrayscale draws the picture in shades of gray.
func TestGrayscale(t *testing.T) {
	tt := ui.NewTester(frame(Props{Src: redBitmap(), Width: 50, Height: 50, Grayscale: true}), 200, 100)
	img := tt.Image()
	c := img.RGBAAt(25, 33)
	if c.R != c.G || c.G != c.B {
		t.Fatalf("a grayscale picture drew %v, not gray", c)
	}
}

// TestSVGShows draws an SVG in its own colors.
func TestSVGShows(t *testing.T) {
	svg := ui.MustParseSVG([]byte(`<svg viewBox="0 0 10 10"><rect width="10" height="10" fill="#00f"/></svg>`))
	tt := ui.NewTester(frame(Props{Src: svg, Width: 50, Height: 50}), 200, 100)
	img := tt.Image()
	if c := img.RGBAAt(25, 33); c.B != 255 || c.R != 0 {
		t.Fatalf("the SVG did not draw: %v", c)
	}
}
