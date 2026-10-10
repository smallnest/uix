// Command snapshot renders the overlays example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/overlays/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/overlays/view"
)

func main() {
	// The plain shot draws the page in its initial state.
	view.Reset()
	save(ui.NewTester(view.OverlaysView, view.Width, view.Height), "overlays.png")

	// The sheet shot opens the filters sheet, and the command shot opens
	// the command palette.
	view.Reset()
	view.State.SheetOpen = true
	save(ui.NewTester(view.OverlaysView, view.Width, view.Height), "overlays-sheet.png")
	view.Reset()
	view.State.CommandOpen = true
	save(ui.NewTester(view.OverlaysView, view.Width, view.Height), "overlays-command.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.OverlaysView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "overlays-dark.png")
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
