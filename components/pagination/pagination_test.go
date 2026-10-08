package pagination

import (
	"reflect"
	"testing"

	"github.com/egoist/mygo/ui"
)

// pg is the harness state of the tests: the page, the total, and the
// pages the OnChange callback reported.
type pg struct {
	page    int
	pages   int
	changed []int
}

// view is the harness the tests draw the pagination in, with its state.
func view(s *pg) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Pagination(c, Props{
				Page:       s.page,
				TotalPages: s.pages,
				OnChange:   func(p int) { s.changed = append(s.changed, p) },
			})
		})
	}
}

// TestShows draws the pagination for five pages with the end buttons and
// every page cell.
func TestShows(t *testing.T) {
	s := &pg{page: 3, pages: 5}
	tt := ui.NewTester(view(s), 420, 80)
	for _, label := range []string{"Previous", "Next", "Go to page 1",
		"Go to page 2", "Go to page 3", "Go to page 4", "Go to page 5"} {
		if !tt.HasText(label) {
			t.Fatalf("missing %q in %q", label, tt.Texts())
		}
	}
}

// TestHidden draws nothing for a single page, as BoardUI hides the
// pagination then.
func TestHidden(t *testing.T) {
	s := &pg{page: 1, pages: 1}
	tt := ui.NewTester(view(s), 420, 80)
	if tt.HasText("Previous") || tt.HasText("Next") {
		t.Fatalf("pagination drawn for one page: %q", tt.Texts())
	}
}

// TestNext reports the page after the current one.
func TestNext(t *testing.T) {
	s := &pg{page: 3, pages: 5}
	tt := ui.NewTester(view(s), 420, 80)
	if err := tt.Click("Next"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.changed, []int{4}) {
		t.Fatalf("changed = %v, want [4]", s.changed)
	}
}

// TestPrev reports the page before the current one.
func TestPrev(t *testing.T) {
	s := &pg{page: 3, pages: 5}
	tt := ui.NewTester(view(s), 420, 80)
	if err := tt.Click("Previous"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.changed, []int{2}) {
		t.Fatalf("changed = %v, want [2]", s.changed)
	}
}

// TestEnds ignores the Previous button on the first page and the Next on
// the last, so the range never leaves it.
func TestEnds(t *testing.T) {
	s := &pg{page: 1, pages: 5}
	tt := ui.NewTester(view(s), 420, 80)
	tt.Click("Previous")
	if len(s.changed) != 0 {
		t.Fatalf("Previous reported on the first page: %v", s.changed)
	}
	s.page = 5
	tt = ui.NewTester(view(s), 420, 80)
	tt.Click("Next")
	if len(s.changed) != 0 {
		t.Fatalf("Next reported on the last page: %v", s.changed)
	}
}

// TestCell jumps to a page through its cell, from a range with dots.
func TestCell(t *testing.T) {
	s := &pg{page: 7, pages: 10}
	tt := ui.NewTester(view(s), 560, 80)
	if err := tt.Click("Go to page 9"); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s.changed, []int{9}) {
		t.Fatalf("changed = %v, want [9]", s.changed)
	}
}

// TestLabels renames the end buttons.
func TestLabels(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Pagination(c, Props{
				Page: 2, TotalPages: 3,
				PreviousLabel: "Back", NextLabel: "Forward",
				OnChange: func(int) {},
			})
		})
	}, 420, 80)
	for _, label := range []string{"Back", "Forward"} {
		if !tt.HasText(label) {
			t.Fatalf("missing %q in %q", label, tt.Texts())
		}
	}
}

// TestRange keeps the window of pages around the current one and puts
// dots where the rest collapsed, as BoardUI collapses them.
func TestRange(t *testing.T) {
	for _, tc := range []struct {
		current, total, sibling int
		want                    []int
	}{
		{3, 5, 1, []int{1, 2, 3, 4, 5}},               // every page fits
		{5, 10, 1, []int{1, DOTS, 4, 5, 6, DOTS, 10}}, // dots both sides
		{2, 10, 1, []int{1, 2, 3, 4, 5, DOTS, 10}},    // dots on the right
		{9, 10, 1, []int{1, DOTS, 6, 7, 8, 9, 10}},    // dots on the left
		{5, 20, 2, []int{1, DOTS, 3, 4, 5, 6, 7, DOTS, 20}},
	} {
		got := pageRange(tc.current, tc.total, tc.sibling)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("pageRange(%d, %d, %d) = %v, want %v",
				tc.current, tc.total, tc.sibling, got, tc.want)
		}
	}
}

// TestPageRangeFmt checks the dots render as an ellipsis, so the test
// above and the component agree on what DOTS means.
func TestPageRangeFmt(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Pagination(c, Props{
				Page: 5, TotalPages: 10,
				OnChange: func(int) {},
			})
		})
	}, 560, 80)
	if !tt.HasText("…") {
		t.Fatalf("missing the ellipsis in %q", tt.Texts())
	}
	// Page 2 is a dot in this range, so it must not be a button.
	if tt.HasText("Go to page 2") {
		t.Fatalf("the dot page is a button in %q", tt.Texts())
	}
}
