// Command snapshot renders the events example headless into PNG files, so
// the components can be seen without opening a window:
//
//	go run ./examples/events/cmd/snapshot
package main

import (
	"image/png"
	"os"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/events/view"
)

func main() {
	// The plain shot draws the event in its initial state.
	view.Reset()
	save(ui.NewTester(view.EventView, view.Width, view.Height), "events.png")

	// The changed shot sets another date, time, color and a reminder
	// before the frame, as the user would with the controls.
	view.Reset()
	view.State.When = time.Date(2026, 10, 20, 15, 45, 0, 0, time.Local)
	view.State.Color = ui.Hex("#3b82f6")
	view.State.Remind = true
	save(ui.NewTester(view.EventView, view.Width, view.Height), "events-changed.png")

	// The dark shot draws the same panel under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.EventView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "events-dark.png")
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
