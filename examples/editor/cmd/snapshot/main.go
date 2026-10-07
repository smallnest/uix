// Command snapshot renders the editor example headless into PNG files, so
// the components can be seen without opening a window:
//
//	go run ./examples/editor/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/editor/view"
)

func main() {
	// The plain shot draws the editor in its initial state.
	view.Reset()
	save(ui.NewTester(view.EditorView, view.Width, view.Height), "editor.png")

	// The style shot edits another file and turns several styles on, as
	// the user would with the controls.
	view.Reset()
	view.State.Place = "todo"
	view.State.Bold = true
	view.State.Italic = true
	view.State.Underline = true
	view.State.Status = "Saved."
	save(ui.NewTester(view.EditorView, view.Width, view.Height), "editor-style.png")

	// The dark shot draws the same editor under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.EditorView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "editor-dark.png")
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
