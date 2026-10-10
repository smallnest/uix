package empty

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestShows draws the empty state: the title, the description and the
// action.
func TestShows(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Empty(c, Props{
				Title:       "No results",
				Description: "Try another search.",
				ActionLabel: "Clear search",
			})
		})
	}, 320, 220)
	for _, want := range []string{"No results", "Try another search.", "Clear search"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestAction runs the action for a click on the button.
func TestAction(t *testing.T) {
	ran := false
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Empty(c, Props{
				Title:       "No results",
				ActionLabel: "Clear search",
				OnAction:    func() { ran = true },
			})
		})
	}, 320, 220)
	if err := tt.Click("Clear search"); err != nil {
		t.Fatal(err)
	}
	if !ran {
		t.Fatal("the action did not run")
	}
}

// TestNoAction draws the state without the button.
func TestNoAction(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Empty(c, Props{Title: "Nothing here"})
		})
	}, 320, 220)
	if tt.HasText("Clear search") {
		t.Fatalf("an action drew without a label: %q", tt.Texts())
	}
	if !tt.HasText("Nothing here") {
		t.Fatalf("missing the title in %q", tt.Texts())
	}
}
