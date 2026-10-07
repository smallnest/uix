// Example docreader opens a window of MyGo native UI that shows the
// richtext, editabletext, findbar and segmented components together in
// a document reader. The segmented control switches between the reader
// and the outline, the title renames in place, and the find bar
// searches the text, whose matches highlight and which it steps
// through.
//
//	go run ./examples/docreader                open the window
//	go test ./examples/docreader/...           render and check it headless
//	go run ./examples/docreader/cmd/snapshot   write docreader.png, docreader-find.png and docreader-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/docreader/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix docreader",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.DocView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
