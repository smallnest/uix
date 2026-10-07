package barchart

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// series are a year of monthly orders: a rising current year against a
// flatter year before, so the totals and the deltas are known.
var series = []Point{
	{Label: "Jan", Current: 1820, Previous: 1650},
	{Label: "Feb", Current: 1750, Previous: 1700},
	{Label: "Mar", Current: 1980, Previous: 1800},
	{Label: "Apr", Current: 1930, Previous: 1850},
	{Label: "May", Current: 2100, Previous: 1900},
	{Label: "Jun", Current: 2210, Previous: 2000},
	{Label: "Jul", Current: 2290, Previous: 2080},
	{Label: "Aug", Current: 2380, Previous: 2150},
	{Label: "Sep", Current: 2460, Previous: 2240},
	{Label: "Oct", Current: 2590, Previous: 2350},
	{Label: "Nov", Current: 2710, Previous: 2480},
	{Label: "Dec", Current: 2790, Previous: 2560},
}

// view is the harness the tests draw the chart in.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			BarChart(c, p)
		})
	}
}

// TestRenders shows the headline of the whole year: the title, the
// summed orders with the delta against the year before, and the legend.
func TestRenders(t *testing.T) {
	selected := -1
	tt := ui.NewTester(view(Props{Data: series, Selected: &selected, Title: "Orders"}), 760, 320)
	for _, s := range []string{"Orders", "27,010", "+9.1%", "24,760 last year",
		"This year", "Last year"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSelect follows the pointer: the headline shows the month and its
// year-earlier figure.
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
	for _, s := range []string{"August", "2,380", "2,150 last year", "+10.7%"} {
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
	if !tt.HasText("Orders") {
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
	if got := grouped(27010); got != "27,010" {
		t.Fatalf("grouped(27010) = %q; want 27,010", got)
	}
	if got := grouped(385); got != "385" {
		t.Fatalf("grouped(385) = %q; want 385", got)
	}
	if got := formatK(14000); got != "$14k" {
		t.Fatalf("formatK(14000) = %q; want $14k", got)
	}
}
