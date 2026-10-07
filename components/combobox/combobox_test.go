package combobox

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(city *string, p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Combobox(c, Props{Value: city, Options: p.Options, Disabled: p.Disabled}).Label("City")
		})
	}
}

// open clicks the field's text input, which opens the popup.
func open(t *testing.T, tt *ui.Tester) {
	t.Helper()
	r, ok := tt.Find("City")
	if !ok {
		t.Fatal("field not found")
	}
	tt.ClickAt(r.X+30, r.Y+r.H/2)
}

func cityProps(options []string) (string, Props) {
	var city string
	return city, Props{Value: &city, Options: options}
}

// TestOpens shows every option when the field is clicked.
func TestOpens(t *testing.T) {
	city, p := cityProps([]string{"Beijing", "Shanghai", "Chengdu"})
	tt := ui.NewTester(frame(&city, p), 300, 160)
	open(t, tt)
	for _, o := range []string{"Beijing", "Shanghai", "Chengdu"} {
		if !tt.HasText(o) {
			t.Fatalf("missing %q", o)
		}
	}
}

// TestFilters keeps the options containing what is typed.
func TestFilters(t *testing.T) {
	city, p := cityProps([]string{"Beijing", "Shanghai", "Chengdu"})
	tt := ui.NewTester(frame(&city, p), 300, 160)
	open(t, tt)
	tt.Type("be")
	if !tt.HasText("Beijing") {
		t.Fatal("typing filters to the matching option")
	}
	for _, o := range []string{"Shanghai", "Chengdu"} {
		if tt.HasText(o) {
			t.Fatalf("typing did not filter out %q", o)
		}
	}
}

// TestChooses picks an option with a click, and closes the popup.
func TestChooses(t *testing.T) {
	city, p := cityProps([]string{"Beijing", "Shanghai", "Chengdu"})
	tt := ui.NewTester(frame(&city, p), 300, 160)
	open(t, tt)
	if err := tt.Click("Shanghai"); err != nil {
		t.Fatal(err)
	}
	if city != "Shanghai" {
		t.Fatalf("chose %q, want Shanghai", city)
	}
	if tt.HasText("Beijing") {
		t.Fatal("the popup still shows after the choice")
	}
}

// TestKeyboard moves with the arrows and chooses with Enter.
func TestKeyboard(t *testing.T) {
	city, p := cityProps([]string{"Beijing", "Shanghai", "Chengdu"})
	tt := ui.NewTester(frame(&city, p), 300, 160)
	open(t, tt)
	tt.Key(0, ui.KeyDown)
	tt.Key(0, ui.KeyDown)
	tt.Key(0, ui.KeyEnter)
	if city != "Shanghai" {
		t.Fatalf("chose %q, want Shanghai", city)
	}
}

// TestDisabled opens no popup, and types nothing.
func TestDisabled(t *testing.T) {
	city, p := cityProps([]string{"Beijing", "Shanghai", "Chengdu"})
	p.Disabled = true
	tt := ui.NewTester(frame(&city, p), 300, 160)
	open(t, tt)
	if tt.HasText("Beijing") {
		t.Fatal("a disabled field opens its popup")
	}
}

// TestDarkMode draws the combobox and its popup under the dark
// appearance too.
func TestDarkMode(t *testing.T) {
	city, p := cityProps([]string{"Beijing", "Shanghai", "Chengdu"})
	tt := ui.NewTester(frame(&city, p), 300, 160)
	tt.SetDark(true)
	open(t, tt)
	if !tt.HasText("Beijing") {
		t.Fatal("popup missing under dark mode")
	}
}
