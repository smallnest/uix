// The gallery example shows the components of the three libraries —
// uix's own, mygo's primitives, and onevra's — in a shell like onevra's
// showcase: a sidebar of every component, one page each, with a badge
// naming the library the page shows. Where the same component exists in
// several libraries, the gallery keeps one, by priority uix, mygo,
// onevra. A search box above the sidebar filters the items.
//
//	go run ./examples/gallery
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/gallery/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Gallery",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.GalleryView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
