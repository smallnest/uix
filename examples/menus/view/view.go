// Package view builds the menus example: a window of MyGo native UI
// that shows the context menu, menu bar and navigation menu components
// of uix together in one page. The menu bar holds the File and Edit
// menus across the top, the navigation menu holds the Docs and About
// links, and the box in the middle opens its context menu for a
// right-click. Choosing an item runs its action and names it at the
// bottom. The main package shows it in a window; the tests and the
// snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/contextmenu"
	"github.com/smallnest/uix/components/menubar"
	"github.com/smallnest/uix/components/navigationmenu"
)

// Width and Height are the size of the example window.
const Width, Height = 520, 420

// State is the state of the example; the controls edit it in place.
var State = MenusState{}

// MenusState holds the example: the navigation bar, the dark toggle and
// the last action that ran.
type MenusState struct {
	// Nav is where the navigation menu is.
	Nav navigationmenu.State
	// Dark is whether the copy is bold, as a checked item shows.
	Dark bool
	// Last is the name of the last action that ran.
	Last string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = MenusState{} }

// MenusView draws the example: the three menus, each under its label.
func MenusView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Menus").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Menu bar, navigation menu and context menu, from uix.").TextColor(t.TextMuted)
		})
		section(c, t, "Menu bar")
		menubarCard(c, t)
		section(c, t, "Navigation menu")
		navCard(c, t)
		section(c, t, "Context menu")
		contextCard(c, t)
		ui.Text(c, "The last action: "+State.Last).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// section is the label above a component.
func section(c *ui.Context, t *ui.Theme, name string) {
	ui.Text(c, name).FontSize(t.FontSize * 0.9).FontWeight(600).TextColor(t.TextMuted)
}

// menubarCard is the menu bar demo: the bar with the File and Edit
// menus, and a named action.
func menubarCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		menubar.Menubar(c, menubar.Props{
			Menus: []menubar.Menu{
				{Label: "File", Items: []menubar.Item{
					{Label: "New", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyN, Action: func() { State.Last = "new file" }},
					{Separator: true},
					{Label: "Save", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyS, Action: func() { State.Last = "saved" }},
					{Label: "Quit", Action: func() { State.Last = "quit" }},
				}},
				{Label: "Edit", Items: []menubar.Item{
					{Label: "Undo", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyZ, Action: func() { State.Last = "undone" }},
					{Label: "Bold", Checked: State.Dark, Action: func() { State.Dark = !State.Dark; State.Last = "bold" }},
				}},
			},
		})
	})
}

// navCard is the navigation menu demo: the bar of links, of which Docs
// opens a panel.
func navCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		navigationmenu.NavigationMenu(c, navigationmenu.Props{
			State: &State.Nav,
			Items: []navigationmenu.Item{
				{Label: "Docs", Children: []navigationmenu.Item{
					{Label: "Getting started", OnClick: func() { State.Last = "getting started" }},
					{Label: "Guides", OnClick: func() { State.Last = "guides" }},
					{Label: "Reference", OnClick: func() { State.Last = "reference" }},
				}},
				{Label: "About", OnClick: func() { State.Last = "about" }},
			},
		})
	})
}

// contextCard is the context menu demo: a box whose right-click opens
// the menu.
func contextCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		contextmenu.ContextMenu(c, contextmenu.Props{
			Items: []contextmenu.Item{
				{Label: "Rename", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyR, Action: func() { State.Last = "renamed" }},
				{Separator: true},
				{Label: "Copy", ShortcutMods: ui.Cmd, ShortcutKey: ui.KeyC, Action: func() { State.Last = "copied" }},
				{Label: "Delete", Disabled: true, Action: func() { State.Last = "deleted" }},
			},
		}).Fill().Height(120).Background(t.Surface).Border(1, t.Border).Radius(t.Radius).
			Center().Children(func() {
			ui.Text(c, "Right-click here").TextColor(t.TextMuted)
		})
	})
}
