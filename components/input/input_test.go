package input

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(props ...Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			for _, p := range props {
				Input(c, p)
			}
		})
	}
}

// The value of an input is drawn by its editor, not by a text node, so
// these tests check that a frame renders and that the modes of the field
// do not break it.
func TestRenders(t *testing.T) {
	value := "Ada"
	tt := ui.NewTester(frame(Props{Value: &value, Placeholder: "a name"}), 240, 100)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

func TestDisabled(t *testing.T) {
	value := "locked"
	tt := ui.NewTester(frame(Props{Value: &value, Disabled: true}), 240, 100)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

func TestDarkMode(t *testing.T) {
	value := "Ada"
	tt := ui.NewTester(frame(Props{Value: &value, Placeholder: "a name"}), 240, 100)
	tt.SetDark(true)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
