package statcard

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			StatCards(c, p)
		})
	}
}

var plain = []Stat{
	{Icon: IconUsers, Label: "Customers", Value: "14,592", Delta: "+5.3%", DeltaColor: Up},
	{Icon: IconBox, Label: "Unit sold", Value: "385", Delta: "-2.1%", DeltaColor: Down},
	{Icon: IconBasket, Label: "Orders", Value: "1,394", Delta: "0.00%", DeltaColor: Flat},
	{Icon: IconChat, Label: "Support tickets", Value: "708", Delta: "+12.8%", DeltaColor: Up},
}

// TestRenders shows the plain cards: the label, the value and the delta
// of every metric.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Stats: plain}), 760, 200)
	for _, s := range []string{"Customers", "14,592", "+5.3%", "Unit sold", "385",
		"Orders", "1,394", "Support tickets", "708"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestFooter shows the display card: the label, the value, the
// comparison caption and the delta pill.
func TestFooter(t *testing.T) {
	tt := ui.NewTester(frame(Props{Variant: Footer, Stats: []Stat{
		{Icon: IconCoins, Label: "Total revenue", Value: "$152,313.92",
			Delta: "16%", DeltaColor: Up, Tone: Blue, Caption: "Against last month"},
	}}), 360, 240)
	for _, s := range []string{"Total revenue", "$152,313.92", "16%", "Against last month"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestHint names the info glyph of a footer card with a hint.
func TestHint(t *testing.T) {
	tt := ui.NewTester(frame(Props{Variant: Footer, Stats: []Stat{
		{Icon: IconCoins, Label: "Total revenue", Value: "$1",
			Delta: "1%", DeltaColor: Up, Hint: "Gross revenue this month."},
	}}), 360, 240)
	if !tt.HasText("About Total revenue") {
		t.Fatalf("the hint glyph is missing from %q", tt.Texts())
	}
}

// TestDraws renders every tone and delta without a blank frame.
func TestDraws(t *testing.T) {
	var tones []Stat
	for _, tone := range []Tone{Blue, Orange, Purple, Pink, Sky, Emerald} {
		tones = append(tones, Stat{Icon: IconSparkle, Label: "Metric", Value: "1",
			Delta: "1%", DeltaColor: Flat, Tone: tone})
	}
	tt := ui.NewTester(frame(Props{Variant: Footer, Stats: tones}), 760, 240)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
