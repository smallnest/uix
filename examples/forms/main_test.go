package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/forms/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() { view.Form = view.FormState{Volume: 60} }

// bg returns the color of the window background, where the theme shows.
func bg(tt *ui.Tester) (uint8, uint8, uint8) {
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		return 0, 0, 0
	}
	r, g, b, _ := img.At(2, 2).RGBA()
	return uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)
}

// TestRenders draws the form headless and checks that every control came
// through: MyGo renders frames in memory, no window needed.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FormsView, view.Width, view.Height)
	for _, s := range []string{
		"Sign up", "Name", "How people find you.", "Email", "Country",
		"Choose a country", "Message", "What is your project about?",
		"Volume", "60%", "Dark mode", "I agree to the terms", "Create account",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestMessage types into the message text area and checks the shared
// value changed.
func TestMessage(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FormsView, view.Width, view.Height)
	tt.ClickAt(40, 330)
	tt.Type("hello")
	if view.Form.Message != "hello" {
		t.Fatalf("message %q, want hello", view.Form.Message)
	}
}

// TestSwitchDarkens proves the dark-mode switch works: it restyles the
// form from light to dark and back.
func TestSwitchDarkens(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FormsView, view.Width, view.Height)
	if r, _, _ := bg(tt); r < 200 {
		t.Fatalf("starts dark (r=%d), want light", r)
	}
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if r, _, _ := bg(tt); r > 60 {
		t.Fatalf("switch did not darken the form (r=%d)", r)
	}
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if r, _, _ := bg(tt); r < 200 {
		t.Fatalf("switch did not lighten the form (r=%d)", r)
	}
}

// TestValidation submits an empty form and checks that the required
// fields say why they fail, then fills the form and submits it clear.
func TestValidation(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FormsView, view.Width, view.Height)
	if err := tt.Click("Create account"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"Name is required.", "Enter a valid email.", "You must agree to the terms."} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
	if view.Form.Submitted {
		t.Fatal("an empty form was accepted")
	}

	view.Form.Name = "Ada"
	view.Form.Email = "ada@example.com"
	view.Form.Agree = true
	if err := tt.Click("Create account"); err != nil {
		t.Fatal(err)
	}
	if !view.Form.Submitted {
		t.Fatal("a clear form was rejected")
	}
	if !tt.HasText("Welcome!") {
		t.Fatal("success message missing")
	}
	if tt.HasText("Name is required.") {
		t.Fatal("the errors did not clear")
	}
}

// TestInteractions drives the form the way a user would: toggle the
// switch and the check box, choose a country, fill the required fields,
// submit.
func TestInteractions(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FormsView, view.Width, view.Height)
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if !view.Form.Dark {
		t.Fatal("switch did not toggle")
	}
	if err := tt.Click("I agree to the terms"); err != nil {
		t.Fatal(err)
	}
	if !view.Form.Agree {
		t.Fatal("check box did not toggle")
	}
	if err := tt.Click("Choose a country"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Japan"); err != nil {
		t.Fatal(err)
	}
	if view.Form.Country != "Japan" {
		t.Fatalf("selected %q, want Japan", view.Form.Country)
	}
	view.Form.Name = "Ada"
	view.Form.Email = "ada@example.com"
	if err := tt.Click("Create account"); err != nil {
		t.Fatal(err)
	}
	if !view.Form.Submitted {
		t.Fatal("submit did not fire")
	}
	if !tt.HasText("Welcome!") {
		t.Fatal("success message missing")
	}
}

// TestDarkMode draws the form under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.FormsView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Sign up") {
		t.Fatal("dark form missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
