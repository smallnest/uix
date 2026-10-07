// Command snapshot renders the agent example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/agent/cmd/snapshot
package main

import (
	"image/png"
	"os"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/agent/view"
)

func main() {
	// The plain shot draws the chat in its initial state.
	view.Reset()
	save(ui.NewTester(view.AgentView, view.Width, view.Height), "agent.png")

	// The reply shot sends a suggestion and lets the demo model stream
	// part of its answer.
	view.ThinkPace, view.StreamPace = 50*time.Millisecond, 30*time.Millisecond
	view.Reset()
	tt := ui.NewTester(view.AgentView, view.Width, view.Height)
	tt.Click("Explain this starter")
	for i := 0; i < 8; i++ {
		time.Sleep(60 * time.Millisecond)
		tt.Frame()
	}
	save(tt, "agent-reply.png")

	// The log shot shows the agent's work mid-run, three steps done.
	view.Reset()
	view.State.Place = "log"
	view.State.Revealed = 3
	save(ui.NewTester(view.AgentView, view.Width, view.Height), "agent-log.png")

	// The notifications shot shows the panel with its tabs and items.
	view.Reset()
	view.State.Place = "notifications"
	save(ui.NewTester(view.AgentView, view.Width, view.Height), "agent-notifications.png")

	// The dark shot draws the chat under the dark appearance.
	view.Reset()
	tt2 := ui.NewTester(view.AgentView, view.Width, view.Height)
	tt2.SetDark(true)
	save(tt2, "agent-dark.png")
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
