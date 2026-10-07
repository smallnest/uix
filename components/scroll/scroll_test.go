package scroll

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Scroll(c, p).Grow(1)
		})
	}
}

func rows(n int) func(c *ui.Context) {
	return func(c *ui.Context) {
		for i := 0; i < n; i++ {
			ui.Text(c, "row "+fmt.Sprint(i))
		}
	}
}

// TestRenders shows the rows that fit in the container.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: rows(5)}), 300, 200)
	for i := 0; i < 5; i++ {
		if !tt.HasText("row " + fmt.Sprint(i)) {
			t.Fatalf("missing row %d in %q", i, tt.Texts())
		}
	}
}

// TestScrolls moves the content up as the pointer wheels: a row rises
// within the container, and a row below comes into view.
func TestScrolls(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: rows(30)}), 300, 200)
	r, _ := tt.Find("row 5")
	y0 := r.Y
	tt.Scroll(r.X, r.Y, 0, 40)
	r, _ = tt.Find("row 5")
	if r.W == 0 || r.Y >= y0 {
		t.Fatalf("the scroll did not move the content up: y %g from %g", r.Y, y0)
	}
}
