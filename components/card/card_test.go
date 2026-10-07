package card

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame() func(c *ui.Context) {
	return func(c *ui.Context) {
		Card(c, Props{
			Title:       "Account",
			Description: "How people reach you.",
			Footer: func() {
				ui.Button(c, "Save").Clicked()
			},
		}, func() {
			ui.Text(c, "Name")
		})
	}
}

// TestRenders shows the header, the content and the footer of a card.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(), 300, 200)
	for _, s := range []string{"Account", "How people reach you.", "Name", "Save"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
}

// TestDarkMode draws the card under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(), 300, 200)
	tt.SetDark(true)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
