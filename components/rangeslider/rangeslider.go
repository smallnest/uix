// Package rangeslider provides a shadcn/ui-style range slider: a track
// with two knobs that set the low and the high of a range, with a label
// and a live readout, from the RangeSlider of MyGo.
//
//	rangeslider.RangeSlider(c, rangeslider.Props{
//		Low:    &minPrice,
//		High:   &maxPrice,
//		Min:    0,
//		Max:    1000,
//		Step:   50,
//		Label:  "Price",
//		Format: func(lo, hi float64) string { return fmt.Sprintf("$%.0f – $%.0f", lo, hi) },
//	})
package rangeslider

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// Props describes the range slider to draw.
type Props struct {
	// Low and High are the bounds of the range, which the knobs set; Low
	// never passes High.
	Low, High *float64
	// Min and Max bound the whole range.
	Min, Max float64
	// Step snaps the values to multiples of it from Min, and shows tick
	// marks; 0 is continuous.
	Step float64
	// Label is the text above the track, next to the readout. It also
	// names the knobs, as "Price minimum" and "Price maximum".
	Label string
	// Format renders the readout from the two values; the default prints
	// them rounded.
	Format func(lo, hi float64) string
}

// RangeSlider draws the label, the readout and the track, and returns
// the column. The readout follows the knobs as the user drags, because
// the view builds every frame.
func RangeSlider(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	format := p.Format
	if format == nil {
		format = func(lo, hi float64) string { return fmt.Sprintf("%.0f – %.0f", lo, hi) }
	}
	return ui.Column(c).Gap(t.Space(0.5)).Children(func() {
		if p.Label != "" {
			ui.Row(c).Gap(t.Space(2)).Children(func() {
				ui.Text(c, p.Label).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
				ui.Text(c, format(*p.Low, *p.High)).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			})
		}
		r := ui.RangeSlider(c, p.Low, p.High, p.Min, p.Max, p.Step)
		r.Label(p.Label).FillWidth()
	})
}
