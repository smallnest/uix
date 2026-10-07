// The dashboard example shows the appshell, statcard, linechart and
// barchart components together in a store overview: the metric cards up
// top, the two charts side by side, and the display cards under them.
// The charts follow the pointer, swapping their headline for the month
// under it.
//
//	go run ./examples/dashboard
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/dashboard/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Dashboard",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.DashboardView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
