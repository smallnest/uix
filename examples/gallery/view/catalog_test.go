package view

import (
	"testing"

	"github.com/egoist/mygo/ui"
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

// TestStartupStateRenders draws every page from the state of a window
// that forgot Reset: zero choices, only the buffer the package inits.
// Clicking mygo jumped to TextInputBuffer in that state once and the
// page panicked on the nil buffer; this keeps the zero state safe.
func TestStartupStateRenders(t *testing.T) {
	State = GalleryState{}
	Demo = DemoState{Buffer: ui.NewTextBuffer("")}
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
