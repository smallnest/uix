package button

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame builds a test frame that shows the buttons for the given props.
func frame(props ...Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Center().Gap(8).Children(func() {
			for _, p := range props {
				Button(c, p)
			}
		})
	}
}

func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Save", Variant: Primary}), 200, 100)
	if !tt.HasText("Save") {
		t.Fatal("button text missing")
	}
}

func TestClick(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(frame(Props{
		Label:   "Go",
		Variant: Primary,
		OnClick: func() { clicks++ },
	}), 200, 100)
	if err := tt.Click("Go"); err != nil {
		t.Fatal(err)
	}
	if clicks != 1 {
		t.Fatalf("OnClick ran %d times, want 1", clicks)
	}
}

func TestDisabled(t *testing.T) {
	clicks := 0
	tt := ui.NewTester(frame(Props{
		Label:    "No",
		Disabled: true,
		OnClick:  func() { clicks++ },
	}), 200, 100)
	if err := tt.Click("No"); err != nil {
		t.Fatal(err)
	}
	if clicks != 0 {
		t.Fatalf("disabled button fired OnClick %d times, want 0", clicks)
	}
}

// TestAllVariantsAndSizes renders every variant in every size and checks
// that a frame comes out: a smoke test over the whole matrix.
func TestAllVariantsAndSizes(t *testing.T) {
	var props []Props
	for _, v := range []Variant{Primary, Secondary, Outline, Ghost, Destructive, Link} {
		for _, s := range []Size{Sm, Md, Lg} {
			props = append(props, Props{Label: "x", Variant: v, Size: s})
		}
	}
	tt := ui.NewTester(frame(props...), 400, 300)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 || img.Bounds().Dy() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestDarkMode renders the variants under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(
		Props{Label: "Primary", Variant: Primary},
		Props{Label: "Ghost", Variant: Ghost},
		Props{Label: "Danger", Variant: Destructive},
	), 240, 120)
	tt.SetDark(true)
	if !tt.HasText("Danger") {
		t.Fatal("dark frame missing a button")
	}
}
