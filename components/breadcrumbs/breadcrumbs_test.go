package breadcrumbs

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Breadcrumbs(c, p)
		})
	}
}

// path is the path the tests draw, with the FAQ item last, where the
// path is.
var path = []string{"Home", "Docs", "FAQ"}

// TestRenders draws every item of the path.
func TestRenders(t *testing.T) {
	chosen := 1
	tt := ui.NewTester(frame(Props{Items: path, Chosen: &chosen}), 400, 80)
	for _, s := range path {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestChooses sets the chosen item when a link is clicked.
func TestChooses(t *testing.T) {
	chosen := -1
	tt := ui.NewTester(frame(Props{Items: path, Chosen: &chosen}), 400, 80)
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if chosen != 1 {
		t.Fatalf("chosen %d, want the Docs item", chosen)
	}
}

// TestLastItem keeps the choice when the last item is clicked, as it is
// not a link.
func TestLastItem(t *testing.T) {
	chosen := -1
	tt := ui.NewTester(frame(Props{Items: path, Chosen: &chosen}), 400, 80)
	if err := tt.Click("FAQ"); err != nil {
		t.Fatal(err)
	}
	if chosen != -1 {
		t.Fatalf("chosen %d, want no choice", chosen)
	}
}

// TestDisabled keeps the items from being chosen.
func TestDisabled(t *testing.T) {
	chosen := -1
	tt := ui.NewTester(frame(Props{Items: path, Chosen: &chosen, Disabled: true}), 400, 80)
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if chosen != -1 {
		t.Fatalf("chosen %d, want no choice", chosen)
	}
}

// TestDarkMode draws the path under the dark appearance too.
func TestDarkMode(t *testing.T) {
	chosen := 1
	tt := ui.NewTester(frame(Props{Items: path, Chosen: &chosen}), 400, 80)
	tt.SetDark(true)
	if !tt.HasText("Home") {
		t.Fatal("breadcrumbs missing under dark mode")
	}
}
