package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/studio/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.StudioView, view.Width, view.Height)
}

// TestRenders draws the studio headless: the title, the mark, the color
// well with its hex, the photo and the preview.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Studio", "Brand color", "#3b82f6", "Photo", "Preview",
		"Pick a brand color, or like the mark.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestPhotoFillsCard draws the photo filling its 160x107 card, not at
// its own small size in the corner: the sky, sun and hill of the picture
// cover most of the card.
func TestPhotoFillsCard(t *testing.T) {
	tt := newTester(t)
	img := tt.Image()
	b := img.Bounds()
	minX, minY, maxX, maxY := b.Max.X, b.Max.Y, -1, -1
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			r, g, bl, _ := img.At(x, y).RGBA()
			R, G, B := int(r>>8), int(g>>8), int(bl>>8)
			if (R == 135 && G == 206 && B >= 180 && B <= 240) || // the sky
				(R == 255 && G == 200 && B == 60) || // the sun
				(R == 60 && G >= 120 && G <= 220 && B == 40) { // the hill
				minX, minY = min(minX, x), min(minY, y)
				maxX, maxY = max(maxX, x), max(maxY, y)
			}
		}
	}
	if maxX < 0 {
		t.Fatal("no photo drew")
	}
	if w, h := maxX-minX+1, maxY-minY+1; w < 150 || h < 100 {
		t.Fatalf("photo drew at %dx%d, want it to fill the 160x107 card", w, h)
	}
}

// TestLikes toggles the like with a click, which the status names and
// the icon shows filled.
func TestLikes(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Like"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Liked {
		t.Fatal("the like did not toggle on")
	}
	if !tt.HasText("Mark liked.") {
		t.Fatalf("the status does not name the like: %q", tt.Texts())
	}
	tt.Click("Like")
	if view.State.Liked {
		t.Fatal("the like did not toggle off")
	}
	if !tt.HasText("Mark unliked.") {
		t.Fatalf("the status does not name the unlike: %q", tt.Texts())
	}
}

// TestPicksColor opens the color well and picks a swatch, which edits
// the accent and the status names it.
func TestPicksColor(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Accent"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Red") {
		t.Fatal("the picker did not open")
	}
	tt.Click("Red")
	if view.State.Accent != ui.Hex("#ef4444") {
		t.Fatalf("the picker chose %v, want red", view.State.Accent)
	}
	if !tt.HasText("#ef4444") {
		t.Fatalf("the hex does not show the new color: %q", tt.Texts())
	}
	if !tt.HasText("Accent #ef4444.") {
		t.Fatalf("the status does not name the accent: %q", tt.Texts())
	}
}

// TestEscCloses closes the picker with Escape, leaving the accent.
func TestEscCloses(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Accent"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyEscape)
	if tt.HasText("Red") {
		t.Fatal("Escape left the picker open")
	}
	if view.State.Accent != ui.Hex("#3b82f6") {
		t.Fatalf("Escape changed the accent to %v", view.State.Accent)
	}
}

// TestDarkMode draws the studio under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Studio") {
		t.Fatal("dark studio missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
