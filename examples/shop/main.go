// Example shop opens a window of MyGo native UI that shows the combobox,
// numberinput, searchfield and rating components together: a product
// catalog that a search box, a category box, a limit and a minimum
// rating all narrow at once.
//
//	go run ./examples/shop                open the window
//	go test ./examples/shop/...           render and check it headless
//	go run ./examples/shop/cmd/snapshot   write shop.png, shop-filtered.png and shop-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/shop/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix shop",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.ShopView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
