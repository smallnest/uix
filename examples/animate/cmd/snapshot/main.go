// Command snapshot renders the animate example headless into PNG files,
// so the animations can be seen at rest without opening a window:
//
//	go run ./examples/animate/cmd/snapshot
//
// Reduce motion completes each one-shot entrance at its final value, so
// the snapshots show the resting look; the loops show wherever their
// period was at the frame.
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/animate/view"
)

func main() {
	for _, dark := range []bool{false, true} {
		view.Reset()
		tt := ui.NewTester(view.ExampleView, view.Width, view.Height)
		tt.SetPreferences(ui.Preferences{ReduceMotion: true})
		tt.SetDark(dark)
		name := "light.png"
		if dark {
			name = "dark.png"
		}
		save(tt, name)
	}
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
