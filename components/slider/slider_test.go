package slider

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Slider(c, p)
		})
	}
}

func TestRenders(t *testing.T) {
	value := 60.0
	tt := ui.NewTester(frame(Props{
		Value:  &value,
		Min:    0,
		Max:    100,
		Label:  "Volume",
		Format: func(v float64) string { return fmt.Sprintf("%.0f%%", v) },
	}), 300, 120)
	for _, s := range []string{"Volume", "60%"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

func TestDefaultFormat(t *testing.T) {
	value := 3.0
	tt := ui.NewTester(frame(Props{Value: &value, Min: 0, Max: 10, Label: "Level"}), 300, 120)
	if !tt.HasText("3") {
		t.Fatal("default readout missing")
	}
}

func TestStepSlider(t *testing.T) {
	value := 1.0
	tt := ui.NewTester(frame(Props{Value: &value, Min: 0, Max: 4, Step: 1}), 300, 120)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
