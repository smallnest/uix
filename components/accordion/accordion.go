// Package accordion provides a shadcn/ui-style Accordion: a column of
// sections in a bordered box, one above the other, that a click opens and
// closes, as do Enter and Space, and that the Up and Down arrows move
// between. Each section keeps its own open state, so several may be open
// at once.
//
//	accordion.Accordion(c, accordion.Props{
//		Items: []accordion.Item{
//			{Title: "General", Open: &app.general, Build: buildGeneral},
//			{Title: "Privacy", Open: &app.privacy, Build: buildPrivacy},
//		},
//	})
//
// To keep one section open at a time, close the others as one opens, as
// this registry does in its examples: build the sections with
// AccordionItem and an index of the one open:
//
//	for i, s := range sections {
//		open := app.section == i
//		if accordion.AccordionItem(c, accordion.Item{Title: s, Open: &open}).Changed() {
//			app.section = -1
//			if open {
//				app.section = i
//			}
//		}
//	}
package accordion

import "github.com/egoist/mygo/ui"

// Item is one section of an Accordion: a header showing Title, which a
// click opens and closes, and below it what Build draws while *Open is
// true.
type Item struct {
	// Title is the text of the header, such as "General".
	Title string
	// Open is whether the section is open.
	Open *bool
	// Build draws the content of the section while it is open.
	Build func(c *ui.Context)
}

// Props describes the accordion to draw.
type Props struct {
	// Items are the sections of the accordion, in order.
	Items []Item
	// Disabled keeps the sections from opening and closing.
	Disabled bool
}

// Accordion draws the sections in a bordered box and returns it, so a
// view can chain more calls on it.
func Accordion(c *ui.Context, p Props) ui.Element {
	e := ui.Accordion(c, func() {
		for _, it := range p.Items {
			ui.AccordionItem(c, it.Title, it.Open, func() {
				if it.Build != nil {
					it.Build(c)
				}
			})
		}
	})
	if p.Disabled {
		e.Disabled(true)
	}
	return e
}

// AccordionItem builds one section of an Accordion, for a view that
// keeps one open at a time, or that builds its sections itself.
func AccordionItem(c *ui.Context, it Item) ui.Element {
	return ui.AccordionItem(c, it.Title, it.Open, func() {
		if it.Build != nil {
			it.Build(c)
		}
	})
}
