// The conversation example shows the marker and attachment components
// of uix together in a chat with a support agent: the markers name the
// turns, and the attachments are the file of the last turn, uploading
// as the page shows. The Upload button starts a new upload, and the X
// removes the attachment.
//
//	go run ./examples/conversation
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/conversation/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Conversation",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.ConversationView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
