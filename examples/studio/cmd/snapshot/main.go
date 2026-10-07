// Command snapshot renders the studio example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/studio/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/studio/view"
)

func main() {
	// The plain shot draws the studio in its initial state.
	view.Reset()
	save(ui.NewTester(view.StudioView, view.Width, view.Height), "studio.png")

	// The liked shot works the studio as the user would: the mark liked
	// and a new brand color picked.
	view.Reset()
	tt := ui.NewTester(view.StudioView, view.Width, view.Height)
	tt.Click("Like")
	tt.Click("Accent")
	tt.Click("Red")
	tt.Key(0, ui.KeyEscape) // close the picker, to show the result
	save(tt, "studio-liked.png")

	// The dark shot draws the same studio under the dark appearance.
	view.Reset()
	tt2 := ui.NewTester(view.StudioView, view.Width, view.Height)
	tt2.SetDark(true)
	save(tt2, "studio-dark.png")
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
