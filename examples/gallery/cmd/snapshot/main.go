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

// pages are the snapshots: a uix page, a mygo page, two onevra pages —
// the established one and a v0.2026.13 newcomer — and a MujicaUI page,
// by the sidebar ID and library of each. The gallery shows one library
// at a time, so the filter picks the page's library or the gallery lands
// on the library's first page instead.
var pages = []struct {
	name, id string
	lib      view.Library
}{
	{"gallery", "button", view.Uix},
	{"gallery-gauge", "gauge", view.Uix},
	{"gallery-gauge-monitor", "gauge-monitor", view.Uix},
	{"gallery-gauge-health", "gauge-health", view.Uix},
	{"gallery-gauge-time", "gauge-time", view.Uix},
	{"gallery-gauge-weather", "gauge-weather", view.Uix},
	{"gallery-gauge-cockpit", "gauge-cockpit", view.Uix},
	{"gallery-box", "box", view.Mygo},
	{"gallery-loading", "loading", view.Mygo},
	{"gallery-pixelgrid", "pixelgrid", view.Onevra},
	{"gallery-imagecompare", "imagecompare", view.Onevra},
	{"gallery-mujica-appshell", "001", view.Mujica},
}

func main() {
	for _, dark := range []bool{false, true} {
		for _, p := range pages {
			view.Reset()
			view.State.Filter = int(p.lib)
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
