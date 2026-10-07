// Package togglegroup provides a shadcn/ui-style ToggleGroup: a row of
// toggle buttons drawn joined as the segments of one control, as Bold,
// Italic and Underline in an editor's toolbar. Each toggle keeps its own
// state, so several may be on at once; they are one stop of Tab, among
// which the arrows move the focus.
//
//	togglegroup.ToggleGroup(c, togglegroup.Props{
//		Items: []togglegroup.Item{
//			{On: &app.bold, Label: "B", Tip: "Bold"},
//			{On: &app.italic, Label: "I", Tip: "Italic"},
//		},
//	})
package togglegroup

import "github.com/egoist/mygo/ui"

// Item is one toggle of the group: a button that stays pressed while
// *On, showing Label.
type Item struct {
	// On is whether the toggle is pressed.
	On *bool
	// Label is the text the toggle shows, such as "B" for Bold.
	Label string
	// Tip is the tooltip of the toggle, shown as the pointer rests on
	// it; empty for none.
	Tip string
}

// Props describes the toggle group to draw.
type Props struct {
	// Items are the toggles of the group, in order.
	Items []Item
	// Disabled keeps the toggles from being pressed.
	Disabled bool
}

// ToggleGroup draws the toggles joined as one control and returns it, so
// a view can chain more calls on it.
func ToggleGroup(c *ui.Context, p Props) *ui.Element {
	e := ui.ToggleGroup(c, func() {
		for _, it := range p.Items {
			tg := ui.Toggle(c, it.On, it.Label)
			if it.Tip != "" {
				tg.Tooltip(it.Tip)
			}
		}
	})
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
