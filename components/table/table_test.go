package table

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(s *ui.ListState, cols []ui.TableColumn, names []string) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Table(c, Props{State: s, Columns: cols, Rows: len(names), Cell: func(row, col int) {
				if col == 0 {
					ui.Text(c, names[row]).SingleLine()
				} else {
					ui.Text(c, "1 KB")
				}
			}}).Grow(1)
		})
	}
}

// TestRenders shows the headers and the rows.
func TestRenders(t *testing.T) {
	s := ui.ListState{}
	cols := []ui.TableColumn{{Title: "Name"}, {Title: "Size", Width: 80}}
	tt := ui.NewTester(frame(&s, cols, []string{"go.mod", "main.go"}), 400, 240)
	for _, x := range []string{"Name", "Size", "go.mod", "main.go"} {
		if !tt.HasText(x) {
			t.Fatalf("missing %q", x)
		}
	}
}

// TestSorts sorts the table with a click on a header: the first click
// chooses the column, the second reverses the order.
func TestSorts(t *testing.T) {
	sort := ui.SortOrder{Column: "Size"}
	sel := -1
	s := ui.ListState{Sort: &sort, Selected: &sel}
	cols := []ui.TableColumn{{Title: "Name", Sortable: true}, {Title: "Size", Width: 80, Sortable: true}}
	tt := ui.NewTester(frame(&s, cols, []string{"a.txt", "b.txt"}), 400, 240)
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	if sort != (ui.SortOrder{Column: "Name"}) {
		t.Fatalf("first click sorted by %+v, want Name", sort)
	}
	if err := tt.Click("Name"); err != nil {
		t.Fatal(err)
	}
	if sort.Column != "Name" || !sort.Descending {
		t.Fatalf("second click gave %+v, want Name descending", sort)
	}
}

// TestSelects chooses the row clicked, which shows in the accent color.
func TestSelects(t *testing.T) {
	sel := -1
	s := ui.ListState{Selected: &sel}
	tt := ui.NewTester(frame(&s, []ui.TableColumn{{Title: "Name"}}, []string{"Ada", "Grace"}), 400, 240)
	if err := tt.Click("Grace"); err != nil {
		t.Fatal(err)
	}
	if sel != 1 {
		t.Fatalf("chose %d, want 1", sel)
	}
}

// TestSubmits opens the row chosen with Enter, as a double click does.
func TestSubmits(t *testing.T) {
	sel, opened := -1, -1
	s := ui.ListState{Selected: &sel}
	names := []string{"Ada", "Grace"}
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			tbl := Table(c, Props{State: &s, Columns: []ui.TableColumn{{Title: "Name"}}, Rows: len(names), Cell: func(row, col int) {
				ui.Text(c, names[row]).SingleLine()
			}})
			tbl.Grow(1)
			if tbl.Submitted() {
				opened = sel
			}
		})
	}, 400, 240)
	if err := tt.Click("Grace"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyEnter)
	if opened != 1 {
		t.Fatalf("Enter opened %d, want 1", opened)
	}
}

// TestDarkMode draws the table under the dark appearance too.
func TestDarkMode(t *testing.T) {
	s := ui.ListState{}
	tt := ui.NewTester(frame(&s, []ui.TableColumn{{Title: "Name"}}, []string{"Ada"}), 400, 240)
	tt.SetDark(true)
	if !tt.HasText("Name") {
		t.Fatal("dark table missing the header")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
