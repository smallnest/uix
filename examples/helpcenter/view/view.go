// Package view builds the help center example: a window of MyGo native
// UI that shows the breadcrumbs, accordion, collapsible and popover
// components together in a help page. The breadcrumbs say where the user
// is, the accordion opens the answers to the questions, a collapsible
// hides more options, and a popover jumps to a section. The main package
// shows it in a window; the tests and the snapshot command draw it
// headless.
package view

import (
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/accordion"
	"github.com/smallnest/uix/components/breadcrumbs"
	"github.com/smallnest/uix/components/collapsible"
	"github.com/smallnest/uix/components/popover"
	"github.com/smallnest/uix/components/separator"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 560

// State is the state of the example; the controls edit it in place.
var State = HelpState{}

// HelpState holds the help page, as the controls edit it.
type HelpState struct {
	// Chosen is where the user is: an index into path.
	Chosen int
	// FAQ1, FAQ2 and FAQ3 open the sections of the accordion.
	FAQ1, FAQ2, FAQ3 bool
	// More opens the collapsible under the accordion.
	More bool
	// HelpOpen shows the popover of the help button.
	HelpOpen bool
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = HelpState{Chosen: 2, FAQ1: true, More: true}
}

// path is the breadcrumb trail of the page.
var path = []string{"Home", "Docs", "FAQ"}

// faqs are the sections of the accordion, in order.
var faqs = []struct {
	Title string
	Open  *bool
	Body  string
}{
	{"What is uix?", &State.FAQ1, "uix is a component registry for MyGo apps. It copies components into your project with uix add, as shadcn/ui does for web apps."},
	{"How do I add components?", &State.FAQ2, "Run uix add button in your Go module to copy a component and what it depends on. Each component imports only mygo, so you own the code."},
	{"How do I build a theme?", &State.FAQ3, "Pick a palette with uix theme set, or set a custom one with uix theme set custom and a hex color of your own."},
}

// HelpView draws the example: a help page whose controls all work.
func HelpView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Help Center").FontSize(24).Bold()
		breadcrumbs.Breadcrumbs(c, breadcrumbs.Props{
			Items:  path,
			Chosen: &State.Chosen,
		})
		ui.Text(c, "You are at "+strings.Join(path[:State.Chosen+1], " › ")).TextColor(t.TextMuted)
		separator.Separator(c, separator.Props{})
		accordion.Accordion(c, accordion.Props{
			Items: []accordion.Item{
				{Title: faqs[0].Title, Open: faqs[0].Open, Build: body(c, t, faqs[0].Body)},
				{Title: faqs[1].Title, Open: faqs[1].Open, Build: body(c, t, faqs[1].Body)},
				{Title: faqs[2].Title, Open: faqs[2].Open, Build: body(c, t, faqs[2].Body)},
			},
		})
		collapsible.Collapsible(c, collapsible.Props{
			Label: "More help options",
			Open:  &State.More,
			Children: func(c *ui.Context) {
				ui.Text(c, "Email support at help@uix.dev, or open an issue on GitHub.").TextColor(t.TextMuted)
			},
		})
		// The popover jumps to a section of the page; its links close it.
		popover.Popover(c, popover.Props{
			Open: &State.HelpOpen,
			Trigger: func(c *ui.Context) ui.Element {
				return ui.Button(c, "Help")
			},
			Content: func(c *ui.Context) {
				ui.Column(c).Gap(4).Children(func() {
					if ui.Button(c, "Go to Docs").Clicked() {
						State.HelpOpen = false
						State.Chosen = 1
					}
					if ui.Button(c, "Go to FAQ").Clicked() {
						State.HelpOpen = false
						State.Chosen = 2
					}
				})
			},
		})
	})
}

// body returns a Build that draws the body of a section, in the muted
// text of the theme.
func body(c *ui.Context, t *ui.Theme, s string) func(*ui.Context) {
	return func(c *ui.Context) {
		ui.Text(c, s).TextColor(t.TextMuted)
	}
}
