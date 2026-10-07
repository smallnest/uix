// Example feedback opens a window of MyGo native UI that shows the
// feedback controls of uix together: badge, progress, toast and dialog,
// plus the button they build on.
//
//	go run ./examples/feedback                open the window
//	go test ./examples/feedback/...           render and check it headless
//	go run ./examples/feedback/cmd/snapshot   write light.png, dark.png and dialog.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/feedback/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix feedback",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.FeedbackView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
