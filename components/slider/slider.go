// Package slider provides a shadcn/ui-style slider: a track and a knob
// that set a number, with a label and a live readout above it.
//
//	slider.Slider(c, slider.Props{
//		Value:  &volume,
//		Min:    0,
//		Max:    100,
//		Label:  "Volume",
//		Format: func(v float64) string { return fmt.Sprintf("%.0f%%", v) },
//	})
package slider

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Props describes the slider to draw.
type Props struct {
	// Value is the number being set.
	Value *float64
	// Min and Max bound the value.
	Min, Max float64
	// Step snaps the value to multiples of it from Min; 0 is continuous.
	Step float64
	// Label is the text above the track, next to the readout.
	Label string
	// Format renders the readout; the default prints the rounded value.
	Format func(v float64) string
}

// Slider draws the label, the readout and the track, and returns the
// column. The readout follows the value as the user drags, because the
// view builds every frame.
func Slider(c *ui.Context, p Props) *ui.Element {
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
		var s *ui.Element
		if p.Step > 0 {
			s = ui.StepSlider(c, p.Value, p.Min, p.Max, p.Step)
		} else {
			s = ui.Slider(c, p.Value, p.Min, p.Max)
		}
		s.FillWidth()
	})
}
