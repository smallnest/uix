package mujica

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// currentPage is the page the shared tester draws; the loop swaps it and
// asks for a frame, so one tester renders every page.
var currentPage Page

// TestEveryPageRenders draws every page of the ported catalog headless,
// as the uix gallery will. A demo that panics after the mygo v0.5.5
// migration fails here, so a broken page never reaches the gallery.
func TestEveryPageRenders(t *testing.T) {
	pages := Pages()
	currentPage = pages[0] // the tester draws its first frame at construction
	tt := ui.NewTester(func(c *ui.Context) {
		Use(c)
		ui.Scroll(c).Fill().Children(func() {
			ui.Column(c).FillWidth().MinWidth(0).Padding(16).Gap(12).Children(func() {
				if currentPage.Demo != nil {
					currentPage.Demo(c, func(string) {})
				}
			})
		})
	}, 960, 600)
	for _, p := range pages {
		currentPage = p
		func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("%s (%s): demo panicked: %v", p.Name, p.ID, r)
				}
			}()
			tt.Frame()
		}()
	}
}

// TestPageCount pins the ported catalog to MujicaUI's: 487 pages in 30
// categories, IDs "001".."500". A change upstream would fail here, not
// silently shrink the gallery.
func TestPageCount(t *testing.T) {
	pages := Pages()
	if len(pages) != 487 {
		t.Fatalf("Pages() returns %d, want 487", len(pages))
	}
	seen := map[string]bool{}
	for _, p := range pages {
		if seen[p.ID] {
			t.Errorf("duplicate ID %q", p.ID)
		}
		seen[p.ID] = true
		if p.Name == "" || p.ZhName == "" || p.Category == "" {
			t.Errorf("%s: page lacks a name, zh name or category", p.ID)
		}
	}
}
