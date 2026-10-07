// Package view builds the files example: a window of MyGo native UI that
// shows the table component at work, as a small file browser. The main
// package shows it in a window; the tests and the snapshot command draw
// it headless.
package view

import (
	"fmt"
	"slices"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/table"
)

// Width and Height are the size of the example window.
const Width, Height = 480, 520

// File is a row of the table.
type File struct {
	Name, Kind string
	Size       int
}

// Files are the rows the table shows, in their original order.
var Files = []File{
	{"README.md", "Markdown", 3124},
	{"main.go", "Go", 2840},
	{"go.mod", "Module", 84},
	{"internal/ui/components/table", "Directory", 0},
	{"LICENSE", "Text", 1065},
	{"go.sum", "Checksum", 12340},
	{"Makefile", "Text", 512},
}

// State is the state of the example; the controls edit it in place.
var State = TableState{Selected: -1, List: &ui.ListState{}}

// TableState holds the choice, the sort and the row opened, plus the
// ListState that keeps where the rows are as they scroll.
type TableState struct {
	Selected int
	Sort     ui.SortOrder
	List     *ui.ListState
	Opened   string
}

// Sorted returns the rows in the order the sort asks for, or the
// original order while no column is sorted.
func Sorted() []File {
	rows := slices.Clone(Files)
	if State.Sort.Column == "" {
		return rows
	}
	desc := State.Sort.Descending
	slices.SortFunc(rows, func(a, b File) int {
		switch State.Sort.Column {
		case "Size":
			if a.Size != b.Size {
				return cmp(a.Size, b.Size, desc)
			}
		case "Kind":
			if a.Kind != b.Kind {
				return cmpStr(a.Kind, b.Kind, desc)
			}
		}
		return cmpStr(a.Name, b.Name, desc)
	})
	return rows
}

// cmp orders ints ascending, or descending when asked.
func cmp(a, b int, desc bool) int {
	if desc {
		return b - a
	}
	return a - b
}

// cmpStr orders strings ascending, or descending when asked.
func cmpStr(a, b string, desc bool) int {
	switch {
	case a < b:
		if desc {
			return 1
		}
		return -1
	case a > b:
		if desc {
			return -1
		}
		return 1
	}
	return 0
}

// FilesView draws the example: a table of files under a title, and a
// line that names the row chosen or opened.
func FilesView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	rows := Sorted()
	State.List.Selected = &State.Selected
	State.List.Sort = &State.Sort
	ui.Column(c).Fill().Padding(24).Gap(14).Children(func() {
		ui.Text(c, "Files").FontSize(24).Bold()
		tbl := table.Table(c, table.Props{
			State: State.List,
			Columns: []ui.TableColumn{
				{Title: "Name", Sortable: true},
				{Title: "Kind", Width: t.Space(26), Sortable: true},
				{Title: "Size", Width: t.Space(20), Align: ui.End, Sortable: true},
			},
			Rows: len(rows),
			Cell: func(row, col int) {
				f := rows[row]
				switch col {
				case 0:
					ui.Text(c, f.Name).SingleLine()
				case 1:
					ui.Text(c, f.Kind).SingleLine()
				case 2:
					ui.Text(c, fmt.Sprintf("%d B", f.Size))
				}
			},
		})
		tbl.Grow(1)
		if tbl.Submitted() && State.Selected >= 0 {
			State.Opened = rows[State.Selected].Name
		}
		switch {
		case State.Opened != "":
			ui.Text(c, "Opened "+State.Opened).TextColor(t.Success).Bold()
		case State.Selected >= 0 && State.Selected < len(rows):
			ui.Text(c, "Selected: "+rows[State.Selected].Name).TextColor(t.TextMuted)
		default:
			ui.Text(c, "No file selected").TextColor(t.TextMuted)
		}
	})
}
