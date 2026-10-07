package badge

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(props ...Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			for _, p := range props {
				Badge(c, p)
			}
		})
	}
}

// TestRenders shows the label of a badge.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Beta", Variant: Secondary}), 200, 80)
	if !tt.HasText("Beta") {
		t.Fatal("label missing")
	}
}

// TestAllVariants renders every variant: a smoke test over the matrix.
func TestAllVariants(t *testing.T) {
	tt := ui.NewTester(frame(
		Props{Label: "Default"},
		Props{Label: "Secondary", Variant: Secondary},
		Props{Label: "Destructive", Variant: Destructive},
		Props{Label: "Outline", Variant: Outline},
	), 320, 120)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
