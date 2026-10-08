// Package pagination provides the pagination of BoardUI for the MyGo
// native toolkit: previous and next buttons and the page numbers, with
// the overflow of a long list of pages collapsing into dots around the
// current page.
//
//	pagination.Pagination(c, pagination.Props{
//		Page:       state.Page,
//		TotalPages: state.Pages,
//		OnChange:   func(p int) { state.Page = p },
//	})
//
// The previous and next buttons sit at the ends, disable at the first
// and last page, and read "Previous" and "Next" unless the app says
// otherwise. The page cells are 32 by 32; the current page is a cell
// with a face, the others fill on hover, and the dots are not buttons.
// The component is controlled: it shows what the app passes and reports
// a click through OnChange, so the app owns the page.
package pagination

import (
	"fmt"

	"github.com/egoist/mygo/ui"
)

// DOTS marks a gap in the page numbers, drawn as an ellipsis.
const DOTS = -1

// Props describes the pagination to draw.
type Props struct {
	// Page is the page shown, from 1.
	Page int
	// TotalPages is how many pages there are.
	TotalPages int
	// OnChange reports the page a button or cell chose.
	OnChange func(page int)
	// SiblingCount is how many pages show on each side of the current
	// page; 1 when zero.
	SiblingCount int
	// PreviousLabel and NextLabel name the end buttons; "Previous" and
	// "Next" when empty.
	PreviousLabel, NextLabel string
}

// Pagination draws the pagination and returns it, so a view can chain
// more calls on it. Nothing is drawn for one page, as BoardUI hides it.
func Pagination(c *ui.Context, p Props) ui.Element {
	if p.TotalPages <= 1 {
		return ui.Box(c)
	}
	t := c.Theme()
	sibling := p.SiblingCount
	if sibling == 0 {
		sibling = 1
	}
	prev, next := p.PreviousLabel, p.NextLabel
	if prev == "" {
		prev = "Previous"
	}
	if next == "" {
		next = "Next"
	}
	return ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Center).
		Gap(t.Space(2)).Children(func() {
		pagerButton(c, t, prev, false, p.Page <= 1, func() { p.OnChange(p.Page - 1) })
		ui.Row(c).Gap(2).Children(func() {
			for _, item := range pageRange(p.Page, p.TotalPages, sibling) {
				if item == DOTS {
					ui.Box(c).Size(32, 32).Center().Children(func() {
						ui.Text(c, "…").TextColor(t.TextMuted).SingleLine()
					})
					continue
				}
				cell(c, t, item, item == p.Page, func() { p.OnChange(item) })
			}
		})
		pagerButton(c, t, next, true, p.Page >= p.TotalPages, func() { p.OnChange(p.Page + 1) })
	})
}

// cell is one page number: 32 by 32 with rounded corners, the current
// page a raised cell and the others a ghost that fills on hover.
func cell(c *ui.Context, t *ui.Theme, page int, active bool, on func()) {
	b := ui.ButtonBase(c).Size(32, 32).Radius(t.Radius).Center()
	b.Label(fmt.Sprintf("Go to page %d", page))
	fg := t.Text
	if active {
		b.Background(t.Surface).Border(1, t.Border)
	} else {
		// The ghost fills as the pointer rests on it, read each frame so
		// the hover works without a rebuild.
		b.Draw(func(pp *ui.Painter, r ui.Rect) {
			if b.Hovered() {
				pp.Fill(r, t.SurfaceHover, t.Radius)
			}
		})
	}
	b.Children(func() {
		ui.Text(c, fmt.Sprintf("%d", page)).TextColor(fg).SingleLine()
	})
	if b.Clicked() {
		on()
	}
}

// pagerButton is the previous or next button: a chevron with the label,
// a face that follows the pointer, disabled at its end of the pages.
func pagerButton(c *ui.Context, t *ui.Theme, label string, next bool, disabled bool, on func()) {
	b := ui.ButtonBase(c)
	b.Label(label)
	fg := t.Text
	if disabled {
		fg = t.TextMuted
	}
	b.TextColor(fg).Padding(t.Space(1), t.Space(2.5)).Gap(t.Space(1.5)).
		Radius(t.Radius).Border(1, t.Border)
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if b.IsDisabled() {
			return
		}
		bg := t.Surface
		switch {
		case b.Pressed():
			bg = t.SurfacePressed
		case b.Hovered():
			bg = t.SurfaceHover
		}
		pp.Fill(r, bg, t.Radius)
	})
	b.Children(func() {
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1)).Children(func() {
			if next {
				ui.Text(c, label).SingleLine()
				chevron(c, t, fg, true)
			} else {
				chevron(c, t, fg, false)
				ui.Text(c, label).SingleLine()
			}
		})
	})
	if disabled {
		b.Disabled(true)
	}
	if b.Clicked() {
		on()
	}
}

// chevron is the arrow of the previous and next buttons.
func chevron(c *ui.Context, t *ui.Theme, color ui.Color, right bool) {
	ui.Box(c).Size(16, 16).Draw(func(pt *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		var p ui.Path
		if right {
			p.MoveTo(cx-3, cy-3.5).LineTo(cx+3, cy).LineTo(cx-3, cy+3.5)
		} else {
			p.MoveTo(cx+3, cy-3.5).LineTo(cx-3, cy).LineTo(cx+3, cy+3.5)
		}
		pt.StrokePath(&p, 1.8, color)
	})
}

// pageRange is the pages to show: the first, the last, the current with
// its siblings, and dots where the rest collapsed, as BoardUI collapses
// them.
func pageRange(current, total, sibling int) []int {
	if sibling*2+5 >= total {
		out := make([]int, total)
		for i := range out {
			out[i] = i + 1
		}
		return out
	}
	left := max(current-sibling, 1)
	right := min(current+sibling, total)
	showLeft := left > 2
	showRight := right < total-2
	switch {
	case !showLeft && showRight:
		return append(inclusive(1, 3+2*sibling), DOTS, total)
	case showLeft && !showRight:
		out := []int{1, DOTS}
		return append(out, inclusive(total-(2+2*sibling), total)...)
	default:
		out := []int{1, DOTS}
		out = append(out, inclusive(left, right)...)
		return append(out, DOTS, total)
	}
}

// inclusive is the pages from lo to hi, both ends in.
func inclusive(lo, hi int) []int {
	out := make([]int, hi-lo+1)
	for i := range out {
		out[i] = lo + i
	}
	return out
}
