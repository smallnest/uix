package main

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/settings/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() { view.State = view.SettingsState{Theme: "system", Language: "English", StatusBar: true} }

// TestRenders draws the settings screen headless and checks that the
// title, the tabs and the controls of the General pane came through:
// MyGo renders frames in memory, no window needed.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	for _, s := range []string{
		"Settings", "General", "Appearance", "About",
		"Language", "English", "Theme preference", "Light", "Dark", "System",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestTabs switches the pane with a click on a tab.
func TestTabs(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Appearance"); err != nil {
		t.Fatal(err)
	}
	if view.State.Tab != 1 {
		t.Fatalf("tab = %d, want 1", view.State.Tab)
	}
	for _, s := range []string{"Dark mode", "Show the status bar"} {
		if !tt.HasText(s) {
			t.Fatalf("appearance pane lacks %q", s)
		}
	}
	if tt.HasText("Theme preference") {
		t.Fatal("the general pane is still shown")
	}
}

// TestRadio chooses a theme with a click on its label.
func TestRadio(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Dark"); err != nil {
		t.Fatal(err)
	}
	if view.State.Theme != "dark" {
		t.Fatalf("theme = %q, want dark", view.State.Theme)
	}
}

// TestThemeRadioSwitches draws the screen in the theme the radio chooses:
// a light background under Light, a dark one under Dark, and one that
// follows the system under System.
func TestThemeRadioSwitches(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Dark"); err != nil {
		t.Fatal(err)
	}
	if c := background(tt); c != "#18181b" {
		t.Fatalf("Dark drew background %s, want #18181b", c)
	}
	if err := tt.Click("Light"); err != nil {
		t.Fatal(err)
	}
	if c := background(tt); c != "#ffffff" {
		t.Fatalf("Light drew background %s, want #ffffff", c)
	}

	reset()
	tt = ui.NewTester(view.SettingsView, view.Width, view.Height)
	tt.SetDark(true)
	if c := background(tt); c != "#18181b" {
		t.Fatalf("System under a dark system drew background %s, want #18181b", c)
	}

	reset()
	tt = ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Light"); err != nil {
		t.Fatal(err)
	}
	tt.SetDark(true)
	if c := background(tt); c != "#ffffff" {
		t.Fatalf("Light under a dark system drew background %s, want #ffffff", c)
	}
}

// TestDarkModeSwitch toggles the appearance from the Appearance pane: on
// chooses Dark, off chooses Light, and the radio of the General pane
// follows.
func TestDarkModeSwitch(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("Appearance"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if view.State.Theme != "dark" {
		t.Fatalf("theme = %q, want dark after the switch", view.State.Theme)
	}
	if c := background(tt); c != "#18181b" {
		t.Fatalf("switch on drew background %s, want #18181b", c)
	}
	if err := tt.Click("Dark mode"); err != nil {
		t.Fatal(err)
	}
	if view.State.Theme != "light" {
		t.Fatalf("theme = %q, want light after the switch", view.State.Theme)
	}
	if c := background(tt); c != "#ffffff" {
		t.Fatalf("switch off drew background %s, want #ffffff", c)
	}
}

// TestStatusBar shows and hides the strip at the bottom of the window
// with the check box of the Appearance pane.
func TestStatusBar(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if !tt.HasText("Ready") {
		t.Fatal("the status bar is missing by default")
	}
	if err := tt.Click("Appearance"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Show the status bar"); err != nil {
		t.Fatal(err)
	}
	if view.State.StatusBar {
		t.Fatal("check box did not toggle")
	}
	if tt.HasText("Ready") {
		t.Fatal("the status bar is still shown")
	}
	if err := tt.Click("Show the status bar"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Ready") {
		t.Fatal("the status bar did not come back")
	}
}

// TestCheckForUpdates shows the toast the About pane fires.
func TestCheckForUpdates(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("About"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Check for updates"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("You are up to date") {
		t.Fatal("the toast did not show")
	}
}

// background returns the color of the window's background, in the hex form
// the light and dark appearances set.
func background(tt *ui.Tester) string {
	c := tt.Image().At(2, 2)
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r>>8, g>>8, b>>8)
}

// TestDropdown opens the language menu, chooses a language, and closes
// the menu either way.
func TestDropdown(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	if err := tt.Click("English"); err != nil {
		t.Fatal(err)
	}
	if !view.State.MenuOpen {
		t.Fatal("the menu did not open")
	}
	if !tt.HasText("日本語") {
		t.Fatal("the menu items are missing")
	}
	if err := tt.Click("日本語"); err != nil {
		t.Fatal(err)
	}
	if view.State.Language != "日本語" {
		t.Fatalf("language = %q, want 日本語", view.State.Language)
	}
	if view.State.MenuOpen {
		t.Fatal("the choice did not close the menu")
	}
	if tt.HasText("Add a language…") {
		t.Fatal("the menu is still open")
	}
}

// TestDarkMode draws the settings screen under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.SettingsView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Settings") {
		t.Fatal("dark screen missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
