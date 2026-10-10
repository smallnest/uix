// Package marker provides a shadcn/ui-style Marker: a short inline
// label of a conversation, as "AI" or "You" before a turn of a chat. It
// has three looks: a filled pill, a pill with a border, and a label on
// a line across the turn.
//
//	marker.Marker(c, marker.Props{Label: "AI", Variant: marker.Border})
package marker

import "github.com/egoist/mygo/ui"

// Variant is the look of a marker.
type Variant int

const (
	// Default is a filled pill on a muted face.
	Default Variant = iota
	// Border is a pill with a border and no face.
	Border
	// Separator is a label on a line across the turn.
	Separator
)

// Props describes the marker to draw.
type Props struct {
	// Label is the text of the marker.
	Label string
	// Variant is the look of the marker; Default unless set.
	Variant Variant
}

// Marker draws the label for props. The filled and bordered looks are
// small pills; the Separator look is a line the label sits on, which
// fills the width of the view.
func Marker(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	if p.Variant == Separator {
		// The label on a line: the line grows to each side of it.
		return ui.Row(c).FillWidth().Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
			line(c, t)
			label(c, t, p.Label, ui.Color{}, ui.Color{})
			line(c, t)
		})
	}
	face, border := ui.Color{}, ui.Color{}
	if p.Variant == Default {
		face = t.SurfaceHover
	} else {
		border = t.Border
	}
	e := label(c, t, p.Label, face, border)
	return e
}

// label draws the text of the marker in its pill.
func label(c *ui.Context, t *ui.Theme, s string, face, border ui.Color) ui.Element {
	b := ui.Box(c).Padding(t.Space(0.5), t.Space(1.5)).Radius(999).Background(face)
	if border.A > 0 {
		b.Border(1, border)
	}
	b.Children(func() { ui.Text(c, s).SingleLine() })
	return b
}

// line draws the rule of the Separator look: a thin line in the border
// color.
func line(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Height(1).Grow(1).Background(t.Border)
}
