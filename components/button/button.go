// Package button provides a shadcn/ui-style Button for the MyGo native
// toolkit. It wraps the MyGo primitives in a variant and size system: a
// view asks for a button by its role, not by its colors.
//
//	button.Button(c, button.Props{
//		Label:   "Save",
//		Variant: button.Primary,
//		OnClick: save,
//	})
//
// All colors come from the theme of the Context, so applying the palette
// of the tokens package restyles every button in the app at once.
package button

import "github.com/egoist/mygo/ui"

// Variant is the role of a button: how it looks and what it does.
type Variant int

const (
	// Primary is the main action of a view, in the accent color.
	Primary Variant = iota
	// Secondary is a regular button on a surface.
	Secondary
	// Outline is a button with a border but no face.
	Outline
	// Ghost is a borderless button that fills on hover.
	Ghost
	// Destructive is for actions that cannot be undone, in the danger color.
	Destructive
	// Link is a text button styled as a link.
	Link
)

// Size sets how much room a button takes.
type Size int

const (
	// Sm is a compact button, for toolbars and tables.
	Sm Size = iota
	// Md is the default size.
	Md
	// Lg is a large button, for empty states and forms.
	Lg
)

// Props describes the button to draw.
type Props struct {
	// Label is the button's text.
	Label string
	// Variant is the role of the button; Primary by default.
	Variant Variant
	// Size is how large the button is; Md by default.
	Size Size
	// Disabled keeps the button from acting.
	Disabled bool
	// OnClick runs for the click, on the frame of the click.
	OnClick func()
}

// Button draws a button for props and returns the MyGo element behind it,
// so a view can chain more calls on it. When the user clicks, Button
// calls OnClick and the view rebuilds on the next frame.
func Button(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	switch p.Variant {
	case Link:
		return link(c, p)
	case Primary:
		b := ui.PrimaryButton(c, p.Label)
		paddingSize(c, t, b, p.Size, false)
		return finish(c, b, p)
	}
	// The variants with a custom face: Secondary, Outline, Ghost,
	// Destructive. The face is drawn each frame, so hover and press
	// track the pointer even though the view does not rebuild.
	face, hover, pressed, fg, border := variantColors(c, t, p.Variant)
	if p.Disabled {
		fg = t.TextMuted
	}
	b := ui.ButtonBase(c)
	paddingSize(c, t, b, p.Size, false)
	b.TextColor(fg)
	if border.A > 0 {
		b.Border(1, border)
	}
	radius := t.Radius
	b.Draw(func(pp *ui.Painter, r ui.Rect) {
		bg := face
		if !b.IsDisabled() {
			switch {
			case b.Pressed():
				bg = pressed
			case b.Hovered():
				bg = hover
			}
		}
		if bg.A > 0 {
			pp.Fill(r, bg, radius)
		}
	})
	b.Children(func() { ui.Text(c, p.Label).SingleLine() })
	return finish(c, b, p)
}

// variantColors returns the face, hover, pressed and text colors and the
// border color of a custom-face variant.
func variantColors(c *ui.Context, t *ui.Theme, v Variant) (face, hover, pressed, fg, border ui.Color) {
	switch v {
	case Secondary:
		return t.Surface, t.SurfaceHover, t.SurfacePressed, t.Text, t.Border
	case Outline:
		return ui.Color{}, ui.Color{}, ui.Color{}, t.Text, t.Border
	case Ghost:
		return ui.Color{}, t.SurfaceHover, t.SurfacePressed, t.Text, ui.Color{}
	case Destructive:
		return t.Danger, t.Danger, t.Danger, ui.Hex("#ffffff"), ui.Color{}
	}
	return ui.Color{}, ui.Color{}, ui.Color{}, t.Text, t.Border
}

// link draws a text button: accent text, no face, padded only across.
func link(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	b := ui.ButtonBase(c)
	fg := t.Accent
	if p.Disabled {
		fg = t.TextMuted
	}
	b.TextColor(fg)
	px := t.Space(3.5)
	if p.Size == Sm {
		px = t.Space(2.5)
	}
	if p.Size == Lg {
		px = t.Space(4.5)
	}
	b.PaddingX(px)
	b.Children(func() { ui.Text(c, p.Label).SingleLine() })
	return finish(c, b, p)
}

// paddingSize sets the padding of the button for its size.
func paddingSize(c *ui.Context, t *ui.Theme, b *ui.Element, s Size, link bool) {
	var py, px float32
	switch s {
	case Sm:
		py, px = t.Space(1), t.Space(2.5)
	case Lg:
		py, px = t.Space(2), t.Space(4.5)
	default:
		py, px = t.Space(1.5), t.Space(3.5)
	}
	if link {
		b.PaddingX(px)
	} else {
		b.Padding(py, px)
	}
}

// finish applies Disabled and runs OnClick for the click. Disabled is
// checked here so a disabled button never acts.
func finish(c *ui.Context, b *ui.Element, p Props) *ui.Element {
	if p.Disabled {
		b.Disabled(true)
	}
	if p.OnClick != nil && !p.Disabled && b.Clicked() {
		p.OnClick()
	}
	return b
}
