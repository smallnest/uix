// Package view builds the gallery example: a window of MyGo native UI
// that shows the grid, gridview, scroll and fieldset components together
// in a file gallery. Two field sets filter the files by type and sort
// them by name or date, and a grid view shows the files that remain, of
// which a click chooses one. The main package shows it in a window; the
// tests and the snapshot command draw it headless.
package view

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/fieldset"
	"github.com/smallnest/uix/components/gridview"
	"github.com/smallnest/uix/components/radio"
	"github.com/smallnest/uix/components/separator"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 480

// State is the state of the example; the controls edit it in place.
var State = GalleryState{}

// GalleryState holds the gallery, as the controls edit it.
type GalleryState struct {
	// Kind is the type shown: 0 photos, 1 docs, 2 videos.
	Kind int
	// Sort is the order: 0 by name, 1 by date.
	Sort int
	// Chosen is the item chosen in the grid, -1 for none.
	Chosen int
	// Status is the last thing the controls did, shown at the bottom.
	Status string
}

// gridState keeps the place and the choice of the grid.
var gridState ui.GridState

// lastKind and lastSort let the view tell a filter change from a plain
// rebuild: the radios edit the state in place, and a group has no change
// of its own to report.
var lastKind, lastSort int

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = GalleryState{
		Chosen: -1,
		Status: "Choose an item, or change the type or the sort.",
	}
	lastKind, lastSort = 0, 0
	// The grid keeps its place in a state with the frame it was built
	// in; a fresh frame must not inherit it.
	gridState = ui.GridState{}
}

// file is one item of the gallery.
type file struct {
	name string
	kind int // 0 photo, 1 doc, 2 video
	day  int // a day of the month, which the date sort orders by
}

// files is the gallery: photos, docs and videos in a jumbled order, so
// the two sorts move them around.
var files = []file{
	{"Alpine.jpg", 0, 12},
	{"Budget.pdf", 1, 5},
	{"Beach.png", 0, 21},
	{"Notes.txt", 1, 3},
	{"Clip.mp4", 2, 9},
	{"Dune.mp4", 2, 17},
	{"City.jpg", 0, 4},
	{"Report.pdf", 1, 14},
	{"Wave.mp4", 2, 2},
	{"Meadow.jpg", 0, 25},
	{"Forest.png", 0, 16},
	{"Spec.pdf", 1, 10},
	{"River.jpg", 0, 8},
	{"Reel.mp4", 2, 22},
	{"Snow.png", 0, 29},
	{"Sunset.jpg", 0, 19},
}

// visible returns the files of the chosen kind, in the chosen order.
func visible() []file {
	out := make([]file, 0, len(files))
	for _, f := range files {
		if f.kind == State.Kind {
			out = append(out, f)
		}
	}
	if State.Sort == 1 {
		slices.SortFunc(out, func(a, b file) int { return cmp.Compare(a.day, b.day) })
	} else {
		slices.SortFunc(out, func(a, b file) int { return strings.Compare(a.name, b.name) })
	}
	return out
}

// GalleryView draws the example: a file gallery whose filters and sorts
// all work.
func GalleryView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	if State.Kind != lastKind || State.Sort != lastSort {
		lastKind, lastSort = State.Kind, State.Sort
		State.Chosen = -1
		State.Status = filterStatus()
	}
	shown := visible()
	gridState.Selected = &State.Chosen
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).PaddingX(16).PaddingY(10).Gap(12).AlignItems(ui.Center).Children(func() {
			ui.Text(c, "Gallery").Bold().Grow(1)
			ui.Text(c, fmt.Sprintf("%d items", len(shown))).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		})
		separator.Separator(c, separator.Props{})
		ui.Row(c).PaddingX(24).PaddingY(12).Gap(24).AlignItems(ui.Start).Children(func() {
			fieldset.Fieldset(c, fieldset.Props{
				Legend: "Type",
				Children: func(c *ui.Context) {
					radio.Group(c, radio.Props[int]{
						Selected: &State.Kind,
						Options: []radio.Option[int]{
							{Value: 0, Label: "Photos"},
							{Value: 1, Label: "Docs"},
							{Value: 2, Label: "Videos"},
						},
					})
				},
			})
			fieldset.Fieldset(c, fieldset.Props{
				Legend: "Sort",
				Children: func(c *ui.Context) {
					radio.Group(c, radio.Props[int]{
						Selected: &State.Sort,
						Options: []radio.Option[int]{
							{Value: 0, Label: "Name"},
							{Value: 1, Label: "Date"},
						},
					})
				},
			})
		})
		separator.Separator(c, separator.Props{})
		gridview.GridView(c, gridview.Props{
			State:    &gridState,
			Count:    len(shown),
			MinWidth: 130,
			Height:   96,
			Item: func(c *ui.Context, i int) {
				cell(c, shown[i])
			},
			OnChoose: func(i int) {
				State.Chosen = i
				State.Status = "Chosen " + shown[i].name + "."
			},
		}).PaddingX(24).PaddingY(12).Grow(1)
		separator.Separator(c, separator.Props{})
		ui.Row(c).PaddingX(24).PaddingY(10).Children(func() {
			ui.Text(c, State.Status).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		})
	})
}

// filterStatus names the filters the status line shows.
func filterStatus() string {
	kinds := []string{"Photos", "Docs", "Videos"}
	sorts := []string{"name", "date"}
	return "Showing " + kinds[State.Kind] + " by " + sorts[State.Sort] + "."
}

// cell draws one file of the grid: a tile in a color for its type, and
// its name. The cell does not draw the choice; the grid styles the
// chosen cell itself.
func cell(c *ui.Context, f file) {
	t := c.Theme()
	ui.Column(c).AlignItems(ui.Center).Gap(6).Children(func() {
		ui.Box(c).Size(56, 56).Radius(t.Space(2)).Background(tileColor(c, f.kind))
		ui.Text(c, f.name).FontSize(t.FontSize * 0.85).SingleLine()
	})
}

// tileColor returns the tile color of a type, from the theme: the accent
// for photos, the success for docs, the warning for videos.
func tileColor(c *ui.Context, kind int) ui.Color {
	t := c.Theme()
	switch kind {
	case 0:
		return t.Accent
	case 1:
		return t.Success
	default:
		return t.Warning
	}
}
