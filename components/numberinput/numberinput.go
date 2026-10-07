// Package numberinput provides a shadcn/ui-style NumberInput: a field
// for a number between bounds, which the arrows and the buttons beside
// it step by step.
//
//	numberinput.NumberInput(c, numberinput.Props{
//		Value: &copies,
//		Lo:    1, Hi: 99, Step: 1,
//	})
package numberinput

import "github.com/egoist/mygo/ui"

// Props describes the number input to draw.
type Props struct {
	// Value is the number being edited.
	Value *float64
	// Lo is the smallest value allowed.
	Lo float64
	// Hi is the largest value allowed.
	Hi float64
	// Step is how much the arrows and the buttons change the value.
	Step float64
	// Disabled keeps the field from being edited.
	Disabled bool
}

// NumberInput draws a field editing *Value as a number between Lo and
// Hi, stepped by Step, and returns it, so a view can chain more calls on
// it. What is typed applies as soon as it is a number in range, and
// shows rounded to the decimals of Step once the field loses the focus.
func NumberInput(c *ui.Context, p Props) *ui.Element {
	e := ui.NumberInput(c, p.Value, p.Lo, p.Hi, p.Step)
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
