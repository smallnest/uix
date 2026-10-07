// Package autocomplete provides a field that completes the text as it is
// typed, as a shadcn/ui component for the Autocomplete of MyGo: the
// suggestions containing the text show below it, those starting with it
// first, and Up and Down move among them while Enter takes one.
//
//	autocomplete.Autocomplete(c, autocomplete.Props{
//		Value:       &task,
//		Suggestions: []string{"Water the plants", "Write the report"},
//		Label:       "Task",
//	})
package autocomplete

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the field to draw.
type Props struct {
	// Value is the text being edited, which the field changes in place,
	// as the user types or takes a suggestion.
	Value *string
	// Suggestions are the texts offered, of which those containing the
	// text show below the field, those starting with it first.
	Suggestions []string
	// Label names the field for the assistive technology.
	Label string
}

// Autocomplete draws the field and its suggestions below it, and returns
// it, so a view can chain more calls on it. The field stretches over its
// parent, as the other fields do; a bounded width overrides it in a Row.
func Autocomplete(c *ui.Context, p Props) *ui.Element {
	e := ui.Autocomplete(c, p.Value, p.Suggestions)
	e.FillWidth()
	if p.Label != "" {
		e.Label(p.Label)
	}
	return e
}
