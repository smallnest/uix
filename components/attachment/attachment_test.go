package attachment

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// view draws the attachment of props in a window.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Attachment(c, p)
		})
	}
}

// TestShows draws the name, the description and the state.
func TestShows(t *testing.T) {
	tt := ui.NewTester(view(Props{
		Title:       "report.pdf",
		Description: "1.2 MB · uploading",
		Progress:    0.6,
	}), 400, 160)
	for _, want := range []string{"report.pdf", "1.2 MB · uploading"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestMedia draws the preview instead of the generic file shape.
func TestMedia(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Attachment(c, Props{
				Title: "report.pdf",
				Media: func() {
					ui.Text(c, "SNAPSHOT")
				},
			})
		})
	}, 400, 160)
	if !tt.HasText("SNAPSHOT") {
		t.Fatalf("missing the preview in %q", tt.Texts())
	}
}

// TestRemove runs OnRemove for a click on the remove button, which is at
// the far right of the card.
func TestRemove(t *testing.T) {
	removed := false
	tt := ui.NewTester(view(Props{
		Title:    "report.pdf",
		OnRemove: func() { removed = true },
	}), 400, 160)
	tt.ClickAt(370, 40)
	if !removed {
		t.Fatal("a click on the remove button did not run OnRemove")
	}
}

// TestNoProgress hides the progress bar for a negative value.
func TestNoProgress(t *testing.T) {
	tt := ui.NewTester(view(Props{Title: "report.pdf", Progress: -1}), 400, 160)
	if !tt.HasText("report.pdf") {
		t.Fatalf("texts %q", tt.Texts())
	}
}

// TestDarkMode draws the card under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(view(Props{Title: "report.pdf"}), 400, 160)
	tt.SetDark(true)
	if !tt.HasText("report.pdf") {
		t.Fatal("attachment missing under dark mode")
	}
}
