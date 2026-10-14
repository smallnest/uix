// Package view builds the animate example: a window that shows the
// animate component of uix — cards entering one after another, a pulse
// that breathes, a square that spins, and a list that slides in. The
// replay button plays the entrances again.
package view

import (
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/animate"
	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/button"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 460

// State is the state of the example.
var State = ExampleState{}

// ExampleState holds the values the controls edit.
type ExampleState struct {
	// Run is bumped by the replay button; the entrances replay when it
	// changes.
	Run uint64
}

// Reset restores the initial state of the example.
func Reset() { State = ExampleState{} }

// Title is the title of the example window.
func Title() string { return "Animate" }

// ExampleView draws the example: the four entrances in a row of cards,
// and below them the loops and the slide-in list, over the tokens
// palette.
func ExampleView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(16).Children(func() {
		// The header: the title, a pulse badge, and the replay button.
		ui.Row(c).FillWidth().MinWidth(0).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			ui.Text(c, "Animate").FontSize(24).Bold()
			ui.Spacer(c)
			animate.Pulse(ui.Box(c).Size(10, 10).Radius(5).Background(t.Danger), animate.Options{Duration: 1400 * time.Millisecond})
			ui.Text(c, "Recording").TextColor(t.TextMuted)
			button.Button(c, button.Props{
				Label:   "Replay",
				Variant: button.Outline,
				OnClick: func() { State.Run++ },
			})
		})
		// The four entrances, one card each, entering one after another.
		ui.Row(c).FillWidth().Gap(t.Space(3)).Children(func() {
			for i, label := range []string{"Fade", "Rise", "Slide", "Scale"} {
				animate.ScaleIn(card(c, t, label),
					animate.Options{Delay: time.Duration(i) * 120 * time.Millisecond, Run: State.Run})
			}
		})
		ui.Row(c).FillWidth().Gap(t.Space(6)).Children(func() {
			// The spin: a square turning around its center.
			ui.Column(c).Grow(1).Gap(t.Space(2)).Children(func() {
				caption(c, t, "Spin")
				animate.Spin(ui.Box(c).Size(48, 48).Radius(t.Radius).Background(t.Accent), animate.Options{})
			})
			// The slide-in list, rows coming in from the left one after
			// another.
			ui.Column(c).Grow(2).Gap(t.Space(2)).Children(func() {
				caption(c, t, "Slides in")
				for i, item := range []string{"First row", "Second row", "Third row"} {
					animate.SlideIn(ui.Row(c).FillWidth().MinWidth(0).Padding(t.Space(1.5), t.Space(2.5)).
						Radius(t.Radius).Background(t.Surface).Border(1, t.Border).Children(func() {
						ui.Text(c, item)
					}), animate.Options{Delay: time.Duration(i) * 100 * time.Millisecond, Run: State.Run}, animate.Left)
				}
			})
		})
	})
}

// card is one entrance of the row: a labeled tile.
func card(c *ui.Context, t *ui.Theme, label string) ui.Element {
	return ui.Column(c).Width(108).Gap(t.Space(1)).Padding(t.Space(2)).Radius(t.Radius).
		Background(t.Surface).Border(1, t.Border).Children(func() {
		ui.Text(c, label).Bold()
		ui.Text(c, "enters").FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
	})
}

// caption labels a group of the example.
func caption(c *ui.Context, t *ui.Theme, text string) {
	ui.Text(c, text).Bold().TextColor(t.TextMuted).FontSize(t.FontSize * 0.875)
}
