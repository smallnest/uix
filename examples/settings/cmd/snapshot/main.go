// Command snapshot renders the settings example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/settings/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/settings/view"
)

func main() {
	for _, dark := range []bool{false, true} {
		view.State = view.SettingsState{Theme: "system", Language: "English", StatusBar: true}
		tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
		tt.SetDark(dark)
		name := "light.png"
		if dark {
			name = "dark.png"
		}
		save(tt, name)
	}

	// The dropdown shot opens the language menu of the General pane.
	view.State = view.SettingsState{Theme: "system", Language: "English", StatusBar: true}
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("English"); err != nil {
		panic(err)
	}
	save(tt, "dropdown.png")

	// The appearance shot switches to the Appearance pane.
	view.State = view.SettingsState{Theme: "system", Language: "English", StatusBar: true}
	tt = ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Appearance"); err != nil {
		panic(err)
	}
	save(tt, "appearance.png")

	// The theme shot chooses Dark in the theme radio, which redraws the
	// screen in the dark appearance on its own.
	view.State = view.SettingsState{Theme: "system", Language: "English", StatusBar: true}
	tt = ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Dark"); err != nil {
		panic(err)
	}
	save(tt, "theme-dark.png")
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
