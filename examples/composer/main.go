// The composer example shows the autocomplete, tokenfield, calendar and
// checkboxgroup components together in a form creating an appointment:
// the task field completes what is typed, the tag field keeps a set of
// tags, the calendar picks the due day, and the reminder group chooses
// how to be reminded.
//
//	go run ./examples/composer
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/composer/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Composer",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.ComposerView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
