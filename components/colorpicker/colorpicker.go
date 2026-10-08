// Package colorpicker provides a shadcn/ui-style ColorPicker: a swatch
// of the color, whose picker below it changes it — a square of
// saturation and brightness, sliders of the hue and the opacity, the
// color in hex, and swatches of common colors.
//
//	colorpicker.ColorPicker(c, colorpicker.Props{Value: &accent})
package colorpicker

import "github.com/egoist/mygo/ui"

// Props describes the color picker to draw.
type Props struct {
	// Value is the color being edited.
	Value *ui.Color
	// Disabled keeps the swatch from opening its picker.
	Disabled bool
}

// ColorPicker draws a swatch of *Value, which a picker below it
// changes, and returns it, so a view can chain more calls on it.
func ColorPicker(c *ui.Context, p Props) ui.Element {
	e := ui.ColorWell(c, p.Value)
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
