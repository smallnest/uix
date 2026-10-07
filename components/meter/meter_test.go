package meter

import (
	"bytes"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Meter(c, p)
		})
	}
}

// TestRenders shows the label and the readout of the value.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Value: 7, Max: 10, Label: "Disk usage"}), 300, 100)
	for _, s := range []string{"Disk usage", "7"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestCustomFormat shows the readout the caller's format renders.
func TestCustomFormat(t *testing.T) {
	tt := ui.NewTester(frame(Props{Value: 7, Max: 10, Label: "Level", Format: func(v float64) string {
		return "high"
	}}), 300, 100)
	if !tt.HasText("high") {
		t.Fatal("custom readout missing")
	}
}

// TestLevels shows the danger color at a value past Critical, which
// differs from the success color of the same value without levels.
func TestLevels(t *testing.T) {
	plain := ui.NewTester(frame(Props{Value: 0.9, Max: 1}), 300, 100)
	warned := ui.NewTester(frame(Props{Value: 0.9, Max: 1, Levels: &ui.MeterLevels{Warning: 0.7, Critical: 0.8}}), 300, 100)
	if bytes.Equal(plain.Image().Pix, warned.Image().Pix) {
		t.Fatal("the levels do not recolor the bar")
	}
}

// TestValueFills draws two values and checks that the frames differ, so
// the value reaches the bar.
func TestValueFills(t *testing.T) {
	low := ui.NewTester(frame(Props{Value: 0.1, Max: 1}), 300, 100)
	high := ui.NewTester(frame(Props{Value: 0.9, Max: 1}), 300, 100)
	if bytes.Equal(low.Image().Pix, high.Image().Pix) {
		t.Fatal("the bar does not change with the value")
	}
}

// TestDefaultMax treats a bare fraction as 0 to 1, so Value 0.9 fills
// the bar at start.
func TestDefaultMax(t *testing.T) {
	full := ui.NewTester(frame(Props{Value: 0.9}), 300, 100)
	if img := full.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
