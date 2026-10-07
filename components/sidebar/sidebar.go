// Package sidebar provides a shadcn/ui-style Sidebar: a rail of sections
// and items at the side of a window, as Finder's sidebar shows one. A
// click on an item chooses it, setting *Selected to its ID; the arrows,
// Home and End and the first letters of an item move the choice among the
// items. A section's title shows and hides its items with a click, unless
// its Open is nil, which keeps it open. The item chosen takes the focus
// and reads selected to assistive technology.
//
//	sidebar.Sidebar(c, sidebar.Props{
//		Selected: &app.place,
//		Sections: []sidebar.Section{
//			{Title: "Favorites", Items: []sidebar.Item{
//				{ID: "recents", Label: "Recents"},
//				{ID: "desktop", Label: "Desktop"},
//			}},
//		},
//	})
package sidebar

import "github.com/egoist/mygo/ui"

// Item is one entry of a Section, choosing ID when it is clicked.
type Item struct {
	// ID is what *Selected holds while the item is chosen.
	ID string
	// Label is the text of the item, such as "Recents".
	Label string
}

// Section is a group of Items under a title, which a click on the title
// shows and hides.
type Section struct {
	// Title is the text of the section's title.
	Title string
	// Open is whether the items show; nil keeps the section open.
	Open *bool
	// Items are the entries of the section, in order.
	Items []Item
}

// Props describes the sidebar to draw.
type Props struct {
	// Selected is set to the ID of the item a click chooses.
	Selected *string
	// Sections are the sections of the sidebar, in order.
	Sections []Section
	// Disabled keeps the items from being chosen.
	Disabled bool
}

// Sidebar draws the rail and returns it, so a view can chain more calls
// on it.
func Sidebar(c *ui.Context, p Props) *ui.Element {
	e := ui.Sidebar(c, p.Selected, func() {
		for _, s := range p.Sections {
			ui.SidebarSection(c, s.Title, s.Open, func() {
				for _, it := range s.Items {
					ui.SidebarItem(c, it.ID, nil, it.Label)
				}
			})
		}
	})
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
