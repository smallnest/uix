// Package view builds the explorer example: a window of MyGo native UI
// that shows the toolbar, tree, split and list components together in a
// small file browser. The toolbar creates files and opens the chosen
// one, the tree chooses a folder, the list shows its files, and the
// divider between them is draggable. The main package shows it in a
// window; the tests and the snapshot command draw it headless.
package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/list"
	"github.com/smallnest/uix/components/menu"
	"github.com/smallnest/uix/components/separator"
	"github.com/smallnest/uix/components/split"
	"github.com/smallnest/uix/components/toolbar"
	"github.com/smallnest/uix/components/tree"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 480

// State is the state of the example; the controls edit it in place.
var State = ExplorerState{}

// ExplorerState holds the explorer, as the controls edit it.
type ExplorerState struct {
	// SrcOpen, LibOpen and DocsOpen open the branches of the tree.
	SrcOpen, LibOpen, DocsOpen bool
	// Folder is the folder whose files the list shows: "src", "lib" or
	// "docs"; empty shows the files at the root.
	Folder string
	// TreeChosen is the label of the chosen node of the tree.
	TreeChosen string
	// Chosen is the row chosen in the list, -1 for none.
	Chosen int
	// Divider is the width of the tree pane, which the divider edits.
	Divider float32
	// ShowSizes shows the size of a file in its row.
	ShowSizes bool
	// NewCount counts the files New created, to name them.
	NewCount int
	// Status is the last thing the controls did, shown under the list;
	// empty for none.
	Status string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = ExplorerState{
		SrcOpen: true, Folder: "src", TreeChosen: "src", Chosen: 0,
		Divider: 180, ShowSizes: true, Status: "Showing src.",
	}
	files = make(map[string][]file, len(initial))
	for k, v := range initial {
		files[k] = append([]file(nil), v...)
	}
	root = append([]file(nil), initialRoot...)
	// The list state remembers the frame it showed, which a new tester
	// would mistake for a second list of the same frame, so start it
	// over with the state.
	listState = ui.ListState{}
}

// file is one row of the list: a name and a size.
type file struct {
	Name string
	Size int
}

// initial are the files of each folder at start, and initialRoot the
// files at the root, which Reset restores after New has added to them.
var initial = map[string][]file{
	"src":  {{Name: "main.go", Size: 1204}, {Name: "app.go", Size: 892}},
	"lib":  {{Name: "util.go", Size: 421}},
	"docs": {{Name: "guide.md", Size: 1874}, {Name: "README.md", Size: 2331}},
}
var initialRoot = []file{{Name: "go.mod", Size: 96}, {Name: "Makefile", Size: 310}}

// files are the files of each folder, and root the files at the root;
// New adds to them, and Reset restores them.
var files = map[string][]file{}
var root = []file{}

// filesFor returns the files the list shows: of the folder chosen, or
// the root when none is.
func filesFor() []file {
	if State.Folder == "" {
		return root
	}
	return files[State.Folder]
}

// listState keeps the place and the choice of the file list.
var listState ui.ListState

// ExplorerView draws the example: a file browser whose controls all
// work.
func ExplorerView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Children(func() {
		toolbar.Toolbar(c, toolbar.Props{
			Label: "Explorer",
			Children: func(c *ui.Context) {
				if ui.Button(c, "New").Clicked() {
					State.NewCount++
					f := file{Name: fmt.Sprintf("new%d.txt", State.NewCount), Size: 16}
					if State.Folder == "" {
						root = append(root, f)
					} else {
						files[State.Folder] = append(files[State.Folder], f)
					}
					State.Status = "Created " + f.Name + "."
				}
				menu.Menu(c, menu.Props{
					Label: "View",
					Items: []menu.Item{
						{Label: "Show sizes", Checked: State.ShowSizes, Action: func() {
							State.ShowSizes = !State.ShowSizes
							if State.ShowSizes {
								State.Status = "Sizes shown."
							} else {
								State.Status = "Sizes hidden."
							}
						}},
					},
				})
				ui.Spacer(c)
				open := ui.Button(c, "Open")
				if open.Clicked() {
					if State.Chosen >= 0 {
						State.Status = "Opened " + filesFor()[State.Chosen].Name + "."
					} else {
						State.Status = "Choose a file first."
					}
				}
			},
		})
		separator.Separator(c, separator.Props{})
		split.Split(c, split.Props{
			Size: &State.Divider,
			First: func(c *ui.Context) {
				// The pane fills the split, so the tree spans its width
				// and its chosen node highlights across the pane.
				ui.Column(c).Fill().Padding(8).Children(func() {
					tree.Tree(c, tree.Props{Items: treeItems(), Chosen: &State.TreeChosen}).Fill()
				})
			},
			Second: func(c *ui.Context) {
				// The pane fills the split, so the list below the status
				// line spans its height.
				ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
					listState.Selected = &State.Chosen
					list.List(c, list.Props{
						State:    &listState,
						Items:    names(),
						Row:      row,
						OnChoose: func(i int) { State.Status = "Chosen " + filesFor()[i].Name + "." },
					}).Grow(1)
					ui.Text(c, State.Status).TextColor(t.TextMuted)
				})
			},
		}).Grow(1)
	})
}

// treeItems returns the tree of the explorer: a branch per folder, with
// the files of the folder inside, and the files of the root as leaves.
func treeItems() []tree.Item {
	return []tree.Item{
		{Label: "src", Open: &State.SrcOpen, Action: func() { chooseFolder("src") }, Children: []tree.Item{
			{Label: "main.go", Action: func() { State.Status = "Opened main.go." }},
			{Label: "app.go", Action: func() { State.Status = "Opened app.go." }},
		}},
		{Label: "lib", Open: &State.LibOpen, Action: func() { chooseFolder("lib") }, Children: []tree.Item{
			{Label: "util.go", Action: func() { State.Status = "Opened util.go." }},
		}},
		{Label: "docs", Open: &State.DocsOpen, Action: func() { chooseFolder("docs") }, Children: []tree.Item{
			{Label: "guide.md", Action: func() { State.Status = "Opened guide.md." }},
			{Label: "README.md", Action: func() { State.Status = "Opened README.md." }},
		}},
		{Label: "go.mod", Action: func() { State.Status = "Opened go.mod." }},
		{Label: "Makefile", Action: func() { State.Status = "Opened Makefile." }},
	}
}

// chooseFolder shows the files of a folder and names the choice.
func chooseFolder(folder string) {
	State.Folder = folder
	State.Chosen = -1
	State.Status = "Showing " + folder + "."
}

// names returns the names of the files the list shows, which also count
// its rows.
func names() []string {
	fs := filesFor()
	out := make([]string, len(fs))
	for i, f := range fs {
		out[i] = f.Name
	}
	return out
}

// row draws one row of the list: the name of the file, and its size
// when ShowSizes is on. The chosen row shows its size in the accent
// text, which reads on the accent of the row.
func row(c *ui.Context, i int) {
	f := filesFor()[i]
	ui.Row(c).Height(32).PaddingX(12).Children(func() {
		ui.Text(c, f.Name).Grow(1)
		if State.ShowSizes {
			size := ui.Text(c, fmt.Sprintf("%d B", f.Size))
			if i != State.Chosen {
				size.TextColor(c.Theme().TextMuted)
			}
		}
	})
}
