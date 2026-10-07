// Command snapshot renders the composer example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/composer/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/composer/view"
)

func main() {
	// The plain shot draws the composer in its initial state.
	view.Reset()
	save(ui.NewTester(view.ComposerView, view.Width, view.Height), "composer.png")

	// The filled shot works the form as the user would: a title typed
	// and added, two tags, a due day picked, and two reminders checked.
	view.Reset()
	tt := ui.NewTester(view.ComposerView, view.Width, view.Height)
	tt.Click("Task title")
	tt.Type("Call the client")
	tt.Key(0, ui.KeyEnter)
	tt.Click("Add a tag")
	tt.Type("urgent,")
	tt.Type("call")
	tt.Key(0, ui.KeyEnter)
	tt.Click("October 20, 2026")
	tt.Click("Mail")
	tt.Click("Calendar")
	save(tt, "composer-filled.png")

	// The dark shot draws the same composer under the dark appearance.
	view.Reset()
	tt2 := ui.NewTester(view.ComposerView, view.Width, view.Height)
	tt2.SetDark(true)
	save(tt2, "composer-dark.png")
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
