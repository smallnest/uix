// Package view builds the workspace example: a window of MyGo native UI
// that shows the form, scrollhorizontal, scrollboth and link components
// of uix together. The main package shows it in a window; the tests and
// the snapshot command draw it headless.
package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/form"
	"github.com/smallnest/uix/components/input"
	"github.com/smallnest/uix/components/link"
	"github.com/smallnest/uix/components/scrollboth"
	"github.com/smallnest/uix/components/scrollhorizontal"
	selector "github.com/smallnest/uix/components/select"
)

// Width and Height are the size of the example window.
const Width, Height = 520, 580

// State is the state of the example; the controls edit it in place.
var State = WorkspaceState{Day: -1}

// WorkspaceState holds the values the controls edit, and the status line.
type WorkspaceState struct {
	Name   string
	Owner  string
	Sprint string
	Day    int
	Cell   string
	Dark   bool
	Status string
}

// sprints are the choices of the sprint field.
var sprints = []string{"Sprint 1", "Sprint 2", "Sprint 3"}

// Reset restores the initial state of the example.
func Reset() {
	State = WorkspaceState{Sprint: "Sprint 1", Day: -1, Status: "Pick a day of the sprint."}
}

// WorkspaceView draws the example: a project panel whose form lines up
// its fields, whose days scroll sideways, whose board scrolls both ways,
// and whose header links open the guide. The form edits the project, a
// click on a day picks it, a click on a cell of the board names it, and
// the status line shows what changed.
func WorkspaceView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark || State.Dark))
	t := c.Theme()
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		// The header: a title, a spacer, and two links.
		ui.Row(c).Fill().AlignItems(ui.Center).Children(func() {
			ui.Text(c, "Workspace").FontSize(24).Bold()
			ui.Spacer(c)
			link.Link(c, link.Props{Label: "Open the guide", URL: "https://example.com/guide"})
			ui.Text(c, " · ").TextColor(t.TextMuted)
			link.Link(c, link.Props{Label: "Report an issue", URL: "https://example.com/issues"})
		})
		ui.Row(c).Fill().Gap(20).Children(func() {
			// The left column: the macOS-style form of the project.
			ui.Column(c).Grow(1).Gap(6).Children(func() {
				caption(c, "Project")
				form.Form(c, form.Props{Children: func(c *ui.Context) {
					ui.Field(c, "Project name", func() {
						input.Input(c, input.Props{Value: &State.Name, Placeholder: "My project"})
					})
					ui.Field(c, "Owner", func() {
						input.Input(c, input.Props{Value: &State.Owner, Placeholder: "you@example.com"})
					})
					ui.Field(c, "Sprint", func() {
						selector.Select(c, selector.Props{Selected: &State.Sprint, Options: sprints})
					})
				}})
			})
			// The right column: the days that scroll sideways, above the
			// board that scrolls both ways.
			ui.Column(c).Grow(1).Fill().Gap(6).Children(func() {
				caption(c, "Days")
				scrollhorizontal.ScrollHorizontal(c, scrollhorizontal.Props{Children: func(c *ui.Context) {
					ui.Row(c).Gap(6).Children(func() {
						for d := 1; d <= 12; d++ {
							day := dayChip(c, t, d)
							if day.Clicked() {
								State.Day = d
								State.Status = fmt.Sprintf("Day %d of the sprint.", d)
							}
						}
					})
				}}).Height(48)
				caption(c, "Board")
				scrollboth.ScrollBoth(c, scrollboth.Props{Children: func(c *ui.Context) {
					ui.Row(c).Gap(6).Children(func() {
						for r := 0; r < 10; r++ {
							ui.Column(c).Gap(6).Children(func() {
								for col := 0; col < 10; col++ {
									cell := ui.Box(c).Size(52, 28).Radius(4).
										Background(t.Surface).Border(1, t.Border).Center().
										Children(func() {
											ui.Text(c, fmt.Sprintf("r%dc%d", r, col)).FontSize(12).TextColor(t.TextMuted)
										})
									if cell.Clicked() {
										State.Cell = fmt.Sprintf("r%dc%d", r, col)
										State.Status = fmt.Sprintf("Cell r%dc%d chosen.", r, col)
									}
								}
							})
						}
					})
				}}).Grow(1)
			})
		})
		// The status line shows what changed.
		ui.Text(c, State.Status).FontSize(13).TextColor(t.TextMuted).Grow(1)
	})
}

// dayChip draws one day of the sprint: a chip that highlights when it is
// the chosen one.
func dayChip(c *ui.Context, t *ui.Theme, d int) *ui.Element {
	chosen := State.Day == d
	bg, fg, bd := t.Surface, t.TextMuted, t.Border
	if chosen {
		bg, fg, bd = t.Accent, t.AccentText, t.Accent
	}
	return ui.Box(c).Size(56, 34).Radius(6).Background(bg).Border(1, bd).Center().
		Children(func() {
			ui.Text(c, fmt.Sprintf("Day %d", d)).FontSize(12).Bold().TextColor(fg)
		})
}

// caption draws the small muted label of a section.
func caption(c *ui.Context, s string) {
	t := c.Theme()
	ui.Text(c, s).FontSize(t.FontSize * 0.875).Bold().TextColor(t.TextMuted)
}

// Title returns the title of the window.
func Title() string {
	return "Workspace"
}
