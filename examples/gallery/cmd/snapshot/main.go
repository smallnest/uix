// Command snapshot renders the gallery example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/gallery/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/gallery/view"
)

func main() {
	// The plain shot draws the gallery in its initial state: the photos
	// by name.
	view.Reset()
	save(ui.NewTester(view.GalleryView, view.Width, view.Height), "gallery.png")

	// The filtered shot shows the docs by date; a click chooses the last
	// one and focuses the grid, which rings it in the accent.
	view.Reset()
	view.State.Kind = 1
	view.State.Sort = 1
	tt := ui.NewTester(view.GalleryView, view.Width, view.Height)
	if err := tt.Click("Report.pdf"); err != nil {
		panic(err)
	}
	save(tt, "gallery-docs.png")

	// The dark shot draws the same gallery under the dark appearance.
	view.Reset()
	tt2 := ui.NewTester(view.GalleryView, view.Width, view.Height)
	tt2.SetDark(true)
	save(tt2, "gallery-dark.png")
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
