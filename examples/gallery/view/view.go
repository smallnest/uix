// Package view builds the gallery example: a shell like onevra's
// showcase, with a sidebar of the components and one page each. Every
// page shows a component of one of the three libraries — uix's own,
// mygo's primitives, or onevra's — with a badge naming the library;
// where the same component exists in several, the gallery keeps the
// one of uix, then mygo, then onevra. A search box above the sidebar
// filters the items. The main package shows it in a window; the tests
// and the snapshot command draw it headless.
package view

import (
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/breadcrumbs"
	"github.com/smallnest/uix/components/scroll"
	"github.com/smallnest/uix/components/searchfield"
	"github.com/smallnest/uix/components/sidebar"
)

// Width and Height are the size of the example window.
const Width, Height = 960, 600

// GalleryView draws the gallery shell: the sidebar of components on the
// left, the page of the chosen one on the right.
func GalleryView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
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
		ui.Text(c, "Gallery").FontSize(t.FontSize * 1.2).Bold()
		ui.Text(c, "uix · mygo · onevra").FontSize(t.FontSize * 0.8).TextColor(t.TextMuted)
		searchfield.SearchField(c, searchfield.Props{Value: &State.Query})
		sections := sectionsFor(State.Query)
		if len(sections) == 0 {
			ui.Text(c, "No component matches.").FontSize(t.FontSize * 0.85).TextColor(t.TextMuted)
			return
		}
		// The sidebar scrolls its items in the rail's remaining height, as
		// the settings rail does with its categories.
		sidebar.Sidebar(c, sidebar.Props{Selected: &State.Selected, Sections: sections}).FillHeight()
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

// sectionsFor returns the sidebar sections whose items match the query:
// every item when the query is empty, else those whose name or ID
// contains it, case-insensitively.
func sectionsFor(query string) []sidebar.Section {
	q := strings.ToLower(strings.TrimSpace(query))
	var out []sidebar.Section
	for _, s := range catalog {
		var items []sidebar.Item
		for _, e := range s.Entries {
			if q == "" || strings.Contains(strings.ToLower(e.Name), q) || strings.Contains(strings.ToLower(e.ID), q) {
				items = append(items, sidebar.Item{ID: e.ID, Label: e.Name})
			}
		}
		if len(items) > 0 {
			out = append(out, sidebar.Section{Title: s.Title, Items: items})
		}
	}
	return out
}
