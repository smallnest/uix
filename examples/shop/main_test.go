package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/shop/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() {
	view.State.Query, view.State.Category = "", ""
	view.State.Limit, view.State.MinRate = 0, 0
}

// TestRenders draws the catalog headless: every filter control and every
// product of it, with the count line.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	for _, s := range []string{
		"Shop", "Search", "Category", "Limit", "Min rate",
		"Wrench", "Hammer", "Drill", "Notepad", "Desk lamp", "Coffee mug", "Kettle",
		"$12", "$45", "7 products",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestSearches types a query into the search box, which leaves the
// matching products only.
func TestSearches(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	if err := tt.Click("Search"); err != nil {
		t.Fatal(err)
	}
	tt.Type("drill")
	if view.State.Query != "drill" {
		t.Fatalf("query %q, want drill", view.State.Query)
	}
	if !tt.HasText("1 products") || !tt.HasText("Drill") {
		t.Fatal("the search did not narrow the catalog to the drill")
	}
	if tt.HasText("Wrench") {
		t.Fatal("the search left a product that does not match")
	}
}

// TestCategory picks a category from the box, which leaves the products
// of that category only. A click on the field's label focuses it, and
// Down opens its popup.
func TestCategory(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	if err := tt.Click("Category"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyDown)
	if err := tt.Click("Tools"); err != nil {
		t.Fatal(err)
	}
	if view.State.Category != "Tools" {
		t.Fatalf("category %q, want Tools", view.State.Category)
	}
	if !tt.HasText("3 products") || !tt.HasText("Wrench") {
		t.Fatal("the category did not narrow the catalog to the tools")
	}
	if tt.HasText("Kettle") {
		t.Fatal("the category left a product of another category")
	}
}

// TestRating raises the minimum rating with the arrows, which leaves the
// products rated at least three. A click on the field's label focuses
// the row.
func TestRating(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	if err := tt.Click("Min rate"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyRight)
	tt.Key(0, ui.KeyRight)
	tt.Key(0, ui.KeyRight)
	if view.State.MinRate != 3 {
		t.Fatalf("minimum rating %d, want 3", view.State.MinRate)
	}
	if !tt.HasText("6 products") || !tt.HasText("Wrench") {
		t.Fatal("the rating did not narrow the catalog to the rated products")
	}
	if tt.HasText("Desk lamp") {
		t.Fatal("the rating left a product rated below three")
	}
}

// TestLimit raises the limit with the plus button, which shows the first
// products only.
func TestLimit(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	if err := tt.Click("+"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("+"); err != nil {
		t.Fatal(err)
	}
	if view.State.Limit != 2 {
		t.Fatalf("limit %g, want 2", view.State.Limit)
	}
	if !tt.HasText("2 products") || !tt.HasText("Wrench") {
		t.Fatal("the limit did not narrow the catalog to the first products")
	}
	if tt.HasText("Drill") {
		t.Fatal("the limit left a product past it")
	}
}

// TestDarkMode draws the catalog under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ShopView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Shop") {
		t.Fatal("dark catalog missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
