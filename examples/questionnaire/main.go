// The questionnaire example shows the questionnaire component of uix
// alone: an onboarding form of three questions, one at a time, with the
// progress bar up top and Back and Next under the choices. The Done
// button of the last step names the answers at the bottom.
//
//	go run ./examples/questionnaire
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/questionnaire/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Questionnaire",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.QuestionnaireView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
