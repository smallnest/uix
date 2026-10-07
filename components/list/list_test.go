package list

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame draws the list of p across a padded window.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			List(c, p).Grow(1)
		})
	}
}

// TestRenders draws the rows of the list.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Items: []string{"Alpha", "Beta", "Gamma"}}), 300, 200)
	for _, s := range []string{"Alpha", "Beta", "Gamma"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestChooses picks the row clicked, and runs OnChoose with it.
func TestChooses(t *testing.T) {
	var st ui.ListState
	chosen := -1
	st.Selected = &chosen
	picked := -1
	p := Props{State: &st, Items: []string{"Alpha", "Beta", "Gamma"}, OnChoose: func(i int) { picked = i }}
	tt := ui.NewTester(frame(p), 300, 200)
	if err := tt.Click("Beta"); err != nil {
		t.Fatal(err)
	}
	if chosen != 1 {
		t.Fatalf("chosen %d, want 1", chosen)
	}
	if picked != 1 {
		t.Fatalf("picked %d, want 1", picked)
	}
}

// TestKeys moves the choice with the arrow keys after a click.
func TestKeys(t *testing.T) {
	var st ui.ListState
	chosen := -1
	st.Selected = &chosen
	tt := ui.NewTester(frame(Props{State: &st, Items: []string{"Alpha", "Beta", "Gamma"}}), 300, 200)
	if err := tt.Click("Alpha"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyDown)
	if chosen != 1 {
		t.Fatalf("chosen %d, want the arrow to move it to 1", chosen)
	}
	tt.Key(0, ui.KeyEnd)
	if chosen != 2 {
		t.Fatalf("chosen %d, want End to move it to 2", chosen)
	}
}

// TestRow draws the rows of a custom builder.
func TestRow(t *testing.T) {
	var st ui.ListState
	p := Props{State: &st, Items: []string{"a", "b"}, Row: func(c *ui.Context, i int) {
		ui.Row(c).Height(32).PaddingX(12).Children(func() {
			ui.Text(c, "Item "+string(rune('1'+i))).Grow(1)
		})
	}}
	tt := ui.NewTester(frame(p), 300, 200)
	if !tt.HasText("Item 1") || !tt.HasText("Item 2") {
		t.Fatalf("texts %q, want the custom rows", tt.Texts())
	}
}

// TestDisabled keeps the list from choosing rows.
func TestDisabled(t *testing.T) {
	var st ui.ListState
	chosen := -1
	st.Selected = &chosen
	tt := ui.NewTester(frame(Props{State: &st, Items: []string{"Alpha", "Beta"}, Disabled: true}), 300, 200)
	if err := tt.Click("Beta"); err != nil {
		t.Fatal(err)
	}
	if chosen != -1 {
		t.Fatalf("chosen %d, want the list disabled to keep it -1", chosen)
	}
}

// TestManyRows scrolls a list longer than the window, which builds only
// the rows in view.
func TestManyRows(t *testing.T) {
	var st ui.ListState
	items := make([]string, 500)
	for i := range items {
		items[i] = "Row"
	}
	tt := ui.NewTester(frame(Props{State: &st, Items: items}), 300, 200)
	texts := tt.Texts()
	if len(texts) == 0 || len(texts) > 50 {
		t.Fatalf("drew %d rows, want only those in view", len(texts))
	}
	if first, last := st.Visible(); first != 0 || last < 5 {
		t.Fatalf("visible %d..%d, want the list at its start", first, last)
	}
}
