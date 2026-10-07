package progress

import (
	"bytes"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Progress(c, p)
		})
	}
}

// TestRenders shows the label and the readout of the value.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Value: 0.64, Label: "Disk usage"}), 300, 100)
	for _, s := range []string{"Disk usage", "64%"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestCustomFormat shows the readout the caller's format renders.
func TestCustomFormat(t *testing.T) {
	tt := ui.NewTester(frame(Props{Value: 0.5, Label: "Level", Format: func(float64) string { return "half" }}), 300, 100)
	if !tt.HasText("half") {
		t.Fatal("custom readout missing")
	}
}

// TestIndeterminate hides the readout and still renders a frame.
func TestIndeterminate(t *testing.T) {
	tt := ui.NewTester(frame(Props{Value: 0, Label: "Indexing", Indeterminate: true}), 300, 100)
	if tt.HasText("0%") {
		t.Fatal("indeterminate shows a readout")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestValueFills draws two values and checks that the frames differ, so
// the value reaches the track.
func TestValueFills(t *testing.T) {
	low := ui.NewTester(frame(Props{Value: 0.1}), 300, 100)
	high := ui.NewTester(frame(Props{Value: 0.9}), 300, 100)
	if bytes.Equal(low.Image().Pix, high.Image().Pix) {
		t.Fatal("the bar does not change with the value")
	}
}
