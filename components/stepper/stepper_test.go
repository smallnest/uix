package stepper

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws a stepper without a label, so the control sits at a known
// place: the top padding 16, the arrows 14 high each, the up arrow at
// x 16..36 and y 16..30 of the 300 wide window.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Stepper(c, p)
		})
	}
}

// TestRenders shows the label and the readout of the value.
func TestRenders(t *testing.T) {
	v := 8080.0
	tt := ui.NewTester(frame(Props{Value: &v, Min: 1024, Max: 65535, Step: 1, Label: "Port"}), 260, 100)
	for _, s := range []string{"Port", "8080"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestKeys steps the value with the keyboard: Tab enters the stepper,
// the Up key raises it, the Down key lowers it, and Home and End jump
// to the bounds.
func TestKeys(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Min: 0, Max: 99, Step: 1}), 260, 100)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyUp)
	if v != 43 {
		t.Fatalf("the Up key set %v, want 43", v)
	}
	tt.Key(0, ui.KeyDown)
	if v != 42 {
		t.Fatalf("the Down key set %v, want 42", v)
	}
	tt.Key(0, ui.KeyHome)
	if v != 0 {
		t.Fatalf("Home set %v, want 0", v)
	}
	tt.Key(0, ui.KeyEnd)
	if v != 99 {
		t.Fatalf("End set %v, want 99", v)
	}
}

// TestClamps keeps the value inside the bounds.
func TestClamps(t *testing.T) {
	v := 99.0
	tt := ui.NewTester(frame(Props{Value: &v, Min: 0, Max: 99, Step: 1}), 260, 100)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyUp)
	if v != 99 {
		t.Fatalf("the Up key went past the top to %v", v)
	}
	v = 0.0
	tt = ui.NewTester(frame(Props{Value: &v, Min: 0, Max: 99, Step: 1}), 260, 100)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyDown)
	if v != 0 {
		t.Fatalf("the Down key went past the bottom to %v", v)
	}
}

// TestArrowClick steps the value with the up arrow, a press on the top
// half of the control.
func TestArrowClick(t *testing.T) {
	v := 42.0
	tt := ui.NewTester(frame(Props{Value: &v, Min: 0, Max: 99, Step: 1}), 260, 100)
	tt.ClickAt(26, 23)
	if v != 43 {
		t.Fatalf("the up arrow set %v, want 43", v)
	}
}
