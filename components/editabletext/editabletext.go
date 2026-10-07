// Package editabletext provides a text the user edits in place, as a
// shadcn/ui component for the EditableText of MyGo: a double click on
// it, or Enter while it has the focus, shows a field of it, with the
// text before the extension chosen, ready to type over.
//
//	editabletext.EditableText(c, editabletext.Props{
//		Value:    &title,
//		OnChange: func(v string) { app.rename(v) },
//	})
package editabletext

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the editable text to draw.
type Props struct {
	// Value is the text shown and edited.
	Value *string
	// OnChange runs when the user commits a new text, by Enter or by
	// moving the focus away.
	OnChange func(v string)
}

// EditableText draws the value as a text the user edits in place, and
// returns it. A double click on it, or Enter while it has the focus,
// shows a field with the text before its extension chosen; Enter or
// moving the focus away keeps what was typed, and Escape goes back.
func EditableText(c *ui.Context, p Props) *ui.Element {
	e := ui.EditableText(c, p.Value)
	if e.Changed() && p.OnChange != nil {
		p.OnChange(*p.Value)
	}
	return e
}
