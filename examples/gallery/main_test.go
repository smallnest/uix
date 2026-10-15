package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/gallery/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.GalleryView, view.Width, view.Height)
}

// countExact counts the texts of the frame equal to s. The library
// filter in the rail names the three libraries, so a library's text
// appears there even on a page of another library: once when only the
// filter names it, twice when the page's badge names it too.
func countExact(tt *ui.Tester, s string) int {
	n := 0
	for _, x := range tt.Texts() {
		if x == s {
			n++
		}
	}
	return n
}

// TestRenders draws the gallery headless: the shell, the default page
// (Button) with its uix badge, and the six faces of the button.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"组件画廊", "全部", "操作与导航",
		"Button", "主要", "次要", "描边", "幽灵", "危险", "禁用",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	if countExact(tt, "uix") < 2 {
		t.Fatal("the Button page lacks the uix badge beside the filter")
	}
	if countExact(tt, "onevra") != 1 {
		t.Fatal("the Button page shows a onevra badge")
	}
}

// reach filters the sidebar to the item, which is deep in the catalog
// and out of the window otherwise, and clicks it. The search box is the
// way to an item out of view, as a user would use it.
func reach(t *testing.T, tt *ui.Tester, query, item string) {
	t.Helper()
	view.State.Query = query
	tt.Frame()
	if err := tt.Click(item); err != nil {
		t.Fatal(err)
	}
}

// TestNavigate clicks a sidebar item and the page switches to it. The
// onevra pages are deep in the catalog, so the search box brings the
// item into view first.
func TestNavigate(t *testing.T) {
	tt := newTester(t)
	reach(t, tt, "pixel", "PixelGrid")
	if view.State.Selected != "pixelgrid" {
		t.Fatalf("the click chose %q, want pixelgrid", view.State.Selected)
	}
	if !tt.HasText("方格背景") {
		t.Fatalf("the page does not show the pixel grid demo: %q", tt.Texts())
	}
}

// TestLibraries shows the badge of each library on one of its pages: a
// mygo page carries the mygo badge, an onevra page the onevra badge.
func TestLibraries(t *testing.T) {
	tt := newTester(t)
	reach(t, tt, "box", "Box")
	if countExact(tt, "mygo") < 2 {
		t.Fatal("the Box page lacks the mygo badge beside the filter")
	}
	reach(t, tt, "donut", "DonutChart")
	if countExact(tt, "onevra") < 2 {
		t.Fatal("the DonutChart page lacks the onevra badge")
	}
	if countExact(tt, "uix") != 1 {
		t.Fatal("the DonutChart page shows a uix badge")
	}
}

// TestLibraryFilter clicks a library and the sidebar shows that
// library's components alone, and the page lands on one of them; the
// 全部 segment restores the mixed sidebar.
func TestLibraryFilter(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("mygo"); err != nil {
		t.Fatal(err)
	}
	if view.State.Selected != "textinputbuffer" {
		t.Fatalf("the mygo filter chose %q, want textinputbuffer", view.State.Selected)
	}
	if !tt.HasText("TextInputBuffer") || !tt.HasText("Box") {
		t.Fatalf("the mygo sidebar misses its components: %q", tt.Texts())
	}
	if tt.HasText("Button") {
		t.Fatal("the mygo filter still shows a uix component")
	}
	if err := tt.Click("uix"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Button") {
		t.Fatalf("the uix filter lost its components: %q", tt.Texts())
	}
	if tt.HasText("Box") {
		t.Fatal("the uix filter still shows a mygo component")
	}
}

// TestSearch filters the sidebar with a query, and a click on a match
// opens its page.
func TestSearch(t *testing.T) {
	tt := newTester(t)
	view.State.Query = "chart"
	tt.Frame()
	if !tt.HasText("DonutChart") {
		t.Fatalf("the sidebar does not show the charts: %q", tt.Texts())
	}
	if tt.HasText("Breadcrumbs") {
		t.Fatal("the sidebar still shows an item that does not match")
	}
	if err := tt.Click("DonutChart"); err != nil {
		t.Fatal(err)
	}
	if view.State.Selected != "donutchart" {
		t.Fatalf("the click chose %q, want donutchart", view.State.Selected)
	}
}

// TestDialog opens a surface from its page: the search box brings the
// trigger page into view, then the trigger button opens the dialog over
// the window.
func TestDialog(t *testing.T) {
	tt := newTester(t)
	reach(t, tt, "dialog", "Dialog")
	if err := tt.Click("打开对话框"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("要删除你的账户吗？") {
		t.Fatalf("the dialog did not open: %q", tt.Texts())
	}
}

// TestDarkMode draws the gallery under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("组件画廊") {
		t.Fatal("dark gallery missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
