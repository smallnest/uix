package marker

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// view draws the marker of props in a window.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Marker(c, p)
		})
	}
}

// TestShows draws the label.
func TestShows(t *testing.T) {
	tt := ui.NewTester(view(Props{Label: "AI"}), 300, 100)
	if !tt.HasText("AI") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestVariants draw the label in every look, and the Separator look
// draws its lines across the window.
func TestVariants(t *testing.T) {
	for _, v := range []Variant{Default, Border, Separator} {
		tt := ui.NewTester(view(Props{Label: "AI", Variant: v}), 300, 100)
		if !tt.HasText("AI") {
			t.Fatalf("variant %d: missing the label in %q", v, tt.Texts())
		}
	}
}

// TestDarkMode draws the marker under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(view(Props{Label: "AI", Variant: Separator}), 300, 100)
	tt.SetDark(true)
	if !tt.HasText("AI") {
		t.Fatal("marker missing under dark mode")
	}
}
