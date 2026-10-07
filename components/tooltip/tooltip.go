// Package tooltip provides a shadcn/ui-style Tooltip: a hint that shows
// once the pointer has rested on the anchor for a moment, and goes as
// the pointer leaves it or presses it, and with Escape, until the pointer
// or the focus comes back. The anchor is any element a view has already
// built; Tooltip attaches the hint to it.
//
//	save := ui.Button(c, "Save")
//	tooltip.Tooltip(c, tooltip.Props{Anchor: save, Text: "Save the file"})
package tooltip

import "github.com/egoist/mygo/ui"

// Props describes the tooltip to draw.
type Props struct {
	// Anchor is the element the tip belongs to.
	Anchor *ui.Element
	// Text is what the tip says.
	Text string
	// Disabled keeps the tip away, even as the pointer rests.
	Disabled bool
}

// Tooltip attaches a hint to the anchor and returns it, so a view can
// chain more calls on it.
func Tooltip(c *ui.Context, p Props) *ui.Element {
	if p.Disabled {
		return p.Anchor
	}
	return p.Anchor.Tooltip(p.Text)
}
