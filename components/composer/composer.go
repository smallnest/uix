// Package composer provides the free AgentComposer of BoardUI for the
// MyGo native toolkit: a pill with a text field and circular controls,
// and the thin status row beneath it.
//
//	composer.Composer(c, composer.Props{
//		Value:    &message,
//		OnSubmit: send,
//		Provider: "Anthropic",
//		Model:    "openai/gpt-5-nano",
//	})
//
// The pill is 52 points tall with 36-point circles flush inside it, like
// the Pro composer's, so a starter reads as the same product. What the
// free composer leaves out are the Pro surfaces: the model and effort
// menus, the attachment plugin panel and the voice equalizer. The
// attachment slot stays, plain. The status row shows what is actually
// true of the running app — which provider answered, which model, how
// many messages are in the thread — rather than decorative chrome.
package composer

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/egoist/mygo/ui"
)

// Props describes the composer to draw.
type Props struct {
	// Value is the text being typed; the composer clears it when it
	// submits.
	Value *string
	// Busy keeps the send button a stop button, and lights the pill.
	Busy bool
	// OnSubmit runs when Enter is pressed or the send button is clicked,
	// with the text as typed, trimmed. The view sends it to the model.
	OnSubmit func(text string)
	// OnStop runs when the stop button is clicked while Busy.
	OnStop func()
	// OnAttachment runs when the attachment button is clicked; nil
	// leaves the slot inert, as the free composer's is.
	OnAttachment func()
	// Model is the model id answering, such as "openai/gpt-5-nano"; the
	// chip shows the name after the last slash. Empty draws no chip.
	Model string
	// Provider is the name of the provider, for the status row.
	Provider string
	// MessageCount is the number of messages in the thread, for the
	// status row; 0 reads "New chat".
	MessageCount int
}

// Composer draws the pill and the status row and returns them, so a view
// can chain more calls on the column.
func Composer(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	return ui.Column(c).FillWidth().Gap(t.Space(2.5)).Children(func() {
		pill(c, t, p)
		statusRow(c, t, p)
	})
}

// pill is the 52-point form: an attachment slot, the text field, the
// model chip and the send or stop control, all inside the pill.
func pill(c *ui.Context, t *ui.Theme, p Props) {
	e := ui.Row(c).FillWidth().Height(52).Radius(26).Padding(8).
		Gap(t.Space(2.5)).AlignItems(ui.Center).Background(t.Background)
	if p.Busy {
		// The light sweeps the pill while the model works, the native
		// stand-in for the liquid-glass loader of the Pro composer.
		e.Draw(func(pt *ui.Painter, r ui.Rect) {
			ms := float64(pt.Now().UnixMilli() % 1400)
			pos := ms / 1400
			cx := r.X + r.W*float32(pos)
			pt.Clip(r, 26, func() {
				// Thin bars, fading each side of the sweep, cheaply a
				// moving gradient.
				for i := -40; i <= 40; i++ {
					d := math.Abs(float64(i)) / 12.0
					a := float32(math.Max(0, 1-d)) * 0.3
					if a <= 0 {
						continue
					}
					x := cx + float32(i)*3
					if x < r.X || x > r.X+r.W {
						continue
					}
					pt.Fill(ui.Rect{X: x, Y: r.Y, W: 3, H: r.H}, t.Accent.Alpha(a), 0)
				}
				pt.After(50 * time.Millisecond)
			})
		})
	}
	e.Children(func() {
		// The attachment slot, with the add-button tokens of the Pro
		// composer so the two read as the same control.
		slot := circleButton(c, t, t.SurfaceHover)
		if p.OnAttachment != nil && slot.Clicked() {
			p.OnAttachment()
		}
		slot.Label("Add attachment")
		slot.Children(func() { paperclip(c, t.TextMuted) })

		in := ui.TextInput(c, p.Value).Grow(1).MinWidth(0).Padding(0).
			Border(0, ui.Color{}).Background(ui.Color{}).
			Placeholder("Ask me anything").Label("Message")
		if in.Submitted() {
			submit(c, p)
		}

		if p.Model != "" {
			modelChip(c, t, p.Model)
		}

		if p.Busy {
			stop := circleButton(c, t, t.Surface)
			stop.Label("Stop generating")
			if stop.Clicked() && p.OnStop != nil {
				p.OnStop()
			}
			stop.Children(func() {
				ui.Box(c).Size(10, 10).Radius(2).Background(t.TextMuted)
			})
		} else {
			send := circleButton(c, t, t.Accent)
			disabled := p.Value == nil || strings.TrimSpace(*p.Value) == ""
			send.Disabled(disabled)
			send.Label("Send message")
			if send.Clicked() {
				submit(c, p)
			}
			send.Children(func() { arrowUp(c, t.AccentText) })
		}
	})
}

// circleButton is a 36-point round button, as the composer's controls
// sit flush inside the pill.
func circleButton(c *ui.Context, t *ui.Theme, bg ui.Color) *ui.Element {
	return ui.ButtonBase(c).Size(36, 36).Radius(18).Background(bg).Center()
}

// submit runs OnSubmit with the trimmed text and clears the field, as
// the composer does when a message goes out.
func submit(c *ui.Context, p Props) {
	if p.OnSubmit == nil || p.Value == nil {
		return
	}
	text := strings.TrimSpace(*p.Value)
	if text == "" {
		return
	}
	*p.Value = ""
	p.OnSubmit(text)
}

// modelChip shows which model answers, the name after the last slash:
// "openai/gpt-5-nano" reads better in a short slot as "gpt-5-nano".
func modelChip(c *ui.Context, t *ui.Theme, model string) {
	ui.Row(c).Height(32).Gap(t.Space(1)).Radius(12).
		Padding(0, t.Space(2)).AlignItems(ui.Center).Children(func() {
		sparkle(c, 14, t.TextMuted)
		ui.Text(c, shortModel(model)).FontSize(t.FontSize * 0.85).
			TextColor(t.TextMuted).SingleLine()
	})
}

// statusRow is the thin row under the pill: the provider on the left,
// the message count on the right.
func statusRow(c *ui.Context, t *ui.Theme, p Props) {
	ui.Row(c).FillWidth().Height(26).Justify(ui.SpaceBetween).
		AlignItems(ui.Center).Children(func() {
		provider := p.Provider
		if provider == "" {
			provider = "Not configured"
		}
		statusItem(c, t, func() { infinity(c, t.TextMuted) }, provider)
		count := plural(p.MessageCount)
		if p.MessageCount == 0 {
			count = "New chat"
		}
		statusItem(c, t, func() { sparkle(c, 14, t.TextMuted) }, count)
	})
}

// statusItem is one label with a small icon before it.
func statusItem(c *ui.Context, t *ui.Theme, icon func(), label string) {
	ui.Row(c).Gap(t.Space(1)).AlignItems(ui.Center).Children(func() {
		icon()
		ui.Text(c, label).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted).SingleLine()
	})
}

func plural(n int) string {
	if n == 1 {
		return "1 message"
	}
	return fmt.Sprintf("%d messages", n)
}

// shortModel returns model without the provider prefix: the text after
// the last slash, or the whole string when there is none.
func shortModel(model string) string {
	if i := strings.LastIndex(model, "/"); i >= 0 {
		return model[i+1:]
	}
	return model
}

// arrowUp draws the send arrow, a shaft with a head.
func arrowUp(c *ui.Context, color ui.Color) {
	ui.Box(c).Size(36, 36).Draw(func(pt *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		var p ui.Path
		p.MoveTo(cx, cy-8).LineTo(cx-5, cy-3)
		p.MoveTo(cx, cy-8).LineTo(cx+5, cy-3)
		p.MoveTo(cx, cy-8).LineTo(cx, cy+9)
		pt.StrokePath(&p, 2, color)
	})
}

// paperclip draws the attachment slot's clip.
func paperclip(c *ui.Context, color ui.Color) {
	ui.Box(c).Size(36, 36).Draw(func(pt *ui.Painter, r ui.Rect) {
		x0, y0, s := r.X, r.Y, r.W
		var p ui.Path
		p.MoveTo(x0+s*0.36, y0+s*0.6)
		p.LineTo(x0+s*0.36, y0+s*0.4)
		p.QuadTo(x0+s*0.36, y0+s*0.28, x0+s*0.5, y0+s*0.28)
		p.QuadTo(x0+s*0.64, y0+s*0.28, x0+s*0.64, y0+s*0.42)
		p.LineTo(x0+s*0.64, y0+s*0.62)
		p.QuadTo(x0+s*0.64, y0+s*0.74, x0+s*0.5, y0+s*0.74)
		p.QuadTo(x0+s*0.38, y0+s*0.74, x0+s*0.38, y0+s*0.62)
		p.QuadTo(x0+s*0.38, y0+s*0.55, x0+s*0.47, y0+s*0.55)
		p.QuadTo(x0+s*0.55, y0+s*0.55, x0+s*0.55, y0+s*0.62)
		pt.StrokePath(&p, 1.6, color)
	})
}

// sparkle draws a four-point star, the mark of an AI model.
func sparkle(c *ui.Context, size float32, color ui.Color) *ui.Element {
	return ui.Box(c).Size(size, size).Shrink(0).Draw(func(pt *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		arm := r.W / 2
		var p ui.Path
		p.MoveTo(cx, cy-arm)
		p.QuadTo(cx+arm*0.3, cy-arm*0.3, cx+arm, cy)
		p.QuadTo(cx+arm*0.3, cy+arm*0.3, cx, cy+arm)
		p.QuadTo(cx-arm*0.3, cy+arm*0.3, cx-arm, cy)
		p.QuadTo(cx-arm*0.3, cy-arm*0.3, cx, cy-arm)
		p.Close()
		pt.FillPath(&p, color)
	})
}

// infinity draws two tangent circles that read as the loop of a run.
func infinity(c *ui.Context, color ui.Color) *ui.Element {
	return ui.Box(c).Size(14, 14).Shrink(0).Draw(func(pt *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		rad := r.H / 4
		var p ui.Path
		p.Circle(cx-rad, cy, rad)
		p.Circle(cx+rad, cy, rad)
		pt.StrokePath(&p, 1.6, color)
	})
}
