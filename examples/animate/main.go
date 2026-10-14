// The animate example shows the animate component of uix: cards entering
// one after another, a pulse that breathes, a square that spins, and a
// list that slides in, all over the tokens palette. The replay button
// plays the entrances again.
//
//	go run ./examples/animate
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/animate/view"
)

func main() {
	view.Reset()
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   view.Title(),
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.ExampleView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
