// Command snapshot renders the inputs example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/inputs/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/inputs/view"
)

func main() {
	// The plain shot draws the page in its initial state.
	view.Reset()
	save(ui.NewTester(view.InputsView, view.Width, view.Height), "inputs.png")

	// The filled shot types an amount and a code, so the cells fill.
	view.Reset()
	view.State.Amount = "12.50"
	view.State.Code.Code = "483920"
	save(ui.NewTester(view.InputsView, view.Width, view.Height), "inputs-filled.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.InputsView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "inputs-dark.png")
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
