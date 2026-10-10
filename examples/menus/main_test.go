package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/menus/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.MenusView, view.Width, view.Height)
}

// TestRenders draws the example headless: the heading, the three menus
// and their labels.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Menus",
		"Menu bar", "File", "Edit",
		"Navigation menu", "Docs", "About",
		"Context menu", "Right-click here",
		"The last action: ",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// open opens the menu of the File button, as the pointer goes down on
// it.
func open(t *testing.T, tt *ui.Tester, label string) {
	t.Helper()
	r, ok := tt.Find(label)
	if !ok {
		t.Fatalf("no %s button in %q", label, tt.Texts())
	}
	tt.Press(r.X+r.W/2, r.Y+r.H/2)
	tt.Release(r.X+r.W/2, r.Y+r.H/2)
}

// TestMenubar chooses an item of the File menu.
func TestMenubar(t *testing.T) {
	tt := newTester(t)
	open(t, tt, "File")
	if err := tt.ChooseMenuItem("Save"); err != nil {
		t.Fatal(err)
	}
	if view.State.Last != "saved" {
		t.Fatalf("Last = %q, want saved", view.State.Last)
	}
}

// TestMenubarChecked toggles the checked item of the Edit menu.
func TestMenubarChecked(t *testing.T) {
	tt := newTester(t)
	open(t, tt, "Edit")
	if err := tt.ChooseMenuItem("Bold"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Dark {
		t.Fatal("the checked item did not toggle")
	}
}

// TestNav opens the panel of the Docs link and chooses Guides.
func TestNav(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Docs"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Getting started") || !tt.HasText("Guides") {
		t.Fatalf("missing the panel in %q", tt.Texts())
	}
	if err := tt.Click("Guides"); err != nil {
		t.Fatal(err)
	}
	if view.State.Last != "guides" {
		t.Fatalf("Last = %q, want guides", view.State.Last)
	}
}

// TestContext chooses an item of the context menu of the box.
func TestContext(t *testing.T) {
	tt := newTester(t)
	if err := tt.RightClick("Right-click here"); err != nil {
		t.Fatal(err)
	}
	if err := tt.ChooseMenuItem("Rename"); err != nil {
		t.Fatal(err)
	}
	if view.State.Last != "renamed" {
		t.Fatalf("Last = %q, want renamed", view.State.Last)
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	for _, s := range []string{"Menus", "File", "Docs"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in the dark frame %q", s, tt.Texts())
		}
	}
}
