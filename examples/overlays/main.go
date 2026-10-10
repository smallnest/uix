// The overlays example shows the sheet, drawer and command components of
// uix together in one page: the side panel of filters, the drawer that
// confirms an action, and the command palette that runs a command. The
// buttons open each, and the command palette opens for ⌘K too.
//
//	go run ./examples/overlays
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/overlays/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Overlays",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.OverlaysView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
