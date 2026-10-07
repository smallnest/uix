// Example editor opens a window of MyGo native UI that shows the
// sidebar, togglegroup, tooltip and menu components together in a note
// editor. The sidebar chooses the file being edited, the toggles restyle
// its text, tips name the controls, and the file menu opens a file or
// saves.
//
//	go run ./examples/editor                open the window
//	go test ./examples/editor/...           render and check it headless
//	go run ./examples/editor/cmd/snapshot   write editor.png, editor-style.png and editor-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/editor/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix editor",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.EditorView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
