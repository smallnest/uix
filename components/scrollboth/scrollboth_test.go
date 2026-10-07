package scrollboth

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			ScrollBoth(c, p).Grow(1)
		})
	}
}

func tiles(n int) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Row(c).Gap(8).Children(func() {
			for i := 0; i < n; i++ {
				ui.Column(c).Gap(8).Children(func() {
					for j := 0; j < n; j++ {
						ui.Text(c, fmt.Sprintf("%dx%d", i, j))
					}
				})
			}
		})
	}
}

// TestRenders shows the tiles that fit in the container.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: tiles(3)}), 300, 200)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			s := fmt.Sprintf("%dx%d", i, j)
			if !tt.HasText(s) {
				t.Fatalf("missing %q in %q", s, tt.Texts())
			}
		}
	}
}

// TestScrollsBothWays moves the content up and sideways as the pointer
// wheels: a tile moves up and left within the container.
func TestScrollsBothWays(t *testing.T) {
	tt := ui.NewTester(frame(Props{Children: tiles(12)}), 300, 200)
	r, _ := tt.Find("5x5")
	x0, y0 := r.X, r.Y
	tt.Scroll(r.X, r.Y, 40, 40)
	r, _ = tt.Find("5x5")
	if r.W == 0 || r.X >= x0 || r.Y >= y0 {
		t.Fatalf("the scroll did not move the content: x %g from %g, y %g from %g", r.X, x0, r.Y, y0)
	}
}
