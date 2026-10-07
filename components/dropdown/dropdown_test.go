package dropdown

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Dropdown(c, p)
		})
	}
}

// TestRenders shows the trigger of the menu.
func TestRenders(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{Trigger: "Account", Open: &open, Items: []Item{{Label: "Sign out"}}}), 300, 100)
	if !tt.HasText("Account") {
		t.Fatal("trigger missing")
	}
	if open {
		t.Fatal("the menu opened on its own")
	}
}

// TestOpensAndChooses opens the menu, chooses an item, and runs its
// OnClick; the menu closes either way.
func TestOpensAndChooses(t *testing.T) {
	open := false
	chosen := ""
	tt := ui.NewTester(frame(Props{Trigger: "Account", Open: &open, Items: []Item{
		{Label: "View profile", OnClick: func() { chosen = "profile" }},
		{Separator: true},
		{Label: "Sign out", OnClick: func() { chosen = "out" }},
	}}), 300, 140)
	if err := tt.Click("Account"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("the trigger did not open the menu")
	}
	if !tt.HasText("Sign out") {
		t.Fatal("the items of the menu are missing")
	}
	if err := tt.Click("Sign out"); err != nil {
		t.Fatal(err)
	}
	if chosen != "out" {
		t.Fatalf("chosen %q, want out", chosen)
	}
	if open {
		t.Fatal("the choice did not close the menu")
	}
}

// TestDisabledItem keeps a disabled item inert: it neither runs its
// OnClick nor closes the menu.
func TestDisabledItem(t *testing.T) {
	open := false
	chosen := 0
	tt := ui.NewTester(frame(Props{Trigger: "Menu", Open: &open, Items: []Item{
		{Label: "One", OnClick: func() { chosen = 1 }},
		{Separator: true},
		{Label: "Two", Disabled: true, OnClick: func() { chosen = 2 }},
	}}), 300, 140)
	if err := tt.Click("Menu"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Two"); err != nil {
		t.Fatal(err)
	}
	if chosen != 0 {
		t.Fatal("a disabled item acted")
	}
	if !open {
		t.Fatal("a disabled item closed the menu")
	}
}

// TestTriggerToggles closes an open menu with a second press of the
// trigger, the frame the popover takes the press.
func TestTriggerToggles(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{Trigger: "Menu", Open: &open, Items: []Item{{Label: "Item"}}}), 300, 100)
	if err := tt.Click("Menu"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("the trigger did not open the menu")
	}
	if err := tt.Click("Menu"); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Fatal("the trigger did not close the menu")
	}
}

// TestClosesFromOutside closes the menu on a press outside it.
func TestClosesFromOutside(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{Trigger: "Menu", Open: &open, Items: []Item{{Label: "Item"}}}), 300, 100)
	if err := tt.Click("Menu"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("the trigger did not open the menu")
	}
	tt.ClickAt(5, 5)
	if open {
		t.Fatal("an outside press did not close the menu")
	}
}
