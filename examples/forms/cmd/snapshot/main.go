// Command snapshot renders the forms example headless into PNG files, so
// the components can be seen without opening a window:
//
//	go run ./examples/forms/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/forms/view"
)

func main() {
	for _, dark := range []bool{false, true} {
		tt := ui.NewTester(view.FormsView, view.Width, view.Height)
		tt.SetDark(dark)
		name := "light.png"
		if dark {
			name = "dark.png"
		}
		f, err := os.Create(name)
		if err != nil {
			panic(err)
		}
		if err := png.Encode(f, tt.Image()); err != nil {
			panic(err)
		}
		f.Close()
	}
}
