// Package label provides a standalone label of a control, for the Label
// of shadcn/ui: a line of text in the medium weight, as a form's field
// names its field. The field component draws the label of a field; use
// Label for one without a field.
//
//	label.Label(c, label.Props{Text: "Name"})
package label

import "github.com/egoist/mygo/ui"

// Props describes the label to draw.
type Props struct {
	// Text is the label.
	Text string
	// Disabled draws the label in the muted color, for the control it
	// names that is off.
	Disabled bool
}

// Label draws the text and returns it, so a view can chain more calls on
// it.
func Label(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	e := ui.Text(c, p.Text).FontWeight(500).SingleLine()
	if p.Disabled {
		e.TextColor(t.TextMuted)
	}
	return e
}
