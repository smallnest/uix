// Package view builds the editor example: a window of MyGo native UI
// that shows the sidebar, togglegroup, tooltip and menu components
// together in a note editor. The sidebar chooses the file being edited,
// the toggles restyle its text, tips name the controls, and the file
// menu opens a file or saves. The main package shows it in a window; the
// tests and the snapshot command draw it headless.
package view

import (
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/menu"
	"github.com/smallnest/uix/components/separator"
	"github.com/smallnest/uix/components/sidebar"
	"github.com/smallnest/uix/components/togglegroup"
	"github.com/smallnest/uix/components/tooltip"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 480

// State is the state of the example; the controls edit it in place.
var State = EditorState{}

// EditorState holds the editor, as the controls edit it.
type EditorState struct {
	// Place is the file chosen in the sidebar: "notes", "todo",
	// "journal" or "draft".
	Place string
	// FilesOpen opens the Files section of the sidebar.
	FilesOpen bool
	// Bold, Italic, Underline and Strike style the note.
	Bold, Italic, Underline, Strike bool
	// Note replaces the text of the file, for the New note menu item;
	// empty shows the text of the file chosen.
	Note string
	// Status is the last thing the menu or Save did, shown under the
	// editor; empty for none.
	Status string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = EditorState{Place: "notes", FilesOpen: true, Bold: true}
}

// notes are the texts of the files the sidebar names.
var notes = map[string]string{
	"notes":   "Buy milk\nCall the plumber.",
	"todo":    "- finish uix batch 5\n- ship the editor example",
	"journal": "October 7 — shipped batches 1 through 4.",
	"draft":   "An unfinished thought.",
}

// styleOf names the toggles that are on, as a comma-separated list, or
// "none".
func styleOf() string {
	var parts []string
	if State.Bold {
		parts = append(parts, "bold")
	}
	if State.Italic {
		parts = append(parts, "italic")
	}
	if State.Underline {
		parts = append(parts, "underline")
	}
	if State.Strike {
		parts = append(parts, "strike")
	}
	if len(parts) == 0 {
		return "none"
	}
	return strings.Join(parts, ", ")
}

// EditorView draws the example: a note editor whose controls all work.
func EditorView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Row(c).Fill().AlignItems(ui.Stretch).Children(func() {
		sidebar.Sidebar(c, sidebar.Props{
			Selected: &State.Place,
			Sections: []sidebar.Section{
				{Title: "Files", Open: &State.FilesOpen, Items: []sidebar.Item{
					{ID: "notes", Label: "Notes"},
					{ID: "todo", Label: "Todo"},
					{ID: "journal", Label: "Journal"},
				}},
				{Title: "Recent", Items: []sidebar.Item{
					{ID: "draft", Label: "Draft"},
				}},
			},
		}).Width(160)
		ui.Column(c).Grow(1).Padding(24).Gap(14).Children(func() {
			ui.Text(c, "Editor").FontSize(24).Bold()
			ui.Row(c).Gap(8).Children(func() {
				// The menu holds the file actions, as a button that opens
				// a menu below it.
				menu.Menu(c, menu.Props{
					Label: "File",
					Items: []menu.Item{
						{Label: "New note", Action: func() { State.Note = ""; State.Status = "New note." }},
						{Label: "Open todo.md", Action: func() { State.Place = "todo"; State.Note = ""; State.Status = "Opened todo.md." }},
						{Separator: true},
						{Label: "Save", Action: func() { State.Status = "Saved." }},
					},
				})
				togglegroup.ToggleGroup(c, togglegroup.Props{
					Items: []togglegroup.Item{
						{On: &State.Bold, Label: "B", Tip: "Bold"},
						{On: &State.Italic, Label: "I", Tip: "Italic"},
						{On: &State.Underline, Label: "U", Tip: "Underline"},
						{On: &State.Strike, Label: "S", Tip: "Strike"},
					},
				})
				ui.Spacer(c)
				// The save button shows the tooltip component; the toggles
				// carry their own tips.
				save := ui.Button(c, "Save")
				tooltip.Tooltip(c, tooltip.Props{Anchor: save, Text: "Save the note"})
				if save.Clicked() {
					State.Status = "Saved."
				}
			})
			separator.Separator(c, separator.Props{})
			// The canvas shows the note of the file chosen, bold while the
			// Bold toggle is on.
			ui.Box(c).Grow(1).Radius(10).Border(1, t.Border).Padding(16).Children(func() {
				ui.Text(c, "Editing "+State.Place+".md").TextColor(t.TextMuted)
				line := ui.Text(c, textOf())
				if State.Bold {
					line.FontWeight(700)
				}
			})
			ui.Text(c, "Style: "+styleOf()).TextColor(t.TextMuted)
			if State.Status != "" {
				ui.Text(c, State.Status).TextColor(t.TextMuted)
			}
		})
	})
}

// textOf returns the text of the note, the file's or the new note's.
func textOf() string {
	if State.Note != "" {
		return State.Note
	}
	return notes[State.Place]
}
