// Package appshell provides the AppShell of BoardUI for the MyGo native
// toolkit: the page frame of a starter, with the sidebar, the breadcrumb
// trail, the page heading and its actions, and the content below.
//
//	appshell.AppShell(c, appshell.Props{
//		Title:    "Dashboard",
//		Team:     "Board team",
//		Member:   "Mertcan",
//		Selected: &place,
//		Sections: sections,
//		Content:  func(c *ui.Context) { ... },
//	})
//
// The shell builds on the sidebar and breadcrumbs components; install
// them with it. It fills the window it is given, so a window can size
// it with Grow.
package appshell

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/breadcrumbs"
	"github.com/smallnest/uix/components/sidebar"
)

// Props describes the shell to draw.
type Props struct {
	// Title is the page's name: the last breadcrumb item and the
	// default heading.
	Title string
	// Team and Member are the ancestors of the trail, such as the team
	// and the member of the app; empty ones are left out.
	Team, Member string
	// Heading is the big text over the content; the Title when empty.
	Heading string
	// Selected is the sidebar's choice, by item ID, as the sidebar
	// sets it.
	Selected *string
	// Sections are the sidebar's sections.
	Sections []sidebar.Section
	// Actions draw the header's buttons on the right; nil draws none.
	Actions func(c *ui.Context)
	// Content is the page below the header; nil draws nothing.
	Content func(c *ui.Context)
}

// AppShell draws the page frame and returns it, so a view can chain more
// calls on it.
func AppShell(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	return ui.Row(c).Fill().Gap(t.Space(4)).Padding(t.Space(3)).Children(func() {
		ui.Box(c).Width(260).FillHeight().Shrink(0).Background(t.Surface).
			Radius(t.Radius * 3).Padding(t.Space(2)).Children(func() {
			sidebar.Sidebar(c, sidebar.Props{
				Selected: p.Selected,
				Sections: p.Sections,
			})
		})
		ui.Column(c).Grow(1).MinWidth(0).FillHeight().Gap(t.Space(2.5)).
			Children(func() {
				breadcrumbRow(c, t, p)
				headingRow(c, t, p)
				if p.Content != nil {
					ui.Column(c).FillWidth().Grow(1).MinHeight(0).Gap(t.Space(4)).
						Children(func() {
							p.Content(c)
						})
				}
			})
	})
}

// breadcrumbRow is the trail to the page: the team, the member and the
// page, the last one where the path is.
func breadcrumbRow(c *ui.Context, t *ui.Theme, p Props) {
	var items []string
	if p.Team != "" {
		items = append(items, p.Team)
	}
	if p.Member != "" {
		items = append(items, p.Member)
	}
	items = append(items, p.Title)
	chosen := len(items) - 1
	breadcrumbs.Breadcrumbs(c, breadcrumbs.Props{Items: items, Chosen: &chosen})
}

// headingRow is the page heading on the left, with the actions on the
// right.
func headingRow(c *ui.Context, t *ui.Theme, p Props) {
	heading := p.Heading
	if heading == "" {
		heading = p.Title
	}
	ui.Row(c).FillWidth().Justify(ui.SpaceBetween).AlignItems(ui.Center).
		Gap(t.Space(4)).Children(func() {
		ui.Text(c, heading).FontSize(t.FontSize * 1.5).FontWeight(600).TextColor(t.Text)
		if p.Actions != nil {
			ui.Row(c).Gap(t.Space(2.5)).Children(func() { p.Actions(c) })
		}
	})
}
