package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/questionnaire/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.QuestionnaireView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading and the first
// question with its choices.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Questionnaire",
		"What are you building?", "A website", "An app", "A bot",
		"1 of 3",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSteps answers all three questions and reports the answers.
func TestSteps(t *testing.T) {
	tt := newTester(t)
	tt.Click("An app")
	tt.Click("Next")
	tt.Click("Just me")
	tt.Click("Next")
	tt.Click("Word of mouth")
	if err := tt.Click("Done"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Done: building an app with a team of one, found by word of mouth") {
		t.Fatalf("missing the answers in %q", tt.Texts())
	}
}

// TestBack goes back to a previous step and changes the answer.
func TestBack(t *testing.T) {
	tt := newTester(t)
	tt.Click("An app")
	tt.Click("Next")
	if err := tt.Click("Back"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("What are you building?") {
		t.Fatalf("Back did not return to the first step: %q", tt.Texts())
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Questionnaire") {
		t.Fatal("questionnaire missing under dark mode")
	}
}
