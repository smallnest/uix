// Package empty provides an empty state, for the Empty of shadcn/ui: an
// icon over a title and a description, with an optional action, shown
// where a list or a search found nothing.
//
//	empty.Empty(c, empty.Props{
//		Title:       "No results",
//		Description: "Try another search.",
//		ActionLabel: "Clear search",
//		OnAction:    clear,
//	})
package empty

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/button"
)

// question is the icon of the empty state when Props.Icon is nil: a
// circle with a question mark, in the current color.
var question = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3"/><path d="M12 17h.01"/></svg>`))

// Props describes the empty state to draw.
type Props struct {
	// Icon is the shape of the state, an SVG in currentColor; a circle
	// with a question mark when nil.
	Icon *ui.SVG
	// Title is the heading of the state.
	Title string
	// Description is the text under the title.
	Description string
	// ActionLabel labels the action button; an empty label hides it.
	ActionLabel string
	// OnAction runs when the action is clicked.
	OnAction func()
}

// Empty draws the state and returns it, so a view can chain more calls
// on it. The state centers in the space of its container.
func Empty(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	icon := p.Icon
	if icon == nil {
		icon = question
	}
	e := ui.Column(c).FillWidth().Grow(1).Center().Gap(t.Space(2))
	e.Children(func() {
		ui.Box(c).Size(48, 48).Radius(24).Background(t.SurfaceHover).Center().Children(func() {
			ui.Icon(c, icon).Size(24, 24)
		})
		ui.Column(c).Center().Gap(t.Space(0.5)).Children(func() {
			ui.Text(c, p.Title).FontWeight(600)
			if p.Description != "" {
				ui.Text(c, p.Description).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
			}
		})
		if p.ActionLabel != "" {
			button.Button(c, button.Props{Label: p.ActionLabel, Size: button.Sm, OnClick: p.OnAction})
		}
	})
	return e
}
