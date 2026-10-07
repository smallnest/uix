// Command snapshot renders the explorer example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/explorer/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/explorer/view"
)

func main() {
	// The plain shot draws the explorer in its initial state.
	view.Reset()
	save(ui.NewTester(view.ExplorerView, view.Width, view.Height), "explorer.png")

	// The changed shot opens another folder and lib, chooses its
	// README, widens the tree pane and hides the sizes, as the user
	// would with the controls.
	view.Reset()
	view.State.Folder = "docs"
	view.State.TreeChosen = "docs"
	view.State.Chosen = 1
	view.State.LibOpen = true
	view.State.Divider = 220
	view.State.ShowSizes = false
	view.State.Status = "Opened README.md."
	save(ui.NewTester(view.ExplorerView, view.Width, view.Height), "explorer-changed.png")

	// The dark shot draws the same explorer under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.ExplorerView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "explorer-dark.png")
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
