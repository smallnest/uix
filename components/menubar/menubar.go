// Package menubar provides a shadcn/ui-style Menu Bar: a row of flat
// buttons, each opening a menu below it, as the pointer goes down on it
// or for Enter, Space or Down while it has the focus. It is the bar of
// an app across the top of a view, as File, Edit and View are. The items
// of the menus come from the Props; an item may show a check and a
// shortcut and be disabled, and a Separator item draws a line.
//
//	menubar.Menubar(c, menubar.Props{
//		Menus: []menubar.Menu{
//			{Label: "File", Items: []menubar.Item{
//				{Label: "New", ShortcutKey: ui.KeyN, Action: app.newFile},
//				{Separator: true},
//				{Label: "Quit", Action: app.quit},
//			}},
//			{Label: "Edit", Items: []menubar.Item{
//				{Label: "Undo", ShortcutKey: ui.KeyZ, Action: app.undo},
//			}},
//		},
//	})
package menubar

import "github.com/egoist/mygo/ui"

// Item is one entry of a menu.
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

// Menu is one menu of the bar.
type Menu struct {
	// Label is the text of the button that opens the menu, such as
	// "File".
	Label string
	// Items are the entries of the menu, in order.
	Items []Item
}

// Props describes the bar to draw.
type Props struct {
	// Menus are the menus of the bar, in order.
	Menus []Menu
}

// Menubar draws the bar of menus and returns it, so a view can chain
// more calls on it.
func Menubar(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	bar := ui.Row(c).FillWidth().Gap(t.Space(1)).PaddingX(t.Space(2))
	bar.Children(func() {
		for _, m := range p.Menus {
			button(c, t, m)
		}
	})
	return bar
}

// button draws one button of the bar: a flat face that fills on hover
// and press, and the menu it opens below it.
func button(c *ui.Context, t *ui.Theme, m Menu) {
	b := ui.ButtonBase(c).Padding(t.Space(1), t.Space(2.5)).Radius(t.Radius)
	b.Menu(func(mm *ui.Menu) {
		for _, it := range m.Items {
			if it.Separator {
				mm.Separator()
				continue
			}
			item := mm.Item(it.Label)
			item.Checked(it.Checked).Disabled(it.Disabled)
			if it.ShortcutKey != 0 {
				item.Shortcut(it.ShortcutMods, it.ShortcutKey)
			}
			if item.Chosen() && it.Action != nil {
				it.Action()
			}
		}
	})
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		if !b.IsDisabled() && (b.Hovered() || b.Pressed()) {
			pp.Fill(r, t.SurfaceHover, t.Radius)
		}
	})
	b.Children(func() {
		ui.Row(c).Gap(t.Space(1.5)).Children(func() {
			ui.Text(c, m.Label).SingleLine()
			chevron(c, t)
		})
	})
}

// chevron is the arrow that says the button opens a menu, drawn in the
// muted color of the theme.
func chevron(c *ui.Context, t *ui.Theme) {
	ui.Icon(c, chevronDown).Size(t.Space(3), t.Space(3)).TextColor(t.TextMuted)
}

var chevronDown = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="m6 9 6 6 6-6"/></svg>`))
