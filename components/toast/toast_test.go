package toast

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// view builds a frame with the toast viewport and a "Notify" button
// that pushes the toast fn makes.
func view(fn func(c *ui.Context) Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		Viewport(c)
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			b := ui.ButtonBase(c)
			b.Children(func() { ui.Text(c, "Notify") })
			if b.Clicked() {
				Push(c, fn(c))
			}
		})
	}
}

// TestPushShows a toast with its title and description.
func TestPushShows(t *testing.T) {
	tt := ui.NewTester(view(func(c *ui.Context) Props {
		return Props{Title: "Saved", Description: "Took effect.", Type: Success}
	}), 320, 200)
	if err := tt.Click("Notify"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"Saved", "Took effect."} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestAction runs OnAction for the action button, which closes the
// toast.
func TestAction(t *testing.T) {
	undone := 0
	tt := ui.NewTester(view(func(c *ui.Context) Props {
		return Props{Title: "Deleted", Action: "Undo", OnAction: func() { undone++ }}
	}), 320, 200)
	if err := tt.Click("Notify"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Deleted") {
		t.Fatal("toast did not show")
	}
	if err := tt.Click("Undo"); err != nil {
		t.Fatal(err)
	}
	if undone != 1 {
		t.Fatalf("OnAction ran %d times, want 1", undone)
	}
	if tt.HasText("Deleted") {
		t.Fatal("action did not close the toast")
	}
}

// TestClose closes the toast with its close button.
func TestClose(t *testing.T) {
	tt := ui.NewTester(view(func(c *ui.Context) Props {
		return Props{Title: "Note"}
	}), 320, 200)
	if err := tt.Click("Notify"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Note") {
		t.Fatal("toast did not show")
	}
	if err := tt.Click("✕"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("Note") {
		t.Fatal("close did not hide the toast")
	}
}
