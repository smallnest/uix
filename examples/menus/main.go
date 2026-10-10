// The menus example shows the context menu, menu bar and navigation menu
// components of uix together in one page: the bar across the top with
// the File and Edit menus, the navigation bar with the panel of links,
// and the box whose right-click opens a context menu. Choosing an item
// runs its action; the last action is named at the bottom.
//
//	go run ./examples/menus
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/menus/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Menus",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.MenusView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
