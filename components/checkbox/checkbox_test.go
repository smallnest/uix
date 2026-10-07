package checkbox

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Checkbox(c, p)
		})
	}
}

func TestRenders(t *testing.T) {
	checked := false
	tt := ui.NewTester(frame(Props{Checked: &checked, Label: "I agree"}), 240, 100)
	if !tt.HasText("I agree") {
		t.Fatal("label missing")
	}
}

func TestClickToggles(t *testing.T) {
	checked := false
	tt := ui.NewTester(frame(Props{Checked: &checked, Label: "I agree"}), 240, 100)
	if err := tt.Click("I agree"); err != nil {
		t.Fatal(err)
	}
	if !checked {
		t.Fatal("click did not check the box")
	}
	if err := tt.Click("I agree"); err != nil {
		t.Fatal(err)
	}
	if checked {
		t.Fatal("second click did not uncheck the box")
	}
}

func TestDisabled(t *testing.T) {
	checked := false
	tt := ui.NewTester(frame(Props{Checked: &checked, Label: "No", Disabled: true}), 240, 100)
	if err := tt.Click("No"); err != nil {
		t.Fatal(err)
	}
	if checked {
		t.Fatal("disabled box was toggled")
	}
}
