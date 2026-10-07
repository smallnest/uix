package field

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Field(c, p, func() { ui.Text(c, "the control") })
		})
	}
}

func TestLabelHint(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Name", Hint: "How people find you."}), 240, 120)
	for _, s := range []string{"Name", "How people find you.", "the control"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

func TestRequired(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Email", Required: true}), 240, 120)
	if !tt.HasText("*") {
		t.Fatal("required field has no star")
	}
}

func TestErrorReplacesHint(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Email", Hint: "helper", Error: "not an email"}), 240, 120)
	if !tt.HasText("not an email") {
		t.Fatal("error message missing")
	}
	if tt.HasText("helper") {
		t.Fatal("hint shows while there is an error")
	}
}
