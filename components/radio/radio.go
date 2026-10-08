// Package radio provides a shadcn/ui-style radio group: a column of
// options, of which one is chosen. The group is one stop of Tab; the
// arrows and Home and End choose among the options, as a radio group of
// the system does.
//
//	radio.Group(c, radio.Props[string]{
//		Selected: &size,
//		Options: []radio.Option[string]{
//			{Value: "sm", Label: "Small"},
//			{Value: "lg", Label: "Large"},
//		},
//	})
package radio

import "github.com/egoist/mygo/ui"

// Option is one choice of a group.
type Option[T comparable] struct {
	// Value is the value chosen into Selected.
	Value T
	// Label is the text next to the circle.
	Label string
}

// Props describes the group to draw.
type Props[T comparable] struct {
	// Selected is the chosen value; the group edits it, and Changed
	// reports a new one.
	Selected *T
	// Options are the choices, in order.
	Options []Option[T]
}

// Group draws a column of the options for props, the chosen one in the
// accent color, and returns the group, so a view can chain more calls on
// it.
func Group[T comparable](c *ui.Context, p Props[T]) ui.Element {
	return ui.RadioGroup(c, func() {
		for _, o := range p.Options {
			ui.Radio(c, p.Selected, o.Value, o.Label)
		}
	})
}
