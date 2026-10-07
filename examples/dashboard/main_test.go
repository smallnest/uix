package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/dashboard/view"
)

func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.DashboardView, view.Width, view.Height)
}

// TestOverview shows the whole dashboard: the heading, the metric
// cards, both charts with their yearly headlines, and the display cards.
func TestOverview(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{"Overview", "Customers", "14,592", "Revenue",
		"$152,314", "+16%", "Orders", "25,162", "+20%", "Total revenue",
		"$152,313.92", "Against last month", "This year", "Last year"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestHover follows the pointer on the revenue chart: resting on August
// swaps the headline for that month and its year-earlier figure.
func TestHover(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("Aug")
	if !ok {
		t.Fatalf("no month column found in %q", tt.Texts())
	}
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	if view.State.Revenue != 7 {
		t.Fatalf("Revenue = %d, want 7", view.State.Revenue)
	}
	for _, s := range []string{"August", "$13,120", "+9.1%", "$12,030 last year"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestHoverClears clears the selection when the pointer leaves the
// chart.
func TestHoverClears(t *testing.T) {
	tt := newTester(t)
	r, _ := tt.Find("Aug")
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	tt.Move(1060, 800) // the corner, outside the charts
	tt.Frame()
	if view.State.Revenue != -1 {
		t.Fatalf("Revenue = %d, want -1", view.State.Revenue)
	}
}

// TestOrdersHeadline shows the orders chart's month when the state rests
// on one, as the pointer would set it.
func TestOrdersHeadline(t *testing.T) {
	view.Reset()
	view.State.Orders = 9
	tt := ui.NewTester(view.DashboardView, view.Width, view.Height)
	for _, s := range []string{"October", "2,410", "+22.3%", "1,970 last year"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestDark draws the dashboard under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
	if !tt.HasText("Revenue") {
		t.Fatalf("missing Revenue in %q", tt.Texts())
	}
}
