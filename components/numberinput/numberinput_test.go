package numberinput

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			NumberInput(c, p)
		})
	}
}

// TestRenders shows the stepping buttons beside the field.
func TestRenders(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1}), 260, 100)
	for _, s := range []string{"−", "+"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestStepsUp raises the value with the plus button.
func TestStepsUp(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1}), 260, 100)
	if err := tt.Click("+"); err != nil {
		t.Fatal(err)
	}
	if v != 43 {
		t.Fatalf("plus stepped to %v, want 43", v)
	}
}

// TestStepsDown lowers the value with the minus button.
func TestStepsDown(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1}), 260, 100)
	if err := tt.Click("−"); err != nil {
		t.Fatal(err)
	}
	if v != 41 {
		t.Fatalf("minus stepped to %v, want 41", v)
	}
}

// TestClamps keeps the value inside the bounds.
func TestClamps(t *testing.T) {
	v := 99.0
	tt := ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1}), 260, 100)
	if err := tt.Click("+"); err != nil {
		t.Fatal(err)
	}
	if v != 99 {
		t.Fatalf("plus went past the top to %v", v)
	}
	v = 0.0
	tt = ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1}), 260, 100)
	if err := tt.Click("−"); err != nil {
		t.Fatal(err)
	}
	if v != 0 {
		t.Fatalf("minus went past the bottom to %v", v)
	}
}

// TestDisabled blocks the buttons.
func TestDisabled(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1, Disabled: true}), 260, 100)
	if err := tt.Click("+"); err != nil {
		t.Fatal(err)
	}
	if v != 42 {
		t.Fatalf("disabled field stepped to %v", v)
	}
}

// TestDarkMode draws the number input under the dark appearance too.
func TestDarkMode(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Lo: 0, Hi: 99, Step: 1}), 260, 100)
	tt.SetDark(true)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
