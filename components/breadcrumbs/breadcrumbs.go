// Package breadcrumbs provides a shadcn/ui-style Breadcrumb: a path of
// items separated by chevrons, as Finder's path bar shows one. The items
// but the last are links, which a click chooses, as Enter does, setting
// *Chosen to the index of the item clicked; the last item is where the
// path is, not a link. Long items shorten with an ellipsis as room runs
// short.
//
//	breadcrumbs.Breadcrumbs(c, breadcrumbs.Props{
//		Items:  []string{"Home", "Docs", "Guide"},
//		Chosen: &app.chosen,
//	})
package breadcrumbs

import "github.com/egoist/mygo/ui"

// Props describes the breadcrumb path to draw.
type Props struct {
	// Items are the items of the path, in order; the last is where the
	// path is.
	Items []string
	// Chosen is set to the index of the item a click chooses; -1 while
	// no item has been chosen.
	Chosen *int
	// Disabled keeps the items from being chosen.
	Disabled bool
}

// Breadcrumbs draws the path and returns it, so a view can chain more
// calls on it.
func Breadcrumbs(c *ui.Context, p Props) ui.Element {
	e := ui.Breadcrumbs(c, p.Items, p.Chosen)
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}
