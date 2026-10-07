// Example explorer opens a window of MyGo native UI that shows the
// toolbar, tree, split and list components together in a small file
// browser. The toolbar creates files and opens the chosen one, the tree
// chooses a folder, the list shows its files, and the divider between
// them is draggable.
//
//	go run ./examples/explorer                open the window
//	go test ./examples/explorer/...           render and check it headless
//	go run ./examples/explorer/cmd/snapshot   write explorer.png, explorer-changed.png and explorer-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/explorer/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix explorer",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.ExplorerView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
