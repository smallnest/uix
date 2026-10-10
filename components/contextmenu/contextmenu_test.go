package contextmenu

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

// view draws the box of the context menu across the window, with a
// child to right-click on.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			ui.Box(c).Fill().Children(func() {
				ContextMenu(c, p).Fill().Children(func() {
					ui.Text(c, "Right-click me")
				})
			})
		})
	}
}

// TestRenders draws the box and its child.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(view(Props{}), 300, 200)
	if !tt.HasText("Right-click me") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestOpens shows the items for a right-click, with "-" for the
// separator.
func TestOpens(t *testing.T) {
	tt := ui.NewTester(view(Props{Items: []Item{
		{Label: "Rename"},
		{Separator: true},
		{Label: "Delete"},
	}}), 300, 200)
	if err := tt.RightClick("Right-click me"); err != nil {
		t.Fatal(err)
	}
	if got := tt.Menu(); !slices.Equal(got, []string{"Rename", "-", "Delete"}) {
		t.Fatalf("menu %q, want the items", got)
	}
}

// TestChooses runs the action of an item when it is chosen.
func TestChooses(t *testing.T) {
	acted := ""
	tt := ui.NewTester(view(Props{Items: []Item{
		{Label: "Rename", Action: func() { acted = "rename" }},
		{Label: "Delete", Action: func() { acted = "delete" }},
	}}), 300, 200)
	if err := tt.RightClick("Right-click me"); err != nil {
		t.Fatal(err)
	}
	if err := tt.ChooseMenuItem("Delete"); err != nil {
		t.Fatal(err)
	}
	if acted != "delete" {
		t.Fatalf("acted %q, want delete", acted)
	}
}

// TestDisabled keeps an item from being chosen.
func TestDisabled(t *testing.T) {
	acted := false
	tt := ui.NewTester(view(Props{Items: []Item{
		{Label: "Delete", Disabled: true, Action: func() { acted = true }},
	}}), 300, 200)
	if err := tt.RightClick("Right-click me"); err != nil {
		t.Fatal(err)
	}
	if err := tt.ChooseMenuItem("Delete"); err == nil {
		t.Fatal("a disabled item was chosen")
	}
	if acted {
		t.Fatal("a disabled item acted")
	}
}

// TestChecked shows a checked item as a check box, which the system
// draws beside its label.
func TestChecked(t *testing.T) {
	tt := ui.NewTester(view(Props{Items: []Item{
		{Label: "Bold", Checked: true},
	}}), 300, 200)
	if err := tt.RightClick("Right-click me"); err != nil {
		t.Fatal(err)
	}
	if got := tt.Menu(); !slices.Equal(got, []string{"Bold"}) {
		t.Fatalf("menu %q, want the item", got)
	}
}
