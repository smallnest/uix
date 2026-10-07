// Example settings opens a window of MyGo native UI that shows the
// navigation controls of uix together: tabs, dropdown and radio, plus
// the switch, check box and badge they sit with.
//
//	go run ./examples/settings                open the window
//	go test ./examples/settings/...           render and check it headless
//	go run ./examples/settings/cmd/snapshot   write light.png, dark.png, dropdown.png, appearance.png and theme-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/settings/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix settings",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.SettingsView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
