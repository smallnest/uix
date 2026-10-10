package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/inputs/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.InputsView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading and the four
// controls.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Inputs",
		"Button group", "Left", "Center", "Right",
		"Input group", "Amount", "USD",
		"One-time password", "Type the code from the message.",
		"Shortcuts", "New", "Save", "Find",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestAlign chooses the Center button of the group.
func TestAlign(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Center"); err != nil {
		t.Fatal(err)
	}
	if len(view.State.Align) != 1 || view.State.Align[0] != "center" {
		t.Fatalf("Align = %v, want [center]", view.State.Align)
	}
}

// TestAmount types an amount into the field, on the row of its
// currency tag.
func TestAmount(t *testing.T) {
	tt := newTester(t)
	usd, _ := tt.Find("USD")
	tt.ClickAt(usd.X-100, usd.Y+usd.H/2)
	tt.Type("12.50")
	if view.State.Amount != "12.50" {
		t.Fatalf("Amount = %q, want 12.50", view.State.Amount)
	}
}

// TestCode fills the code cells and reports the message of the
// completion.
func TestCode(t *testing.T) {
	tt := newTester(t)
	// The cells are under the One-time password label; the first cell
	// starts at its left edge.
	heading, _ := tt.Find("One-time password")
	tt.ClickAt(heading.X+40, heading.Y+heading.H+20)
	tt.Type("123456")
	if view.State.Code.Code != "123456" {
		t.Fatalf("Code = %q, want 123456", view.State.Code.Code)
	}
	if !tt.HasText("Code 123456 verified") {
		t.Fatalf("missing the verification in %q", tt.Texts())
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	for _, s := range []string{"Inputs", "Amount", "Shortcuts"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in the dark frame %q", s, tt.Texts())
		}
	}
}
