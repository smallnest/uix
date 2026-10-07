package editabletext

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(value *string, onChange func(string)) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			EditableText(c, Props{Value: value, OnChange: onChange})
		})
	}
}

// TestRenames edits the value in place: a double click shows a field,
// typing over the text before the extension, and Tab keeps what was
// typed.
func TestRenames(t *testing.T) {
	title := "Untitled.md"
	changes := 0
	tt := ui.NewTester(frame(&title, func(v string) { changes++ }), 400, 300)
	r, ok := tt.Find("Untitled.md")
	if !ok {
		t.Fatal("no title")
	}
	tt.ClickAt(r.X+5, r.Y+5)
	tt.ClickAt(r.X+5, r.Y+5)
	tt.Type("Notes")
	tt.Key(0, ui.KeyTab)
	if title != "Notes.md" {
		t.Fatalf("after Tab the title is %q, want Notes.md", title)
	}
	if changes != 1 {
		t.Fatalf("OnChange ran %d times, want 1", changes)
	}
}

// TestEnterKeeps commits with Enter, which edits a text that a click
// has chosen, over the text before the extension.
func TestEnterKeeps(t *testing.T) {
	title := "Untitled.md"
	changes := 0
	tt := ui.NewTester(frame(&title, func(v string) { changes++ }), 400, 300)
	if err := tt.Click("Untitled.md"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyEnter)
	tt.Type("Plan")
	tt.Key(0, ui.KeyEnter)
	if title != "Plan.md" {
		t.Fatalf("after Enter the title is %q, want Plan.md", title)
	}
	if changes != 1 {
		t.Fatalf("OnChange ran %d times, want 1", changes)
	}
}
