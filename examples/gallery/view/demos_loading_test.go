package view

import (
	"testing"

	"github.com/egoist/mygo/plugins/loading"
	"github.com/egoist/mygo/ui"
)

// TestLoadingShowsEveryIndicator requires the page to name every
// indicator the plugin ships. It reads loading.All() rather than a
// copy of it, so a variant the plugin adds fails here until the page
// shows it, and a cell that the layout drops is caught too.
func TestLoadingShowsEveryIndicator(t *testing.T) {
	Reset()
	filterTo("loading")
	State.Selected = "loading"
	tt := ui.NewTester(GalleryView, Width, Height)
	tt.Frame()
	all := loading.All()
	if len(all) == 0 {
		t.Fatal("loading.All() is empty: the page would show nothing")
	}
	for _, s := range all {
		if !tt.HasText(s.Name) {
			t.Errorf("the page does not show %q", s.Name)
		}
	}
}
