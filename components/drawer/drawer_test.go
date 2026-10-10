package drawer

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// view draws a short drawer in a window, with the app's open flag.
func view(open *bool) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Drawer(c, Props{
				Open:   open,
				Title:  "Confirm",
				Height: 200,
				Children: func() {
					ui.Text(c, "Delete the file?")
				},
			})
		})
	}
}

// TestShows draws the drawer with its title and its child.
func TestShows(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 400, 320)
	if !tt.HasText("Confirm") || !tt.HasText("Delete the file?") {
		t.Fatalf("missing the drawer in %q", tt.Texts())
	}
}

// TestBackdropCloses closes for a click on the dimmed window, above the
// drawer.
func TestBackdropCloses(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 400, 320)
	tt.ClickAt(200, 20)
	if open {
		t.Fatal("a backdrop click did not close the drawer")
	}
}

// TestEscapeCloses closes for Escape.
func TestEscapeCloses(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 400, 320)
	tt.Key(0, ui.KeyEscape)
	if open {
		t.Fatal("Escape did not close the drawer")
	}
}

// TestPlacesBottom draws the drawer along the bottom of the window: a
// point on it differs from one over the dimmed window above.
func TestPlacesBottom(t *testing.T) {
	open := true
	tt := ui.NewTester(view(&open), 400, 320)
	img := tt.Image()
	a := img.RGBAAt(200, 300) // on the drawer
	b := img.RGBAAt(200, 20)  // over the dimmed window
	if a == b {
		t.Fatalf("the drawer does not cover the bottom: %v both", a)
	}
}
