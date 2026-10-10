// Package view builds the overlays example: a window of MyGo native UI
// that shows the sheet, drawer and command components of uix together in
// one page. The buttons open the side panel of filters and the drawer
// that confirms an action, and the command palette runs a command, for
// a click or for ⌘K. The main package shows it in a window; the tests
// and the snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/command"
	"github.com/smallnest/uix/components/drawer"
	"github.com/smallnest/uix/components/field"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/sheet"
	switcher "github.com/smallnest/uix/components/switch"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 420

// State is the state of the example; the controls edit it in place.
var State = OverlaysState{}

// OverlaysState holds the example: what shows, the filter fields and
// the last command that ran.
type OverlaysState struct {
	// SheetOpen shows the sheet when true.
	SheetOpen bool
	// DrawerOpen shows the drawer when true.
	DrawerOpen bool
	// CommandOpen shows the command palette when true.
	CommandOpen bool
	// Command is where the command palette is.
	Command command.State
	// Min is the minimum price of the filter sheet.
	Min string
	// Kind is the kind of the filter sheet.
	Kind bool
	// Last is the name of the last command that ran.
	Last string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = OverlaysState{} }

// OverlaysView draws the example: the three buttons that open the
// overlays, and the overlays themselves.
func OverlaysView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	// ⌘K opens the command palette from anywhere in the window.
	if c.Shortcut(ui.Cmd, ui.KeyK) {
		State.CommandOpen = true
	}
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Overlays").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Sheet, drawer and command palette, from uix.").TextColor(t.TextMuted)
		})
		ui.Column(c).FillWidth().Gap(t.Space(3)).Children(func() {
			button.Button(c, button.Props{
				Label:   "Open filters",
				Variant: button.Secondary,
				OnClick: func() { State.SheetOpen = true },
			})
			button.Button(c, button.Props{
				Label:   "Delete the project",
				Variant: button.Secondary,
				OnClick: func() { State.DrawerOpen = true },
			})
			button.Button(c, button.Props{
				Label:   "Run a command  ⌘K",
				Variant: button.Secondary,
				OnClick: func() { State.CommandOpen = true },
			})
			ui.Text(c, "The last command: "+State.Last).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		})
		sheet.Sheet(c, sheet.Props{
			Open:     &State.SheetOpen,
			Title:    "Filters",
			Side:     sheet.Right,
			Children: func() { filterForm(c, t) },
		})
		drawer.Drawer(c, drawer.Props{
			Open:  &State.DrawerOpen,
			Title: "Delete the project",
			Children: func() {
				ui.Text(c, "This removes the project and its files. This cannot be undone.").
					TextColor(t.TextMuted)
				ui.Row(c).FillWidth().Justify(ui.End).Gap(t.Space(1.5)).Children(func() {
					button.Button(c, button.Props{
						Label:   "Cancel",
						Variant: button.Secondary,
						Size:    button.Sm,
						OnClick: func() { State.DrawerOpen = false },
					})
					button.Button(c, button.Props{
						Label:   "Delete",
						Variant: button.Destructive,
						Size:    button.Sm,
						OnClick: func() {
							State.DrawerOpen = false
							State.Last = "project deleted"
						},
					})
				})
			},
		})
		command.Command(c, command.Props{
			Open:  &State.CommandOpen,
			State: &State.Command,
			Items: []command.Item{
				{Label: "New file", Keywords: []string{"create"}, OnSelect: func() { State.Last = "new file" }},
				{Label: "Save", OnSelect: func() { State.Last = "saved" }},
				{Label: "Delete the project", OnSelect: func() {
					State.Last = "project deleted"
					State.DrawerOpen = true
				}},
			},
		})
	})
}

// filterForm is the body of the sheet: two fields of the filter.
func filterForm(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(3)).Children(func() {
		field.Field(c, field.Props{Label: "Minimum price"}, func() {
			input.Input(c, input.Props{Value: &State.Min, Placeholder: "0"})
		})
		ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Center).Children(func() {
			switcher.Switch(c, switcher.Props{On: &State.Kind, Label: "In stock only"})
		})
	})
}
