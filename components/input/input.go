// Package input provides a shadcn/ui-style text field for the MyGo native
// toolkit: a single-line TextInput. The multi-line field is the separate
// textarea component, so a form says which it means.
//
//	input.Input(c, input.Props{
//		Value:       &email,
//		Placeholder: "you@example.com",
//	})
package input

import "github.com/egoist/mygo/ui"

// Props describes the input to draw.
type Props struct {
	// Value is the text being edited.
	Value *string
	// Placeholder shows in the empty field, in the muted color.
	Placeholder string
	// Disabled keeps the field from being edited.
	Disabled bool
}

// Input draws a field editing *Value and returns it, so a view can chain
// more calls on it. It fills the width of its container.
func Input(c *ui.Context, p Props) *ui.Element {
	e := ui.TextInput(c, p.Value)
	e.Placeholder(p.Placeholder)
	e.FillWidth()
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
