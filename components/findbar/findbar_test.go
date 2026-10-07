package findbar

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			FindBar(c, p)
		})
	}
}

// TestCounts shows the current match and the count, "n of m".
func TestCounts(t *testing.T) {
	open, query, current := true, "the", 0
	tt := ui.NewTester(frame(Props{Open: &open, Query: &query, Matches: 3, Current: &current}), 400, 120)
	if !tt.HasText("1 of 3") {
		t.Fatalf("missing the count in %q", tt.Texts())
	}
}

// TestNoMatches names a query with no matches.
func TestNoMatches(t *testing.T) {
	open, query, current := true, "xyz", 0
	tt := ui.NewTester(frame(Props{Open: &open, Query: &query, Matches: 0, Current: &current}), 400, 120)
	if !tt.HasText("No matches") {
		t.Fatalf("missing the no-match note in %q", tt.Texts())
	}
}

// TestSteps moves the current match with the next and the previous
// buttons, wrapping at the ends.
func TestSteps(t *testing.T) {
	open, query, current := true, "the", 1
	tt := ui.NewTester(frame(Props{Open: &open, Query: &query, Matches: 3, Current: &current}), 400, 120)
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if current != 2 {
		t.Fatalf("Next set %d, want 2", current)
	}
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if current != 0 {
		t.Fatalf("Next wrapped to %d, want 0", current)
	}
	if err := tt.Click("Previous"); err != nil {
		t.Fatal(err)
	}
	if current != 2 {
		t.Fatalf("Previous set %d, want 2", current)
	}
}

// TestTypes edits the query in the field, which a bar that opens
// focuses and selects.
func TestTypes(t *testing.T) {
	open, query, current := true, "", 0
	tt := ui.NewTester(frame(Props{Open: &open, Query: &query, Matches: 0, Current: &current}), 400, 120)
	tt.Type("ui")
	if query != "ui" {
		t.Fatalf("the field has %q, want ui", query)
	}
	if !tt.HasText("No matches") {
		t.Fatalf("missing the no-match note in %q", tt.Texts())
	}
}

// TestDone closes the bar with its Done button.
func TestDone(t *testing.T) {
	open, query, current := true, "the", 0
	tt := ui.NewTester(frame(Props{Open: &open, Query: &query, Matches: 3, Current: &current}), 400, 120)
	if err := tt.Click("Done"); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Fatal("Done did not close the bar")
	}
}

// TestEscape closes the bar with the Escape key.
func TestEscape(t *testing.T) {
	open, query, current := true, "the", 0
	tt := ui.NewTester(frame(Props{Open: &open, Query: &query, Matches: 3, Current: &current}), 400, 120)
	tt.Key(0, ui.KeyEscape)
	if open {
		t.Fatal("Escape did not close the bar")
	}
}
