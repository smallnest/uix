package tree

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws the tree of p in a padded column.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Tree(c, p)
		})
	}
}

// items returns a small tree: a branch "src" with two leaves, and a
// leaf "go.mod".
func items(src bool) []Item {
	return []Item{
		{Label: "src", Open: &src, Children: []Item{
			{Label: "main.go"},
			{Label: "app.go"},
		}},
		{Label: "go.mod"},
	}
}

// TestRenders draws the labels of the tree.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Items: items(false)}), 300, 200)
	for _, s := range []string{"src", "go.mod"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	if tt.HasText("main.go") {
		t.Fatal("a closed branch shows its leaves")
	}
}

// TestOpens opens a branch with the Right key, which focuses its item.
func TestOpens(t *testing.T) {
	tt := ui.NewTester(frame(Props{Items: items(false)}), 300, 200)
	if err := tt.Click("src"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("main.go") {
		t.Fatal("the branch was open at first")
	}
	tt.Key(0, ui.KeyRight)
	if !tt.HasText("main.go") {
		t.Fatalf("Right did not open the branch: %q", tt.Texts())
	}
	tt.Key(0, ui.KeyLeft)
	if tt.HasText("main.go") {
		t.Fatal("Left did not close the branch")
	}
}

// TestChooses sets Chosen and runs the action of the node clicked.
func TestChooses(t *testing.T) {
	items := items(true)
	picked := ""
	items[1].Action = func() { picked = "root" }
	chosen := ""
	tt := ui.NewTester(frame(Props{Items: items, Chosen: &chosen}), 300, 200)
	if err := tt.Click("go.mod"); err != nil {
		t.Fatal(err)
	}
	if chosen != "go.mod" {
		t.Fatalf("chosen %q, want go.mod", chosen)
	}
	if picked != "root" {
		t.Fatalf("picked %q, want root", picked)
	}
}

// TestDisabled keeps the tree from choosing or opening nodes.
func TestDisabled(t *testing.T) {
	chosen := ""
	tt := ui.NewTester(frame(Props{Items: items(false), Chosen: &chosen, Disabled: true}), 300, 200)
	if err := tt.Click("src"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyRight)
	if chosen != "" || tt.HasText("main.go") {
		t.Fatalf("a disabled tree chose %q or opened", chosen)
	}
}

// TestDarkMode draws the tree under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Items: items(false)}), 300, 200)
	tt.SetDark(true)
	if !tt.HasText("src") {
		t.Fatal("tree missing under dark mode")
	}
}
