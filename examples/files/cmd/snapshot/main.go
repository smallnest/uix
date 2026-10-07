// Command snapshot renders the files example headless into PNG files, so
// the table can be seen without opening a window:
//
//	go run ./examples/files/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/files/view"
)

func main() {
	// The plain shot draws the table in the original order.
	view.State = view.TableState{Selected: -1, List: &ui.ListState{}}
	save(ui.NewTester(view.FilesView, view.Width, view.Height), "files.png")

	// The sorted shot clicks the Size header, which reorders the rows by
	// size and shows the header arrow.
	view.State = view.TableState{Selected: -1, List: &ui.ListState{}}
	tt := ui.NewTester(view.FilesView, view.Width, view.Height)
	if err := tt.Click("Size"); err != nil {
		panic(err)
	}
	save(tt, "files-sorted.png")

	// The dark shot draws the same table under the dark appearance.
	view.State = view.TableState{Selected: -1, List: &ui.ListState{}}
	tt = ui.NewTester(view.FilesView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "files-dark.png")
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
