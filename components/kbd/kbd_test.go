package kbd

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestShows draws the key with its label.
func TestShows(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Kbd(c, Props{Label: "⌘K"})
		})
	}, 200, 80)
	if !tt.HasText("⌘K") {
		t.Fatalf("missing the label in %q", tt.Texts())
	}
}

// TestChip draws the key on a chip: a point inside the chip is not the
// window background, sampled far from it.
func TestChip(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Kbd(c, Props{Label: "Esc"})
		})
	}, 200, 80)
	img := tt.Image()
	bg := img.RGBAAt(190, 70)
	if c := img.RGBAAt(24, 24); c == bg {
		t.Fatalf("no chip drew at 24,24: %v", c)
	}
}
