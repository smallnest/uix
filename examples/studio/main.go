// The studio example shows the icon, image and colorwell components
// together in a brand workbench: the mark of the studio and a like
// toggle, a preview of a photo, and a color well that picks the brand
// color.
//
//	go run ./examples/studio
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/studio/view"
)

func main() {
	view.Reset()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   view.Title(),
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.StudioView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
