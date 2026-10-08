// Package popover provides a shadcn/ui-style Popover: a panel that shows
// below the element Trigger draws, while *Open is true, and that a click
// on the trigger opens and closes. Clicking outside the panel or pressing
// Escape closes it.
//
//	popover.Popover(c, popover.Props{
//		Open: &app.help,
//		Trigger: func(c *ui.Context) ui.Element {
//			return ui.Button(c, "Help")
//		},
//		Content: func(c *ui.Context) {
//			ui.Text(c, "Press the key for what you want.")
//		},
//	})
package popover

import "github.com/egoist/mygo/ui"

// Props describes the popover to draw.
type Props struct {
	// Open is whether the panel shows.
	Open *bool
	// Trigger draws the element the panel belongs to; a click on it
	// opens and closes the panel.
	Trigger func(c *ui.Context) ui.Element
	// Content draws the panel, below the trigger.
	Content func(c *ui.Context)
	// Disabled keeps the trigger from opening the panel.
	Disabled bool
}

// Popover draws the trigger, and while *Open is true the panel below it.
// It returns the panel while it shows, and nil while it is closed, so a
// view can chain more calls on the panel.
func Popover(c *ui.Context, p Props) ui.Element {
	anchor := p.Trigger(c)
	if p.Disabled {
		anchor.Disabled(true)
	} else if anchor.Clicked() {
		*p.Open = !*p.Open
	}
	return ui.Popover(c, anchor, p.Open, func() {
		if p.Content != nil {
			p.Content(c)
		}
	})
}
