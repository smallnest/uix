package switcher

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Switch(c, p)
		})
	}
}

func TestRenders(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Dark mode"}), 240, 100)
	if !tt.HasText("Dark mode") {
		t.Fatal("label missing")
	}
}

func TestClickToggles(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Dark mode"}), 240, 100)
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if !on {
		t.Fatal("click did not turn the switch on")
	}
}

func TestChangedReports(t *testing.T) {
	on, reported := false, false
	tt := ui.NewTester(frame(Props{On: &on, Label: "Dark mode", Changed: func(v bool) { reported = v }}), 240, 100)
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if !reported {
		t.Fatal("Changed did not report the toggle")
	}
}

func TestDisabled(t *testing.T) {
	on := false
	tt := ui.NewTester(frame(Props{On: &on, Label: "No", Disabled: true}), 240, 100)
	if err := tt.Click("No"); err != nil {
		t.Fatal(err)
	}
	if on {
		t.Fatal("disabled switch was toggled")
	}
}
