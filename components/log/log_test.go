package log

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Log(c, p)
		})
	}
}

// TestReveals shows only the items the app has revealed.
func TestReveals(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Items:    []string{"Planning", "Searching", "Writing"},
		Revealed: 2,
	}), 320, 160)
	if !tt.HasText("Planning") || !tt.HasText("Searching") {
		t.Fatalf("revealed rows missing from %q", tt.Texts())
	}
	if tt.HasText("Writing") {
		t.Fatalf("an unrevealed row shows in %q", tt.Texts())
	}
}

// TestClamps keeps Revealed within the items, as a log that over-runs
// cannot show more than it has.
func TestClamps(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Items:    []string{"Planning"},
		Revealed: 9,
	}), 320, 160)
	if !tt.HasText("Planning") {
		t.Fatalf("the item is missing from %q", tt.Texts())
	}
}

// TestRuns shows the working row with its label, below the revealed
// items.
func TestRuns(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Items:    []string{"Planning", "Searching"},
		Revealed: 2,
		Running:  true,
		Working:  "Writing the summary",
	}), 320, 160)
	if !tt.HasText("Writing the summary") {
		t.Fatalf("the working row is missing from %q", tt.Texts())
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestNone draws an empty log for nothing revealed.
func TestNone(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Items:    []string{"Planning", "Searching"},
		Revealed: 0,
	}), 320, 160)
	for _, s := range []string{"Planning", "Searching"} {
		if tt.HasText(s) {
			t.Fatalf("a hidden item shows in %q", tt.Texts())
		}
	}
}
