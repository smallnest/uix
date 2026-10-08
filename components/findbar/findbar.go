// Package findbar provides a bar to find within a document, as a
// shadcn/ui component for the FindBar of MyGo: a search field, the
// count and the current match, and buttons to step to the next and the
// previous. The app counts the matches and shows them highlighted.
//
//	findbar.FindBar(c, findbar.Props{
//		Open: &open, Query: &query, Matches: 3, Current: &current,
//	})
package findbar

import (
	"github.com/egoist/mygo/ui"
)

// Props describes the find bar to draw.
type Props struct {
	// Open shows the bar when true; the Escape key and the Done button
	// set it false.
	Open *bool
	// Query is the text being searched for, which the field edits.
	Query *string
	// Matches is the number of times Query occurs, which the app
	// counts; the bar shows "n of m" and steps within it.
	Matches int
	// Current is the index of the match shown, which the bar steps as
	// the user asks.
	Current *int
}

// FindBar draws the search field, the match count, the step buttons and
// the Done button, and returns the bar. Enter and Cmd+G (F3 elsewhere)
// step to the next match, Shift adds the previous, and Escape or Done
// closes. While the bar is closed it draws nothing.
func FindBar(c *ui.Context, p Props) ui.Element {
	return ui.FindBar(c, p.Open, p.Query, p.Matches, p.Current)
}
