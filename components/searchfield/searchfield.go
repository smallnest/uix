// Package searchfield provides a shadcn/ui-style SearchField: a field
// for searching, with a magnifying glass and a button that clears the
// text. The field shows "Search" while empty.
//
//	searchfield.SearchField(c, searchfield.Props{Value: &query})
package searchfield

import "github.com/egoist/mygo/ui"

// Props describes the search field to draw.
type Props struct {
	// Value is the text being searched.
	Value *string
	// Disabled keeps the field from being edited.
	Disabled bool
}

// SearchField draws a field editing *Value and returns it, so a view can
// chain more calls on it. Label names it for assistive technology. A
// clear button shows while the text holds any, which clears it, as
// Escape does then; Enter reports Submitted, and any change reports
// Changed. It fills the width of its container.
func SearchField(c *ui.Context, p Props) *ui.Element {
	e := ui.SearchField(c, p.Value)
	e.FillWidth()
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
