// Package kbd provides a key of the keyboard shown as a chip, for the
// Kbd of shadcn/ui: the key's label in a small rounded box, as in the
// shortcut hint "⌘K" of a menu item.
//
//	kbd.Kbd(c, kbd.Props{Label: "⌘K"})
package kbd

import "github.com/egoist/mygo/ui"

// Props describes the key to draw.
type Props struct {
	// Label is the text of the key, such as "⌘K", "Esc" or "Enter".
	Label string
}

// Kbd draws the key and returns it, so a view can chain more calls on
// it.
func Kbd(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	e := ui.Box(c).Padding(t.Space(0.75), t.Space(1.5)).Radius(t.Radius-2).
		Background(t.SurfaceHover).Border(1, t.Border)
	e.Children(func() {
		ui.Text(c, p.Label).FontSize(t.FontSize * 0.8).TextColor(t.TextMuted).SingleLine()
	})
	return e
}
