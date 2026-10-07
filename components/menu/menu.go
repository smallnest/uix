// Package menu provides a shadcn/ui-style Menu: a button showing a label
// and an arrow, which opens a menu below it, as the pointer goes down on
// it, or for Enter, Space or Down while it has the focus. The menu holds
// the items of the Props, one above the other; choosing one runs its
// Action and closes the menu. An item may show a check and be disabled;
// a Separator item draws a line between the items.
//
//	menu.Menu(c, menu.Props{
//		Label: "File",
//		Items: []menu.Item{
//			{Label: "New", Action: func() { app.new() }},
//			{Separator: true},
//			{Label: "Save", Checked: app.dirty, Action: func() { app.save() }},
//		},
//	})
package menu

import "github.com/egoist/mygo/ui"

// Item is one entry of a Menu.
type Item struct {
	// Label is the text of the item.
	Label string
	// Checked shows a check beside the label, as for a setting that is
	// on.
	Checked bool
	// Disabled keeps the item from being chosen.
	Disabled bool
	// Separator draws a line instead of an item, as between groups of
	// items; the other fields are ignored.
	Separator bool
	// Action runs when the item is chosen; nil for none.
	Action func()
}

// Props describes the menu to draw.
type Props struct {
	// Label is the text of the button, such as "File".
	Label string
	// Items are the entries of the menu, in order.
	Items []Item
	// Disabled keeps the menu from opening.
	Disabled bool
}

// Menu draws the button, and the menu it opens, and returns the button,
// so a view can chain more calls on it.
func Menu(c *ui.Context, p Props) *ui.Element {
	e := ui.MenuButton(c, p.Label, func(m *ui.Menu) {
		for _, it := range p.Items {
			if it.Separator {
				m.Separator()
				continue
			}
			item := m.Item(it.Label)
			item.Checked(it.Checked).Disabled(it.Disabled)
			if item.Chosen() && it.Action != nil {
				it.Action()
			}
		}
	})
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
