package theme

import (
	"fmt"

	"github.com/lucasb-eyer/go-colorful"
)

// CustomPalette is the custom-accent entry uix theme list shows. It has
// no fixed colors of its own: Custom derives them from the hex that uix
// theme set custom names.
var CustomPalette = Palette{
	Name:        "custom",
	Description: "your own accent color, from uix theme set custom <hex>",
	Radius:      8,
}

// Custom derives a palette from one accent color, the way the shadcn/ui
// theme builder turns a single swatch into a theme. The light appearance
// darkens the accent for hover and press; the dark appearance lightens
// it, and a very light accent keeps its color, so its hue does not wash
// out. The text on the accent reads white or near-black by its
// luminance, at the same split the built-in palettes use.
//
// Why not hand-roll the math as parseHex does: blending and luminance
// are real color math, and go-colorful is the mature, standard library
// for it; parseHex stays hand-rolled because eight lines of parsing beat
// a dependency.
func Custom(hex string) (*Palette, error) {
	r, g, b, err := parseHex(hex)
	if err != nil {
		return nil, err
	}
	a := colorful.Color{R: float64(r) / 255, G: float64(g) / 255, B: float64(b) / 255}
	light := Scale{
		Accent:        a.Hex(),
		AccentHover:   a.BlendRgb(colorful.Color{}, 0.12).Hex(),
		AccentPressed: a.BlendRgb(colorful.Color{}, 0.2).Hex(),
		AccentText:    textOn(a),
		Danger:        "#e11d48",
	}
	// The dark appearance lightens a mid or dark accent so it reads on
	// the dark background; a very light one stays as it is. The steps
	// keep the order of the built-in palettes: accent lightest, hover
	// one step down, pressed darkest.
	darkAccent := a
	darkHover, darkPressed := a.BlendRgb(colorful.Color{}, 0.12), a.BlendRgb(colorful.Color{}, 0.2)
	if luminance(a) <= 0.75 {
		darkAccent = a.BlendRgb(colorful.Color{R: 1, G: 1, B: 1}, 0.15)
		darkHover = a
		darkPressed = a.BlendRgb(colorful.Color{}, 0.12)
	}
	dark := Scale{
		Accent:        darkAccent.Hex(),
		AccentHover:   darkHover.Hex(),
		AccentPressed: darkPressed.Hex(),
		AccentText:    textOn(darkAccent),
		Danger:        "#f43f5e",
	}
	return &Palette{
		Name:        CustomPalette.Name,
		Description: fmt.Sprintf("your own accent, %s", a.Hex()),
		Radius:      CustomPalette.Radius,
		Light:       light,
		Dark:        dark,
	}, nil
}

// textOn returns the text color that reads on the accent: white on a
// dark accent, near-black on a light one, at the luminance split the
// built-in palettes already draw.
func textOn(a colorful.Color) string {
	if luminance(a) < 0.4 {
		return "#ffffff"
	}
	return "#0f172a"
}

// luminance returns the relative luminance of the color, the WCAG
// definition: the linear channels weighted by human sensitivity.
func luminance(a colorful.Color) float64 {
	r, g, b := a.LinearRgb()
	return 0.2126*r + 0.7152*g + 0.0722*b
}
