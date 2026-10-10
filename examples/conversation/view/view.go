// Package view builds the conversation example: a window of MyGo native
// UI that shows the marker and attachment components of uix together in
// a chat with a support agent. The markers name the turns, and the
// attachment is the file of the last turn, which uploads as the page
// shows; the Upload button starts another, and the X removes it. The
// upload is simulated, so the example runs anywhere. The main package
// shows it in a window; the tests and the snapshot command draw it
// headless.
package view

import (
	"time"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/attachment"
	"github.com/smallnest/uix/components/button"
	"github.com/smallnest/uix/components/marker"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 520

// State is the state of the example; the controls edit it in place.
var State = ConversationState{}

// ConversationState holds the example: the attachment and the share of
// its upload.
type ConversationState struct {
	// File is the name of the attachment shown.
	File string
	// Desc is the line under the name, such as the size.
	Desc string
	// Progress is the share of the upload done, from 0 to 1.
	Progress float64
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = ConversationState{
		File:     "report.pdf",
		Desc:     "1.2 MB · uploading",
		Progress: 0.35,
	}
}

// ConversationView draws the example: the chat with the markers and the
// attachment, and the controls under it.
func ConversationView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Conversation").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Markers and attachments, from uix.").TextColor(t.TextMuted)
		})
		chatCard(c, t)
		controls(c, t)
	})
}

// chatCard is the conversation: the turns of the agent and the user,
// the second holding the attachment.
func chatCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(3)).Background(t.Surface).Radius(t.Radius).
		Border(1, t.Border).Padding(t.Space(3), t.Space(4)).Children(func() {
		turn(c, t, marker.Border, "AI", "Thanks for the report. I have looked at the numbers, and the growth is steady.")
		ui.Row(c).FillWidth().Gap(t.Space(2)).Children(func() {
			ui.Column(c).Grow(1).Gap(t.Space(2)).Children(func() {
				ui.Text(c, "Here is the file you asked for.").TextColor(t.Text)
				if State.File != "" {
					attachment.Attachment(c, attachment.Props{
						Title:       State.File,
						Description: State.Desc,
						Progress:    State.Progress,
						OnRemove: func() {
							State.File = ""
							State.Desc = ""
							State.Progress = -1
						},
					})
				}
			})
			ui.Column(c).AlignItems(ui.Center).Gap(t.Space(1)).Children(func() {
				marker.Marker(c, marker.Props{Label: "You"})
			})
		})
	})
}

// turn is one message of the chat: the marker and the text under it.
func turn(c *ui.Context, t *ui.Theme, v marker.Variant, who, text string) {
	ui.Column(c).FillWidth().Gap(t.Space(1)).Children(func() {
		marker.Marker(c, marker.Props{Label: who, Variant: v})
		ui.Text(c, text).TextColor(t.Text)
	})
}

// controls are the buttons under the chat: one to start an upload, and
// the note that it is simulated.
func controls(c *ui.Context, t *ui.Theme) {
	ui.Row(c).FillWidth().Justify(ui.End).Gap(t.Space(3)).Children(func() {
		button.Button(c, button.Props{
			Label:   "Upload a file",
			Size:    button.Sm,
			OnClick: func() { startUpload() },
		})
	})
	ui.Text(c, "The upload is simulated, so the example runs anywhere.").
		FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
}

// startUpload begins the upload of a sample file, as if the user had
// picked it; the progress grows as time passes.
func startUpload() {
	State.File = "sales-report.xlsx"
	State.Desc = "1.2 MB · uploading"
	State.Progress = 0
}

// Tick advances the simulated upload, for the tests and the running
// example.
func Tick(now time.Time) {
	if State.Progress >= 0 && State.Progress < 1 {
		State.Progress += 0.05
		if State.Progress >= 1 {
			State.Progress = 1
			State.Desc = "1.2 MB · uploaded"
		}
	}
}
