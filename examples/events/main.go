// Example events opens a window of MyGo native UI that shows the
// datepicker, timeinput, colorpicker and toggle components together: an
// event whose date, time and color the pickers set and whose reminder a
// toggle arms, with a summary line that reflects every change.
//
//	go run ./examples/events                open the window
//	go test ./examples/events/...           render and check it headless
//	go run ./examples/events/cmd/snapshot   write events.png, events-changed.png and events-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/events/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix events",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.EventView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
