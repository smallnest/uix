// The gallery example shows the grid, gridview, scroll and fieldset
// components together in a file gallery: two field sets filter the files
// by type and sort them by name or date, and a grid view shows the files
// that remain, of which a click chooses one.
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
