// Package fieldset provides a group of fields under a legend, as a
// shadcn/ui component for the Fieldset of MyGo.
//
//	fieldset.Fieldset(c, fieldset.Props{
//		Legend: "Type",
//		Children: func(c *ui.Context) {
//			radio.Radio(c, radio.Props{...})
//		},
//	})
package fieldset

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the field set to draw.
type Props struct {
	// Legend is the bold text above the fields.
	Legend string
	// Children are the fields of the group.
	Children func(c *ui.Context)
}

// Fieldset draws the legend and the fields as a group, and returns the
// column. A margin sits above it when another group is above it.
func Fieldset(c *ui.Context, p Props) *ui.Element {
	return ui.Fieldset(c, p.Legend, func() {
		if p.Children != nil {
			p.Children(c)
		}
	})
}
