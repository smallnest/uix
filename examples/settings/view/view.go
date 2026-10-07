// Package view builds the settings example: a window of MyGo native UI
// that shows the navigation controls of uix together. The main package
// shows it in a window; the tests and the snapshot command draw it
// headless.
package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/badge"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/checkbox"
	"github.com/smallnest/uix/components/dropdown"
	"github.com/smallnest/uix/components/field"
	"github.com/smallnest/uix/components/radio"
	switcher "github.com/smallnest/uix/components/switch"
	"github.com/smallnest/uix/components/tabs"
	"github.com/smallnest/uix/components/toast"
)

// Width and Height are the size of the example window.
const Width, Height = 440, 520

// State is the state of the example; the controls edit it in place.
var State = SettingsState{Theme: "system", Language: "English", StatusBar: true}

// SettingsState holds the values the controls edit.
type SettingsState struct {
	Tab       int
	Theme     string
	Language  string
	Dark      bool
	StatusBar bool
	MenuOpen  bool
}

// SettingsView draws the example: a settings screen with a title, a tab
// list and the pane of the chosen tab. The General pane has a language
// dropdown and a theme radio group; the Appearance pane has a dark-mode
// switch and a status-bar check box; the About pane has the version. The
// theme preference applies at once: choosing Light, Dark or System in the
// radio, or toggling the dark-mode switch, redraws the screen in the
// chosen appearance.
func SettingsView(c *ui.Context) {
	systemDark := c.Theme().Dark
	pref := State.Theme
	dark := darkFor(pref, systemDark)
	// Keep the dark-mode switch in step with the preference.
	State.Dark = dark
	c.SetTheme(tokens.For(dark))
	t := c.Theme()
	ui.Column(c).Fill().Children(func() {
		ui.Column(c).Grow(1).FillWidth().Padding(24).Gap(14).Children(func() {
			ui.Text(c, "Settings").FontSize(24).Bold()
			tabs.Tabs(c, tabs.Props{Selected: &State.Tab, Labels: []string{"General", "Appearance", "About"}})
			switch State.Tab {
			case 0:
				general(c, t)
			case 1:
				appearance(c, t)
			case 2:
				about(c, t)
			}
		})
		// The check box of the Appearance pane shows the status bar.
		if State.StatusBar {
			statusBar(c, t)
		}
		toast.Viewport(c)
	})
	// A control changed the preference after the theme above was set: ask
	// for another frame, so the new appearance shows.
	if State.Theme != pref {
		c.Invalidate()
	}
}

// darkFor returns whether the app draws dark after the theme preference:
// a light or dark choice decides, and system follows the system.
func darkFor(pref string, systemDark bool) bool {
	switch pref {
	case "light":
		return false
	case "dark":
		return true
	}
	return systemDark
}

// general draws the pane of the General tab: a language dropdown and a
// theme radio group, each in a field.
func general(c *ui.Context, t *ui.Theme) {
	field.Field(c, field.Props{Label: "Language"}, func() {
		dropdown.Dropdown(c, dropdown.Props{
			Trigger: State.Language,
			Open:    &State.MenuOpen,
			Items: []dropdown.Item{
				{Label: "English", OnClick: func() { State.Language = "English" }},
				{Label: "简体中文", OnClick: func() { State.Language = "简体中文" }},
				{Label: "日本語", OnClick: func() { State.Language = "日本語" }},
				{Separator: true},
				{Label: "Add a language…", Disabled: true},
			},
		}).MinWidth(t.Space(35))
	})
	field.Field(c, field.Props{Label: "Theme preference", Hint: "How the app looks to you."}, func() {
		radio.Group(c, radio.Props[string]{
			Selected: &State.Theme,
			Options: []radio.Option[string]{
				{Value: "light", Label: "Light"},
				{Value: "dark", Label: "Dark"},
				{Value: "system", Label: "System"},
			},
		})
	})
}

// appearance draws the pane of the Appearance tab. Toggling the dark-mode
// switch chooses the Dark or Light preference, as a shortcut next to the
// theme radio of the General pane.
func appearance(c *ui.Context, t *ui.Theme) {
	switcher.Switch(c, switcher.Props{
		On:    &State.Dark,
		Label: "Dark mode",
		Changed: func(on bool) {
			if on {
				State.Theme = "dark"
			} else {
				State.Theme = "light"
			}
		},
	})
	checkbox.Checkbox(c, checkbox.Props{Checked: &State.StatusBar, Label: "Show the status bar"})
}

// about draws the pane of the About tab.
func about(c *ui.Context, t *ui.Theme) {
	ui.Text(c, "uix 0.1.0").Bold()
	ui.Row(c).Gap(t.Space(2)).Children(func() {
		badge.Badge(c, badge.Props{Label: "MyGo", Variant: badge.Secondary})
		badge.Badge(c, badge.Props{Label: "Native", Variant: badge.Outline})
	})
	ui.Text(c, "A component registry for native UI, in the way of shadcn/ui.").
		TextColor(t.TextMuted).MaxWidth(t.Space(45))
	button.Button(c, button.Props{
		Label:   "Check for updates",
		Variant: button.Outline,
		OnClick: func() { toast.Push(c, toast.Props{Title: "You are up to date", Type: toast.Success}) },
	})
}

// statusBar draws the strip the check box of the Appearance pane shows
// and hides, at the bottom of the window.
func statusBar(c *ui.Context, t *ui.Theme) {
	ui.Row(c).FillWidth().Padding(t.Space(2), t.Space(3)).Gap(t.Space(2)).Background(t.Surface).
		Draw(func(pp *ui.Painter, r ui.Rect) { pp.Fill(ui.Rect{X: r.X, Y: r.Y, W: r.W, H: 1}, t.Border, 0) }).
		Children(func() {
			ui.Text(c, "Ready").FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			ui.Box(c).Grow(1)
			ui.Text(c, "uix 0.1.0").FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		})
}
