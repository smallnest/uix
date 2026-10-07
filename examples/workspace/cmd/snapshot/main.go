// Command snapshot renders the workspace example headless into PNG
// files, so the components can be seen without opening a window:
//
//	go run ./examples/workspace/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/workspace/view"
)

func main() {
	// The plain shot draws the workspace in its initial state.
	view.Reset()
	save(ui.NewTester(view.WorkspaceView, view.Width, view.Height), "workspace.png")

	// The filled shot works the workspace as the user would: a project
	// named, a day picked, and a cell of the board chosen.
	view.Reset()
	tt := ui.NewTester(view.WorkspaceView, view.Width, view.Height)
	tt.Click("Project name")
	tt.Type("Aurora")
	tt.Click("Owner")
	tt.Type("nora")
	tt.Click("Day 2")
	tt.Click("r0c0")
	save(tt, "workspace-filled.png")

	// The dark shot draws the same workspace under the dark appearance.
	view.Reset()
	tt2 := ui.NewTester(view.WorkspaceView, view.Width, view.Height)
	tt2.SetDark(true)
	save(tt2, "workspace-dark.png")
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
