// Package attachment provides a shadcn/ui-style Attachment: a card for
// a file of a conversation, with a preview, its name, a line about it,
// the progress of its upload, and a button to remove it. It is the file
// of a chat or a form, which uploads as the card shows.
//
//	attachment.Attachment(c, attachment.Props{
//		Title:       "report.pdf",
//		Description: "1.2 MB · uploading",
//		Progress:    0.6,
//		OnRemove:    app.removeFile,
//	})
package attachment

import "github.com/egoist/mygo/ui"

// file is the generic shape of a file, for an attachment without a
// preview of its own.
var file = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M15 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7Z"/><path d="M14 2v4a2 2 0 0 0 2 2h4"/></svg>`))

// x is the button that removes the attachment.
var x = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>`))

// Props describes the attachment to draw.
type Props struct {
	// Title is the name of the file.
	Title string
	// Description is a line under the title, such as the size and the
	// state of the upload.
	Description string
	// Media draws the preview of the file on the left; nil draws a
	// generic file shape.
	Media func()
	// Progress is the share of the upload done, from 0 to 1, while the
	// file uploads; a value below 0 hides the progress bar.
	Progress float64
	// OnRemove runs for the remove button; nil hides the button.
	OnRemove func()
}

// Attachment draws the card of the file and returns it, so a view can
// chain more calls on it. The title grows across the card, and the
// remove button sits at the far right.
func Attachment(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	return ui.Row(c).FillWidth().Gap(t.Space(2.5)).Padding(t.Space(2)).
		Background(t.Surface).Border(1, t.Border).Radius(t.Radius).Children(func() {
		ui.Box(c).Size(t.Space(10), t.Space(10)).Shrink(0).Radius(t.Radius - 1).
			Background(t.SurfaceHover).Clip().Center().Children(func() {
			if p.Media != nil {
				p.Media()
			} else {
				ui.Icon(c, file).Size(t.Space(5), t.Space(5)).TextColor(t.TextMuted)
			}
		})
		ui.Column(c).Grow(1).Gap(t.Space(0.5)).Children(func() {
			ui.Text(c, p.Title).SingleLine()
			if p.Description != "" {
				ui.Text(c, p.Description).SingleLine().FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			}
			if p.Progress >= 0 {
				ui.Progress(c, p.Progress).FillWidth()
			}
		})
		if p.OnRemove != nil {
			b := ui.ButtonBase(c).Size(t.Space(6), t.Space(6)).Shrink(0).Radius(t.Radius)
			b.Draw(func(pp *ui.Painter, r ui.Rect) {
				if b.Hovered() || b.Pressed() {
					pp.Fill(r, t.SurfaceHover, t.Radius)
				}
			})
			b.Children(func() {
				ui.Icon(c, x).Size(t.Space(4), t.Space(4)).TextColor(t.TextMuted)
			})
			if b.Clicked() {
				p.OnRemove()
			}
		}
	})
}
