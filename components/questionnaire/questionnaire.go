// Package questionnaire provides a shadcn/ui-style Questionnaire: a
// form of questions, one at a time, as onboarding asks them. A progress
// bar shows how far along it is, each step offers its choices as
// selectable rows, and Back and Next move through the steps. The last
// step's button says Done, and runs OnDone.
//
//	q := &questionnaire.State{}
//	questionnaire.Questionnaire(c, questionnaire.Props{
//		State: q,
//		Steps: []questionnaire.Step{
//			{Title: "What are you building?",
//				Choices: []questionnaire.Choice{
//					{Label: "A website", Value: "site"},
//					{Label: "An app", Value: "app"},
//				}},
//			{Title: "How big is the team?",
//				Choices: []questionnaire.Choice{
//					{Label: "Just me", Value: "solo"},
//					{Label: "Two to ten", Value: "small"},
//				}},
//		},
//		OnDone: app.done,
//	})
package questionnaire

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Choice is one answer a step offers.
type Choice struct {
	// Label is the text of the answer.
	Label string
	// Value is what the app sees as the answer.
	Value string
}

// Step is one question of the form.
type Step struct {
	// Title is the question.
	Title string
	// Description is the help under the title.
	Description string
	// Choices are the answers offered, one under the other.
	Choices []Choice
}

// State is where the form is, kept by the app.
type State struct {
	// Step is the question shown, from 0.
	Step int
	// Answers holds the chosen Value of each step, "" for a step not
	// yet answered.
	Answers []string
}

// Props describes the form to draw.
type Props struct {
	// State is where the form is, for the app to keep.
	State *State
	// Steps are the questions, in order.
	Steps []Step
	// OnDone runs when the last step's Done is chosen.
	OnDone func()
}

// Questionnaire draws the form for props and returns the card of it, so
// a view can chain more calls on it. Back and Next sit at the bottom,
// right-aligned; Next is disabled until the step has an answer.
func Questionnaire(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	st := p.State
	if st == nil {
		st = &State{}
	}
	if len(st.Answers) != len(p.Steps) {
		st.Answers = make([]string, len(p.Steps))
	}
	last := len(p.Steps) - 1
	step := st.Step
	if step > last {
		step = last
	}
	return ui.Column(c).FillWidth().Gap(t.Space(3)).Padding(t.Space(4)).
		Background(t.Background).Border(1, t.Border).Radius(t.Radius).Children(func() {
		// The progress bar, and the count beside it.
		ui.Row(c).FillWidth().Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
			ui.Progress(c, float64(step+1)/float64(len(p.Steps))).Grow(1)
			ui.Text(c, questionCount(step, len(p.Steps))).SingleLine().
				FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		})
		// The question, and the choices, of which the selected one shows
		// in the accent color.
		ui.Column(c).Gap(t.Space(2)).Children(func() {
			ui.Text(c, p.Steps[step].Title).Bold()
			if p.Steps[step].Description != "" {
				ui.Text(c, p.Steps[step].Description).TextColor(t.TextMuted)
			}
			choiceList(c, t, st, p.Steps[step].Choices)
		})
		// Back and Next; Next on the last step is Done.
		ui.Row(c).Justify(ui.End).Gap(t.Space(1.5)).Children(func() {
			if step > 0 {
				backButton(c, t, st)
			}
			nextButton(c, t, st, p.Steps[step].Choices, step == last, p.OnDone)
		})
	})
}

// questionCount is the text of the count, as "1 of 3".
func questionCount(step, total int) string {
	return fmt.Sprintf("%d of %d", step+1, total)
}

// choiceList draws the choices of the step, of which one may be
// selected: it shows in the accent color, and a click on it answers the
// step.
func choiceList(c *ui.Context, t *ui.Theme, st *State, choices []Choice) {
	ui.Column(c).Gap(t.Space(1)).Children(func() {
		for _, ch := range choices {
			ch := ch
			on := st.Answers[st.Step] == ch.Value
			b := ui.ButtonBase(c).FillWidth().Justify(ui.Start).
				Padding(t.Space(1.5), t.Space(2.5)).Radius(t.Radius)
			b.Draw(func(pp *ui.Painter, r ui.Rect) {
				face := t.Surface
				if on {
					face = t.Accent
				} else if b.Hovered() || b.Pressed() {
					face = t.SurfaceHover
				}
				pp.Fill(r, face, t.Radius)
			})
			fg := t.Text
			if on {
				fg = t.AccentText
			}
			b.Children(func() { ui.Text(c, ch.Label).SingleLine().TextColor(fg) })
			if b.Clicked() {
				st.Answers[st.Step] = ch.Value
			}
		}
	})
}

// backButton draws the Back button: it goes to the previous step.
func backButton(c *ui.Context, t *ui.Theme, st *State) {
	b := ui.ButtonBase(c).Padding(t.Space(1), t.Space(3)).Radius(t.Radius).
		Background(t.Surface)
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if b.Hovered() || b.Pressed() {
			pp.Fill(r, t.SurfaceHover, t.Radius)
		}
	})
	b.Children(func() { ui.Text(c, "Back").SingleLine() })
	if b.Clicked() && st.Step > 0 {
		st.Step--
	}
}

// nextButton draws the Next button, which says Done on the last step
// and then runs OnDone. It stays disabled while the step has no answer.
func nextButton(c *ui.Context, t *ui.Theme, st *State, choices []Choice, last bool, onDone func()) {
	label := "Next"
	if last {
		label = "Done"
	}
	b := ui.PrimaryButton(c, label)
	if st.Answers[st.Step] == "" {
		b.Disabled(true)
	}
	if b.Clicked() && !b.IsDisabled() {
		if last {
			if onDone != nil {
				onDone()
			}
		} else {
			st.Step++
		}
	}
}
