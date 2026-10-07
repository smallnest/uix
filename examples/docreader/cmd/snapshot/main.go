// Command snapshot renders the docreader example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/docreader/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/docreader/view"
)

func main() {
	// The plain shot draws the reader in its initial state.
	view.Reset()
	save(ui.NewTester(view.DocView, view.Width, view.Height), "docreader.png")

	// The find shot opens the find bar on a query, whose matches
	// highlight and the first of which is current, and renames the
	// title, as the user would with the controls.
	view.Reset()
	view.State.Title = "mygo notes"
	view.State.FindOpen = true
	view.State.Query = "the"
	view.State.Current = 0
	save(ui.NewTester(view.DocView, view.Width, view.Height), "docreader-find.png")

	// The dark shot draws the same reader under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.DocView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "docreader-dark.png")
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
