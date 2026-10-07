// Package segmented provides a segmented control: a row of segments of
// which one is chosen, as a shadcn/ui component for the Segmented of
// MyGo, a switch between views.
//
//	segmented.Segmented(c, segmented.Props{
//		Selected: &view,
//		Labels:   []string{"List", "Grid"},
//	})
package segmented

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the segmented control to draw.
type Props struct {
	// Selected is the index of the segment chosen, which a click or the
	// arrow keys change.
	Selected *int
	// Labels are the segments.
	Labels []string
	// OnChange runs when a click or the arrows choose another segment.
	OnChange func(i int)
	// Label names the control to assistive technology.
	Label string
}

// Segmented draws the segments and returns the control. A click chooses
// a segment, and the arrow keys move the choice while the control has
// the focus; the chosen segment reads on a raised background.
func Segmented(c *ui.Context, p Props) *ui.Element {
	e := ui.Segmented(c, p.Selected, p.Labels...)
	if p.Label != "" {
		e.Label(p.Label)
	}
	if e.Changed() && p.OnChange != nil {
		p.OnChange(*p.Selected)
	}
	return e
}
