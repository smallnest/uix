package fieldset

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Fieldset(c, p)
		})
	}
}

// TestRenders shows the legend above the fields, which are a group.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Legend: "Type",
		Children: func(c *ui.Context) {
			ui.Text(c, "Photos")
			ui.Text(c, "Docs")
		},
	}), 300, 200)
	for _, s := range []string{"Type", "Photos", "Docs"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	legend, _ := tt.Find("Type")
	photo, _ := tt.Find("Photos")
	if legend.Y >= photo.Y {
		t.Fatal("the legend is not above the fields")
	}
}

// TestGap sets a margin above a group when another is above it, so two
// groups in a column do not touch.
func TestGap(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Fieldset(c, Props{
				Legend: "Type",
				Children: func(c *ui.Context) {
					ui.Text(c, "Photos")
				},
			})
			Fieldset(c, Props{
				Legend: "Sort",
				Children: func(c *ui.Context) {
					ui.Text(c, "Name")
				},
			})
		})
	}, 300, 200)
	photo, _ := tt.Find("Photos")
	sort, _ := tt.Find("Sort")
	if sort.Y-photo.Y < 8 {
		t.Fatalf("the groups are too close: %g", sort.Y-photo.Y)
	}
}
