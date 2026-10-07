// Command snapshot renders the profile example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/profile/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/profile/view"
)

func main() {
	// The plain shot draws the profile page before the follow click.
	view.State.Following = false
	save(ui.NewTester(view.ProfileView, view.Width, view.Height), "profile.png")

	// The following shot clicks the follow button, which swaps its label
	// and shows the success alert.
	view.State.Following = false
	tt := ui.NewTester(view.ProfileView, view.Width, view.Height)
	if err := tt.Click("Follow"); err != nil {
		panic(err)
	}
	save(tt, "profile-following.png")

	// The dark shot draws the same page under the dark appearance.
	view.State.Following = false
	tt = ui.NewTester(view.ProfileView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "profile-dark.png")
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
