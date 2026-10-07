package spinner

import (
	"bytes"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Center().Children(func() {
			Spinner(c, p)
		})
	}
}

// TestRenders draws the spinner in both appearances.
func TestRenders(t *testing.T) {
	for _, dark := range []bool{false, true} {
		tt := ui.NewTester(frame(Props{Label: "Scanning"}), 200, 100)
		tt.SetDark(dark)
		if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
			t.Fatalf("no frame rendered (dark %v)", dark)
		}
	}
}

// TestTurns draws the spinner at two moments, which differ, so it turns.
func TestTurns(t *testing.T) {
	tt := ui.NewTester(frame(Props{}), 200, 100)
	first := tt.Image()
	time.Sleep(100 * time.Millisecond)
	tt.Frame()
	second := tt.Image()
	if bytes.Equal(first.Pix, second.Pix) {
		t.Fatal("the spinner does not turn")
	}
}
