// Package view builds the docreader example: a window of MyGo native UI
// that shows the richtext, editabletext, findbar and segmented
// components together in a document reader. The segmented control
// switches between the reader and the outline, the title renames in
// place, and the find bar searches the text, whose matches highlight
// and which it steps through. The main package shows it in a window;
// the tests and the snapshot command draw it headless.
package view

import (
	"fmt"
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/editabletext"
	"github.com/smallnest/uix/components/findbar"
	"github.com/smallnest/uix/components/richtext"
	"github.com/smallnest/uix/components/segmented"
	"github.com/smallnest/uix/components/separator"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 480

// State is the state of the example; the controls edit it in place.
var State = DocState{}

// DocState holds the document reader, as the controls edit it.
type DocState struct {
	// Title is the name of the document, which edits in place.
	Title string
	// Pane is the pane shown: 0 the reader, 1 the outline.
	Pane int
	// FocusPara is the paragraph the outline jumped to, -1 for none;
	// its heading shows in the accent in the reader.
	FocusPara int
	// FindOpen shows the find bar.
	FindOpen bool
	// Query is the text being searched for, which the find bar edits.
	Query string
	// Current is the index of the match shown, which the find bar
	// steps.
	Current int
	// Status is the last thing the controls did, shown at the bottom,
	// unless the find bar is searching.
	Status string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = DocState{
		Title:     "uix notes",
		Pane:      0,
		FocusPara: -1,
		Status:    "Double click the title to rename it, or press Find to search.",
	}
}

// para is a paragraph of the document: a heading and a body.
type para struct {
	Heading string
	Body    string
}

// doc is the document the reader shows.
var doc = []para{
	{Heading: "Welcome", Body: "uix is a component registry for the MyGo native UI toolkit. You copy a component into your project and own it, as shadcn/ui does for the web."},
	{Heading: "Find", Body: "Press Find, and type a word. The matches highlight in the text, the bar counts them, and Next and Previous step to each in turn."},
	{Heading: "Rename", Body: "Double click the title to rename the document in place, as a file in Finder. Enter keeps the new name, and Escape goes back."},
	{Heading: "Outline", Body: "The segmented control switches between the reader and the outline of the headings. A click on a heading jumps back to it in the reader."},
}

// findRanges returns the byte ranges of the occurrences of q in s,
// without regard to case.
func findRanges(s, q string) [][2]int {
	if q == "" {
		return nil
	}
	low, lq := strings.ToLower(s), strings.ToLower(q)
	var out [][2]int
	for i := 0; ; {
		j := strings.Index(low[i:], lq)
		if j < 0 {
			return out
		}
		lo := i + j
		out = append(out, [2]int{lo, lo + len(lq)})
		i = lo + len(lq)
	}
}

// paraMatches returns the matches of the query in the body of each
// paragraph, and the total across them.
func paraMatches() (ranges [][][2]int, total int) {
	ranges = make([][][2]int, len(doc))
	for i, p := range doc {
		ranges[i] = findRanges(p.Body, State.Query)
		total += len(ranges[i])
	}
	return ranges, total
}

// paragraphOf returns the index of the paragraph holding match current.
func paragraphOf(ranges [][][2]int, current int) int {
	n := 0
	for i, rs := range ranges {
		if current < n+len(rs) {
			return i
		}
		n += len(rs)
	}
	return len(ranges) - 1
}

// DocView draws the example: a document reader whose controls all work.
func DocView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ranges, total := paraMatches()
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).PaddingX(16).PaddingY(10).Gap(12).AlignItems(ui.Center).Children(func() {
			editabletext.EditableText(c, editabletext.Props{
				Value: &State.Title,
				OnChange: func(v string) {
					State.Status = "Renamed to “" + v + "”."
				},
			})
			ui.Spacer(c)
			segmented.Segmented(c, segmented.Props{
				Selected: &State.Pane,
				Labels:   []string{"Reader", "Outline"},
				Label:    "Pane",
				OnChange: func(i int) {
					if i == 0 {
						State.Status = "Switched to the reader."
					} else {
						State.Status = "Switched to the outline."
					}
				},
			})
			find := ui.Button(c, "Find")
			if find.Clicked() {
				State.FindOpen = !State.FindOpen
				if State.FindOpen {
					State.Status = "Type a word to find it."
				}
			}
		})
		findbar.FindBar(c, findbar.Props{
			Open: &State.FindOpen, Query: &State.Query,
			Matches: total, Current: &State.Current,
		})
		separator.Separator(c, separator.Props{})
		if State.Pane == 0 {
			ui.Scroll(c).Grow(1).Children(func() {
				reader(c, ranges)
			})
		} else {
			ui.Scroll(c).Grow(1).Children(func() {
				outline(c)
			})
		}
		separator.Separator(c, separator.Props{})
		ui.Row(c).PaddingX(24).PaddingY(10).Children(func() {
			ui.Text(c, statusText(ranges, total)).TextColor(t.TextMuted)
		})
	})
}

// reader draws the document: a heading and a body per paragraph, the
// matches of the query highlighted, the current one in the accent.
// The column sizes to its content, so the scroll can move it.
func reader(c *ui.Context, ranges [][][2]int) {
	t := c.Theme()
	ui.Column(c).Padding(24).Gap(10).Children(func() {
		for i, p := range doc {
			heading := richtext.RichText(c, richtext.Props{Spans: []ui.Span{
				{Text: p.Heading, Weight: 600, Size: t.FontSize * 1.2},
			}})
			if i == State.FocusPara {
				heading.TextColor(t.Accent)
			}
			richtext.RichText(c, richtext.Props{Spans: bodySpans(c, i, ranges)})
		}
	})
}

// bodySpans returns the spans of the body of paragraph i, its matches
// of the query highlighted: the current one in the accent, the rest on
// a muted ground.
func bodySpans(c *ui.Context, i int, ranges [][][2]int) []ui.Span {
	body := doc[i].Body
	rs := ranges[i]
	t := c.Theme()
	if len(rs) == 0 {
		return []ui.Span{{Text: body}}
	}
	before := 0
	for k := 0; k < i; k++ {
		before += len(ranges[k])
	}
	currentLocal := -1
	if State.Current >= before && State.Current < before+len(rs) {
		currentLocal = State.Current - before
	}
	var spans []ui.Span
	pos := 0
	for k, r := range rs {
		if r[0] > pos {
			spans = append(spans, ui.Span{Text: body[pos:r[0]]})
		}
		sp := ui.Span{Text: body[r[0]:r[1]], Background: t.SurfaceHover}
		if k == currentLocal {
			sp.Background = t.Accent
			sp.Color = t.AccentText
		}
		spans = append(spans, sp)
		pos = r[1]
	}
	if pos < len(body) {
		spans = append(spans, ui.Span{Text: body[pos:]})
	}
	return spans
}

// outline draws the headings of the document as buttons; a click jumps
// back to the paragraph in the reader.
func outline(c *ui.Context) {
	ui.Column(c).Padding(24).Gap(8).Children(func() {
		for i, p := range doc {
			if ui.Button(c, p.Heading).Clicked() {
				State.Pane = 0
				State.FocusPara = i
				State.Status = "Jumped to “" + p.Heading + "”."
			}
		}
	})
}

// statusText returns the status line: the find progress while the bar
// searches, otherwise what the last control did.
func statusText(ranges [][][2]int, total int) string {
	if State.FindOpen && State.Query != "" {
		if total > 0 {
			return fmt.Sprintf("Match %d of %d, in “%s”.", State.Current+1, total, doc[paragraphOf(ranges, State.Current)].Heading)
		}
		return "No matches for “" + State.Query + "”."
	}
	return State.Status
}
