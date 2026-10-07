package autocomplete

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Autocomplete(c, p)
		})
	}
}

// TestRenders shows the field under its label, holding its initial text.
func TestRenders(t *testing.T) {
	city := "Paris"
	tt := ui.NewTester(frame(Props{
		Value: &city, Suggestions: []string{"Paris", "Parma", "Lyon"}, Label: "City",
	}), 400, 200)
	if !tt.HasText("City") {
		t.Fatalf("missing the label in %q", tt.Texts())
	}
	if city != "Paris" {
		t.Fatalf("the field changed its text to %q", city)
	}
}

// TestSuggests shows the suggestions that contain the text typed, and
// none before anything was typed.
func TestSuggests(t *testing.T) {
	city := ""
	tt := ui.NewTester(frame(Props{
		Value: &city, Suggestions: []string{"Paris", "Parma", "Lyon", "Comparis"}, Label: "City",
	}), 400, 300)
	if err := tt.Click("City"); err != nil {
		t.Fatal(err)
	}
	if tt.HasText("Paris") {
		t.Fatal("suggestions before anything was typed")
	}
	tt.Type("par")
	for _, s := range []string{"Paris", "Parma", "Comparis"} {
		if !tt.HasText(s) {
			t.Fatalf("missing suggestion %q in %q", s, tt.Texts())
		}
	}
	if tt.HasText("Lyon") {
		t.Fatal("Lyon does not contain par")
	}
}

// TestTakes takes a suggestion with the arrows and Enter, which changes
// the value.
func TestTakes(t *testing.T) {
	city := ""
	tt := ui.NewTester(frame(Props{
		Value: &city, Suggestions: []string{"Paris", "Parma"}, Label: "City",
	}), 400, 300)
	if err := tt.Click("City"); err != nil {
		t.Fatal(err)
	}
	tt.Type("p")
	tt.Key(0, ui.KeyDown) // the first suggestion, Paris
	tt.Key(0, ui.KeyEnter)
	if city != "Paris" {
		t.Fatalf("the keys took %q, want Paris", city)
	}
}

// TestSubmits submits what was typed when Enter presses on no
// suggestion.
func TestSubmits(t *testing.T) {
	city := ""
	submitted := 0
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			e := Autocomplete(c, Props{
				Value: &city, Suggestions: []string{"Paris", "Parma"}, Label: "City",
			})
			if e.Submitted() {
				submitted++
			}
		})
	}, 400, 300)
	if err := tt.Click("City"); err != nil {
		t.Fatal(err)
	}
	tt.Type("lon")
	tt.Key(0, ui.KeyEnter)
	if city != "lon" || submitted != 1 {
		t.Fatalf("Enter submitted %q (%d), want lon (1)", city, submitted)
	}
}
