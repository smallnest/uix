// Package field wraps a form control with its label and helper text, as
// the FormItem of shadcn/ui does. A field is a small column:
//
//	field.Field(c, field.Props{
//		Label: "Name",
//		Hint:  "How people find you.",
//	}, func() {
//		input.Input(c, input.Props{Value: &name})
//	})
//
// The label and the messages come from the theme, so the palette of the
// tokens package restyles every field at once.
package field

import "github.com/egoist/mygo/ui"

// Props describes the field to draw.
type Props struct {
	// Label is the name of the control.
	Label string
	// Required marks the field as mandatory, with a star.
	Required bool
	// Hint is the helper text under the control.
	Hint string
	// Error is the validation message under the control. It replaces
	// Hint and shows in the danger color.
	Error string
}

// Field draws the label, the control that fn builds, and the message
// under it, and returns the column.
func Field(c *ui.Context, p Props, fn func()) *ui.Element {
	t := c.Theme()
	return ui.Column(c).Gap(t.Space(0.5)).Children(func() {
		if p.Label != "" {
			// A required label ends in a star in the danger color.
			ui.Row(c).Gap(t.Space(0.5)).Children(func() {
				ui.Text(c, p.Label).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
				if p.Required {
					ui.Text(c, "*").FontSize(t.FontSize * 0.875).Bold().TextColor(t.Danger)
				}
			})
		}
		fn()
		switch {
		case p.Error != "":
			ui.Text(c, p.Error).FontSize(t.FontSize * 0.875).TextColor(t.Danger)
		case p.Hint != "":
			ui.Text(c, p.Hint).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		}
	})
}
