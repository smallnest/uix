package searchfield

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func searchFrame(q *string, disabled bool) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			SearchField(c, Props{Value: q, Disabled: disabled}).Label("Search mail")
		})
	}
}

// TestRenders draws the field; the query lives in an editor, so the
// check is that a frame renders.
func TestRenders(t *testing.T) {
	q := ""
	tt := ui.NewTester(searchFrame(&q, false), 260, 100)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestTypesAndClears types a query, then takes it out with the clear
// button.
func TestTypesAndClears(t *testing.T) {
	q := ""
	tt := ui.NewTester(searchFrame(&q, false), 260, 100)
	if err := tt.Click("Search mail"); err != nil {
		t.Fatal(err)
	}
	tt.Type("ada")
	if q != "ada" {
		t.Fatalf("query %q, want ada", q)
	}
	if !tt.HasText("Clear") {
		t.Fatal("clear button missing while the query holds")
	}
	if err := tt.Click("Clear"); err != nil {
		t.Fatal(err)
	}
	if q != "" {
		t.Fatalf("query %q, want empty after clear", q)
	}
	if tt.HasText("Clear") {
		t.Fatal("clear button still shows after clearing")
	}
}

// TestSubmits reports Enter as a submission.
func TestSubmits(t *testing.T) {
	q := ""
	submitted := false
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			el := SearchField(c, Props{Value: &q}).Label("Search mail")
			if el.Submitted() {
				submitted = true
			}
		})
	}, 260, 100)
	if err := tt.Click("Search mail"); err != nil {
		t.Fatal(err)
	}
	tt.Type("ada")
	tt.Key(0, ui.KeyEnter)
	if !submitted {
		t.Fatal("Enter did not submit the search")
	}
}

// TestDisabled keeps the field from being typed into.
func TestDisabled(t *testing.T) {
	q := ""
	tt := ui.NewTester(searchFrame(&q, true), 260, 100)
	if err := tt.Click("Search mail"); err != nil {
		t.Fatal(err)
	}
	tt.Type("ada")
	if q != "" {
		t.Fatalf("disabled field took %q", q)
	}
}

// TestDarkMode draws the search field under the dark appearance too.
func TestDarkMode(t *testing.T) {
	q := ""
	tt := ui.NewTester(searchFrame(&q, false), 260, 100)
	tt.SetDark(true)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
