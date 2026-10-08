// Package card provides a shadcn/ui-style Card: a bordered rounded
// container with an optional header and footer, for grouping related
// content on a surface.
//
//	card.Card(c, card.Props{
//		Title:       "Account",
//		Description: "How people reach you.",
//		Footer: func() {
//			button.Button(c, button.Props{Label: "Save", Variant: button.Primary})
//		},
//	}, func() {
//		field.Field(c, field.Props{Label: "Name"}, func() {
//			input.Input(c, input.Props{Value: &name})
//		})
//	})
package card

import "github.com/egoist/mygo/ui"

// Props describes the card to draw.
type Props struct {
	// Title is the heading of the card.
	Title string
	// Description is the text under the title.
	Description string
	// Footer draws into the action row at the bottom, right-aligned.
	Footer func()
}

// Card draws a bordered rounded container: the header, then the content
// that content draws, then the footer row. It returns the container, so
// a view can size it (for example, to grow in a column).
func Card(c *ui.Context, p Props, content func()) ui.Element {
	t := c.Theme()
	return ui.Column(c).Background(t.Background).Border(1, t.Border).Radius(t.Radius).
		Padding(t.Space(4)).Gap(t.Space(3)).Children(func() {
		if p.Title != "" || p.Description != "" {
			ui.Column(c).Gap(t.Space(1)).Children(func() {
				if p.Title != "" {
					ui.Text(c, p.Title).Bold()
				}
				if p.Description != "" {
					ui.Text(c, p.Description).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
				}
			})
		}
		content()
		if p.Footer != nil {
			ui.Row(c).Justify(ui.End).Gap(t.Space(1.5)).Children(p.Footer)
		}
	})
}
