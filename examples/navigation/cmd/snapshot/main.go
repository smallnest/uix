// Command snapshot renders the navigation example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/navigation/cmd/snapshot
package main

import (
	"image/png"
	"os"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/fileupload"
	"github.com/smallnest/uix/examples/navigation/view"
)

func main() {
	// The plain shot draws the page in its initial state.
	view.Reset()
	save(ui.NewTester(view.NavigationView, view.Width, view.Height), "navigation.png")

	// The upload shot begins a sample upload already a third in and
	// rests the pagination on page five, so the dots show.
	view.Reset()
	view.State.Page = 5
	view.State.Upload.Start("sales-report.xlsx", 1258291,
		time.Now().Add(-time.Duration(0.35*float64(fileupload.UploadDuration))))
	save(ui.NewTester(view.NavigationView, view.Width, view.Height), "navigation-upload.png")
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
