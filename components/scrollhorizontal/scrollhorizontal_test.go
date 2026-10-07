package scrollhorizontal

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			ScrollHorizontal(c, p).FillWidth()
		})
	}
}

func cells(n int) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Row(c).Gap(8).Children(func() {
			for i := 0; i < n; i++ {
				ui.Box(c).Size(120, 48).Radius(4).Background(ui.RGB(200, 200, 200)).Children(func() {
					ui.Text(c, "cell "+fmt.Sprint(i))
				})
			}
		})
	}
}

// TestRenders shows the cells that fit in the container.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: cells(3)}), 300, 100)
	for i := 0; i < 3; i++ {
		if !tt.HasText("cell " + fmt.Sprint(i)) {
			t.Fatalf("missing cell %d in %q", i, tt.Texts())
		}
	}
}

// TestScrolls moves the content left as the pointer wheels: a cell moves
// sideways within the container, and a cell to the right comes into view.
func TestScrolls(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: cells(20)}), 300, 100)
	r, _ := tt.Find("cell 2")
	x0 := r.X
	tt.Scroll(r.X, r.Y, 40, 0)
	r, _ = tt.Find("cell 2")
	if r.W == 0 || r.X >= x0 {
		t.Fatalf("the scroll did not move the content left: x %g from %g", r.X, x0)
	}
}
