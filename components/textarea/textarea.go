// Package textarea provides a shadcn/ui-style multi-line text field for
// the MyGo native toolkit.
//
//	textarea.Textarea(c, textarea.Props{
//		Value:       &notes,
//		Placeholder: "Write your notes here…",
//	})
package textarea

import "github.com/egoist/mygo/ui"

// Props describes the text area to draw.
type Props struct {
	// Value is the text being edited.
	Value *string
	// Placeholder shows in the empty field, in the muted color.
	Placeholder string
	// Disabled keeps the field from being edited.
	Disabled bool
	// Rows is the height of the field in lines of text; the default is 4.
	Rows int
}

// Textarea draws a multi-line field editing *Value and returns it, so a
// view can chain more calls on it. It fills the width of its container.
func Textarea(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	if p.Rows <= 0 {
		p.Rows = 4
	}
	e := ui.TextArea(c, p.Value)
	e.Placeholder(p.Placeholder)
	e.FillWidth()
	// MyGo's text area comes with a floor of 20 units, near four lines;
	// MinHeight replaces it, so Rows can ask for any height.
	//
	// Why not the line height of the text engine: it is a metric of the
	// font, internal to MyGo, so we approximate it with CSS's standard
	// line height of 1.4em, as Rem returns. The box grows by the padding
	// of the field, 1.5 units top and bottom.
	e.MinHeight(float32(p.Rows)*t.Rem(1.4) + t.Space(3))
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
