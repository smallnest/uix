// Package collapsible provides a shadcn/ui-style Collapsible: a
// disclosure with a label and an arrow, which a click opens and closes,
// as do Enter and Space, showing below it what Children draws while
// *Open is true. The arrow turns and the content grows into view as it
// opens.
//
//	collapsible.Collapsible(c, collapsible.Props{
//		Label:    "Advanced",
//		Open:     &app.advanced,
//		Children: func(c *ui.Context) { ui.Text(c, "Verbose logging") },
//	})
package collapsible

import "github.com/egoist/mygo/ui"

// Props describes the collapsible to draw.
type Props struct {
	// Label is the text of the disclosure, such as "Advanced".
	Label string
	// Open is whether the content is shown.
	Open *bool
	// Children draws the content of the collapsible while it is open.
	Children func(c *ui.Context)
	// Disabled keeps the collapsible from opening and closing.
	Disabled bool
}

// Collapsible draws the disclosure and returns it, so a view can chain
// more calls on it.
func Collapsible(c *ui.Context, p Props) *ui.Element {
	e := ui.Collapsible(c, p.Label, p.Open, func() {
		if p.Children != nil {
			p.Children(c)
		}
	})
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
