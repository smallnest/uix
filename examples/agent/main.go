// The agent example shows the appshell, chat, composer, log,
// notificationcenter and thinking components together in a starter for
// an agentic app: a demo model streams an answer in the chat, the log
// reveals its steps, and the notifications filter by tab.
//
//	go run ./examples/agent
package main

import (
	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/agent/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "Agent",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.AgentView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		panic(err)
	}
}
