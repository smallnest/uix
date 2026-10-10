package label

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// TestShows draws the label with its text.
func TestShows(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Label(c, Props{Text: "Full name"})
		})
	}, 200, 80)
	if !tt.HasText("Full name") {
		t.Fatalf("missing the text in %q", tt.Texts())
	}
}

// TestDisabled draws the label in the muted color: a disabled label
// shows the text but not in the foreground color.
func TestDisabled(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Label(c, Props{Text: "Name", Disabled: true})
		})
	}, 200, 80)
	if !tt.HasText("Name") {
		t.Fatalf("missing the text in %q", tt.Texts())
	}
}
