package gridview

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			GridView(c, p).Grow(1)
		})
	}
}

func items(n int) func(c *ui.Context, i int) {
	return func(c *ui.Context, i int) {
		ui.Text(c, "Item "+fmt.Sprint(i))
	}
}

// TestRenders shows the items in the grid, all at once when they fit.
func TestRenders(t *testing.T) {
	var gs ui.GridState
	tt := ui.NewTester(frame(Props{
		State:    &gs,
		Count:    8,
		MinWidth: 80,
		Height:   60,
		Item:     items(8),
	}), 400, 240)
	for i := 0; i < 8; i++ {
		if !tt.HasText("Item " + fmt.Sprint(i)) {
			t.Fatalf("missing Item %d in %q", i, tt.Texts())
		}
	}
}

// TestChooses selects an item with a click, which runs OnChoose, and the
// chosen cell shows its choice.
func TestChooses(t *testing.T) {
	var gs ui.GridState
	chosen := -1
	gs.Selected = &chosen
	changes := 0
	tt := ui.NewTester(frame(Props{
		State:    &gs,
		Count:    4,
		MinWidth: 80,
		Height:   60,
		Item:     items(4),
		OnChoose: func(i int) { changes++ },
	}), 400, 240)
	if err := tt.Click("Item 2"); err != nil {
		t.Fatal(err)
	}
	if chosen != 2 || changes != 1 {
		t.Fatalf("the click chose %d (%d changes), want 2 (1)", chosen, changes)
	}
}

// TestKeys moves the choice with the arrows while the grid has the focus:
// the first arrow picks the first item, then Right steps by a column.
func TestKeys(t *testing.T) {
	var gs ui.GridState
	chosen := -1
	gs.Selected = &chosen
	tt := ui.NewTester(frame(Props{
		State:    &gs,
		Count:    4,
		MinWidth: 80,
		Height:   60,
		Item:     items(4),
	}), 400, 240)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if chosen != 0 {
		t.Fatalf("the first Right key chose %d, want 0", chosen)
	}
	tt.Key(0, ui.KeyRight)
	if chosen != 1 {
		t.Fatalf("the Right key chose %d, want 1", chosen)
	}
}
