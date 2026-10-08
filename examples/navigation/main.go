// The navigation example shows the carousel, pagination and file upload
// components of BoardUI together in one page: the slides to flip
// through, the pages to step through, and the file zone to fill in. The
// upload is simulated, so the zone needs no storage; the Upload sample
// button starts one without the dialog.
//
//	go run ./examples/navigation
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/navigation/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Navigation",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.NavigationView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
