package rating

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Rating(c, p).Label("Rating")
		})
	}
}

// starX returns the x of the center of the i-th star of the row at r.
// The row stretches, and its stars sit at its start, as high as they are
// wide, so the row height is the star size; the gap between them is half
// a space of the default theme.
func starX(r ui.Rect, i, max int) float32 {
	step := r.H + float32(2)
	return r.X + float32(i)*step + r.H/2
}

func TestRenders(t *testing.T) {
	v := 3
	tt := ui.NewTester(frame(Props{Value: &v, Max: 5}), 300, 100)
	if _, ok := tt.Find("Rating"); !ok {
		t.Fatal("rating not found")
	}
}

// TestClicks sets the score to the star clicked.
func TestClicks(t *testing.T) {
	v := 0
	tt := ui.NewTester(frame(Props{Value: &v, Max: 5}), 300, 100)
	r, ok := tt.Find("Rating")
	if !ok {
		t.Fatal("rating not found")
	}
	tt.ClickAt(starX(r, 2, 5), r.Y+r.H/2)
	if v != 3 {
		t.Fatalf("click on the third star: %d, want 3", v)
	}
}

// TestClears a click on the set star takes the score back to zero.
func TestClears(t *testing.T) {
	v := 3
	tt := ui.NewTester(frame(Props{Value: &v, Max: 5}), 300, 100)
	r, _ := tt.Find("Rating")
	tt.ClickAt(starX(r, 2, 5), r.Y+r.H/2)
	if v != 0 {
		t.Fatalf("click on the set star: %d, want 0", v)
	}
}

// TestKeyboard moves the score with the arrows once the row is focused.
func TestKeyboard(t *testing.T) {
	v := 0
	tt := ui.NewTester(frame(Props{Value: &v, Max: 5}), 300, 100)
	r, _ := tt.Find("Rating")
	// A click on a star focuses the row; the arrows then move the score.
	tt.ClickAt(starX(r, 0, 5), r.Y+r.H/2)
	tt.Key(0, ui.KeyRight)
	tt.Key(0, ui.KeyRight)
	if v != 3 {
		t.Fatalf("Right twice: %d, want 3", v)
	}
	tt.Key(0, ui.KeyEnd)
	if v != 5 {
		t.Fatalf("End: %d, want 5", v)
	}
}

// TestReadOnly draws the stars without taking a click or a key.
func TestReadOnly(t *testing.T) {
	v := 3
	tt := ui.NewTester(frame(Props{Value: &v, Max: 5, ReadOnly: true}), 300, 100)
	r, ok := tt.Find("Rating")
	if !ok {
		t.Fatal("rating not found")
	}
	tt.ClickAt(starX(r, 2, 5), r.Y+r.H/2)
	tt.Key(0, ui.KeyRight)
	if v != 3 {
		t.Fatalf("read-only rating changed to %d", v)
	}
}

// TestDarkMode draws the rating under the dark appearance too.
func TestDarkMode(t *testing.T) {
	v := 3
	tt := ui.NewTester(frame(Props{Value: &v, Max: 5}), 300, 100)
	tt.SetDark(true)
	if _, ok := tt.Find("Rating"); !ok {
		t.Fatal("rating missing under dark mode")
	}
}
