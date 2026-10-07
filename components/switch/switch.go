// Package switcher provides a shadcn/ui-style switch with a label, from
// the Switch of MyGo. The package is named switcher because switch is a
// Go keyword.
//
//	switcher.Switch(c, switcher.Props{
//		On:    &dark,
//		Label: "Dark mode",
//	})
package switcher

import "github.com/egoist/mygo/ui"

// Props describes the switch to draw.
type Props struct {
	// On is the bool the switch toggles.
	On *bool
	// Label is the text beside the switch. It also names the switch for
	// tests: Click("Dark mode") toggles it.
	Label string
	// Disabled keeps the switch from being toggled.
	Disabled bool
	// Changed reports the new On, when the user toggled the switch.
	Changed func(on bool)
}

// Switch draws a row with the switch and its label, and returns it. The
// switch itself is clickable; the label is not.
func Switch(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	row := ui.Row(c).Gap(t.Space(2)).Children(func() {
		s := ui.Switch(c, p.On)
		if p.Label != "" {
			s.Label(p.Label)
		}
		if p.Disabled {
			s.Disabled(true)
		}
		if p.Changed != nil && s.Changed() {
			p.Changed(*p.On)
		}
		if p.Label != "" {
			ui.Text(c, p.Label)
		}
	})
	return row
}
