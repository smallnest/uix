// Package view builds the studio example: a window of MyGo native UI
// that shows the icon, image and colorwell components of uix together.
// The main package shows it in a window; the tests and the snapshot
// command draw it headless.
package view

import (
	"fmt"
	stdimage "image"
	"image/color"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/colorwell"
	"github.com/smallnest/uix/components/icon"
	"github.com/smallnest/uix/components/image"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 400

// State is the state of the example; the controls edit it in place.
var State = StudioState{Accent: ui.Hex("#3b82f6")}

// StudioState holds the values the controls edit, and the status line.
type StudioState struct {
	Accent ui.Color
	Liked  bool
	Status string
}

// Reset restores the initial state of the example.
func Reset() {
	State = StudioState{Accent: ui.Hex("#3b82f6"), Status: "Pick a brand color, or like the mark."}
}

// Logo is the mark of the studio, three squares in its own colors.
var Logo = ui.MustParseSVG([]byte(`<svg viewBox="0 0 30 10"><rect width="9" height="10" fill="#ef4444"/><rect x="11" width="9" height="10" fill="#3b82f6"/><rect x="22" width="9" height="10" fill="#22c55e"/></svg>`))

// heart outlines a heart; filled fills it, for the like toggle.
var (
	heartOutline = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>`))
	heartFilled = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="currentColor"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/></svg>`))
)

// photo is a small picture to show: a sky over a hill, as a gradient.
func photo() *ui.Bitmap {
	src := stdimage.NewRGBA(stdimage.Rect(0, 0, 60, 40))
	for y := 0; y < 40; y++ {
		for x := 0; x < 60; x++ {
			switch {
			case y < 18: // the sky, lighter up top
				b := uint8(180 + y*3)
				src.SetRGBA(x, y, color.RGBA{R: 135, G: 206, B: b, A: 255})
			case y < 26: // the sun
				src.SetRGBA(x, y, color.RGBA{R: 255, G: 200, B: 60, A: 255})
			default: // the hill, greener down
				g := uint8(120 + (y-26)*2)
				src.SetRGBA(x, y, color.RGBA{R: 60, G: g, B: 40, A: 255})
			}
		}
	}
	return ui.NewBitmap(src)
}

// hexOf renders a color as "#rrggbb", as the picker shows it.
func hexOf(c ui.Color) string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

// StudioView draws the example: the mark of the studio and a like
// toggle, a preview of the photo, and a color well that picks the brand
// color. The status line shows what changed.
func StudioView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		// The header: the mark, a title and the like toggle.
		ui.Row(c).Fill().AlignItems(ui.Center).Gap(10).Children(func() {
			icon.Icon(c, icon.Props{SVG: Logo, Width: 30, Height: 10, Label: "Studio mark"})
			ui.Text(c, "Studio").FontSize(24).Bold()
			ui.Spacer(c)
			heart := ui.Box(c).Size(32, 32).Radius(8).Background(t.Surface).
				Border(1, t.Border).Center().Cursor(ui.CursorPointer).Children(func() {
					h := heartOutline
					if State.Liked {
						h = heartFilled
					}
					icon.Icon(c, icon.Props{SVG: h, Width: 18, Height: 18, Label: "Like"})
				})
			if heart.Clicked() {
				State.Liked = !State.Liked
				if State.Liked {
					State.Status = "Mark liked."
				} else {
					State.Status = "Mark unliked."
				}
			}
		})
		ui.Row(c).Fill().Gap(20).Children(func() {
			// The left column: the brand color and the photo.
			ui.Column(c).Grow(1).Gap(6).Children(func() {
				caption(c, "Brand color")
				ui.Row(c).Gap(8).AlignItems(ui.Center).Children(func() {
					well := colorwell.ColorWell(c, colorwell.Props{Color: &State.Accent, Label: "Accent"})
					if well.Changed() {
						State.Status = "Accent " + hexOf(State.Accent) + "."
					}
					ui.Text(c, hexOf(State.Accent)).FontSize(13).TextColor(t.TextMuted)
				})
				caption(c, "Photo")
				// The photo fills its card: Contain scales the small
				// source up to the box, where ScaleDown would leave it at
				// its own tiny size in the corner.
				image.Image(c, image.Props{Src: photo(), Width: 160, Height: 107}).
					Radius(8).Border(1, t.Border).Clip()
			})
			// The right column: the mark large, on the brand color.
			ui.Column(c).Grow(1).Gap(6).Children(func() {
				caption(c, "Preview")
				ui.Box(c).Fill().Radius(12).Background(State.Accent.Alpha(0.15)).Border(1, t.Border).Center().
					Children(func() {
						ui.Column(c).Gap(8).AlignItems(ui.Center).Children(func() {
							// The mark in its own colors: the picture
							// component keeps them.
							image.Image(c, image.Props{Src: Logo, Width: 90, Height: 30})
							ui.Text(c, "Studio").FontSize(18).Bold()
						})
					})
			})
		})
		// The status line shows what changed.
		ui.Text(c, State.Status).FontSize(13).TextColor(t.TextMuted).Grow(1)
	})
}

// caption draws the small muted label of a section.
func caption(c *ui.Context, s string) {
	t := c.Theme()
	ui.Text(c, s).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
}

// Title returns the title of the window.
func Title() string {
	return "Studio"
}
