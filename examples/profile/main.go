// Example profile opens a window of MyGo native UI that shows the card,
// avatar, separator and alert components together: a profile page with a
// follow button that swaps its label and shows a success alert.
//
//	go run ./examples/profile                open the window
//	go test ./examples/profile/...           render and check it headless
//	go run ./examples/profile/cmd/snapshot   write profile.png, profile-following.png and profile-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/profile/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix profile",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.ProfileView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
