// Example files opens a window of MyGo native UI that shows the table
// component at work: a file browser whose columns sort, whose rows are
// chosen, and whose chosen row opens with Enter or a double click.
//
//	go run ./examples/files                open the window
//	go test ./examples/files/...           render and check it headless
//	go run ./examples/files/cmd/snapshot   write files.png, files-sorted.png and files-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/files/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix files",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.FilesView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
