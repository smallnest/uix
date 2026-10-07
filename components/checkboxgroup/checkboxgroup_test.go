package checkboxgroup

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/checkbox"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			CheckboxGroup(c, p)
		})
	}
}

func notifications(mail, cal *bool) func(c *ui.Context) {
	return func(c *ui.Context) {
		checkbox.Checkbox(c, checkbox.Props{Checked: mail, Label: "Mail"})
		checkbox.Checkbox(c, checkbox.Props{Checked: cal, Label: "Calendar"})
	}
}

// TestRenders shows the label and the boxes.
func TestRenders(t *testing.T) {
	mail, cal := false, false
	tt := ui.NewTester(frame(Props{
		Label:    "Notifications",
		Children: notifications(&mail, &cal),
	}), 300, 200)
	for _, s := range []string{"Notifications", "Mail", "Calendar"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestChecksAll checks all the boxes with a click on the group, and a
// second clears them all.
func TestChecksAll(t *testing.T) {
	mail, cal := false, false
	tt := ui.NewTester(frame(Props{
		Label:    "Notifications",
		Children: notifications(&mail, &cal),
	}), 300, 200)
	if err := tt.Click("Notifications"); err != nil {
		t.Fatal(err)
	}
	if !mail || !cal {
		t.Fatalf("the click did not check all: %v %v", mail, cal)
	}
	if err := tt.Click("Notifications"); err != nil {
		t.Fatal(err)
	}
	if mail || cal {
		t.Fatalf("the second click did not clear all: %v %v", mail, cal)
	}
}

// TestMixed checks one box, so the group is mixed, and the click on the
// group then checks all.
func TestMixed(t *testing.T) {
	mail, cal := false, false
	tt := ui.NewTester(frame(Props{
		Label:    "Notifications",
		Children: notifications(&mail, &cal),
	}), 300, 200)
	if err := tt.Click("Mail"); err != nil {
		t.Fatal(err)
	}
	if !mail || cal {
		t.Fatal("the box did not check on its own")
	}
	if err := tt.Click("Notifications"); err != nil {
		t.Fatal(err)
	}
	if !mail || !cal {
		t.Fatalf("the group did not check the rest: %v %v", mail, cal)
	}
}
