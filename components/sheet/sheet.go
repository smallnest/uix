// Package sheet provides a shadcn/ui-style Sheet: a panel that slides
// over the window from an edge, with the rest of the window dimmed and
// inert while it shows. It is for a secondary task next to the page, as
// a side panel of filters, where a dialog would cover too much.
//
//	sheet.Sheet(c, sheet.Props{
//		Open:  &open,
//		Title: "Filters",
//		Side:  sheet.Right,
//		Children: func() { filterForm(c) },
//	})
package sheet

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/animate"
)

// Side is the edge the sheet comes from.
type Side int

// The edges of the window.
const (
	// Right slides from the right, Left from the left, Top from the
	// top, and Bottom from the bottom.
	Right Side = iota
	Left
	Top
	Bottom
)

// Props describes the sheet to draw.
type Props struct {
	// Open is the bool that shows and hides the sheet; the backdrop
	// and Escape set it to false.
	Open *bool
	// Side is the edge the sheet comes from; Right when zero.
	Side Side
	// Title is the heading of the sheet.
	Title string
	// Size is the width of a side sheet, or the height of a top or
	// bottom one; 320 when zero.
	Size float32
	// Children draw in the body of the sheet.
	Children func()
}

// Sheet draws the sheet for props, or nothing while it is closed, and
// returns the panel. The backdrop dims the window; clicking it or
// pressing Escape closes the sheet.
func Sheet(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	size := p.Size
	if size <= 0 {
		size = 320
	}
	return ui.DialogBase(c, p.Open, func(back, panel ui.Element) {
		// The sheet animates through the backdrop, whose parent (the
		// window's overlay) predates the sheet, so its Enter fires and
		// the panel inherits it; a panel Enter would never run.
		back.Background(ui.RGBA(0, 0, 0, 0.4)).
			Transition(animate.FromEdge(directionFor(p.Side), animate.Options{Distance: 24}))
		panel.Padding(t.Space(5)).Radius(t.Radius+2).Background(t.Background).
			Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3))
		switch p.Side {
		case Left:
			panel.Absolute().Left(0).Top(0).Bottom(0).Width(size)
		case Top:
			panel.Absolute().Left(0).Top(0).Right(0).Height(size)
		case Bottom:
			panel.Absolute().Left(0).Right(0).Bottom(0).Height(size)
		default:
			panel.Absolute().Right(0).Top(0).Bottom(0).Width(size)
		}
		panel.Children(func() {
			ui.Column(c).Fill().Gap(t.Space(2)).Children(func() {
				if p.Title != "" {
					ui.Text(c, p.Title).Bold()
				}
				if p.Children != nil {
					p.Children()
				}
			})
		})
	})
}

// directionFor maps the side the sheet lives on to the edge it slides
// over from: a right sheet comes from the right, and so on.
func directionFor(s Side) animate.Direction {
	switch s {
	case Left:
		return animate.Left
	case Top:
		return animate.Up
	case Bottom:
		return animate.Down
	}
	return animate.Right
}
