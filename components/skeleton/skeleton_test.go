package skeleton

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// tester draws a skeleton bar in a padded column and returns its tester
// and the window background, sampled far from the bar.
func tester(p Props) (*ui.Tester, interface{}) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Skeleton(c, p)
		})
	}, 240, 80)
	return tt, tt.Image().RGBAAt(230, 75)
}

// TestShows draws a bar of the given size: a point inside it is not the
// window background.
func TestShows(t *testing.T) {
	tt, bg := tester(Props{Width: 120, Height: 24})
	if c := tt.Image().RGBAAt(30, 27); c == bg {
		t.Fatalf("no bar drew at 30,27: %v", c)
	}
}

// TestFills draws a bar without a size that fills the column.
func TestFills(t *testing.T) {
	tt, bg := tester(Props{})
	if c := tt.Image().RGBAAt(200, 27); c == bg {
		t.Fatalf("the bar did not fill the width: %v at 200,27", c)
	}
}

// TestPulses moves the bar's tone over time, as the pulse animates it.
func TestPulses(t *testing.T) {
	tt, _ := tester(Props{Width: 120, Height: 24})
	first := tt.Image().RGBAAt(30, 27)
	// Let the pulse move: frames over a good part of the period, with
	// time passing between them, as the animation reads the clock.
	for i := 0; i < 30; i++ {
		time.Sleep(10 * time.Millisecond)
		tt.Frame()
	}
	second := tt.Image().RGBAAt(30, 27)
	if first == second {
		t.Fatalf("the pulse did not move the tone: %v both frames", first)
	}
}
