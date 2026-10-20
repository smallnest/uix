package gauge

import (
	"strconv"

	"github.com/egoist/mygo/ui"
)

// TextAnchor is where a label sits relative to its position.
type TextAnchor int

const (
	// TextStart puts the label's left edge at its position.
	TextStart TextAnchor = iota
	// TextMiddle centers the label on its position.
	TextMiddle
	// TextEnd puts the label's right edge at its position.
	TextEnd
)

// FontFamily is the family a label is drawn in.
type FontFamily int

const (
	// FontSans is the theme's own family and the zero value, so a label
	// with no family asked for follows the app.
	FontSans FontFamily = iota
	// FontMono is a monospaced family, which holds the digits of a
	// readout in place while its value changes.
	FontMono
	// FontRounded is a rounded family, the softer face of a dial's large
	// reading. Not every platform has one, so this asks for the nearest.
	FontRounded
)

// family is the family list this font asks the shaper for. The shaper
// takes a list, so each falls back through what the platform is likely
// to have.
func (f FontFamily) family(theme string) string {
	switch f {
	case FontMono:
		return "monospace"
	case FontRounded:
		return "SF Pro Rounded, ui-rounded, Nunito, Verdana, sans-serif"
	default:
		return theme
	}
}

// textAt draws s, its center at the device point (x, y), at the design
// font size, measured and drawn by the same shaping so the label is
// exactly centered. A value in tabular figures, so digits line up as the
// needle moves.
func (g *Canvas) textAt(s string, x, y float32, size float32, weight int, anchor TextAnchor, color ui.Color, f FontFamily) {
	if s == "" {
		return
	}
	fs := g.fontSize(size)
	layout := ui.ShapeText(s, ui.Font{
		Family:   f.family(g.t.Font),
		Size:     fs,
		Weight:   weight,
		Features: "tnum",
	}, 0)
	w, h := layout.Size()
	tx := x
	switch anchor {
	case TextMiddle:
		tx = x - w/2
	case TextEnd:
		tx = x - w
	}
	g.pt.TextLayout(layout, tx, y-h/2, color)
}

// TextOptions draws a free-standing label anywhere on the gauge.
type TextOptions struct {
	// X and Y position the label relative to the gauge center, in design
	// units; positive Y is down.
	X, Y float32
	// FontSize is the label size in design units; 24 when zero.
	FontSize float32
	// Font is the family; the theme's when zero.
	Font FontFamily
	// Color is the label; the theme's text when zero.
	Color ui.Color
	// Weight is the font weight; 500 when zero.
	Weight int
	// Anchor is which edge of the label sits at its position.
	Anchor TextAnchor
	// Text is the label.
	Text string
}

// Text draws a free-standing label anywhere on the gauge, as a gauge-ui
// GaugeText.
func (g *Canvas) Text(o TextOptions) {
	size := o.FontSize
	if size == 0 {
		size = 24
	}
	weight := o.Weight
	if weight == 0 {
		weight = 500
	}
	x, y := g.xy(o.X, o.Y)
	g.textAt(o.Text, x, y, size, weight, o.Anchor, g.color(o.Color), o.Font)
}

// ValueOptions draws the current value as text.
type ValueOptions struct {
	// X and Y position the readout relative to the gauge center, in
	// design units; positive Y is down.
	X, Y float32
	// FontSize is the readout size in design units; 96 when zero.
	FontSize float32
	// Font is the family; the theme's when zero. A readout that holds
	// still as it changes asks for FontMono.
	Font FontFamily
	// Color is the readout; the theme's text when zero.
	Color ui.Color
	// Weight is the font weight; 600 when zero.
	Weight int
	// Decimals rounds the readout; 0 when zero.
	Decimals int
	// Format renders the readout in place of the rounded value; a range
	// formats its two ends apart.
	Format func(value float64) string
	// Separator joins the two ends of a range; "–" when empty.
	Separator string
	// End shows only one end of a range rather than both.
	End End
}

// Value draws the current value as text, or a range as its two ends, as
// a gauge-ui GaugeValue.
func (g *Canvas) Value(o ValueOptions) {
	size := o.FontSize
	if size == 0 {
		size = 96
	}
	weight := o.Weight
	if weight == 0 {
		weight = 600
	}
	separator := o.Separator
	if separator == "" {
		separator = "–"
	}
	text := func(v float64) string {
		if o.Format != nil {
			return o.Format(v)
		}
		return strconv.FormatFloat(v, 'f', o.Decimals, 64)
	}
	var s string
	switch {
	case o.End == EndFrom:
		s = text(g.from)
	case g.rangeMode:
		s = text(g.from) + separator + text(g.value)
	default:
		s = text(g.value)
	}
	x, y := g.xy(o.X, o.Y)
	g.textAt(s, x, y, size, weight, TextMiddle, g.color(o.Color), o.Font)
}
