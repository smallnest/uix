// Package spinner provides a shadcn/ui-style spinner: a ring of spokes
// that turns while work goes on, from the Spinner of MyGo.
//
//	spinner.Spinner(c, spinner.Props{Label: "Scanning"})
package spinner

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the spinner to draw.
type Props struct {
	// Label names the spinner to assistive technology. It is not drawn.
	Label string
}

// Spinner draws the turning spokes, and returns the element. It turns on
// its own, repainting without rebuilding the view, so a view that shows
// it does not have to animate.
func Spinner(c *ui.Context, p Props) *ui.Element {
	e := ui.Spinner(c)
	if p.Label != "" {
		e.Label(p.Label)
	}
	return e
}
