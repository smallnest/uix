// Package inputgroup provides a shadcn/ui-style Input Group: a field
// with an icon or a text adornment at either end, as a search field with
// a magnifier, or a currency field with a "$". The field keeps the look
// of the input component, its border turning the accent while it has the
// focus.
//
//	inputgroup.InputGroup(c, inputgroup.Props{
//		Value:     &amount,
//		Leading:   dollar,
//		TrailingText: "USD",
//	})
package inputgroup

import "github.com/egoist/mygo/ui"

// Props describes the field to draw.
type Props struct {
	// Value is the text being edited.
	Value *string
	// Placeholder shows in the empty field, in the muted color.
	Placeholder string
	// Disabled keeps the field from being edited.
	Disabled bool
	// Leading draws an icon at the start of the field, and LeadingText
	// draws text there instead, such as "$" or "@".
	Leading     *ui.SVG
	LeadingText string
	// Trailing draws an icon at the end of the field, and TrailingText
	// draws text there instead, such as "kg" or "%".
	Trailing     *ui.SVG
	TrailingText string
}

// InputGroup draws the field with its adornments and returns it, so a
// view can chain more calls on it. It fills the width of its container;
// the field inside grows to it, and a click on it edits as an input
// does.
func InputGroup(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	// The field is built among the children, so it grows with the row:
	// an element created before the row would not take part in its
	// layout. Its focus, read after the frame, turns the border accent.
	var in ui.Element
	root := ui.Row(c).FillWidth().AlignItems(ui.Center).Gap(t.Space(2)).
		Padding(t.Space(1.5), t.Space(2.5)).Radius(t.Radius).Background(t.Surface).
		Border(1, t.Border)
	root.DrawOver(func(pp *ui.Painter, r ui.Rect) {
		if !p.Disabled && in.Valid() && in.Focused() {
			pp.Stroke(r, t.Accent, t.Radius, 1)
		}
	})
	root.Children(func() {
		if p.Leading != nil {
			ui.Icon(c, p.Leading).TextColor(t.TextMuted)
		} else if p.LeadingText != "" {
			ui.Text(c, p.LeadingText).TextColor(t.TextMuted)
		}
		in = ui.TextInputBase(c, p.Value).Grow(1)
		if p.Disabled {
			in.Disabled(true)
		}
		if p.Placeholder != "" {
			in.Placeholder(p.Placeholder)
		}
		if p.Trailing != nil {
			ui.Icon(c, p.Trailing).TextColor(t.TextMuted)
		} else if p.TrailingText != "" {
			ui.Text(c, p.TrailingText).TextColor(t.TextMuted)
		}
	})
	return root
}
