// Package toolbar provides a row of actions along the top of a window,
// from the mygo Toolbar widget. The controls are one stop of Tab; those
// that do not fit go, from the last, into a menu at the end.
package toolbar

import (
	"github.com/egoist/mygo/ui"
)

// Props configure the toolbar.
type Props struct {
	// Label names the toolbar for assistive technology.
	Label string
	// Children builds the controls: buttons, toggles, toggle groups
	// and menu buttons, which take no face of their own until hovered.
	// A Spacer pushes the controls after it to the end.
	Children func(c *ui.Context)
	// Disabled disables the controls of the toolbar.
	Disabled bool
}

// Toolbar draws the controls in a row, one stop of Tab, with Left and
// Right moving the focus among them, and Home and End to the first and
// the last.
func Toolbar(c *ui.Context, p Props) *ui.Element {
	tb := ui.Toolbar(c, func() {
		if p.Children != nil {
			p.Children(c)
		}
	})
	if p.Label != "" {
		tb.Label(p.Label)
	}
	if p.Disabled {
		tb.Disabled(true)
	}
	return tb
}
