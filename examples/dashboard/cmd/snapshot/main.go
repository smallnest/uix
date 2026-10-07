// Command snapshot renders the dashboard example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/dashboard/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/dashboard/view"
)

func main() {
	// The plain shot draws the overview in its initial state.
	view.Reset()
	save(ui.NewTester(view.DashboardView, view.Width, view.Height), "dashboard.png")

	// The hover shot rests the revenue chart on August and the orders
	// chart on October, as the pointer would.
	view.Reset()
	view.State.Revenue = 7
	view.State.Orders = 9
	save(ui.NewTester(view.DashboardView, view.Width, view.Height), "dashboard-hover.png")

	// The dark shot draws the overview under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.DashboardView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "dashboard-dark.png")
}

// save writes the frame of the tester into a PNG file.
func save(tt *ui.Tester, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, tt.Image()); err != nil {
		panic(err)
	}
}
