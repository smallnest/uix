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

// TestRenders draws the gallery headless: the title, the filters, the
// photos of the initial view, and the status line. The docs and the
// videos do not show, because the type filter shows photos only.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Gallery", "8 items", "Type", "Sort",
		"Photos", "Docs", "Videos", "Name", "Date",
		"Alpine.jpg", "Beach.png", "City.jpg", "Forest.png",
		"Meadow.jpg", "River.jpg", "Snow.png", "Sunset.jpg",
		"Choose an item, or change the type or the sort.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	for _, s := range []string{"Budget.pdf", "Clip.mp4"} {
		if tt.HasText(s) {
			t.Fatalf("the type filter shows %q, which is no photo", s)
		}
	}
}

// TestFilters switches the type filter to the docs, which the grid shows
// and the count follows.
func TestFilters(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if view.State.Kind != 1 {
		t.Fatal("the Docs radio did not switch the type")
	}
	if view.State.Chosen != -1 {
		t.Fatal("the filter did not drop the choice")
	}
	if !tt.HasText("4 items") || !tt.HasText("Budget.pdf") {
		t.Fatalf("the grid does not show the docs: %q", tt.Texts())
	}
	if tt.HasText("Alpine.jpg") {
		t.Fatal("the grid still shows a photo")
	}
	if !tt.HasText("Showing Docs by name.") {
		t.Fatal("the status does not name the filter")
	}
}

// TestSorts switches the sort to the date: the earliest photo, City,
// moves to the first column, before the alphabetically first, Alpine.
func TestSorts(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Date"); err != nil {
		t.Fatal(err)
	}
	if view.State.Sort != 1 {
		t.Fatal("the Date radio did not switch the sort")
	}
	city, _ := tt.Find("City.jpg")
	alpine, _ := tt.Find("Alpine.jpg")
	if city.X >= alpine.X {
		t.Fatalf("the date sort did not move City (%g) before Alpine (%g)", city.X, alpine.X)
	}
}

// TestChooses chooses a file with a click, which the status names and
// the state records.
func TestChooses(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Beach.png"); err != nil {
		t.Fatal(err)
	}
	if view.State.Chosen != 1 {
		t.Fatalf("the click chose %d, want 1 (Beach.png)", view.State.Chosen)
	}
	if !tt.HasText("Chosen Beach.png.") {
		t.Fatalf("the status does not name the choice: %q", tt.Texts())
	}
}

// TestKeys chooses a file with the arrows: a Tab focuses the grid, and
// the Right key picks the first photo.
func TestKeys(t *testing.T) {
	tt := newTester(t)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyTab)
	tt.Key(0, ui.KeyRight)
	if view.State.Chosen != 0 {
		t.Fatalf("the arrows chose %d, want 0", view.State.Chosen)
	}
}

// TestSecondWindow draws the gallery again after a choice, as a second
// window would; the reset must give the grid a fresh state, or the old
// frame would not build again.
func TestSecondWindow(t *testing.T) {
	tt1 := newTester(t)
	if err := tt1.Click("Beach.png"); err != nil {
		t.Fatal(err)
	}
	tt2 := newTester(t)
	if !tt2.HasText("Alpine.jpg") || !tt2.HasText("8 items") {
		t.Fatalf("the second window does not show the gallery: %q", tt2.Texts())
	}
	if tt2.HasText("Chosen Beach.png.") {
		t.Fatal("the second window kept the status of the first")
	}
}

// TestDarkMode draws the gallery under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("Gallery") {
		t.Fatal("dark gallery missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
