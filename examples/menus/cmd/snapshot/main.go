// Command snapshot renders the menus example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/menus/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/menus/view"
)

func main() {
	// The plain shot draws the page in its initial state.
	view.Reset()
	save(ui.NewTester(view.MenusView, view.Width, view.Height), "menus.png")

	// The nav shot opens the panel of the Docs link.
	view.Reset()
	tt := ui.NewTester(view.MenusView, view.Width, view.Height)
	if err := tt.Click("Docs"); err != nil {
		panic(err)
	}
	save(tt, "menus-nav.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	dt := ui.NewTester(view.MenusView, view.Width, view.Height)
	dt.SetDark(true)
	save(dt, "menus-dark.png")
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
