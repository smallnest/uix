// Package selector provides a shadcn/ui-style drop-down: a trigger button
// that opens a list of options to choose one from. The package is named
// selector because select is a Go keyword.
//
//	selector.Select(c, selector.Props{
//		Selected:    &country,
//		Options:     []string{"China", "Japan", "Germany"},
//		Placeholder: "Choose a country",
//	})
package selector

import "github.com/egoist/mygo/ui"

// Props describes the select to draw.
type Props struct {
	// Selected is the chosen option; the trigger shows it, or the
	// placeholder while it is empty.
	Selected *string
	// Options are the choices in the popup.
	Options []string
	// Placeholder shows in the trigger while nothing is chosen.
	Placeholder string
	// Disabled keeps the select from opening.
	Disabled bool
}

// Select draws a drop-down for props and returns the trigger element, so
// a view can chain more calls on it. The popup opens under the trigger on
// click; the arrows and Enter choose by keyboard.
func Select(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	sel := ui.SelectBase(c, p.Selected)
	b := sel.Trigger
	b.Padding(t.Space(1.5), t.Space(3.5)).Gap(t.Space(1.5)).Radius(t.Radius).
		Justify(ui.SpaceBetween).MinWidth(t.Space(35)).
		Background(t.Surface).TextColor(t.Text).Border(1, t.Border)
	if p.Disabled {
		b.Disabled(true)
		b.TextColor(t.TextMuted)
	}
	b.Children(func() {
		switch chosen := *p.Selected; {
		case chosen != "":
			ui.Text(c, chosen).SingleLine()
		case p.Placeholder != "":
			ui.Text(c, p.Placeholder).SingleLine().TextColor(t.TextMuted)
		}
		chevron(c, t)
	})
	sel.Popup(func(panel ui.Element) {
		panel.Margin(t.Space(1), 0, 0, 0).Padding(t.Space(1)).Radius(t.Radius+2).
			Background(t.Background).Border(1, t.Border).MinWidth(b.Bounds().W)
		panel.Shadow(0, 6, 20, 0, ui.RGBA(0, 0, 0, 0.18))
		for _, opt := range p.Options {
			item := sel.Item(opt).Padding(t.Space(1.5), t.Space(2.5)).Radius(t.Radius)
			switch {
			case item.Highlighted():
				item.Background(t.Accent).TextColor(t.AccentText)
			case opt == *p.Selected:
				item.Background(t.Surface)
			}
			item.Children(func() { ui.Text(c, opt).SingleLine() })
		}
	})
	return b
}

// chevron draws the arrow of the trigger.
func chevron(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Size(t.Space(2.5), t.Space(2.5)).Shrink(0).Draw(func(p *ui.Painter, r ui.Rect) {
		var path ui.Path
		path.MoveTo(r.X+r.W*0.1, r.Y+r.H*0.3).LineTo(r.X+r.W*0.5, r.Y+r.H*0.7).LineTo(r.X+r.W*0.9, r.Y+r.H*0.3)
		p.StrokePath(&path, 1.5, t.TextMuted)
	})
}
