package radio

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props[string]) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Group(c, p)
		})
	}
}

// TestRenders shows every label of the group.
func TestRenders(t *testing.T) {
	selected := "system"
	tt := ui.NewTester(frame(Props[string]{
		Selected: &selected,
		Options: []Option[string]{
			{Value: "light", Label: "Light"},
			{Value: "dark", Label: "Dark"},
			{Value: "system", Label: "System"},
		},
	}), 300, 120)
	for _, s := range []string{"Light", "Dark", "System"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestChooses selects the option clicked.
func TestChooses(t *testing.T) {
	selected := "light"
	tt := ui.NewTester(frame(Props[string]{
		Selected: &selected,
		Options: []Option[string]{
			{Value: "light", Label: "Light"},
			{Value: "dark", Label: "Dark"},
		},
	}), 300, 100)
	if err := tt.Click("Dark"); err != nil {
		t.Fatal(err)
	}
	if selected != "dark" {
		t.Fatalf("selected %q, want dark", selected)
	}
}

// TestArrows choose with the arrows, which keep the group one stop of
// Tab.
func TestArrows(t *testing.T) {
	selected := "light"
	tt := ui.NewTester(frame(Props[string]{
		Selected: &selected,
		Options: []Option[string]{
			{Value: "light", Label: "Light"},
			{Value: "dark", Label: "Dark"},
		},
	}), 300, 100)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyDown)
	if selected != "dark" {
		t.Fatalf("after the arrows selected %q, want dark", selected)
	}
	tt.Key(0, ui.KeyHome)
	if selected != "light" {
		t.Fatalf("after Home selected %q, want light", selected)
	}
}
