package grid

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Grid(c, p)
		})
	}
}

func letters() func(c *ui.Context) {
	return func(c *ui.Context) {
		for _, s := range []string{"A", "B", "C", "D"} {
			ui.Text(c, s)
		}
	}
}

// TestRenders lays out children in cells of columns, which all show.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Columns: 2, Children: letters()}), 300, 200)
	for _, s := range []string{"A", "B", "C", "D"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestCells places the first two children side by side in the first row,
// and the third under the first: they fill the columns in order.
func TestCells(t *testing.T) {
	tt := ui.NewTester(frame(Props{Columns: 2, Children: letters()}), 300, 200)
	a, _ := tt.Find("A")
	b, _ := tt.Find("B")
	cc, _ := tt.Find("C")
	if a.Y != b.Y {
		t.Fatalf("A and B are not in the same row: y %g vs %g", a.Y, b.Y)
	}
	if b.X <= a.X {
		t.Fatalf("B (%g) is not right of A (%g)", b.X, a.X)
	}
	if cc.Y <= a.Y {
		t.Fatalf("C (%g) is not below A (%g)", cc.Y, a.Y)
	}
}

// TestTracks sizes the columns from tracks: the fixed tracks hold their
// widths, so the three columns of the first row measure 60 and 90.
func TestTracks(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		ColumnTracks: []ui.Track{ui.Fixed(60), ui.Fixed(90), ui.Fixed(120)},
		Children:     letters(),
	}), 300, 200)
	a, _ := tt.Find("A")
	b, _ := tt.Find("B")
	cc, _ := tt.Find("C")
	if b.X-a.X != 60 {
		t.Fatalf("the first column is %g wide, want 60", b.X-a.X)
	}
	if cc.X-b.X != 90 {
		t.Fatalf("the second column is %g wide, want 90", cc.X-b.X)
	}
}
