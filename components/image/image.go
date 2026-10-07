// Package image provides a bitmap or an SVG shown as a picture, for the
// Image of MyGo.
//
//	image.Image(c, image.Props{
//		Src:     photo,
//		Width:   120,
//		Height:  90,
//		Fit:     ui.ScaleDown,
//	})
//
// The picture shows at the size of its source, in DIPs, unless Width and
// Height fit it into another box, scaled by Fit: ScaleDown fits it within
// without enlarging it, and the default, Contain, fits it within whatever
// the size needs, keeping the aspect. Width alone fits it into a box of
// its own aspect, and Height alone into one of the aspect of the source.
// Grayscale draws it without color.
package image

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the picture to draw.
type Props struct {
	// Src is the picture: a *ui.Bitmap, or an *ui.SVG in its own colors.
	Src ui.ImageSource
	// Width and Height, when set, fit the picture into a box of that
	// size. One of them alone keeps the aspect of the source; neither
	// shows the picture at its own size.
	Width, Height float32
	// Fit scales the picture to the box: the zero value, Contain, fits
	// it within keeping the aspect, ScaleDown only where that is
	// smaller, Cover crops to cover it, and FillBox stretches it.
	Fit ui.Fit
	// Grayscale draws the picture in shades of gray.
	Grayscale bool
}

// Image draws the picture and returns it, so a view can chain more calls
// on it.
func Image(c *ui.Context, p Props) *ui.Element {
	e := ui.Image(c, p.Src)
	if p.Width > 0 || p.Height > 0 {
		switch {
		case p.Width > 0 && p.Height > 0:
			e.Size(p.Width, p.Height)
		case p.Width > 0:
			e.Width(p.Width)
		case p.Height > 0:
			e.Height(p.Height)
		}
	}
	if p.Fit != 0 {
		e.Fit(p.Fit)
	}
	if p.Grayscale {
		e.Grayscale()
	}
	return e
}
