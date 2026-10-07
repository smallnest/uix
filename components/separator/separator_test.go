package separator

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(o Orientation, w, h int) *ui.Tester {
	return ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Separator(c, Props{Orientation: o})
		})
	}, w, h)
}

// TestHorizontal draws a 1-point line across the width of the parent,
// inside its padding.
func TestHorizontal(t *testing.T) {
	tt := frame(Horizontal, 200, 80)
	r, ok := tt.Find("separator")
	if !ok {
		t.Fatal("separator not found")
	}
	if r.W != 168 {
		t.Fatalf("horizontal separator width = %v, want 168 (200 minus 32 of padding)", r.W)
	}
	if r.H != 1 {
		t.Fatalf("horizontal separator height = %v, want 1", r.H)
	}
}

// TestVertical draws a 1-point line down the height of the parent,
// inside its padding.
func TestVertical(t *testing.T) {
	tt := frame(Vertical, 200, 80)
	r, ok := tt.Find("separator")
	if !ok {
		t.Fatal("separator not found")
	}
	if r.W != 1 {
		t.Fatalf("vertical separator width = %v, want 1", r.W)
	}
	if r.H != 48 {
		t.Fatalf("vertical separator height = %v, want 48 (80 minus 32 of padding)", r.H)
	}
}

// TestDarkMode draws the separator under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := frame(Horizontal, 200, 80)
	tt.SetDark(true)
	if _, ok := tt.Find("separator"); !ok {
		t.Fatal("separator missing under dark mode")
	}
}
