// Package timeinput provides a shadcn/ui-style TimeInput: a field of the
// time of day, whose hours and minutes each take the focus, stepped by
// the arrows and set by the digits typed.
//
//	timeinput.TimeInput(c, timeinput.Props{Value: &alarm})
package timeinput

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the time input to draw.
type Props struct {
	// Value is the time being edited; its date is kept.
	Value *time.Time
	// Disabled keeps the field from being edited.
	Disabled bool
}

// TimeInput draws a field editing the time of day of *Value and returns
// it, so a view can chain more calls on it. The field keeps the width
// its segments need; it does not stretch.
func TimeInput(c *ui.Context, p Props) *ui.Element {
	e := ui.TimeInput(c, p.Value)
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
