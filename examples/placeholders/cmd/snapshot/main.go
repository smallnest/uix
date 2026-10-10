// Command snapshot renders the placeholders example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/placeholders/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/placeholders/view"
)

func main() {
	// The plain shot draws the page in its initial state.
	view.Reset()
	save(ui.NewTester(view.PlaceholdersView, view.Width, view.Height), "placeholders.png")

	// The empty shot shows the empty state of a search that found
	// nothing.
	view.Reset()
	view.State.Searched = true
	save(ui.NewTester(view.PlaceholdersView, view.Width, view.Height), "placeholders-empty.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.PlaceholdersView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "placeholders-dark.png")
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
