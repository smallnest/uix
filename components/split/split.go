// Package split provides resizable panes, from the mygo Split widget.
// A divider between the panes drags to resize them, or moves with the
// arrows once it has the keyboard focus.
package split

import (
	"github.com/egoist/mygo/ui"
)

// Props configure the split.
type Props struct {
	// Size is the width of the first pane, or its height when the
	// split is Vertical. The divider keeps it between 40 DIPs and the
	// room left; give it a field of the app, which the divider edits
	// as it moves.
	Size *float32
	// Vertical stacks the panes, first above second; a horizontal
	// split puts them side by side.
	Vertical bool
	// First builds the pane before the divider.
	First func(c *ui.Context)
	// Second builds the pane after the divider.
	Second func(c *ui.Context)
}

// Split lays out the two panes with a divider between them that the
// user drags to resize, or moves with the arrows once the divider has
// the keyboard focus. Size it, as a Row or Column, to fill the room it
// shares.
func Split(c *ui.Context, p Props) *ui.Element {
	if p.First == nil {
		p.First = func(c *ui.Context) {}
	}
	if p.Second == nil {
		p.Second = func(c *ui.Context) {}
	}
	// There is no Disabled prop: mygo does not gate the drag of the
	// divider on a disabled flag, so it could not freeze the divider,
	// and a frozen divider is rarely what an app wants anyway.
	if p.Vertical {
		return ui.SplitVertical(c, p.Size, func() { p.First(c) }, func() { p.Second(c) })
	}
	return ui.Split(c, p.Size, func() { p.First(c) }, func() { p.Second(c) })
}
