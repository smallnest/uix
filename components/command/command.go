// Package command provides a shadcn/ui-style Command: a palette that
// opens over the window to run a command by name, as ⌘K does. A field
// filters the commands by their labels and keywords as you type, the
// arrows move the choice down the list that remains, and Enter or a
// click runs the chosen one.
//
//	command.Command(c, command.Props{
//		Open:  &open,
//		Items: []command.Item{
//			{Label: "New file", Keywords: []string{"create"}, OnSelect: app.newFile},
//		},
//	})
package command

import (
	"strings"

	"github.com/egoist/mygo/ui"
)

// magnifier is the icon of the search field, as the SearchField of MyGo
// draws.
var magnifier = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="11" cy="11" r="8"/><path d="m21 21-4.3-4.3"/></svg>`))

// Item is one command of the palette.
type Item struct {
	// Label is the text of the command.
	Label string
	// Keywords are the extra words a search matches, such as synonyms
	// of the command.
	Keywords []string
	// Icon is the shape of the command, an SVG in currentColor; nil
	// for none.
	Icon *ui.SVG
	// OnSelect runs when the command is chosen.
	OnSelect func()
}

// State is where the palette is, kept by the app.
type State struct {
	// Query is the text filtering the commands, cleared as the palette
	// closes.
	Query string

	// ls is the place and the choice of the list, and prev the query
	// the choice last went with.
	ls   ui.ListState
	sel  int
	prev string
}

// Props describes the palette to draw.
type Props struct {
	// Open is the bool that shows and hides the palette; the backdrop
	// and Escape set it to false.
	Open *bool
	// State is where the palette is, for the app to keep.
	State *State
	// Items are the commands offered, in order.
	Items []Item
}

// Command draws the palette for props, or nothing while it is closed. It
// returns the panel, so a view can chain more calls on it. The backdrop
// dims the window; clicking it or pressing Escape closes the palette.
func Command(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	st := p.State
	if st == nil {
		st = &State{}
	}
	if !*p.Open {
		// The palette starts fresh each time it opens.
		st.Query = ""
		return ui.Element{}
	}
	st.ls.Selected = &st.sel
	panel := ui.DialogBase(c, p.Open, func(back, wp ui.Element) {
		back.Background(ui.RGBA(0, 0, 0, 0.4))
		w, h := c.Size()
		wp.Width(min(w-t.Space(12), t.Space(120))).MaxHeight(h-t.Space(8)).
			Radius(t.Radius+2).Background(t.Background).
			Shadow(0, 10, 30, 0, ui.RGBA(0, 0, 0, 0.3))
		wp.Children(func() {
			ui.Column(c).Children(func() {
				in := search(c, t, st)
				// Enter is asked first: asking applies the typing of this
				// frame to the query, so the commands below are the ones
				// the query leaves now, not those of the frame before.
				enter := in.Submitted()
				// The commands the query leaves, and the choice kept
				// within them: a new query chooses the first again, and a
				// shorter list drops a choice past its end.
				shown := visible(p.Items, st.Query)
				if st.Query != st.prev {
					st.prev = st.Query
					st.sel = 0
				}
				if st.sel >= len(shown) {
					st.sel = -1
				}
				// The field keeps the focus while the palette is used, so
				// the arrows and Enter are the field's to pass on to the
				// list, as the field of a combobox does: the list hears
				// them only while it has the focus itself.
				down, up := in.Shortcut(0, ui.KeyDown), in.Shortcut(0, ui.KeyUp)
				if n := len(shown); n > 0 && (down || up) {
					if down {
						st.sel = min(st.sel+1, n-1)
					} else {
						st.sel = max(st.sel-1, 0)
					}
					st.ls.ScrollIntoView(st.sel)
				}
				list(c, t, st, shown, p)
				if enter && st.sel >= 0 {
					run(p, shown[st.sel])
				}
			})
		})
	})
	return panel
}

// search draws the field of the palette, a magnifier and the query, and
// returns the field. The dialog gives its first field the focus as it
// opens, so the typing filters at once. The field does not take the focus
// itself (AutoFocus): that happens before the dialog notes what had the
// focus, and closing the palette would not give it back.
func search(c *ui.Context, t *ui.Theme, st *State) ui.Element {
	var in ui.Element
	ui.Row(c).FillWidth().Padding(t.Space(2), t.Space(3)).Gap(t.Space(2)).Children(func() {
		ui.Icon(c, magnifier).TextColor(t.TextMuted)
		in = ui.TextInputBase(c, &st.Query).Grow(1).Placeholder("Type a command…")
	})
	return in
}

// list draws the commands the query left, of which the arrows choose one
// and Enter or a click runs it.
func list(c *ui.Context, t *ui.Theme, st *State, shown []int, p Props) {
	ls := ui.List(c, &st.ls, len(shown), func(i int) {
		row := ui.Row(c).FillWidth().Height(t.Space(9)).PaddingX(t.Space(3)).Gap(t.Space(2.5)).AlignItems(ui.Center)
		it := p.Items[shown[i]]
		if i == st.sel {
			row.Background(t.Accent).TextColor(t.AccentText)
		}
		row.Children(func() {
			if it.Icon != nil {
				ui.Icon(c, it.Icon).Size(t.Space(4), t.Space(4))
			}
			ui.Text(c, it.Label).SingleLine().Grow(1)
		})
		if row.Clicked() {
			run(p, shown[i])
		}
	}).Height(float32(min(len(shown), 8))*t.Space(9) + 4)
	if len(shown) > 0 && ls.Submitted() {
		run(p, shown[st.sel])
	}
}

// run chooses the command at place i: the palette closes and the item's
// OnSelect runs.
func run(p Props, i int) {
	*p.Open = false
	if p.Items[i].OnSelect != nil {
		p.Items[i].OnSelect()
	}
}

// visible returns the places of the items the query leaves: those whose
// label or keywords hold the query, as lower-case letters.
func visible(items []Item, query string) []int {
	q := strings.ToLower(strings.TrimSpace(query))
	if q == "" {
		out := make([]int, len(items))
		for i := range items {
			out[i] = i
		}
		return out
	}
	var out []int
	for i, it := range items {
		if strings.Contains(strings.ToLower(it.Label), q) {
			out = append(out, i)
			continue
		}
		for _, k := range it.Keywords {
			if strings.Contains(strings.ToLower(k), q) {
				out = append(out, i)
				break
			}
		}
	}
	return out
}
