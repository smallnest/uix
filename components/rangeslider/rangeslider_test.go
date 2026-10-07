package rangeslider

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws a range slider without a label, so the track sits at a
// known place: the top padding 16, the height Space(5) 20, the track
// center at y 26 across x 16..284 of the 300 wide window.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			RangeSlider(c, p)
		})
	}
}

// TestRenders shows the label and the readout of the two values.
func TestRenders(t *testing.T) {
	low, high := 20.0, 80.0
	tt := ui.NewTester(frame(Props{
		Low: &low, High: &high, Min: 0, Max: 100,
		Label:  "Price",
		Format: func(lo, hi float64) string { return fmt.Sprintf("$%.0f – $%.0f", lo, hi) },
	}), 300, 120)
	for _, s := range []string{"Price", "$20 – $80"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestDefaultFormat prints the two values rounded.
func TestDefaultFormat(t *testing.T) {
	low, high := 2.0, 8.0
	tt := ui.NewTester(frame(Props{Low: &low, High: &high, Min: 0, Max: 10, Label: "Range"}), 300, 120)
	if !tt.HasText("2 – 8") {
		t.Fatalf("default readout missing in %q", tt.Texts())
	}
}

// TestDragsLow presses the track at its middle, which moves the nearest
// knob — the low one, as the two are equally near — toward it.
func TestDragsLow(t *testing.T) {
	low, high := 20.0, 80.0
	tt := ui.NewTester(frame(Props{Low: &low, High: &high, Min: 0, Max: 100}), 300, 120)
	tt.Press(150, 26)
	tt.Frame()
	tt.Move(190, 26)
	tt.Frame()
	tt.Release(190, 26)
	tt.Frame()
	if low < 60 || low > 70 || high != 80 {
		t.Fatalf("the drag set low %v high %v, want low about 65, high 80", low, high)
	}
}

// TestDragsHigh presses the track near its right end, which moves the
// high knob toward it, leaving the low one.
func TestDragsHigh(t *testing.T) {
	low, high := 20.0, 80.0
	tt := ui.NewTester(frame(Props{Low: &low, High: &high, Min: 0, Max: 100}), 300, 120)
	tt.Press(270, 26)
	tt.Frame()
	tt.Release(270, 26)
	tt.Frame()
	if low != 20 || high < 90 {
		t.Fatalf("the drag set low %v high %v, want low 20, high about 95", low, high)
	}
}

// TestKeys moves the knobs with the keyboard: Tab enters the low knob,
// the Right key moves it, Tab moves to the high knob, and the Right key
// moves it.
func TestKeys(t *testing.T) {
	low, high := 20.0, 80.0
	tt := ui.NewTester(frame(Props{Low: &low, High: &high, Min: 0, Max: 100, Step: 10}), 300, 120)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if low != 30 {
		t.Fatalf("the low knob moved to %v, want 30", low)
	}
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if high != 90 {
		t.Fatalf("the high knob moved to %v, want 90", high)
	}
}

// TestClamps keeps the low knob from passing the high one: End sets it
// to the high, and the Right key cannot push it past.
func TestClamps(t *testing.T) {
	low, high := 20.0, 80.0
	tt := ui.NewTester(frame(Props{Low: &low, High: &high, Min: 0, Max: 100, Step: 10}), 300, 120)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyEnd)
	if low != 80 || high != 80 {
		t.Fatalf("End set low %v high %v, want both 80", low, high)
	}
	tt.Key(0, ui.KeyRight)
	if low != 80 {
		t.Fatalf("the Right key pushed the low knob to %v, want it clamped at 80", low)
	}
}
