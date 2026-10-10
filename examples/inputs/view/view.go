// Package view builds the inputs example: a window of MyGo native UI
// that shows the button group, input group, one-time password, key and
// label components of uix together in one page. The alignment group
// chooses one of three, the amount field takes money with a currency
// tag, the code field takes a verification code cell by cell, and the
// keys hint at the shortcuts. The main package shows it in a window; the
// tests and the snapshot command draw it headless.
package view

import (
	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/buttongroup"
	"github.com/smallnest/uix/components/inputgroup"
	"github.com/smallnest/uix/components/inputotp"
	"github.com/smallnest/uix/components/kbd"
	"github.com/smallnest/uix/components/label"
)

// Width and Height are the size of the example window.
const Width, Height = 460, 560

// State is the state of the example; the controls edit it in place.
var State = InputsState{Align: []string{"left"}}

// InputsState holds the example: the alignment, the amount, the code and
// the message the last code reported.
type InputsState struct {
	// Align is the chosen alignment, one value.
	Align []string
	// Amount is the money typed into the amount field.
	Amount string
	// Code is where the verification code is.
	Code inputotp.State
	// Verified is the message the last complete code reported.
	Verified string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() { State = InputsState{Align: []string{"left"}} }

// dollar is the currency mark of the amount field.
var dollar = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v20"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>`))

// InputsView draws the example: the four controls, each under its label.
func InputsView(c *ui.Context) {
	t := tokens.For(c.Theme().Dark)
	c.SetTheme(t)
	ui.Column(c).Fill().Padding(24).Gap(t.Space(4)).Children(func() {
		ui.Column(c).FillWidth().Gap(1).Children(func() {
			ui.Text(c, "Inputs").FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
			ui.Text(c, "Button group, input group, code and keys, from uix.").TextColor(t.TextMuted)
		})
		section(c, t, "Button group")
		alignCard(c, t)
		section(c, t, "Input group")
		amountCard(c, t)
		section(c, t, "One-time password")
		codeCard(c, t)
		section(c, t, "Shortcuts")
		keysCard(c, t)
	})
}

// section is the label above a component.
func section(c *ui.Context, t *ui.Theme, name string) {
	ui.Text(c, name).FontSize(t.FontSize * 0.9).FontWeight(600).TextColor(t.TextMuted)
}

// alignCard is the button group demo: the alignment to choose, with the
// chosen one named under it.
func alignCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		buttongroup.ButtonGroup(c, buttongroup.Props{
			Options: []buttongroup.Option{
				{Label: "Left", Value: "left"},
				{Label: "Center", Value: "center"},
				{Label: "Right", Value: "right"},
			},
			Selected: &State.Align,
		})
		ui.Text(c, "The text aligns to the chosen side.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
	})
}

// amountCard is the input group demo: the amount with a dollar mark
// before it and a currency tag after it.
func amountCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		label.Label(c, label.Props{Text: "Amount"})
		inputgroup.InputGroup(c, inputgroup.Props{
			Value:        &State.Amount,
			Placeholder:  "0.00",
			Leading:      dollar,
			TrailingText: "USD",
		})
	})
}

// codeCard is the one-time password demo: the cells that take the code,
// and the message the last complete code reported.
func codeCard(c *ui.Context, t *ui.Theme) {
	ui.Column(c).FillWidth().Gap(t.Space(2)).Children(func() {
		inputotp.OTP(c, inputotp.Props{
			State:      &State.Code,
			Length:     6,
			OnComplete: func(code string) { State.Verified = "Code " + code + " verified" },
		})
		if State.Verified != "" {
			ui.Text(c, State.Verified).FontSize(t.FontSize * 0.85).TextColor(t.Success)
		} else {
			ui.Text(c, "Type the code from the message.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		}
	})
}

// keysCard is the key demo: the shortcut of the actions, key by key.
func keysCard(c *ui.Context, t *ui.Theme) {
	ui.Row(c).FillWidth().Gap(t.Space(2)).AlignItems(ui.Center).Children(func() {
		keyRow(c, t, "New", "⌘N")
		keyRow(c, t, "Save", "⌘S")
		keyRow(c, t, "Find", "⌘F")
	})
}

// keyRow is one shortcut: the action, then its keys.
func keyRow(c *ui.Context, t *ui.Theme, action, keys string) {
	ui.Row(c).Gap(t.Space(1.5)).AlignItems(ui.Center).Children(func() {
		ui.Text(c, action).FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
		kbd.Kbd(c, kbd.Props{Label: keys})
	})
}
