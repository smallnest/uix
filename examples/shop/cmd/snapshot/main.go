// Command snapshot renders the shop example headless into PNG files, so
// the components can be seen without opening a window:
//
//	go run ./examples/shop/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/shop/view"
)

func main() {
	// The plain shot draws the catalog with every filter empty.
	reset()
	save(ui.NewTester(view.ShopView, view.Width, view.Height), "shop.png")

	// The filtered shot sets a query, a category and a minimum rating
	// before the frame, as the user would with the controls.
	reset()
	view.State.Query = "d"
	view.State.Category = "Tools"
	view.State.MinRate = 3
	save(ui.NewTester(view.ShopView, view.Width, view.Height), "shop-filtered.png")

	// The dark shot draws the same catalog under the dark appearance.
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "shop-dark.png")
}

// reset puts the example in its initial state.
func reset() {
	view.State.Query, view.State.Category = "", ""
	view.State.Limit, view.State.MinRate = 0, 0
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
