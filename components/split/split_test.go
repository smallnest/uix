package split

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws the split of p across a padded window.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(8).Children(func() {
			Split(c, p).Grow(1)
		})
	}
}

// newSplit returns a tester of a split whose panes name themselves.
func newSplit(t *testing.T, size float32) (*ui.Tester, *float32) {
	t.Helper()
	p := Props{Size: &size, First: func(c *ui.Context) { ui.Text(c, "Left") }, Second: func(c *ui.Context) { ui.Text(c, "Right") }}
	return ui.NewTester(frame(p), 400, 200), p.Size
}

// TestRenders draws both panes.
func TestRenders(t *testing.T) {
	tt, _ := newSplit(t, 160)
	for _, s := range []string{"Left", "Right"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestDrags moves the divider when the pointer drags it. The window
// pads the split by 8 DIPs, so the divider of a 160 DIP first pane
// sits at x 168.
func TestDrags(t *testing.T) {
	tt, size := newSplit(t, 160)
	tt.Press(168, 100)
	tt.Frame()
	tt.Move(228, 100)
	tt.Frame()
	tt.Release(228, 100)
	tt.Frame()
	if *size < 200 {
		t.Fatalf("size %v, want the drag to widen the first pane", *size)
	}
	if !tt.HasText("Left") || !tt.HasText("Right") {
		t.Fatal("the drag lost a pane")
	}
}

// TestKeepsEdge keeps the divider off the edges of the split, which a
// drag toward one edge clamps.
func TestKeepsEdge(t *testing.T) {
	tt, size := newSplit(t, 160)
	tt.Press(168, 100)
	tt.Frame()
	tt.Move(3, 100)
	tt.Frame()
	tt.Release(3, 100)
	tt.Frame()
	if *size < 40 {
		t.Fatalf("size %v, want the divider kept off the edge", *size)
	}
}

// TestVertical stacks the panes, first above second.
func TestVertical(t *testing.T) {
	size := float32(80)
	p := Props{Size: &size, Vertical: true, First: func(c *ui.Context) { ui.Text(c, "Top") }, Second: func(c *ui.Context) { ui.Text(c, "Bottom") }}
	tt := ui.NewTester(frame(p), 400, 300)
	for _, s := range []string{"Top", "Bottom"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestDarkMode draws the split under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt, _ := newSplit(t, 160)
	tt.SetDark(true)
	if !tt.HasText("Left") {
		t.Fatal("split missing under dark mode")
	}
}
