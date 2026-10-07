package link

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Link(c, p)
		})
	}
}

// TestRenders shows the label of the link.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Read the guide"}), 240, 80)
	if !tt.HasText("Read the guide") {
		t.Fatalf("missing the label in %q", tt.Texts())
	}
}

// TestOpens opens the URL in the browser when the link is clicked, which
// the tester records.
func TestOpens(t *testing.T) {
	tt := ui.NewTester(frame(Props{Label: "Read the guide", URL: "https://example.com/guide"}), 240, 80)
	if err := tt.Click("Read the guide"); err != nil {
		t.Fatal(err)
	}
	urls := tt.OpenedURLs()
	if len(urls) != 1 || urls[0] != "https://example.com/guide" {
		t.Fatalf("the link opened %q, want the guide", urls)
	}
}
