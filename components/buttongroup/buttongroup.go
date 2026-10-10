// Package buttongroup provides a shadcn/ui-style Button Group: a row of
// buttons drawn joined as the segments of one control, of which one is
// chosen, or several with Multiple. It builds on the ToggleGroup of
// MyGo, so the buttons are one stop of Tab and the arrows move the focus
// among them.
//
//	buttongroup.ButtonGroup(c, buttongroup.Props{
//		Options:  []buttongroup.Option{{Label: "Left", Value: "l"}, {Label: "Center", Value: "c"}, {Label: "Right", Value: "r"}},
//		Selected: &align,
//		OnChange: func() { app.alignChanged() },
//	})
package buttongroup

import (
	"slices"

	"github.com/egoist/mygo/ui"
)

// Option is one button of the group.
type Option struct {
	// Label is the text of the button.
	Label string
	// Value identifies the option among the group.
	Value string
	// Disabled keeps the button from being chosen.
	Disabled bool
}

// Props describes the group to draw.
type Props struct {
	// Options are the buttons of the group, in order.
	Options []Option
	// Selected holds the values chosen: one for the single choice, or
	// the ones chosen with Multiple.
	Selected *[]string
	// Multiple lets several buttons stay chosen at once.
	Multiple bool
	// Disabled keeps the buttons from being chosen.
	Disabled bool
	// OnChange runs when a click changes the choice.
	OnChange func()
}

// ButtonGroup draws the buttons joined as one control and returns it, so
// a view can chain more calls on it. A click chooses a button: the one
// alone with the single choice, which stays chosen, or a toggle with
// Multiple. The chosen buttons read pressed on the track.
func ButtonGroup(c *ui.Context, p Props) ui.Element {
	// The choice of each option this frame, read from Selected before
	// the toggles are built, and written back to it after, so a click
	// that Changed one moves to the slice the app owns.
	ons := make([]bool, len(p.Options))
	for i, o := range p.Options {
		ons[i] = slices.Contains(*p.Selected, o.Value)
	}
	changed := false
	g := ui.ToggleGroup(c, func() {
		for i, o := range p.Options {
			tg := ui.Toggle(c, &ons[i], o.Label)
			if o.Disabled {
				tg.Disabled(true)
			}
			if tg.Changed() {
				changed = true
				if !p.Multiple {
					// The toggle flipped before Changed reads it: on
					// means just chosen, so the others let go; off
					// means the chosen button clicked, which stays
					// chosen, as a radio's does.
					if ons[i] {
						for j := range ons {
							if j != i {
								ons[j] = false
							}
						}
					} else {
						ons[i] = true
					}
				}
			}
		}
	})
	if p.Disabled {
		g.Disabled(true)
	}
	if changed {
		next := make([]string, 0, len(p.Options))
		for i, o := range p.Options {
			if ons[i] {
				next = append(next, o.Value)
			}
		}
		*p.Selected = next
		if p.OnChange != nil {
			p.OnChange()
		}
	}
	return g
}
