// Package calendar provides a month's calendar choosing a day, as a
// shadcn/ui component for the Calendar of MyGo, like the graphical date
// picker of SwiftUI.
//
//	calendar.Calendar(c, calendar.Props{Date: &due})
package calendar

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the calendar to draw.
type Props struct {
	// Date is the day chosen, which the calendar changes in place,
	// keeping its time of day and its location.
	Date *time.Time
}

// Calendar draws the month around Date, and returns it, so a view can
// chain more calls on it. A click chooses a day, as the arrows do while
// the calendar has the focus, Page Up and Page Down move by months, and
// Home and End go to the first and the last day of the month.
func Calendar(c *ui.Context, p Props) ui.Element {
	return ui.Calendar(c, p.Date)
}
