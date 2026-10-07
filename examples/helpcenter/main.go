// Example helpcenter opens a window of MyGo native UI that shows the
// breadcrumbs, accordion, collapsible and popover components together in
// a help page. The breadcrumbs say where the user is, the accordion
// opens the answers to the questions, a collapsible hides more options,
// and a popover jumps to a section.
//
//	go run ./examples/helpcenter                open the window
//	go test ./examples/helpcenter/...           render and check it headless
//	go run ./examples/helpcenter/cmd/snapshot   write helpcenter.png, helpcenter-open.png and helpcenter-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/helpcenter/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix help center",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.HelpView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
