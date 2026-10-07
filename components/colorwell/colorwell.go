// Package colorwell provides a swatch of a color that opens a picker
// below it, as AppKit's color well does, for the ColorWell of MyGo.
//
//	well := colorwell.ColorWell(c, colorwell.Props{
//		Color: &tint,
//		Label: "Tint",
//	})
//	if well.Changed() {
//		save()
//	}
//
// A click, Enter or Space opens the picker, and Escape or a click outside
// closes it. The picker edits the color by swatch, by channel, by hex or
// on a square of hue and brightness. Changed reports a new color, so the
// view can save it. Assistive technology sees a color well whose value
// is the color in hex, named by Label.
package colorwell

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the swatch to draw.
type Props struct {
	// Color is the color being set; the well shows it and the picker
	// edits it.
	Color *ui.Color
	// Label names the well for assistive technology.
	Label string
}

// ColorWell draws a swatch of *Color that opens a ColorPicker below it,
// and returns it, so a view can chain more calls on it.
func ColorWell(c *ui.Context, p Props) *ui.Element {
	w := ui.ColorWell(c, p.Color)
	if p.Label != "" {
		w.Label(p.Label)
	}
	return w
}
