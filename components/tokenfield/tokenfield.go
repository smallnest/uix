// Package tokenfield provides a field of tokens, as of tags or the
// recipients of a mail, as a shadcn/ui component for the TokenField of
// MyGo: the tokens show as chips, each with a button taking it out, and
// typing adds one on Enter or a comma, the suggestions of what was typed
// showing below it, of those not already chips.
//
//	tokenfield.TokenField(c, tokenfield.Props{
//		Tokens:      &tags,
//		Suggestions: []string{"urgent", "later", "work"},
//		Label:       "Tags",
//	})
package tokenfield

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the field to draw.
type Props struct {
	// Tokens are the chips of the field, which it changes in place.
	Tokens *[]string
	// Suggestions are the texts offered, of those containing the text
	// typed, and not already chips.
	Suggestions []string
	// Label names the field for the assistive technology.
	Label string
}

// TokenField draws the tokens as chips and the field adding them, and
// returns it, so a view can chain more calls on it. A chip has a button
// taking it out; typing adds a token on Enter or a comma, a suggestion
// adds on Enter or a click, and Backspace in the empty field takes out
// the last.
func TokenField(c *ui.Context, p Props) ui.Element {
	e := ui.TokenField(c, p.Tokens, p.Suggestions)
	if p.Label != "" {
		e.Label(p.Label)
	}
	return e
}
