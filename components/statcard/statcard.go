// Package statcard provides the KPI stat cards of BoardUI for the MyGo
// native toolkit: a grid of metric cards in two looks, the compact plain
// card and the display-size footer card.
//
//	statcard.StatCards(c, statcard.Props{
//		Variant: statcard.Footer,
//		Stats: []statcard.Stat{
//			{Icon: statcard.IconCoins, Label: "Total revenue",
//				Value: "$152,313.92", Delta: "16%", DeltaColor: statcard.Up,
//				Tone: statcard.Blue, Hint: "Gross revenue, before refunds."},
//		},
//	})
//
// The cards sit on the theme's Surface, so a row of them reads as a
// dashboard. Each card is a component of its own; the grid is a row of
// equal cards that wraps when it must. The delta chip tones the change
// by direction, the footer variant tones its icon tile, and the hint
// glyph carries a tooltip.
package statcard

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/tooltip"
)

// Variant is the look of a card.
type Variant int

const (
	// Plain is the compact card: an icon tile, the label, the value and
	// the delta chip.
	Plain Variant = iota
	// Footer adds a tinted icon tile, a display-size value and a footer
	// band that carries the comparison caption and the delta pill.
	Footer
)

// DeltaColor is the tone of the delta chip, by the direction of the
// change.
type DeltaColor int

const (
	// Up is a change for the better, the lime chip.
	Up DeltaColor = iota
	// Down is a change for the worse, the rose chip.
	Down
	// Flat is a change without a direction, the neutral chip.
	Flat
)

// Tone tints the footer variant's icon tile.
type Tone int

const (
	// The tones, from BoardUI's StatTone, as the theme has a color for
	// them: Blue is the Accent, Orange the Warning, Pink the Danger,
	// Emerald the Success; Purple and Sky share the Accent, as the theme
	// has no violet or sky of its own.
	Blue Tone = iota
	Orange
	Purple
	Pink
	Sky
	Emerald
)

// Icon is a card icon, drawn in a box of the card's icon size. The
// exported icons (IconUsers, IconBox, IconBasket, IconChat, IconCoins,
// IconRefund) fit the box; a card can draw its own instead.
type Icon func(c *ui.Context, color ui.Color)

// Stat is one metric card.
type Stat struct {
	// Icon is the mark of the metric, on the plain tile or the footer
	// tile.
	Icon Icon
	// Label is the name of the metric.
	Label string
	// Value is the metric itself, as the app formats it.
	Value string
	// Delta is the change, such as "+5.3%".
	Delta string
	// DeltaColor tones the chip.
	DeltaColor DeltaColor
	// Tone tints the footer variant's icon tile; Blue when zero.
	Tone Tone
	// Caption is the footer band's comparison, "From last month" when
	// empty.
	Caption string
	// Hint shows the info glyph with this text in a tooltip; empty hides
	// the glyph.
	Hint string
}

// Props describes the grid to draw.
type Props struct {
	// Stats are the cards, in order, one per column.
	Stats []Stat
	// Variant is the look of every card.
	Variant Variant
}

// StatCards draws the grid of cards and returns it, so a view can chain
// more calls on it.
func StatCards(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	return ui.Row(c).FillWidth().Wrap().Gap(t.Space(4)).Children(func() {
		for _, s := range p.Stats {
			card(c, t, p.Variant, s)
		}
	})
}

// card is one metric card, in its variant's look.
func card(c *ui.Context, t *ui.Theme, v Variant, s Stat) {
	if v == Footer {
		footerCard(c, t, s)
		return
	}
	plainCard(c, t, s)
}

// plainCard is the compact card: the icon tile on top, the label and the
// value with its delta below, 132 points tall as the BoardUI cards are.
func plainCard(c *ui.Context, t *ui.Theme, s Stat) {
	ui.Box(c).Height(132).Grow(1).Shrink(0).Background(t.Surface).
		Radius(t.Radius * 2).Padding(t.Space(4)).Children(func() {
		tile(c, t, s.Icon, 30, t.SurfaceHover, t.Text)
		ui.Column(c).FillWidth().Gap(2).Children(func() {
			ui.Text(c, s.Label).TextColor(t.TextMuted).SingleLine()
			ui.Row(c).FillWidth().AlignItems(ui.Center).Gap(t.Space(2)).
				Children(func() {
					ui.Text(c, s.Value).FontSize(t.FontSize * 1.15).
						FontWeight(600).TextColor(t.Text).SingleLine()
					deltaChip(c, t, s.Delta, s.DeltaColor)
				})
		})
	})
}

// footerCard is the display card: the tinted icon tile with the hint
// glyph above the label and value, and the footer band with the
// comparison and the delta pill.
func footerCard(c *ui.Context, t *ui.Theme, s Stat) {
	ui.Box(c).Grow(1).Shrink(0).Background(t.Surface).
		Radius(t.Radius * 2).Padding(t.Space(2)).Children(func() {
		ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Start).
			Gap(t.Space(2.5)).Padding(t.Space(2)).Children(func() {
			tile(c, t, s.Icon, 40, toneColor(t, s.Tone), t.InverseText)
			if s.Hint != "" {
				hintGlyph(c, t, s)
			}
		})
		ui.Column(c).Gap(2).Padding(t.Space(2), t.Space(2), t.Space(3.5), t.Space(2)).
			Children(func() {
				ui.Text(c, s.Label).TextColor(t.TextMuted).SingleLine()
				ui.Text(c, s.Value).FontSize(t.FontSize * 1.6).
					FontWeight(600).TextColor(t.Text).SingleLine()
			})
		caption := s.Caption
		if caption == "" {
			caption = "From last month"
		}
		ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Center).
			Gap(t.Space(2)).Background(t.Background).Radius(t.Radius + 2).
			Padding(t.Space(2.5), t.Space(1.5), t.Space(2.5), t.Space(2.5)).
			Children(func() {
				ui.Text(c, caption).TextColor(t.TextMuted).Grow(1).SingleLine()
				deltaChip(c, t, s.Delta, s.DeltaColor)
			})
	})
}

// tile is the rounded box the icon sits in: a plain tile on a muted
// ground, or a footer tile tinted with the tone.
func tile(c *ui.Context, t *ui.Theme, icon Icon, size float32, bg, fg ui.Color) {
	if icon == nil {
		icon = IconSparkle
	}
	ui.Box(c).Size(size, size).Radius(size / 3).Background(bg).Center().
		Children(func() {
			icon(c, fg)
		})
}

// toneColor is the fill of a footer tile for its tone.
func toneColor(t *ui.Theme, tone Tone) ui.Color {
	switch tone {
	case Orange:
		return t.Warning
	case Pink:
		return t.Danger
	case Emerald:
		return t.Success
	}
	return t.Accent
}

// hintGlyph is the info mark with the hint as a tooltip, as the footer
// cards explain themselves.
func hintGlyph(c *ui.Context, t *ui.Theme, s Stat) {
	b := ui.ButtonBase(c).Size(20, 20).Radius(10).Center()
	b.Label("About " + s.Label)
	b.Children(func() {
		info(c, t.TextMuted)
	})
	tooltip.Tooltip(c, tooltip.Props{Anchor: b, Text: s.Hint})
}

// deltaChip is the pill of the change, tinted by its direction with an
// arrow that says the same.
func deltaChip(c *ui.Context, t *ui.Theme, delta string, d DeltaColor) {
	bg, fg := t.SurfaceHover, t.TextMuted
	arrow := arrowFlat
	switch d {
	case Up:
		bg, fg = t.Success.Alpha(0.15), t.Success
		arrow = arrowUp
	case Down:
		bg, fg = t.Danger.Alpha(0.15), t.Danger
		arrow = arrowDown
	}
	ui.Row(c).Shrink(0).Gap(3).Radius(9).Padding(2, 4).
		AlignItems(ui.Center).Background(bg).Children(func() {
		arrow(c, fg)
		ui.Text(c, delta).FontSize(t.FontSize * 0.8).TextColor(fg).SingleLine()
	})
}

// IconUsers draws two people, the customers of a metric.
func IconUsers(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.Circle(cx-4.5, cy-5.5, 2.4)
		p.Circle(cx+4.5, cy-5.5, 2.4)
		pt.FillPath(&p, color)
		var s ui.Path
		s.MoveTo(cx-7.5, cy+3).QuadTo(cx-4.5, cy+0.5, cx-1.5, cy+3)
		s.MoveTo(cx+1.5, cy+3).QuadTo(cx+4.5, cy+0.5, cx+7.5, cy+3)
		pt.StrokePath(&s, 1.6, color)
	})
}

// IconBox draws a package, the units sold of a metric.
func IconBox(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.MoveTo(cx-5, cy-1).LineTo(cx+5, cy-1).LineTo(cx+5, cy+6).
			LineTo(cx-5, cy+6).Close()
		p.MoveTo(cx-5, cy-1).LineTo(cx-5, cy-4).LineTo(cx+5, cy-4).LineTo(cx+5, cy-1)
		p.MoveTo(cx, cy-1).LineTo(cx, cy+6)
		pt.StrokePath(&p, 1.6, color)
	})
}

// IconBasket draws a shopping basket, the orders of a metric.
func IconBasket(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.MoveTo(cx-3.5, cy-1.5).QuadTo(cx-3.5, cy-6, cx, cy-6).
			QuadTo(cx+3.5, cy-6, cx+3.5, cy-1.5)
		p.MoveTo(cx-6, cy-1.5).LineTo(cx+6, cy-1.5).LineTo(cx+4, cy+6).
			LineTo(cx-4, cy+6).Close()
		pt.StrokePath(&p, 1.6, color)
	})
}

// IconChat draws a smiling bubble, the support tickets of a metric.
func IconChat(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.Circle(cx, cy-1.5, 6.5)
		pt.StrokePath(&p, 1.6, color)
		var eyes ui.Path
		eyes.Circle(cx-2.5, cy-2.5, 1.1)
		eyes.Circle(cx+2.5, cy-2.5, 1.1)
		pt.FillPath(&eyes, color)
		var smile ui.Path
		smile.MoveTo(cx-3.5, cy+0.5).QuadTo(cx, cy+3, cx+3.5, cy+0.5)
		pt.StrokePath(&smile, 1.6, color)
	})
}

// IconCoins draws a coin, the money of a metric.
func IconCoins(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.Circle(cx, cy, 6.5)
		pt.StrokePath(&p, 1.6, color)
		pt.Line(cx, cy-4.5, cx, cy+4.5, 1.6, color)
		pt.Line(cx-2.5, cy-1.5, cx+2.5, cy-1.5, 1.6, color)
		pt.Line(cx-2.5, cy+1.5, cx+2.5, cy+1.5, 1.6, color)
	})
}

// IconRefund draws a return arrow, the refunds of a metric.
func IconRefund(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.MoveTo(cx+5.5, cy-2).LineTo(cx+5.5, cy+3).
			QuadTo(cx+5.5, cy+6, cx+2.5, cy+6).LineTo(cx-3.5, cy+6)
		p.QuadTo(cx-6.5, cy+6, cx-6.5, cy+3).LineTo(cx-6.5, cy+1)
		p.MoveTo(cx-4, cy+3.5).LineTo(cx-6.5, cy+1).LineTo(cx-9, cy+3.5)
		pt.StrokePath(&p, 1.6, color)
	})
}

// IconSparkle is the fallback mark of a card without an icon, the
// four-point star.
func IconSparkle(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var p ui.Path
		p.MoveTo(cx, cy-6).QuadTo(cx+2, cy-2, cx+6, cy).
			QuadTo(cx+2, cy+2, cx, cy+6).QuadTo(cx-2, cy+2, cx-6, cy).
			QuadTo(cx-2, cy-2, cx, cy-6).Close()
		pt.FillPath(&p, color)
	})
}

// info draws the i of the hint glyph.
func info(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		var ring ui.Path
		ring.Circle(cx, cy, 5.5)
		pt.StrokePath(&ring, 1.2, color)
		pt.Line(cx, cy-2, cx, cy+1, 1.4, color)
		var dot ui.Path
		dot.Circle(cx, cy+3.2, 0.9)
		pt.FillPath(&dot, color)
	})
}

// arrowUp is the chip's arrow for a better change.
func arrowUp(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		pt.Line(cx, cy+3.5, cx, cy-3.5, 1.6, color)
		var p ui.Path
		p.MoveTo(cx-2.5, cy-1).LineTo(cx, cy-3.5).LineTo(cx+2.5, cy-1)
		pt.StrokePath(&p, 1.6, color)
	})
}

// arrowDown is the chip's arrow for a worse change.
func arrowDown(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		pt.Line(cx, cy-3.5, cx, cy+3.5, 1.6, color)
		var p ui.Path
		p.MoveTo(cx-2.5, cy+1).LineTo(cx, cy+3.5).LineTo(cx+2.5, cy+1)
		pt.StrokePath(&p, 1.6, color)
	})
}

// arrowFlat is the chip's bar for a change without a direction.
func arrowFlat(c *ui.Context, color ui.Color) {
	glyph(c, func(pt *ui.Painter, cx, cy float32) {
		pt.Line(cx-3.5, cy, cx+3.5, cy, 1.6, color)
	})
}

// glyph runs fn over a 20-point drawing centered in its box.
func glyph(c *ui.Context, fn func(pt *ui.Painter, cx, cy float32)) {
	ui.Box(c).Size(20, 20).Draw(func(pt *ui.Painter, r ui.Rect) {
		fn(pt, r.X+r.W/2, r.Y+r.H/2)
	})
}
