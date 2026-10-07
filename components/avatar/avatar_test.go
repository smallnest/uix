package avatar

import (
	"image"
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"
)

func initialsFrame(name string) func(c *ui.Context) {
	return func(c *ui.Context) {
		Avatar(c, Props{Name: name})
	}
}

func imageFrame() func(c *ui.Context) {
	img := image.NewRGBA(image.Rect(0, 0, 40, 40))
	for y := 0; y < 40; y++ {
		for x := 0; x < 40; x++ {
			img.Set(x, y, color.RGBA{200, 30, 60, 255})
		}
	}
	bm := ui.NewBitmap(img)
	return func(c *ui.Context) {
		Avatar(c, Props{Name: "Ada Lovelace", Image: bm})
	}
}

// TestInitials shows the first letters of the first and last words.
func TestInitials(t *testing.T) {
	tt := ui.NewTester(initialsFrame("Ada Lovelace"), 120, 80)
	if !tt.HasText("AL") {
		t.Fatal(`want initials "AL" for "Ada Lovelace"`)
	}
}

// TestSingleWord shows one letter for a single-word name.
func TestSingleWord(t *testing.T) {
	tt := ui.NewTester(initialsFrame("Grace"), 120, 80)
	if !tt.HasText("G") {
		t.Fatal(`want "G" for "Grace"`)
	}
}

// TestEmptyName draws an empty face without crashing.
func TestEmptyName(t *testing.T) {
	tt := ui.NewTester(initialsFrame(""), 120, 80)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestImage draws the photo instead of the initials.
func TestImage(t *testing.T) {
	tt := ui.NewTester(imageFrame(), 120, 80)
	if tt.HasText("AL") {
		t.Fatal("a photo must hide the initials")
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestDarkMode draws the avatar under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(initialsFrame("Ada Lovelace"), 120, 80)
	tt.SetDark(true)
	if !tt.HasText("AL") {
		t.Fatal(`want initials "AL" under dark mode`)
	}
}
