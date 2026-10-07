// Example forms opens a window of MyGo native UI that shows the form
// controls of uix together: field, input, select, slider, switch,
// checkbox and button.
//
//	go run ./examples/forms                open the window
//	go test ./examples/forms/...           render and check it headless
//	go run ./examples/forms/cmd/snapshot   write light.png and dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/forms/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix forms",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.FormsView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
