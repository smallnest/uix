package dialog

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// frame shows an "Open" trigger button and the dialog for the props, as
// a view uses the component.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			trigger := ui.ButtonBase(c)
			trigger.Children(func() { ui.Text(c, "Open") })
			if trigger.Clicked() {
				*p.Open = true
			}
			Dialog(c, p)
		})
	}
}

// TestClosedRendersNothing draws no dialog while it is closed.
func TestClosedRendersNothing(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{Open: &open, Title: "Delete?"}), 300, 200)
	if tt.HasText("Delete?") {
		t.Fatal("closed dialog shows its title")
	}
}

// TestOpensAndCancels opens the dialog, shows its parts, and closes it
// with the cancel button.
func TestOpensAndCancels(t *testing.T) {
	open := false
	tt := ui.NewTester(frame(Props{
		Open:        &open,
		Title:       "Delete?",
		Description: "Cannot be undone.",
		Confirm:     "Delete",
		Cancel:      "Keep",
	}), 300, 200)
	if err := tt.Click("Open"); err != nil {
		t.Fatal(err)
	}
	if !open {
		t.Fatal("dialog did not open")
	}
	for _, s := range []string{"Delete?", "Cannot be undone.", "Keep"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
	if err := tt.Click("Keep"); err != nil {
		t.Fatal(err)
	}
	if open {
		t.Fatal("cancel did not close the dialog")
	}
	if tt.HasText("Delete?") {
		t.Fatal("dialog still shows after cancel")
	}
}

// TestConfirm runs OnConfirm for the confirm button and closes the
// dialog.
func TestConfirm(t *testing.T) {
	open := false
	confirmed := 0
	tt := ui.NewTester(frame(Props{
		Open:        &open,
		Title:       "Delete?",
		Confirm:     "Delete",
		Destructive: true,
		OnConfirm:   func() { confirmed++ },
	}), 300, 200)
	if err := tt.Click("Open"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Delete"); err != nil {
		t.Fatal(err)
	}
	if confirmed != 1 {
		t.Fatalf("OnConfirm ran %d times, want 1", confirmed)
	}
	if open {
		t.Fatal("confirm did not close the dialog")
	}
}
