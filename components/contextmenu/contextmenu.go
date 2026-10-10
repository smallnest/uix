// Package contextmenu provides a shadcn/ui-style Context Menu, on the
// native context menu of MyGo: the system shows it where the element is
// clicked with the secondary button, or Control-clicked on macOS, and
// below the element for the menu key or Shift+F10 while it has the
// focus. The entries of the menu come from the Props.
//
//	box := contextmenu.ContextMenu(c, contextmenu.Props{
//		Items: []contextmenu.Item{
//			{Label: "Rename", Action: app.rename},
//			{Separator: true},
//			{Label: "Delete", Disabled: app.locked, Action: app.delete},
//		},
//	}).Fill()
package contextmenu

import "github.com/egoist/mygo/ui"

// Item is one entry of the menu.
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
	// Shortcut names a key with its modifiers beside the item, as the
	// shortcut that does the same; the view handles the key itself.
	ShortcutMods ui.Modifiers
	ShortcutKey  ui.Key
	// Action runs when the item is chosen.
	Action func()
}

// Props describes the menu to draw.
type Props struct {
	// Items are the entries of the menu, in order.
	Items []Item
}

// ContextMenu gives the box it returns the context menu of the items,
// and returns it, so a view can size it or chain more calls on it. The
// system shows the menu over the box.
func ContextMenu(c *ui.Context, p Props) *ui.Element {
	e := ui.Box(c)
	e.ContextMenu(func(m *ui.Menu) {
		for _, it := range p.Items {
			if it.Separator {
				m.Separator()
				continue
			}
			item := m.Item(it.Label)
			item.Checked(it.Checked).Disabled(it.Disabled)
			if it.ShortcutKey != 0 {
				item.Shortcut(it.ShortcutMods, it.ShortcutKey)
			}
			if item.Chosen() && it.Action != nil {
				it.Action()
			}
		}
	})
	return e
}
