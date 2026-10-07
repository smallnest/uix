package accordion

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// text returns a Build that draws one line of text, to show the content
// of a section is there or gone.
func text(s string) func(c *ui.Context) {
	return func(c *ui.Context) { ui.Text(c, s) }
}

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Accordion(c, p)
		})
	}
}

// TestRenders draws the headers, and the content of the sections open.
func TestRenders(t *testing.T) {
	general, privacy := true, false
	tt := ui.NewTester(frame(Props{Items: []Item{
		{Title: "General", Open: &general, Build: text("Save")},
		{Title: "Privacy", Open: &privacy, Build: text("Tracking")},
	}}), 300, 180)
	if !tt.HasText("General") || !tt.HasText("Privacy") {
		t.Fatalf("texts %q", tt.Texts())
	}
	if !tt.HasText("Save") || tt.HasText("Tracking") {
		t.Fatal("the sections show the wrong content")
	}
}

// TestOpens shows the content of a section when its header is clicked.
func TestOpens(t *testing.T) {
	general, privacy := true, false
	tt := ui.NewTester(frame(Props{Items: []Item{
		{Title: "General", Open: &general, Build: text("Save")},
		{Title: "Privacy", Open: &privacy, Build: text("Tracking")},
	}}), 300, 180)
	if err := tt.Click("Privacy"); err != nil {
		t.Fatal(err)
	}
	if !privacy || !tt.HasText("Tracking") {
		t.Fatal("a click did not open the section")
	}
}

// TestCloses hides the content of a section when its header is clicked
// again.
func TestCloses(t *testing.T) {
	general, privacy := true, false
	tt := ui.NewTester(frame(Props{Items: []Item{
		{Title: "General", Open: &general, Build: text("Save")},
		{Title: "Privacy", Open: &privacy, Build: text("Tracking")},
	}}), 300, 180)
	tt.SetPreferences(ui.Preferences{ReduceMotion: true, TextScale: 1})
	if err := tt.Click("General"); err != nil {
		t.Fatal(err)
	}
	if general || tt.HasText("Save") {
		t.Fatal("a click did not close the section")
	}
}

// TestKeyboard moves between the headers with the arrows, and opens one
// with Enter.
func TestKeyboard(t *testing.T) {
	general, privacy := true, false
	tt := ui.NewTester(frame(Props{Items: []Item{
		{Title: "General", Open: &general, Build: text("Save")},
		{Title: "Privacy", Open: &privacy, Build: text("Tracking")},
	}}), 300, 180)
	tt.Click("General")
	tt.Key(0, ui.KeyDown)
	if !tt.Focused("Privacy") {
		t.Fatal("Down did not move to the next header")
	}
	tt.Key(0, ui.KeyEnter)
	if !privacy || !tt.HasText("Tracking") {
		t.Fatal("Enter did not open the section")
	}
}

// TestDisabled keeps the sections from opening and closing.
func TestDisabled(t *testing.T) {
	general, privacy := true, false
	tt := ui.NewTester(frame(Props{Items: []Item{
		{Title: "General", Open: &general, Build: text("Save")},
		{Title: "Privacy", Open: &privacy, Build: text("Tracking")},
	}, Disabled: true}), 300, 180)
	if err := tt.Click("Privacy"); err != nil {
		t.Fatal(err)
	}
	if privacy {
		t.Fatal("a disabled accordion opened a section")
	}
}

// TestDarkMode draws the accordion under the dark appearance too.
func TestDarkMode(t *testing.T) {
	general, privacy := true, false
	tt := ui.NewTester(frame(Props{Items: []Item{
		{Title: "General", Open: &general, Build: text("Save")},
		{Title: "Privacy", Open: &privacy, Build: text("Tracking")},
	}}), 300, 180)
	tt.SetDark(true)
	if !tt.HasText("General") {
		t.Fatal("accordion missing under dark mode")
	}
}
