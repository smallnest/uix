// Package linechart provides the RevenueChartCard of BoardUI for the
// MyGo native toolkit: a card of a year of monthly values drawn against
// the year before. The current year is the filled area and the solid
// line, the year before the dashed line behind it, so the gap between
// them is the story. Resting the pointer on a month swaps the headline
// for that month and shows what it was a year earlier.
//
//	linechart.LineChart(c, linechart.Props{
//		Data:     months,
//		Selected: &active,
//		Title:    "Revenue",
//	})
//
// The card is a rounded panel on the theme's Surface with the headline
// and legend above the plot. The chart needs no legend text of its own
// — the two series read by color — but the legend names them. The card
// fills the height it is given, so a dashboard row can size it with
// Grow.
package linechart

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/egoist/mygo/ui"
)

// Point is one month of the series.
type Point struct {
	// Label is the short month name, such as "Jan", on the axis.
	Label string
	// Current is the value of the current year.
	Current float64
	// Previous is the value of the year before.
	Previous float64
}

// Props describes the chart to draw.
type Props struct {
	// Data are the months, in order.
	Data []Point
	// Selected is the month under the pointer, by index, -1 for none.
	// The chart keeps it as the pointer moves over the plot and clears
	// it as the pointer leaves; pass nil to draw a chart that does not
	// follow the pointer.
	Selected *int
	// Title is the headline's label when no month is under the pointer.
	Title string
}

// months are the full names of the year, for the headline of the month
// under the pointer.
var months = []string{
	"January", "February", "March", "April", "May", "June",
	"July", "August", "September", "October", "November", "December",
}

// LineChart draws the card and returns it, so a view can chain more
// calls on it.
func LineChart(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	selected := -1
	if p.Selected != nil {
		selected = *p.Selected
	}
	title := p.Title
	if title == "" {
		title = "Revenue"
	}
	label, current, previous, delta, up := headline(t, p.Data, selected, title)
	return ui.Column(c).Fill().Background(t.Surface).Radius(t.Radius * 2).
		Padding(t.Space(4), t.Space(4), t.Space(3), t.Space(4)).
		Gap(t.Space(4)).Children(func() {
		chartHeader(c, t, label, current, previous, delta, up)
		plot(c, t, p)
	})
}

// headline is what the header says: the total, or the month under the
// pointer and its year-earlier figure.
func headline(t *ui.Theme, data []Point, selected int, title string) (label, current, previous, delta string, up bool) {
	if selected >= 0 && selected < len(data) {
		pt := data[selected]
		label = pt.Label
		if len(pt.Label) >= 3 {
			if i := strings.Index("JanFebMarAprMayJunJulAugSepOctNovDec", pt.Label[:3]); i >= 0 {
				label = months[i/3]
			}
		}
		current, previous = formatMoney(pt.Current), formatMoney(pt.Previous)
		delta, up = deltaOf(pt.Current, pt.Previous)
		return
	}
	var sumCur, sumPrev float64
	for _, pt := range data {
		sumCur += pt.Current
		sumPrev += pt.Previous
	}
	label = title
	current, previous = formatMoney(sumCur), formatMoney(sumPrev)
	delta, up = deltaOf(sumCur, sumPrev)
	return
}

// chartHeader is the headline over the plot: the label, the value with
// its delta chip, the comparison under it, and the legend on the right.
func chartHeader(c *ui.Context, t *ui.Theme, label, current, previous, delta string, up bool) {
	ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Start).
		Gap(t.Space(3)).Children(func() {
		ui.Column(c).Gap(3).Children(func() {
			ui.Text(c, label).TextColor(t.TextMuted).SingleLine()
			ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
				ui.Text(c, current).FontSize(t.FontSize * 1.2).FontWeight(600).
					TextColor(t.Text).SingleLine()
				deltaChip(c, t, delta, up)
			})
			ui.Text(c, previous+" last year").FontSize(t.FontSize * 0.8).
				TextColor(t.TextMuted).SingleLine()
		})
		ui.Row(c).Gap(t.Space(4)).Children(func() {
			legend(c, t, t.Success, "This year")
			legend(c, t, t.Text.Alpha(0.3), "Last year")
		})
	})
}

// legend is one named color dot.
func legend(c *ui.Context, t *ui.Theme, color ui.Color, name string) {
	ui.Row(c).Gap(t.Space(1)).AlignItems(ui.Center).Children(func() {
		ui.Box(c).Size(8, 8).Radius(4).Background(color)
		ui.Text(c, name).FontSize(t.FontSize * 0.8).TextColor(t.TextMuted).SingleLine()
	})
}

// plot is the chart itself: the grid, the series, the markers, and the
// invisible columns that feel the pointer.
func plot(c *ui.Context, t *ui.Theme, p Props) {
	plotBox := ui.Box(c).Grow(1).MinHeight(0)
	selected := -1
	if p.Selected != nil {
		hovered := -1
		// The invisible columns, one per month, tile the plot: the
		// pointer over one picks the month, over the margins clears it.
		plotBox.Children(func() {
			ui.Row(c).Fill().Padding(8, 8, 22, 40).Children(func() {
				for i := range p.Data {
					col := ui.Box(c).Grow(1).FillHeight()
					col.Label(p.Data[i].Label)
					if col.Hovered() {
						hovered = i
					}
				}
			})
		})
		if plotBox.Hovered() {
			*p.Selected = hovered
		} else {
			*p.Selected = -1
		}
		selected = *p.Selected
	}
	plotBox.Draw(func(pt *ui.Painter, r ui.Rect) {
		drawPlot(pt, t, p.Data, selected, r)
	})
}

// drawPlot paints the grid, the year before as a dashed line, the
// current year as the filled area and the solid line, and the marker of
// the month under the pointer.
func drawPlot(pt *ui.Painter, t *ui.Theme, data []Point, selected int, r ui.Rect) {
	n := len(data)
	if n == 0 {
		return
	}
	plot := ui.Rect{X: r.X + 40, Y: r.Y + 8, W: r.W - 48, H: r.H - 30}
	var scale float64
	for _, p := range data {
		if m := math.Max(p.Current, p.Previous); m > scale {
			scale = m
		}
	}
	scale *= 1.1
	if scale == 0 {
		scale = 1
	}
	y := func(v float64) float32 {
		return plot.Y + plot.H*float32(1-v/scale)
	}
	// The grid, four lines with their labels right-aligned on the left.
	for i := 0; i < 4; i++ {
		gy := plot.Y + plot.H*float32(i)/3
		pt.Line(plot.X, gy, plot.X+plot.W, gy, 1, t.Border)
		label := formatK(scale * float64(3-i) / 3)
		size := t.FontSize * 0.75
		w := float32(len(label)) * size * 0.55
		pt.Text(plot.X-8-w, gy+size*0.35, label, size, t.TextMuted)
	}
	// The band the plot has, so every column reads its center.
	band := plot.W / float32(n)
	prev := make([]xy, n)
	cur := make([]xy, n)
	for i, p := range data {
		x := plot.X + band*(float32(i)+0.5)
		prev[i] = xy{x, y(p.Previous)}
		cur[i] = xy{x, y(p.Current)}
	}
	// The year before, dashed, as the story behind the area.
	dashLine(pt, prev, 5, 5, 2, t.Text.Alpha(0.3))
	// The current year's area, fading down from the line.
	area(pt, t, cur, plot)
	// The current year's line.
	stroke(pt, cur, 2.5, t.Success)
	// The month under the pointer, marked with a bright dot and a halo.
	if selected >= 0 && selected < n {
		p := cur[selected]
		var halo ui.Path
		halo.Circle(p.X, p.Y, 7)
		pt.FillPath(&halo, t.Success.Alpha(0.25))
		var dot ui.Path
		dot.Circle(p.X, p.Y, 4)
		pt.FillPath(&dot, t.Success)
	}
	// The month names under the axis.
	size := t.FontSize * 0.75
	for i, p := range data {
		w := float32(len(p.Label)) * size * 0.55
		pt.Text(cur[i].X-w/2, plot.Y+plot.H+18, p.Label, size, t.TextMuted)
	}
}

// area fills the space under the current line, fading down: thin bands
// at every step of x, each a little fainter, so the fill reads as a
// vertical gradient.
func area(pt *ui.Painter, t *ui.Theme, pts []xy, plot ui.Rect) {
	const bands = 4
	bottom := plot.Y + plot.H
	for x := plot.X; x <= plot.X+plot.W; x += 2 {
		top := lineY(pts, x)
		for b := 0; b < bands; b++ {
			bTop := top + float32(b)*(plot.H/bands)
			if bTop >= bottom {
				break
			}
			bBottom := min(bTop+plot.H/bands, bottom)
			a := float32(0.3) * (1 - float32(b)/bands)
			pt.Fill(ui.Rect{X: x, Y: bTop, W: 2, H: bBottom - bTop}, t.Success.Alpha(a), 0)
		}
	}
}

// lineY is the height of the line at x, between its points.
func lineY(pts []xy, x float32) float32 {
	if x <= pts[0].X {
		return pts[0].Y
	}
	for i := 1; i < len(pts); i++ {
		if x <= pts[i].X {
			t := (x - pts[i-1].X) / (pts[i].X - pts[i-1].X)
			return pts[i-1].Y + (pts[i].Y-pts[i-1].Y)*t
		}
	}
	return pts[len(pts)-1].Y
}

// stroke joins the points with a line.
func stroke(pt *ui.Painter, pts []xy, width float32, c ui.Color) {
	var p ui.Path
	p.MoveTo(pts[0].X, pts[0].Y)
	for i := 1; i < len(pts); i++ {
		p.LineTo(pts[i].X, pts[i].Y)
	}
	pt.StrokePath(&p, width, c)
}

// dashLine draws the points as a dashed line, the dashes running
// continuously across the points rather than restarting at each one.
func dashLine(pt *ui.Painter, pts []xy, dash, gap, width float32, c ui.Color) {
	var carry float32
	on := true
	for i := 1; i < len(pts); i++ {
		ax, ay := pts[i-1].X, pts[i-1].Y
		bx, by := pts[i].X, pts[i].Y
		seg := float32(math.Hypot(float64(bx-ax), float64(by-ay)))
		if seg == 0 {
			continue
		}
		for t := float32(0); t < seg; {
			step := dash
			if !on {
				step = gap
			}
			run := min(step-carry, seg-t)
			if on {
				x0 := ax + (bx-ax)*(t/seg)
				y0 := ay + (by-ay)*(t/seg)
				x1 := ax + (bx-ax)*((t+run)/seg)
				y1 := ay + (by-ay)*((t+run)/seg)
				pt.Line(x0, y0, x1, y1, width, c)
			}
			t += run
			carry += run
			if carry >= step {
				carry = 0
				on = !on
			}
		}
	}
}

// deltaChip is the pill of the change, green for a gain, red for a
// loss, plain for a tie.
func deltaChip(c *ui.Context, t *ui.Theme, delta string, up bool) {
	flat := delta == "0%" || delta == "New"
	bg, fg := t.SurfaceHover, t.TextMuted
	if up {
		bg, fg = t.Success.Alpha(0.15), t.Success
	} else if !flat {
		bg, fg = t.Danger.Alpha(0.15), t.Danger
	}
	ui.Row(c).Shrink(0).Gap(3).Radius(9).Padding(2, 4).
		AlignItems(ui.Center).Background(bg).Children(func() {
		ui.Box(c).Size(12, 12).Draw(func(pt *ui.Painter, r ui.Rect) {
			cx, cy := r.X+r.W/2, r.Y+r.H/2
			if up {
				pt.Line(cx, cy+2.5, cx, cy-2.5, 1.6, fg)
				var p ui.Path
				p.MoveTo(cx-1.8, cy-0.5).LineTo(cx, cy-2.5).LineTo(cx+1.8, cy-0.5)
				pt.StrokePath(&p, 1.6, fg)
			} else if flat {
				pt.Line(cx-2.5, cy, cx+2.5, cy, 1.6, fg)
			} else {
				pt.Line(cx, cy-2.5, cx, cy+2.5, 1.6, fg)
				var p ui.Path
				p.MoveTo(cx-1.8, cy+0.5).LineTo(cx, cy+2.5).LineTo(cx+1.8, cy+0.5)
				pt.StrokePath(&p, 1.6, fg)
			}
		})
		ui.Text(c, delta).FontSize(t.FontSize * 0.8).TextColor(fg).SingleLine()
	})
}

// deltaOf is the change of current against previous: the label of the
// chip and whether it reads as a gain.
func deltaOf(current, previous float64) (label string, up bool) {
	if previous == 0 {
		return "New", false
	}
	change := (current - previous) / previous * 100
	rounded := math.Round(change*10) / 10
	if rounded == 0 {
		return "0%", false
	}
	if rounded > 0 {
		return fmt.Sprintf("+%s%%", strconv.FormatFloat(rounded, 'f', -1, 64)), true
	}
	return fmt.Sprintf("%s%%", strconv.FormatFloat(rounded, 'f', -1, 64)), false
}

// formatK is the axis label: "$14k" for the thousands, "$0" below.
func formatK(v float64) string {
	if v >= 1000 {
		return fmt.Sprintf("$%dk", int(math.Round(v/1000)))
	}
	return fmt.Sprintf("$%d", int(v))
}

// formatMoney is the headline figure as currency, rounded to the dollar
// with thousands separators: "$14,392".
func formatMoney(v float64) string {
	return "$" + grouped(int64(math.Round(v)))
}

// grouped writes n with a comma before every third digit.
func grouped(n int64) string {
	s := strconv.FormatInt(n, 10)
	start := len(s) % 3
	if start == 0 && len(s) > 3 {
		start = 3
	}
	var b strings.Builder
	for i, r := range s {
		if i > 0 && i == start {
			b.WriteByte(',')
		} else if i > start && (i-start)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(r)
	}
	return b.String()
}

// xy is a point of the plot.
type xy struct{ X, Y float32 }
