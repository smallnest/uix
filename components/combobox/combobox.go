// Package combobox provides a shadcn/ui-style Combobox: a searchable
// drop-down. The field shows the choice; the text filters the options as
// you type, and the arrows or a click choose one.
//
//	combobox.Combobox(c, combobox.Props{
//		Value:   &city,
//		Options: []string{"Beijing", "Shanghai", "Chengdu"},
//	})
package combobox

import "github.com/egoist/mygo/ui"

// Props describes the combobox to draw.
type Props struct {
	// Value is the chosen option; the field shows it.
	Value *string
	// Options are the choices in the popup, which the text filters.
	Options []string
	// Disabled keeps the field from being edited.
	Disabled bool
}

// Combobox draws a searchable drop-down for props and returns the field,
// so a view can chain more calls on it. Label names it for assistive
// technology. The text filters the options as you type, those starting
// with it first; a click or Down shows them all, and Up and Down move
// among them, Enter or a click chooses one, and Escape closes the popup.
// It fills the width of its container.
func Combobox(c *ui.Context, p Props) ui.Element {
	e := ui.Combobox(c, p.Value, p.Options)
	e.FillWidth()
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
