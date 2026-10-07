package togglegroup

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			ToggleGroup(c, p)
		})
	}
}

// items are the toggles of the tests, the middle one carrying a tip.
func items(bold, italic, underline *bool) []Item {
	return []Item{
		{On: bold, Label: "B", Tip: "Bold"},
		{On: italic, Label: "I"},
		{On: underline, Label: "U"},
	}
}

// TestRenders draws every toggle of the group.
func TestRenders(t *testing.T) {
	b, i, u := false, false, false
	tt := ui.NewTester(frame(Props{Items: items(&b, &i, &u)}), 300, 100)
	for _, s := range []string{"B", "I", "U"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestToggles presses a toggle with a click.
func TestToggles(t *testing.T) {
	b, i, u := false, false, false
	tt := ui.NewTester(frame(Props{Items: items(&b, &i, &u)}), 300, 100)
	if err := tt.Click("B"); err != nil {
		t.Fatal(err)
	}
	if !b {
		t.Fatal("the click did not press Bold")
	}
}

// TestSeveral leaves more than one toggle on at once.
func TestSeveral(t *testing.T) {
	b, i, u := false, false, false
	tt := ui.NewTester(frame(Props{Items: items(&b, &i, &u)}), 300, 100)
	if err := tt.Click("B"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("I"); err != nil {
		t.Fatal(err)
	}
	if !b || !i {
		t.Fatal("B and I did not both stay on")
	}
}

// TestTip shows the tip of a toggle as the pointer rests on it.
func TestTip(t *testing.T) {
	b, i, u := false, false, false
	tt := ui.NewTester(frame(Props{Items: items(&b, &i, &u)}), 300, 100)
	r, ok := tt.Find("B")
	if !ok {
		t.Fatal("no Bold toggle")
	}
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	time.Sleep(650 * time.Millisecond)
	tt.Frame()
	if !tt.HasText("Bold") {
		t.Fatal("the tip did not show")
	}
}

// TestDisabled keeps the toggles from being pressed.
func TestDisabled(t *testing.T) {
	b, i, u := false, false, false
	tt := ui.NewTester(frame(Props{Items: items(&b, &i, &u), Disabled: true}), 300, 100)
	if err := tt.Click("B"); err != nil {
		t.Fatal(err)
	}
	if b {
		t.Fatal("a disabled group was pressed")
	}
}

// TestDarkMode draws the group under the dark appearance too.
func TestDarkMode(t *testing.T) {
	b, i, u := false, false, false
	tt := ui.NewTester(frame(Props{Items: items(&b, &i, &u)}), 300, 100)
	tt.SetDark(true)
	if !tt.HasText("B") {
		t.Fatal("group missing under dark mode")
	}
}
