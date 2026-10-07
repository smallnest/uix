// Package dialog provides a shadcn/ui-style modal dialog with a title, a
// description and an action row, on top of the DialogBase of MyGo.
//
//	dialog.Dialog(c, dialog.Props{
//		Open:        &open,
//		Title:       "Delete account?",
//		Description: "This cannot be undone.",
//		Confirm:     "Delete",
//		Destructive: true,
//		OnConfirm:   deleteAccount,
//	})
package dialog

import "github.com/egoist/mygo/ui"

// Props describes the dialog to draw.
type Props struct {
	// Open is the bool that shows and hides the dialog; the backdrop
	// and Escape set it to false.
	Open *bool
	// Title is the heading of the dialog.
	Title string
	// Description is the text under the title.
	Description string
	// Confirm labels the confirm button; an empty label hides it.
	Confirm string
	// Cancel labels the cancel button; "Cancel" when empty.
	Cancel string
	// Destructive colors the confirm button with the danger color.
	Destructive bool
	// OnConfirm runs for the confirm button, which also closes the
	// dialog.
	OnConfirm func()
	// OnCancel runs for the cancel button; the button closes the
	// dialog either way.
	OnCancel func()
}

// Dialog draws the dialog for props, or nothing while it is closed, and
// returns the panel. The backdrop dims the window; clicking it or
// pressing Escape closes the dialog.
func Dialog(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	cancel := p.Cancel
	if cancel == "" {
		cancel = "Cancel"
	}
	return ui.DialogBase(c, p.Open, func(back, panel *ui.Element) {
		back.Background(ui.RGBA(0, 0, 0, 0.4))
		panel.Padding(t.Space(5)).Gap(t.Space(3)).Radius(t.Radius+4).Background(t.Background).
			MinWidth(t.Space(48)).MaxWidth(t.Space(72)).Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3))
		panel.Children(func() {
			ui.Column(c).Gap(t.Space(1.5)).Children(func() {
				ui.Text(c, p.Title).Bold()
				if p.Description != "" {
					ui.Text(c, p.Description).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
				}
			})
			ui.Row(c).Justify(ui.End).Gap(t.Space(1.5)).Children(func() {
				cancelButton(c, t, cancel, p)
				if p.Confirm != "" {
					confirmButton(c, t, p)
				}
			})
		})
	})
}

// cancelButton draws the cancel button, which closes the dialog. Why not
// a hover face: the pointer is already inside the modal, so a solid face
// keeps the dialog calm, as the Modal of MyGo does.
func cancelButton(c *ui.Context, t *ui.Theme, label string, p Props) {
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).
		Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	b.Children(func() { ui.Text(c, label).SingleLine() })
	if b.Clicked() {
		*p.Open = false
		if p.OnCancel != nil {
			p.OnCancel()
		}
	}
}

// confirmButton draws the confirm button, in the accent or the danger
// color, which closes the dialog and runs OnConfirm.
func confirmButton(c *ui.Context, t *ui.Theme, p Props) {
	face, fg := t.Accent, t.AccentText
	if p.Destructive {
		face, fg = t.Danger, ui.Hex("#ffffff")
	}
	b := ui.ButtonBase(c).Padding(t.Space(1.5), t.Space(3.5)).Radius(t.Radius).Background(face).TextColor(fg)
	b.Children(func() { ui.Text(c, p.Confirm).SingleLine() })
	if b.Clicked() {
		*p.Open = false
		if p.OnConfirm != nil {
			p.OnConfirm()
		}
	}
}
