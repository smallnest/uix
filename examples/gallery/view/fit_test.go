package view

import (
	"testing"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
)

// TestLibraryFilterFitsInRail requires the last segment of the library
// filter to stay inside the rail. Five segments once drew past it; the
// gallery dropped the "全部" segment so the four libraries fit, and this
// keeps the filter from growing back out.
func TestLibraryFilterFitsInRail(t *testing.T) {
	Reset()
	tt := ui.NewTester(GalleryView, Width, Height)
	tt.Frame()
	last, ok := tt.Find("mujica")
	if !ok {
		t.Fatalf("no mujica segment drawn")
	}
	// The rail is 244 wide and begins at the window's padding.
	theme := tokens.For(false)
	railRight := theme.Space(3) + 244
	if last.X+last.W > railRight {
		t.Errorf("the mujica segment draws past the rail: right=%v, rail right=%v", last.X+last.W, railRight)
	}
}
