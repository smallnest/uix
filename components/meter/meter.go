// Package meter provides a shadcn/ui-style meter: a bar that shows a
// measure against its range, colored by how near the top it is, with a
// label and a live readout, from the Meter of MyGo.
//
//	meter.Meter(c, meter.Props{
//		Value:  0.78,
//		Max:    1,
//		Label:  "Disk usage",
//		Levels: &ui.MeterLevels{Warning: 0.7, Critical: 0.9},
//	})
package meter

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Props describes the meter to draw.
type Props struct {
	// Value is the measure, between Min and Max.
	Value float64
	// Min and Max bound the measure. A Max at or below Min is taken as 1
	// above Min, so a bare fraction fills the bar.
	Min, Max float64
	// Levels turn the bar the Warning color from its Warning share of the
	// range, and the Danger color from Critical. The zero value keeps the
	// Success color throughout. A Critical below Warning makes low values
	// the bad ones, as a battery's.
	Levels *ui.MeterLevels
	// Label is the text above the bar, next to the readout.
	Label string
	// Format renders the readout; the default prints the rounded value.
	Format func(v float64) string
}

// Meter draws the label, the readout and the bar, and returns the
// column. The readout follows the value as the app changes it, because
// the view builds every frame.
func Meter(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	format := p.Format
	if format == nil {
		format = func(v float64) string { return fmt.Sprintf("%.0f", v) }
	}
	hi := p.Max
	if hi <= p.Min {
		hi = p.Min + 1
	}
	return ui.Column(c).Gap(t.Space(0.5)).Children(func() {
		if p.Label != "" {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, p.Label).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
				ui.Text(c, format(p.Value)).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			})
		}
		ui.Meter(c, p.Value, p.Min, hi, p.Levels).FillWidth()
	})
}
