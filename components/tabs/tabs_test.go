package tabs

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Tabs(c, p)
		})
	}
}

// TestRenders shows every label of the list.
func TestRenders(t *testing.T) {
	selected := 0
	tt := ui.NewTester(frame(Props{Selected: &selected, Labels: []string{"General", "Appearance", "About"}}), 300, 60)
	for _, s := range []string{"General", "Appearance", "About"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestClick chooses a tab with a click, which the view reads from the
// selected index.
func TestClick(t *testing.T) {
	selected := 0
	tt := ui.NewTester(frame(Props{Selected: &selected, Labels: []string{"General", "Appearance"}}), 300, 60)
	if err := tt.Click("Appearance"); err != nil {
		t.Fatal(err)
	}
	if selected != 1 {
		t.Fatalf("selected %d, want 1", selected)
	}
}

// TestKeys choose with the arrows, which keep the list one stop of Tab.
func TestKeys(t *testing.T) {
	selected := 0
	tt := ui.NewTester(frame(Props{Selected: &selected, Labels: []string{"General", "Appearance", "About"}}), 300, 60)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if selected != 1 {
		t.Fatalf("after the arrows selected %d, want 1", selected)
	}
	tt.Key(0, ui.KeyEnd)
	if selected != 2 {
		t.Fatalf("after End selected %d, want 2", selected)
	}
}
