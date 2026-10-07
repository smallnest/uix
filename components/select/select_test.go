package selector

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Select(c, p)
		})
	}
}

func TestRendersPlaceholder(t *testing.T) {
	selected := ""
	tt := ui.NewTester(frame(Props{
		Selected:    &selected,
		Options:     []string{"China", "Japan", "Germany"},
		Placeholder: "Choose a country",
	}), 300, 100)
	if !tt.HasText("Choose a country") {
		t.Fatal("placeholder missing")
	}
}

func TestChooses(t *testing.T) {
	selected := ""
	tt := ui.NewTester(frame(Props{
		Selected:    &selected,
		Options:     []string{"China", "Japan", "Germany"},
		Placeholder: "Choose a country",
	}), 300, 100)
	if err := tt.Click("Choose a country"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Japan"); err != nil {
		t.Fatal(err)
	}
	if selected != "Japan" {
		t.Fatalf("selected %q, want Japan", selected)
	}
	if !tt.HasText("Japan") {
		t.Fatal("trigger does not show the choice")
	}
}

func TestShowsChoice(t *testing.T) {
	selected := "Germany"
	tt := ui.NewTester(frame(Props{
		Selected: &selected,
		Options:  []string{"China", "Japan", "Germany"},
	}), 300, 100)
	if !tt.HasText("Germany") {
		t.Fatal("trigger does not show the choice")
	}
}

func TestDisabled(t *testing.T) {
	selected := ""
	tt := ui.NewTester(frame(Props{
		Selected:    &selected,
		Options:     []string{"China", "Japan"},
		Placeholder: "Choose",
		Disabled:    true,
	}), 300, 100)
	if err := tt.Click("Choose"); err != nil {
		t.Fatal(err)
	}
	if selected != "" {
		t.Fatal("disabled select chose an option")
	}
}
