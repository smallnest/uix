// Package checkbox provides a shadcn/ui-style check box: a box and a
// label that toggle a bool together, from the Checkbox of MyGo.
//
//	checkbox.Checkbox(c, checkbox.Props{
//		Checked: &agree,
//		Label:   "I agree to the terms",
//	})
package checkbox

import "github.com/egoist/mygo/ui"

// Props describes the check box to draw.
type Props struct {
	// Checked is the bool the box toggles.
	Checked *bool
	// Label is the text beside the box.
	Label string
	// Disabled keeps the box from being toggled.
	Disabled bool
}

// Checkbox draws a check box for props, the whole row being clickable.
func Checkbox(c *ui.Context, p Props) ui.Element {
	b := ui.Checkbox(c, p.Checked, p.Label)
	if p.Disabled {
		b.Disabled(true)
	}
	return b
}
