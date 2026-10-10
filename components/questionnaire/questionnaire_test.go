package questionnaire

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// steps are two questions of the tests.
var steps = []Step{
	{Title: "What are you building?", Choices: []Choice{
		{Label: "A website", Value: "site"},
		{Label: "An app", Value: "app"},
	}},
	{Title: "How big is the team?", Choices: []Choice{
		{Label: "Just me", Value: "solo"},
		{Label: "Two to ten", Value: "small"},
	}},
}

// view draws the questionnaire of props in a window.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Questionnaire(c, p)
		})
	}
}

// TestShows draws the first question and its choices.
func TestShows(t *testing.T) {
	tt := ui.NewTester(view(Props{State: &State{}, Steps: steps}), 400, 300)
	for _, want := range []string{"What are you building?", "A website", "An app", "1 of 2"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestNextDisabled keeps Next from moving on while the step has no
// answer.
func TestNextDisabled(t *testing.T) {
	st := &State{}
	tt := ui.NewTester(view(Props{State: st, Steps: steps}), 400, 300)
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if st.Step != 0 {
		t.Fatalf("step %d, want 0", st.Step)
	}
}

// TestAnswers moves on for a choice and Next, and goes back for Back.
func TestAnswers(t *testing.T) {
	st := &State{}
	tt := ui.NewTester(view(Props{State: st, Steps: steps}), 400, 300)
	if err := tt.Click("A website"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if st.Step != 1 {
		t.Fatalf("step %d, want 1", st.Step)
	}
	if got := st.Answers[0]; got != "site" {
		t.Fatalf("answer %q, want site", got)
	}
	if err := tt.Click("Back"); err != nil {
		t.Fatal(err)
	}
	if st.Step != 0 {
		t.Fatalf("step %d, want 0 after Back", st.Step)
	}
}

// TestDone runs OnDone for the last step's answer and its Done button.
func TestDone(t *testing.T) {
	done := false
	st := &State{}
	tt := ui.NewTester(view(Props{State: st, Steps: steps, OnDone: func() { done = true }}), 400, 300)
	tt.Click("A website")
	tt.Click("Next")
	tt.Click("Just me")
	if err := tt.Click("Done"); err != nil {
		t.Fatal(err)
	}
	if !done {
		t.Fatal("Done did not run OnDone")
	}
}

// TestDarkMode draws the form under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(view(Props{State: &State{}, Steps: steps}), 400, 300)
	tt.SetDark(true)
	if !tt.HasText("What are you building?") {
		t.Fatal("questionnaire missing under dark mode")
	}
}
