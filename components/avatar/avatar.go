// Package avatar provides a shadcn/ui-style Avatar: a circle with a
// photo, or the initials of a name on a muted face. Assistive technology
// sees an image labeled with the name.
//
//	avatar.Avatar(c, avatar.Props{Name: "Ada Lovelace"})
package avatar

import (
	"unicode"

	"github.com/egoist/mygo/ui"
)

// Props describes the avatar to draw.
type Props struct {
	// Name is the person's name: its initials fill the circle when there
	// is no photo, and it labels the avatar for assistive technology.
	Name string
	// Image is the photo, cropped to the circle. Nil draws the initials.
	Image *ui.Bitmap
	// Size is the diameter of the circle, in points. 0 uses the default
	// of about twice the font size.
	Size float32
}

// Avatar draws a circle with the photo or the initials, and returns the
// circle.
func Avatar(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	size := p.Size
	if size == 0 {
		size = t.FontSize * 2.25
	}
	e := ui.Box(c).Size(size, size).Radius(size / 2).Clip().Shrink(0).Center().
		Role(ui.RoleImage).Label(p.Name)
	if p.Image != nil {
		e.Children(func() { ui.Image(c, p.Image).Fit(ui.Cover).Size(size, size) })
		return e
	}
	e.Background(t.SurfaceHover).TextColor(t.Text)
	e.Children(func() {
		ui.Text(c, initials(p.Name)).FontWeight(600).FontSize(size * 0.4).SingleLine()
	})
	return e
}

// initials returns the first letters of the first and last words of
// name, like "AL" for "Ada Lovelace". Why not MyGo's helper: it is
// unexported, and the loop is a few lines.
func initials(name string) string {
	var first, last rune
	inWord := false
	for _, r := range name {
		letter := unicode.IsLetter(r) || unicode.IsDigit(r)
		if letter && !inWord {
			if first == 0 {
				first = unicode.ToUpper(r)
			} else {
				last = unicode.ToUpper(r)
			}
		}
		inWord = letter
	}
	switch {
	case first == 0:
		return ""
	case last == 0:
		return string(first)
	}
	return string(first) + string(last)
}
