// Command snapshot renders the help center example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/helpcenter/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/helpcenter/view"
)

func main() {
	// The plain shot draws the page in its initial state.
	view.Reset()
	save(ui.NewTester(view.HelpView, view.Width, view.Height), "helpcenter.png")

	// The open shot sets another trail and opens two sections, as the
	// user would with the controls.
	view.Reset()
	view.State.Chosen = 1
	view.State.FAQ1 = false
	view.State.FAQ2 = true
	view.State.FAQ3 = true
	view.State.More = false
	save(ui.NewTester(view.HelpView, view.Width, view.Height), "helpcenter-open.png")

	// The dark shot draws the same page under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.HelpView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "helpcenter-dark.png")
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
