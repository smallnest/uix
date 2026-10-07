// Package datepicker provides a shadcn/ui-style DatePicker: a field
// showing the date, whose calendar below it changes it, the arrows and
// Enter moving among the days and Page Up and Page Down among the
// months.
//
//	datepicker.DatePicker(c, datepicker.Props{Value: &due})
package datepicker

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the date picker to draw.
type Props struct {
	// Value is the date being edited; its time of day is kept.
	Value *time.Time
	// Disabled keeps the field from opening its calendar.
	Disabled bool
}

// DatePicker draws a field editing *Value as a date, which a calendar
// below it changes, and returns it, so a view can chain more calls on
// it. The field stretches over its parent, as the other fields do; a
// bounded width overrides it in a Row.
func DatePicker(c *ui.Context, p Props) *ui.Element {
	e := ui.DateInput(c, p.Value)
	e.FillWidth()
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
