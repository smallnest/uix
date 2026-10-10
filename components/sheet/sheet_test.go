package sheet

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// view draws a right sheet in a window, with the app's open flag and a
// child to see.
func view(open *bool) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Sheet(c, Props{
				Open:  open,
				Title: "Filters",
				Side:  Right,
				Children: func() {
					ui.Text(c, "Kind")
				},
			})
		})
	}
}

// TestShows draws the sheet with its title and its child.
func TestShows(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 500, 300)
	if !tt.HasText("Filters") || !tt.HasText("Kind") {
		t.Fatalf("missing the sheet in %q", tt.Texts())
	}
}

// TestBackdropCloses closes for a click on the dimmed window, left of
// the panel.
func TestBackdropCloses(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 500, 300)
	tt.ClickAt(50, 150)
	if open {
		t.Fatal("a backdrop click did not close the sheet")
	}
}

// TestEscapeCloses closes for Escape.
func TestEscapeCloses(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 500, 300)
	tt.Key(0, ui.KeyEscape)
	if open {
		t.Fatal("Escape did not close the sheet")
	}
}

// TestLeftPlaces draws a left sheet on the left of the window: a point
// on it differs from one over the dimmed window.
func TestLeftPlaces(t *testing.T) {
	open := true
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Sheet(c, Props{Open: &open, Title: "Filters", Side: Left})
		})
	}, 500, 300)
	img := tt.Image()
	a := img.RGBAAt(50, 150)  // on the panel
	b := img.RGBAAt(450, 150) // over the dimmed window
	if a == b {
		t.Fatalf("the left sheet does not cover the window: %v both", a)
	}
}
