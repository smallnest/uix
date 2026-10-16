package view

import (
	"image/color"
	"testing"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
)

// TestCatalogCompleteness makes sure every entry carries an icon and a
// demo, so the sidebar draws every item with one and every page has
// something to show.
func TestCatalogCompleteness(t *testing.T) {
	for _, s := range catalog {
		for _, e := range s.Entries {
			if e.Icon == nil {
				t.Errorf("%s (%s) has no icon", e.Name, e.ID)
			}
			if e.Build == nil {
				t.Errorf("%s (%s) has no demo", e.Name, e.ID)
			}
			if e.Name == "" || e.ID == "" {
				t.Errorf("an entry of %q lacks a name or an ID", s.Title)
			}
		}
	}
}

// TestEveryPageRenders draws every page headless, so a demo that panics
// fails the gallery. The heading and the sidebar item both carry the
// entry's name, so a page whose demo drew nothing shows it once; a page
// whose heading drew shows it twice.
func TestEveryPageRenders(t *testing.T) {
	tt := ui.NewTester(GalleryView, Width, Height)
	for _, s := range catalog {
		for _, e := range s.Entries {
			Reset()
			State.Selected = e.ID
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%s (%s) panicked: %v", e.Name, e.ID, r)
					}
				}()
				tt.Frame()
			}()
			if n := count(e.Name, tt.Texts()); n < 2 {
				t.Errorf("%s (%s) drew its name %d times, want 2: %q", e.Name, e.ID, n, tt.Texts())
			}
		}
	}
}

// count counts the texts equal to s.
func count(s string, texts []string) int {
	n := 0
	for _, x := range texts {
		if x == s {
			n++
		}
	}
	return n
}

// sameColor reports whether the pixel holds the theme color, ignoring
// alpha, which the image makes opaque.
func sameColor(p color.RGBA, c ui.Color) bool {
	return p.R == c.R && p.G == c.G && p.B == c.B
}

// TestRailStaysInItsPanel draws the whole catalog in the sidebar, in
// light and dark, and makes sure the sidebar clips at the rail's bottom.
// It once grew past the rail and drew its items below the panel; the
// wheel over the rail scrolls it instead. The test finds the panel's
// bottom edge as the deepest row of the rail column holding the surface
// color, then requires that only the window background lies below it.
func TestRailStaysInItsPanel(t *testing.T) {
	for _, dark := range []bool{false, true} {
		Reset()
		tt := ui.NewTester(GalleryView, Width, Height)
		tt.SetDark(dark)
		tt.Frame()
		img := tt.Image()
		b := img.Bounds()
		theme := tokens.For(dark)
		panelBottom := 0
		for y := 0; y < b.Max.Y; y++ {
			for x := 30; x < 240; x++ {
				if sameColor(img.RGBAAt(x, y), theme.Surface) {
					panelBottom = y
				}
			}
		}
		if panelBottom == 0 {
			t.Fatalf("dark=%v: no rail panel found", dark)
		}
		for y := panelBottom + 1; y < b.Max.Y; y++ {
			for x := 30; x < 240; x++ {
				if !sameColor(img.RGBAAt(x, y), theme.Background) {
					t.Fatalf("dark=%v: the sidebar draws at (%d, %d) below the rail bottom %d", dark, x, y, panelBottom)
				}
			}
		}
		// The pages below the fold stay reachable: the wheel scrolls the
		// sidebar, so the last item moves toward the top.
		app, _ := tt.Find("AppShell")
		tt.Move(130, 400)
		tt.Scroll(130, 400, 0, 3000)
		app2, _ := tt.Find("AppShell")
		if app2.Y >= app.Y {
			t.Fatalf("dark=%v: the wheel did not scroll the sidebar (%v -> %v)", dark, app.Y, app2.Y)
		}
	}
}

// TestOnevraPagesDrawContent opens the onevra pages whose demos build
// their own content and makes sure it shows, and that the onevra filter
// lists every onevra page, so the filter is not smaller than the library.
func TestOnevraPagesDrawContent(t *testing.T) {
	tt := ui.NewTester(GalleryView, Width, Height)
	for _, tc := range []struct{ id, want string }{
		{"timepicker", "时间"},
		{"colorview", "颜色"},
		{"tablevirtualized", "文件_1.go"},
		{"treevirtualized", "file_0001.go"},
		{"conversation", "帮我审查这段代码。"},
	} {
		Reset()
		State.Selected = tc.id
		tt.Frame()
		if !tt.HasText(tc.want) {
			t.Errorf("%s: missing %q in %q", tc.id, tc.want, tt.Texts())
		}
	}
	Reset()
	tt.Frame()
	if err := tt.Click("onevra"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"DonutChart", "TimePicker", "ColorView", "ProgressRing", "TableVirtualized", "TreeVirtualized", "Conversation"} {
		if !tt.HasText(s) {
			t.Errorf("the onevra filter misses %q: %q", s, tt.Texts())
		}
	}
}

// TestDemoContentShows opens the demos whose pages build rich content
// and makes sure the content that matters shows. The render-every-page
// test proves a page draws its name; this proves the enrichment drew
// what a visitor should see, from the markdown document to the table
// rows to the pagination readout.
func TestDemoContentShows(t *testing.T) {
	tt := ui.NewTester(GalleryView, Width, Height)
	for _, tc := range []struct{ id, want string }{
		{"markdown", "复制代码"},
		{"markdown", "已完成的任务"},
		{"markdown", "分隔线之后的文字"},
		{"code", "def greet"},
		{"richtext", "波浪线"},
		{"chat", "它有哪些输入组件?"},
		{"message", "系统消息使用次要样式。"},
		{"bubble", "最后一条。"},
		{"table", "Ada Lovelace"},
		{"tree", "catalog.go"},
		{"empty", "清除搜索"},
		{"card", "查看更多"},
		{"slider", "步进"},
		{"rating", "只读的评分"},
		{"pagination", "当前第 3 页"},
		{"tabs", "概览页的内容。"},
		{"menu", "分隔线下有禁用的项"},
		{"link", "onevra"},
		{"grid", "格 1"},
		{"outline", "go.mod"},
		{"toast", "错误"},
	} {
		Reset()
		State.Selected = tc.id
		tt.Frame()
		if !tt.HasText(tc.want) {
			t.Errorf("%s: missing %q in %q", tc.id, tc.want, tt.Texts())
		}
	}
}

// TestStartupStateRenders draws every page from the state of a window
// that forgot Reset: zero choices, only the pointers the package inits
// (the buffer and the history). Clicking mygo jumped to TextInputBuffer
// in that state once and the page panicked on the nil buffer; this
// keeps the zero state safe.
func TestStartupStateRenders(t *testing.T) {
	State = GalleryState{}
	Demo = DemoState{Buffer: ui.NewTextBuffer(""), History: ui.NewRouter("/")}
	tt := ui.NewTester(GalleryView, Width, Height)
	for _, s := range catalog {
		for _, e := range s.Entries {
			State.Selected = e.ID
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Errorf("%s (%s) panicked at startup: %v", e.Name, e.ID, r)
					}
				}()
				tt.Frame()
			}()
		}
	}
}
