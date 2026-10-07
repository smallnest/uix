// Package stepper provides a shadcn/ui-style stepper: a number with two
// arrow buttons that step it, and a label and a live readout, from the
// Stepper of MyGo.
//
//	stepper.Stepper(c, stepper.Props{
//		Value: &port,
//		Min:   1024,
//		Max:   65535,
//		Step:  1,
//		Label: "Port",
//	})
package stepper

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Props describes the stepper to draw.
type Props struct {
	// Value is the number being set.
	Value *float64
	// Min and Max bound the value.
	Min, Max float64
	// Step is how much the arrows and the arrow keys move the value.
	Step float64
	// Label is the text above the control, next to the readout.
	Label string
	// Format renders the readout; the default prints the rounded value.
	Format func(v float64) string
}

// Stepper draws the label, the readout and the arrows, and returns the
// column. The readout follows the value as the user clicks or holds an
// arrow — it repeats, faster after a while — or presses the arrow keys,
// Home and End while the control has the focus, because the view builds
// every frame.
func Stepper(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	format := p.Format
	if format == nil {
		format = func(v float64) string { return fmt.Sprintf("%.0f", v) }
	}
	return ui.Column(c).Gap(t.Space(0.5)).Children(func() {
		if p.Label != "" {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, p.Label).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
				ui.Text(c, format(*p.Value)).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			})
		}
		s := ui.Stepper(c, p.Value, p.Min, p.Max, p.Step)
		s.Label(p.Label)
	})
}
