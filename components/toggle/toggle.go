// Package toggle provides a shadcn/ui-style Toggle: a button that stays
// pressed while it is on, as Bold in an editor's toolbar, which a click
// or Space while it has the focus turns over.
//
//	toggle.Toggle(c, toggle.Props{On: &bold, Label: "B"})
package toggle

import "github.com/egoist/mygo/ui"

// Props describes the toggle to draw.
type Props struct {
	// On is whether the toggle is pressed.
	On *bool
	// Label is the text the toggle shows, such as "Bold".
	Label string
	// Disabled keeps the toggle from being pressed.
	Disabled bool
}

// Toggle draws a button that stays pressed while *On and returns it, so
// a view can chain more calls on it.
func Toggle(c *ui.Context, p Props) *ui.Element {
	e := ui.Toggle(c, p.On, p.Label)
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
