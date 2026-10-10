// The inputs example shows the button group, input group, one-time
// password, key and label components of uix together in one page: the
// alignment to choose, the amount to type with its currency, the
// verification code its cells take, and the shortcut keys the actions
// hint at. The code field calls Verify when it fills.
//
//	go run ./examples/inputs
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/inputs/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Inputs",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.InputsView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
