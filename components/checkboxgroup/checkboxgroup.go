// Package checkboxgroup provides a group of check boxes under a check of
// all of them, as a shadcn/ui component for the CheckboxGroup of MyGo:
// the label's check checks all the boxes, or none when they all are, and
// shows mixed while some are.
//
//	checkboxgroup.CheckboxGroup(c, checkboxgroup.Props{
//		Label: "Notifications",
//		Children: func(c *ui.Context) {
//			checkbox.Checkbox(c, checkbox.Props{Checked: &mail, Label: "Mail"})
//			checkbox.Checkbox(c, checkbox.Props{Checked: &schedule, Label: "Calendar"})
//		},
//	})
package checkboxgroup

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the group to draw.
type Props struct {
	// Label is the name of the group, next to the check of all its
	// boxes.
	Label string
	// Children are the check boxes of the group.
	Children func(c *ui.Context)
}

// CheckboxGroup draws the label with the check of all the boxes, and the
// boxes indented under it, and returns the group. A click on the label
// checks all the boxes, or none when they all are; the check shows mixed
// while some are, and a click then checks them all.
func CheckboxGroup(c *ui.Context, p Props) ui.Element {
	return ui.CheckboxGroup(c, p.Label, func() {
		if p.Children != nil {
			p.Children(c)
		}
	})
}
