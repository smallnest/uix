// Command snapshot renders the feedback example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/feedback/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/feedback/view"
)

func main() {
	for _, dark := range []bool{false, true} {
		view.State = view.FeedbackState{}
		tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
		tt.SetDark(dark)
		name := "light.png"
		if dark {
			name = "dark.png"
		}
		save(tt, name)
	}

	// The dialog shot shows a toast first, then opens the dialog: the
	// toast is pushed before the backdrop appears, because the modal
	// covers the buttons behind it.
	view.State = view.FeedbackState{}
	tt := ui.NewTester(view.FeedbackView, view.Width, view.Height)
	if err := tt.Click("Show toast"); err != nil {
		panic(err)
	}
	if err := tt.Click("Delete account"); err != nil {
		panic(err)
	}
	save(tt, "dialog.png")
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
