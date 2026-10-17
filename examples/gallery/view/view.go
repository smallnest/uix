// Package view builds the gallery example: a shell like onevra's
// showcase, with a sidebar of the components and one page each. Every
// page shows a component of one of the four libraries — uix's own,
// mygo's primitives, onevra's, or MujicaUI's — with a badge
// naming the library; where the same component exists in several, the
// gallery keeps the one of uix, then mygo, then onevra. A search box
// above the sidebar filters the items. The main package shows it in a
// window; the tests and the snapshot command draw it headless.
package view

import (
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/breadcrumbs"
	"github.com/smallnest/uix/components/scroll"
	"github.com/smallnest/uix/components/searchfield"
	"github.com/smallnest/uix/components/segmented"
	"github.com/smallnest/uix/components/sidebar"
)

// Width and Height are the size of the example window.
const Width, Height = 960, 600

// GalleryView draws the gallery shell: the sidebar of components on the
// left, the page of the chosen one on the right.
func GalleryView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	// A library filter that does not show the chosen page lands on the
	// first page of the library it does show.
	if lib, ok := filterLib(); ok && !inLib(lib, State.Selected) {
		if e, ok := firstOf(lib); ok {
			State.Selected = e.ID
		}
	}
	entry, section := current()
	ui.Row(c).Fill().Gap(t.Space(3)).Padding(t.Space(3)).Children(func() {
		rail(c, t)
		page(c, t, entry, section)
	})
}

// rail is the left column: the title, the search box, and the sidebar
// of the components that match the query.
func rail(c *ui.Context, t *ui.Theme) {
	ui.Column(c).Width(244).FillHeight().Shrink(0).Gap(t.Space(2)).
		Background(t.Surface).Radius(t.Radius * 3).Padding(t.Space(2.5)).Children(func() {
		ui.Text(c, "组件画廊").FontSize(t.FontSize * 1.2).Bold()
		// The library filter: 全部 mixes the libraries, and each
		// name shows that library's components alone.
		segmented.Segmented(c, segmented.Props{
			Selected: &State.Filter,
			Labels:   []string{"全部", "uix", "mygo", "onevra", "mujica"},
			Label:    "按库过滤",
		})
		searchfield.SearchField(c, searchfield.Props{Value: &State.Query, Placeholder: "搜索组件…"})
		sections := sectionsFor(State.Query)
		if len(sections) == 0 {
			ui.Text(c, "没有匹配的组件。").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
			return
		}
		// The sidebar scrolls its items in the rail's remaining height; it
		// grows to that, not to a percent of the window, so its scroll
		// clips at the rail's bottom instead of drawing past it.
		sidebar.Sidebar(c, sidebar.Props{Selected: &State.Selected, Sections: sections}).Grow(1)
	})
}

// page is the right column: the breadcrumb, the heading with the
// library badge, and the demo, scrollable.
func page(c *ui.Context, t *ui.Theme, entry Entry, section string) {
	ui.Column(c).Grow(1).MinWidth(0).FillHeight().Gap(t.Space(2)).Children(func() {
		breadcrumbs.Breadcrumbs(c, breadcrumbs.Props{Items: []string{section, entry.Name}})
		ui.Row(c).AlignItems(ui.Center).Gap(t.Space(2)).Children(func() {
			ui.Text(c, entry.Name).FontSize(t.FontSize * 1.5).FontWeight(600)
			libraryBadge(c, t, entry.Lib)
		})
		scroll.Scroll(c, scroll.Props{Children: func(c *ui.Context) {
			ui.Column(c).FillWidth().MinWidth(0).Padding(t.Space(4)).Gap(t.Space(3)).Children(func() {
				entry.Build(c, t)
			})
		}}).Fill().Background(t.Background).Radius(t.Radius*2).Border(1, t.Border)
	})
}

// current returns the chosen entry and its section, or the first entry
// of the catalog when nothing is chosen.
func current() (Entry, string) {
	for _, s := range catalog {
		for _, e := range s.Entries {
			if e.ID == State.Selected {
				return e, s.Title
			}
		}
	}
	e := catalog[0].Entries[0]
	return e, catalog[0].Title
}

// sectionsFor returns the sidebar sections whose items match the query
// and the library filter: every item when the query is empty, else
// those whose name or ID contains it, case-insensitively; and only the
// filtered library's items when the filter is on. Each item carries the
// icon of its component.
func sectionsFor(query string) []sidebar.Section {
	q := strings.ToLower(strings.TrimSpace(query))
	lib, filtered := filterLib()
	var out []sidebar.Section
	for _, s := range catalog {
		var items []sidebar.Item
		for _, e := range s.Entries {
			if filtered && e.Lib != lib {
				continue
			}
			if q == "" || strings.Contains(strings.ToLower(e.Name), q) || strings.Contains(strings.ToLower(e.ID), q) {
				items = append(items, sidebar.Item{ID: e.ID, Label: e.Name, Icon: e.Icon()})
			}
		}
		if len(items) > 0 {
			out = append(out, sidebar.Section{Title: s.Title, Items: items})
		}
	}
	return out
}

// filterLib returns the library the sidebar is filtered to, and whether
// a filter is on; Filter 0 (全部) filters none.
func filterLib() (Library, bool) {
	if State.Filter == 0 {
		return 0, false
	}
	return Library(State.Filter - 1), true
}

// firstOf returns the first catalog entry of the library, for the
// filter to land on a page it shows.
func firstOf(lib Library) (Entry, bool) {
	for _, s := range catalog {
		for _, e := range s.Entries {
			if e.Lib == lib {
				return e, true
			}
		}
	}
	return Entry{}, false
}

// inLib reports whether the page chosen belongs to the library.
func inLib(lib Library, id string) bool {
	for _, s := range catalog {
		for _, e := range s.Entries {
			if e.ID == id {
				return e.Lib == lib
			}
		}
	}
	return false
}
