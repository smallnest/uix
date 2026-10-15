// Command snapshot renders the gallery headless into PNG files, so a
// page of each library can be seen without opening a window:
//
//	go run ./examples/gallery/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/gallery/view"
)

// pages are the snapshots: a uix page, a mygo page and an onevra page,
// by the sidebar ID of each.
var pages = []struct{ name, id string }{
	{"gallery", "button"},
	{"gallery-box", "box"},
	{"gallery-pixelgrid", "pixelgrid"},
}

func main() {
	for _, dark := range []bool{false, true} {
		for _, p := range pages {
			view.Reset()
			view.State.Selected = p.id
			tt := ui.NewTester(view.GalleryView, view.Width, view.Height)
			tt.SetDark(dark)
			suffix := "light.png"
			if dark {
				suffix = "dark.png"
			}
			save(tt, p.name+"-"+suffix)
		}
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
