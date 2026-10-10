// Command snapshot renders the conversation example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/conversation/cmd/snapshot
package main

import (
	"image/png"
	"os"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/conversation/view"
)

func main() {
	// The plain shot draws the page in its initial state: the chat with
	// the attachment a third uploaded.
	view.Reset()
	save(ui.NewTester(view.ConversationView, view.Width, view.Height), "conversation.png")

	// The done shot finishes the upload of the attachment.
	view.Reset()
	for i := 0; i < 14; i++ {
		view.Tick(tickTime)
	}
	save(ui.NewTester(view.ConversationView, view.Width, view.Height), "conversation-done.png")

	// The dark shot draws the page under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.ConversationView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "conversation-dark.png")
}

// tickTime is the time the ticks advance the simulated upload by.
var tickTime = time.Now()

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
