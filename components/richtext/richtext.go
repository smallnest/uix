// Package richtext provides a text of styled runs, as a shadcn/ui
// component for the RichText of MyGo: spans with their own weight,
// color, underline or background, over the style of the text around
// them.
//
//	richtext.RichText(c, richtext.Props{
//		Spans: []ui.Span{
//			{Text: "Saved "},
//			{Text: "report.pdf", Weight: 600},
//			{Text: " to the cloud.", Color: t.TextMuted},
//		},
//	})
package richtext

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the rich text to draw.
type Props struct {
	// Spans are the styled runs of the text, in order. What a span
	// leaves zero takes the style of the text around it.
	Spans []ui.Span
}

// RichText draws the spans as one text, and returns it. Style the whole
// text with the methods of the returned element, as for a Text.
func RichText(c *ui.Context, p Props) *ui.Element {
	return ui.RichText(c, p.Spans...)
}
