// Package view builds the questionnaire example: a window of MyGo
// native UI that shows the questionnaire component of uix alone, as an
// onboarding form of three questions. The progress bar shows how far
// along it is, each question offers its choices as selectable rows, and
// Back and Next move through them; the Done button of the last step
// names the answers at the bottom. The main package shows it in a
// window; the tests and the snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/questionnaire"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 420

// State is the state of the example; the controls edit it in place.
var State = QuestionnaireState{}

// QuestionnaireState holds the example: the form and the answers the
// last completion reported.
type QuestionnaireState struct {
	// Form is where the questionnaire is.
	Form questionnaire.State
	// Done is the message of the last completion.
	Done string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = QuestionnaireState{} }

// steps are the three questions of the example.
var steps = []questionnaire.Step{
	{Title: "What are you building?", Choices: []questionnaire.Choice{
		{Label: "A website", Value: "site"},
		{Label: "An app", Value: "app"},
		{Label: "A bot", Value: "bot"},
	}},
	{Title: "How big is the team?", Description: "You can work alone too.",
		Choices: []questionnaire.Choice{
			{Label: "Just me", Value: "solo"},
			{Label: "Two to ten", Value: "small"},
			{Label: "More than ten", Value: "large"},
		}},
	{Title: "How will people find it?", Choices: []questionnaire.Choice{
		{Label: "Search", Value: "search"},
		{Label: "Word of mouth", Value: "word"},
		{Label: "Advertising", Value: "ads"},
	}},
}

// QuestionnaireView draws the example: the form, and the message of the
// last completion under it.
func QuestionnaireView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Questionnaire").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "A form of questions, from uix.").TextColor(t.TextMuted)
		})
		questionnaire.Questionnaire(c, questionnaire.Props{
			State:  &State.Form,
			Steps:  steps,
			OnDone: func() { State.Done = answersText() },
		})
		if State.Done != "" {
			ui.Text(c, "Done: "+State.Done).FontSize(t.FontSize * 0.85).TextColor(t.Success)
		}
	})
}

// answersText is what the answers name, as "a website for a team of
// one, found by word of mouth".
func answersText() string {
	labels := map[string]string{
		"site": "a website", "app": "an app", "bot": "a bot",
		"solo": "a team of one", "small": "a team of two to ten", "large": "a team of many",
		"search": "search", "word": "word of mouth", "ads": "advertising",
	}
	out := make([]string, 0, 3)
	for _, a := range State.Form.Answers {
		out = append(out, labels[a])
	}
	return "building " + out[0] + " with " + out[1] + ", found by " + out[2]
}
