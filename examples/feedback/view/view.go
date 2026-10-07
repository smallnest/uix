// Package view builds the feedback example: a window of MyGo native UI
// that shows the feedback controls of uix together. The main package
// shows it in a window; the tests and the snapshot command draw it
// headless.
package view

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/badge"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/dialog"
	"github.com/smallnest/uix/components/progress"
	"github.com/smallnest/uix/components/toast"
)

// Width and Height are the size of the example window.
const Width, Height = 440, 480

// State is the state of the example; the controls edit it in place.
var State = FeedbackState{}

// FeedbackState holds the values the feedback controls edit.
type FeedbackState struct {
	// DeleteOpen opens the delete-account dialog.
	DeleteOpen bool
}

// FeedbackView draws the example: a title with badges, two progress
// bars, a button that shows a toast, and a button that opens the
// delete-account dialog.
func FeedbackView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Feedback").FontSize(24).Bold()
		ui.Row(c).Gap(t.Space(2)).Children(func() {
			badge.Badge(c, badge.Props{Label: "v1.2.0", Variant: badge.Outline})
			badge.Badge(c, badge.Props{Label: "Beta", Variant: badge.Secondary})
			badge.Badge(c, badge.Props{Label: "Needs attention", Variant: badge.Destructive})
		})
		progress.Progress(c, progress.Props{Value: 0.64, Label: "Disk usage"})
		progress.Progress(c, progress.Props{Value: 0, Label: "Indexing", Indeterminate: true})
		button.Button(c, button.Props{
			Label:   "Show toast",
			Variant: button.Secondary,
			OnClick: func() {
				toast.Push(c, toast.Props{Title: "Settings saved", Description: "They take effect at once.", Type: toast.Success})
			},
		})
		button.Button(c, button.Props{
			Label:   "Delete account",
			Variant: button.Destructive,
			OnClick: func() { State.DeleteOpen = true },
		})
	})
	toast.Viewport(c)
	dialog.Dialog(c, dialog.Props{
		Open:        &State.DeleteOpen,
		Title:       "Delete your account?",
		Description: "This removes your data and cannot be undone.",
		Confirm:     "Delete",
		Cancel:      "Keep account",
		Destructive: true,
		OnConfirm: func() {
			toast.Push(c, toast.Props{Title: "Account deleted", Type: toast.Error})
		},
		OnCancel: func() {
			toast.Push(c, toast.Props{Title: "Account kept", Type: toast.Info})
		},
	})
}
