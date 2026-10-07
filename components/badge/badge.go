// Package badge provides a shadcn/ui-style badge: a short label in a
// pill, for counts and status.
//
//	badge.Badge(c, badge.Props{Label: "Beta", Variant: badge.Secondary})
package badge

import "github.com/egoist/mygo/ui"

// Variant is the look of a badge.
type Variant int

const (
	// Default is a solid badge in the text color.
	Default Variant = iota
	// Secondary is a badge on a muted face.
	Secondary
	// Destructive is a badge in the danger color.
	Destructive
	// Outline is a bordered badge without a face.
	Outline
)

// Props describes the badge to draw.
type Props struct {
	// Label is the text of the badge.
	Label string
	// Variant is the look of the badge; Default unless set.
	Variant Variant
}

// Badge draws a pill with the label for props.
func Badge(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	face, fg, border := variantColors(t, p.Variant)
	b := ui.Box(c).Padding(t.Space(0.5), t.Space(1.5)).Radius(999).Background(face).TextColor(fg)
	if border.A > 0 {
		b.Border(1, border)
	}
	b.Children(func() { ui.Text(c, p.Label).SingleLine() })
	return b
}

// variantColors returns the face, text and border colors of a variant.
func variantColors(t *ui.Theme, v Variant) (face, fg, border ui.Color) {
	switch v {
	case Secondary:
		return t.SurfaceHover, t.Text, ui.Color{}
	case Destructive:
		return t.Danger, ui.Hex("#ffffff"), ui.Color{}
	case Outline:
		return ui.Color{}, t.Text, t.Border
	}
	// Default is the theme turned over: text on the text color, so it
	// reads in both appearances.
	return t.Text, t.InverseText, ui.Color{}
}
