// Package chat provides a working chat surface for the MyGo native
// toolkit, built from the agent components: a transcript that scrolls,
// the thinking indicator while the model works, and the composer at the
// bottom. It is the AgentChat of BoardUI without the backend: the app
// owns the messages and sends them however it likes.
//
//	chat.Chat(c, chat.Props{
//		Messages: msgs,
//		Value:    &input,
//		Busy:     busy,
//		OnSubmit: send,
//	})
//
// The surface is a rounded card on the theme's Surface, one step darker
// than the window, with the composer's pill on the Background over it.
// A chat spans the width it is given and grows to the height it is
// given, so a window can size it with Grow.
package chat

import (
	"math"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/composer"
	"github.com/smallnest/uix/components/thinking"
)

// Role is who a message speaks for.
type Role int

const (
	// User is the person at the keyboard; their words go in a bubble.
	User Role = iota
	// Assistant is the model; its words read plain on the surface.
	Assistant
)

// Message is one line of the transcript.
type Message struct {
	// Role is who speaks.
	Role Role
	// Text is what they said. The assistant's message the model is
	// still writing grows as the app appends to it.
	Text string
	// Streaming marks the assistant's message as the one being written,
	// which draws a blinking caret after the text.
	Streaming bool
}

// Props describes the chat to draw.
type Props struct {
	// Messages are the lines of the transcript, oldest first. While the
	// model works the app keeps an empty assistant message at the end
	// and sets Thinking.
	Messages []Message
	// Value is the text being typed in the composer.
	Value *string
	// Busy keeps the composer's send a stop button.
	Busy bool
	// Thinking shows the thinking indicator where the reply will land;
	// set while the model works and nothing has been written yet.
	Thinking bool
	// OnSubmit runs when the composer sends, with the text as typed.
	OnSubmit func(text string)
	// OnStop runs when the stop button is clicked while Busy.
	OnStop func()
	// Model and Provider feed the composer's chip and status row.
	Model    string
	Provider string
	// Suggestions are the chips of the empty state, each sending itself.
	Suggestions []string
	// Scroll keeps the transcript's place, as TrackScroll does. Pass
	// one you keep, and the transcript follows the end when it is
	// already there. Nil scrolls without keeping the place.
	Scroll *ui.ScrollState
}

// Chat draws the chat surface and returns it, so a view can chain more
// calls on it.
func Chat(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	return ui.Column(c).Fill().Background(t.Surface).Radius(t.Radius * 3).Clip().
		Children(func() {
			if len(p.Messages) == 0 {
				empty(c, t, p)
			} else {
				transcript(c, t, p)
			}
			composerWrap(c, t, p)
		})
}

// transcript is the scrollable column of messages, with the thinking
// indicator where a reply will land.
func transcript(c *ui.Context, t *ui.Theme, p Props) {
	sc := ui.Scroll(c).Grow(1).MinHeight(0)
	if p.Scroll != nil {
		sc.TrackScroll(p.Scroll)
		// A log that follows its end, unless the user scrolled up from
		// it, as MyGo's Scroll describes.
		if p.Scroll.Y >= p.Scroll.MaxY-t.Space(2) && p.Scroll.MaxY > 0 {
			p.Scroll.Y = math.MaxFloat32
		}
	}
	sc.Children(func() {
		ui.Column(c).FillWidth().Gap(t.Space(5)).
			Padding(t.Space(6), t.Space(4), t.Space(2), t.Space(4)).
			Children(func() {
				for _, m := range p.Messages {
					message(c, t, m)
				}
				if p.Thinking {
					ui.Row(c).Padding(0, t.Space(1)).Children(func() {
						thinking.Thinking(c, thinking.Props{
							Label:   "Thinking",
							Variant: thinking.Wave,
						})
					})
				}
			})
	})
}

// empty is the centered prompt and suggestion chips of a new thread.
func empty(c *ui.Context, t *ui.Theme, p Props) {
	ui.Column(c).Grow(1).FillWidth().Center().Gap(t.Space(4)).
		Padding(0, t.Space(4)).Children(func() {
			ui.Column(c).Center().Gap(t.Space(1)).Children(func() {
				ui.Text(c, "What can I help with?").
					FontSize(t.FontSize * 1.4).FontWeight(600).TextColor(t.Text)
				ui.Text(c, "The chat runs against your own model key. "+
					"History stays in this app.").
					TextColor(t.TextMuted).TextAlign(ui.Center)
			})
			if len(p.Suggestions) > 0 {
				ui.Row(c).Wrap().Justify(ui.Center).Gap(t.Space(2)).Children(func() {
					for _, s := range p.Suggestions {
						chip(c, t, p, s)
					}
				})
			}
			if p.Thinking {
				ui.Row(c).Children(func() {
					thinking.Thinking(c, thinking.Props{
						Label:   "Thinking",
						Variant: thinking.Wave,
					})
				})
			}
		})
}

// chip is one suggestion, a pill that sends itself when clicked.
func chip(c *ui.Context, t *ui.Theme, p Props, s string) {
	b := ui.ButtonBase(c).Radius(t.Radius + 12).Background(t.Background).
		Padding(t.Space(2), t.Space(3.5))
	if b.Clicked() && p.OnSubmit != nil {
		p.OnSubmit(s)
	}
	b.Children(func() {
		ui.Text(c, s).TextColor(t.TextMuted).SingleLine()
	})
}

// message draws one line of the transcript, in its role's own voice.
func message(c *ui.Context, t *ui.Theme, m Message) {
	switch m.Role {
	case User:
		ui.Row(c).FillWidth().Justify(ui.End).Children(func() {
			ui.Box(c).MaxWidthPercent(75).Padding(t.Space(2.5), t.Space(3.5)).
				Radius(t.Radius + 4).Background(t.Accent).Children(func() {
				ui.Text(c, m.Text).TextColor(t.AccentText)
			})
		})
	default:
		ui.Row(c).FillWidth().Justify(ui.Start).AlignItems(ui.End).
			Gap(t.Space(1)).Children(func() {
				ui.Text(c, m.Text).MaxWidthPercent(85).TextColor(t.Text)
				if m.Streaming {
					caret(c, t)
				}
			})
	}
}

// caret is the blinking block that marks the line the model is writing.
func caret(c *ui.Context, t *ui.Theme) *ui.Element {
	size := t.FontSize
	return ui.Box(c).Size(2, size).Shrink(0).Draw(func(pt *ui.Painter, r ui.Rect) {
		phase := float64(pt.Now().UnixMilli()%800) / 800
		// A triangle wave of opacity: fade out then snap back, as a
		// caret does.
		a := float32(math.Abs(1 - 2*phase))
		pt.Fill(ui.Rect{X: r.X, Y: r.Y, W: r.W, H: r.H}, t.Text.Alpha(a), 1)
		pt.After(80 * time.Millisecond)
	})
}

// composerWrap is the composer with the room it keeps from the edge of
// the card.
func composerWrap(c *ui.Context, t *ui.Theme, p Props) {
	ui.Column(c).Padding(0, t.Space(3)).Shrink(0).Children(func() {
		composer.Composer(c, composer.Props{
			Value:        p.Value,
			Busy:         p.Busy,
			OnSubmit:     p.OnSubmit,
			OnStop:       p.OnStop,
			Model:        p.Model,
			Provider:     p.Provider,
			MessageCount: len(p.Messages),
		})
	})
}
