// Package icon provides an SVG shown as an icon, in the color of the
// text around it, for the Icon of MyGo.
//
//	icon.Icon(c, icon.Props{
//		SVG:   save,
//		Label: "Save",
//	})
//
// The icon is as high as the font of the text around it, keeping the
// aspect of the SVG, and does not stretch across a column. The color
// follows the text: the theme of the Context, or a TextColor set on the
// row around it. Size fits it into another box, and Grayscale draws it
// without color. Icons are decorations that assistive technology does
// not see, unless Label names them.
package icon

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the icon to draw.
type Props struct {
	// SVG is the shape of the icon, in currentColor so it takes the
	// color of the text.
	SVG *ui.SVG
	// Label names the icon for assistive technology; without it the
	// icon is a decoration it skips.
	Label string
	// Width and Height, when set, fit the icon into a box of that size,
	// instead of the height of the font.
	Width, Height float32
	// Grayscale draws the icon in shades of gray.
	Grayscale bool
}

// Icon draws the SVG as an icon and returns it, so a view can chain more
// calls on it.
func Icon(c *ui.Context, p Props) *ui.Element {
	e := ui.Icon(c, p.SVG)
	if p.Label != "" {
		e.Label(p.Label)
	}
	if p.Width > 0 && p.Height > 0 {
		e.Size(p.Width, p.Height)
	}
	if p.Grayscale {
		e.Grayscale()
	}
	return e
}
