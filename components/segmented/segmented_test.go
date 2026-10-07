package segmented

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Segmented(c, p)
		})
	}
}

// TestRenders shows the labels, with the first segment chosen.
func TestRenders(t *testing.T) {
	selected := 0
	tt := ui.NewTester(frame(Props{Selected: &selected, Labels: []string{"Reader", "Outline"}}), 300, 100)
	for _, s := range []string{"Reader", "Outline"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestClicks chooses a segment with a click, which runs OnChange.
func TestClicks(t *testing.T) {
	selected := 0
	changes := 0
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Labels:   []string{"Reader", "Outline", "Source"},
		OnChange: func(i int) { changes++ },
	}), 300, 100)
	if err := tt.Click("Outline"); err != nil {
		t.Fatal(err)
	}
	if selected != 1 || changes != 1 {
		t.Fatalf("the click chose %d (%d changes), want 1 (1)", selected, changes)
	}
	if err := tt.Click("Source"); err != nil {
		t.Fatal(err)
	}
	if selected != 2 || changes != 2 {
		t.Fatalf("the click chose %d (%d changes), want 2 (2)", selected, changes)
	}
}

// TestKeys moves the choice with the arrows while the control has the
// focus.
func TestKeys(t *testing.T) {
	selected := 0
	tt := ui.NewTester(frame(Props{Selected: &selected, Labels: []string{"Reader", "Outline"}}), 300, 100)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if selected != 1 {
		t.Fatalf("the Right key chose %d, want 1", selected)
	}
	tt.Key(0, ui.KeyLeft)
	if selected != 0 {
		t.Fatalf("the Left key chose %d, want 0", selected)
	}
}

// TestWraps moves the choice around, as a radio group does: the Right
// key at the end returns to the first segment.
func TestWraps(t *testing.T) {
	selected := 1
	tt := ui.NewTester(frame(Props{Selected: &selected, Labels: []string{"Reader", "Outline"}}), 300, 100)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if selected != 0 {
		t.Fatalf("the Right key wrapped to %d, want 0", selected)
	}
}
