// Package form provides a form whose field labels line up, in the way
// macOS draws them: a column of fields, with the label of each to its
// left, right-aligned to the widest. It wraps the Form of MyGo.
//
//	form.Form(c, form.Props{
//		Children: func(c *ui.Context) {
//			ui.Field(c, "Name", func() { input.Input(c, input.Props{Value: &name}) })
//			ui.Field(c, "Email", func() { input.Input(c, input.Props{Value: &email}) })
//		},
//	})
//
// Build each field with the Field of MyGo (or the Fieldset of uix), whose
// labels the form lines up and names for assistive technology; a click on
// a label focuses its control.
package form

import "github.com/egoist/mygo/ui"

// Props describes the form to draw.
type Props struct {
	// Children build the fields of the form, with the Field or the
	// Fieldset of MyGo, whose labels line up at the widest.
	Children func(c *ui.Context)
}

// Form draws a column of the fields that Children builds, with their
// labels to the left, right-aligned to the widest, and returns it.
func Form(c *ui.Context, p Props) ui.Element {
	if p.Children == nil {
		p.Children = func(c *ui.Context) {}
	}
	return ui.Form(c, func() { p.Children(c) })
}
