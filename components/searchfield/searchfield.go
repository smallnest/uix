// Package searchfield provides a shadcn/ui-style SearchField: a field
// for searching, with a magnifying glass and a button that clears the
// text. The field shows "Search" while empty; Placeholder replaces it,
// so a view in another language can say what it searches.
//
//	searchfield.SearchField(c, searchfield.Props{Value: &query})
//	searchfield.SearchField(c, searchfield.Props{Value: &query, Placeholder: "搜索"})
package searchfield

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/icon"
	"github.com/smallnest/uix/components/lucide"
)

// Props describes the search field to draw.
type Props struct {
	// Value is the text being searched.
	Value *string
	// Placeholder shows in the field while it is empty, in the muted
	// color. "Search" shows when it is empty.
	Placeholder string
	// Disabled keeps the field from being edited.
	Disabled bool
}

// SearchField draws a field editing *Value and returns it, so a view can
// chain more calls on it. Label names it for assistive technology. A
// clear button shows while the text holds any, which clears it, as
// Escape does then; Enter reports Submitted, and any change reports
// Changed. It fills the width of its container.
func SearchField(c *ui.Context, p Props) ui.Element {
	if p.Placeholder == "" {
		e := ui.SearchField(c, p.Value)
		e.FillWidth()
		if p.Disabled {
			e.Disabled(true)
		}
		return e
	}
	return withPlaceholder(c, p)
}

// withPlaceholder draws the field with a custom placeholder. mygo's own
// search field shows "Search" and lets no view change it, so this one is
// built from the public pieces: a magnifying glass before an unstyled
// input in a bordered rounded box, and a clear button as mygo draws one.
func withPlaceholder(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	var in ui.Element
	f := ui.Row(c).AlignItems(ui.Center).Gap(t.Space(1.5)).Padding(t.Space(1.5), t.Space(2.5)).
		Radius(t.Radius).Background(t.Surface).Border(1, t.Border).FillWidth()
	if p.Disabled {
		f = f.Disabled(true)
	}
	f.Children(func() {
		icon.Icon(c, icon.Props{SVG: lucide.Search(), Label: "Search", Width: t.Space(3.75), Height: t.Space(3.75)}).
			Shrink(0).TextColor(t.TextMuted)
		in = ui.TextInputBase(c, p.Value).Grow(1).MinWidth(t.Space(15)).Placeholder(p.Placeholder)
		if *p.Value != "" {
			clear := ui.ButtonBase(c).Size(t.Space(4), t.Space(4)).Radius(t.Space(2)).
				Background(t.TextMuted.Alpha(0.5)).Label("Clear").FocusRing(false)
			clear.Draw(func(p *ui.Painter, r ui.Rect) {
				k := r.W * 0.3
				var x ui.Path
				x.MoveTo(r.X+k, r.Y+k).LineTo(r.X+r.W-k, r.Y+r.H-k).MoveTo(r.X+r.W-k, r.Y+k).LineTo(r.X+k, r.Y+r.H-k)
				p.StrokePath(&x, 1.5, t.Background)
			})
			if clear.Clicked() {
				*p.Value = ""
			}
		}
	})
	// The border follows the focus and the hover, as mygo's field does.
	switch {
	case in.Focused():
		f = f.BorderColor(t.Accent)
	case f.Hovered():
		f = f.BorderColor(t.Border.Mix(t.Text, 0.25))
	}
	return f
}
