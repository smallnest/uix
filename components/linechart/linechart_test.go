package linechart

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// months are the BoardUI revenue series: a rising year against a flat
// year before, so the totals and the deltas are known.
var series = []Point{
	{Label: "Jan", Current: 14218, Previous: 13678},
	{Label: "Feb", Current: 13539, Previous: 13302},
	{Label: "Mar", Current: 15363, Previous: 14610},
	{Label: "Apr", Current: 15231, Previous: 14856},
	{Label: "May", Current: 16507, Previous: 15171},
	{Label: "Jun", Current: 17241, Previous: 16048},
	{Label: "Jul", Current: 18126, Previous: 16748},
	{Label: "Aug", Current: 19067, Previous: 17473},
	{Label: "Sep", Current: 19780, Previous: 18344},
	{Label: "Oct", Current: 21163, Previous: 19303},
	{Label: "Nov", Current: 22473, Previous: 20450},
	{Label: "Dec", Current: 23290, Previous: 21184},
}

// view is the harness the tests draw the chart in.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			LineChart(c, p)
		})
	}
}

// TestRenders shows the headline of the whole year: the title, the
// summed revenue with the delta against the year before, and the legend.
func TestRenders(t *testing.T) {
	selected := -1
	tt := ui.NewTester(view(Props{Data: series, Selected: &selected, Title: "Revenue"}), 760, 320)
	for _, s := range []string{"Revenue", "$215,998", "+7.4%", "$201,167 last year",
		"This year", "Last year"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSelect follows the pointer: the headline shows the month and its
// year-earlier figure, and the month label stays findable.
func TestSelect(t *testing.T) {
	selected := -1
	tt := ui.NewTester(view(Props{Data: series, Selected: &selected}), 760, 320)
	r, ok := tt.Find("Aug")
	if !ok {
		t.Fatalf("no month column found in %q", tt.Texts())
	}
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	if selected != 7 {
		t.Fatalf("Selected = %d, want 7", selected)
	}
	for _, s := range []string{"August", "$19,067", "$17,473 last year", "+9.1%"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestLeave clears the selection when the pointer leaves the plot.
func TestLeave(t *testing.T) {
	selected := -1
	tt := ui.NewTester(view(Props{Data: series, Selected: &selected}), 760, 320)
	r, _ := tt.Find("Aug")
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	tt.Move(750, 310) // the corner, outside the plot
	tt.Frame()
	if selected != -1 {
		t.Fatalf("Selected = %d, want -1", selected)
	}
}

// TestStatic draws a chart that does not follow the pointer.
func TestStatic(t *testing.T) {
	tt := ui.NewTester(view(Props{Data: series}), 760, 320)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
	if !tt.HasText("Revenue") {
		t.Fatalf("missing the default title in %q", tt.Texts())
	}
}

// TestDraws draws the delta helper against known figures.
func TestDraws(t *testing.T) {
	if l, up := deltaOf(100, 80); l != "+25%" || !up {
		t.Fatalf("deltaOf(100, 80) = %q, %v; want +25%%, true", l, up)
	}
	if l, up := deltaOf(80, 100); l != "-20%" || up {
		t.Fatalf("deltaOf(80, 100) = %q, %v; want -20%%, false", l, up)
	}
	if l, _ := deltaOf(100, 100); l != "0%" {
		t.Fatalf("deltaOf(100, 100) = %q; want 0%%", l)
	}
	if l, _ := deltaOf(100, 0); l != "New" {
		t.Fatalf("deltaOf(100, 0) = %q; want New", l)
	}
	if got := formatMoney(213387.0); got != "$213,387" {
		t.Fatalf("formatMoney(213387) = %q; want $213,387", got)
	}
	if got := formatMoney(385); got != "$385" {
		t.Fatalf("formatMoney(385) = %q; want $385", got)
	}
	if got := formatK(14000); got != "$14k" {
		t.Fatalf("formatK(14000) = %q; want $14k", got)
	}
}
