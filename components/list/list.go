// Package list provides a scrolling list of rows, from the mygo List
// widget. The list builds only the rows in view and keeps its place
// among them; choosing a row highlights it in the accent color.
package list

import (
	"github.com/egoist/mygo/ui"
)

// Props configure the list.
type Props struct {
	// State is the place and the choice of the list, kept from frame
	// to frame. Give each list a state of its own, and set its fields
	// before building: to let the user choose a row, set
	// State.Selected = &app.Chosen, where app.Chosen is the row
	// chosen, -1 for none.
	State *ui.ListState
	// Items are the rows, each shown as its text, unless Row draws it.
	// Items also count the rows, so a Row builder changes only the
	// look, not the number.
	Items []string
	// Row draws row i in place of its text; nil shows Items[i] in a
	// line of the theme's height.
	Row func(c *ui.Context, i int)
	// OnChoose runs when the user chooses a row: i is the row chosen.
	// It needs State.Selected, which names the choice.
	OnChoose func(i int)
	// Disabled keeps the list from choosing rows.
	Disabled bool
}

// List draws the items as a scrolling list. A click chooses a row, as
// the arrow keys and Home and End do while the list has the keyboard
// focus; a double click or Enter submits it.
func List(c *ui.Context, p Props) *ui.Element {
	e := ui.List(c, p.State, len(p.Items), func(i int) {
		if p.Row != nil {
			p.Row(c, i)
			return
		}
		// The row shows its text across the list, as a row of the
		// gallery's lists does, so the chosen row highlights along the
		// whole list.
		ui.Row(c).Height(32).PaddingX(12).Children(func() {
			ui.Text(c, p.Items[i]).Grow(1)
		})
	})
	if p.OnChoose != nil && e.Changed() && p.State != nil && p.State.Selected != nil {
		p.OnChoose(*p.State.Selected)
	}
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
