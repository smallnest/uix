package alert

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(props ...Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			for _, p := range props {
				Alert(c, p)
			}
		})
	}
}

// TestRenders shows the title and the description of an alert.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Title:       "Heads up!",
		Description: "You can add components to your app.",
		Variant:     Info,
	}), 300, 120)
	for _, s := range []string{"Heads up!", "You can add components to your app."} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestTitleOnly draws an alert without a description.
func TestTitleOnly(t *testing.T) {
	tt := ui.NewTester(frame(Props{Title: "Saved", Variant: Success}), 300, 80)
	if !tt.HasText("Saved") {
		t.Fatal("title missing")
	}
}

// TestAllVariants renders every variant: a smoke test over the matrix.
func TestAllVariants(t *testing.T) {
	tt := ui.NewTester(frame(
		Props{Title: "Info", Variant: Info},
		Props{Title: "Success", Variant: Success},
		Props{Title: "Warning", Variant: Warning},
		Props{Title: "Danger", Variant: Danger},
	), 300, 240)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestDarkMode draws the alerts under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Title: "Heads up!", Variant: Info}), 300, 120)
	tt.SetDark(true)
	if !tt.HasText("Heads up!") {
		t.Fatal("title missing under dark mode")
	}
}
