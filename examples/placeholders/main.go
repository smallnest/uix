// The placeholders example shows the skeleton, empty and aspect-ratio
// components of uix together in one page: the article loading, the
// search that found nothing, and the video keeping its ratio. The Clear
// button empties the search so the empty state has a way out.
//
//	go run ./examples/placeholders
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/placeholders/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Placeholders",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.PlaceholdersView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
