// Command snapshot renders the questionnaire example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/questionnaire/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/questionnaire/view"
)

func main() {
	// The plain shot draws the page on its first question.
	view.Reset()
	save(ui.NewTester(view.QuestionnaireView, view.Width, view.Height), "questionnaire.png")

	// The second shot draws the second question with its choice made,
	// so the progress bar and the Back button show.
	view.Reset()
	view.State.Form.Step = 1
	view.State.Form.Answers = []string{"app", ""}
	save(ui.NewTester(view.QuestionnaireView, view.Width, view.Height), "questionnaire-step2.png")

	// The done shot answers all three questions and reports the
	// answers.
	view.Reset()
	view.State.Form.Step = 2
	view.State.Form.Answers = []string{"app", "small", "word"}
	view.State.Done = "building an app with a team of two to ten, found by word of mouth"
	save(ui.NewTester(view.QuestionnaireView, view.Width, view.Height), "questionnaire-done.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.QuestionnaireView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "questionnaire-dark.png")
}

// save writes the frame of the tester into a PNG file.
func save(tt *ui.Tester, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, tt.Image()); err != nil {
		panic(err)
	}
}
