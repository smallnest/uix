// Package progress provides a shadcn/ui-style progress bar with a label
// and a live readout, from the Progress of MyGo.
//
//	progress.Progress(c, progress.Props{
//		Value: 0.64,
//		Label: "Disk usage",
//	})
package progress

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Props describes the progress bar to draw.
type Props struct {
	// Value is the share filled, 0 to 1. A negative value shows
	// activity of unknown length, which moves on its own.
	Value float64
	// Label is the text above the bar, next to the readout.
	Label string
	// Format renders the readout; the default prints the percentage.
	Format func(v float64) string
	// Indeterminate shows activity of unknown length instead of Value.
	Indeterminate bool
}

// Progress draws the label, the readout and the bar, and returns the
// column. The readout follows the value as the app changes it, because
// the view builds every frame.
func Progress(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	format := p.Format
	if format == nil {
		format = func(v float64) string { return fmt.Sprintf("%.0f%%", v*100) }
	}
	return ui.Column(c).Gap(t.Space(0.5)).Children(func() {
		if p.Label != "" {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, p.Label).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
				if !p.Indeterminate {
					ui.Text(c, format(p.Value)).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
				}
			})
		}
		value := p.Value
		if p.Indeterminate {
			value = -1
		}
		ui.Progress(c, value).FillWidth()
	})
}
