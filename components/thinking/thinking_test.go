package thinking

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(props ...Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(12).Children(func() {
			for _, p := range props {
				Thinking(c, p)
			}
		})
	}
}

// TestRenders shows the label of the indicator.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Thinking"}), 200, 60)
	if !tt.HasText("Thinking") {
		t.Fatalf("the label is missing from %q", tt.Texts())
	}
}

// TestAllVariants renders every variant and tone: a smoke test over the
// matrix, which must draw without a panic.
func TestAllVariants(t *testing.T) {
	tt := ui.NewTester(frame(
		Props{Label: "Wave", Variant: Wave, Tone: Subtle},
		Props{Label: "Spin", Variant: Spin, Tone: Default},
		Props{Label: "Stars", Variant: Stars, Tone: Primary},
		Props{Label: "Infinity", Variant: Infinity, Tone: Accent},
	), 320, 160)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestTimer runs beside the label: the indicator draws an elapsed time
// and asks for repaints, which the tester settles.
func TestTimer(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Label:     "Thinking",
		ShowTimer: true,
	}), 200, 60)
	if !tt.HasText("Thinking") {
		t.Fatalf("the label is missing from %q", tt.Texts())
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
