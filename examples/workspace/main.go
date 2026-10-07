// The workspace example shows the form, scrollhorizontal, scrollboth and
// link components together in a project panel: the form lines up its
// fields, the days scroll sideways, the board scrolls both ways, and the
// header links open the guide.
//
//	go run ./examples/workspace
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/workspace/view"
)

func main() {
	view.Reset()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   view.Title(),
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.WorkspaceView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
