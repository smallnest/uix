package inputgroup

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// dollar is the leading icon of a currency field.
var dollar = ui.MustParseSVG([]byte(`<svg viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M12 2v20"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>`))

// view draws the field in a padded column, with the app's text.
func view(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			InputGroup(c, p)
		})
	}
}

// TestShows draws the field with its adornments: the trailing "USD"
// reads, and the field itself draws, surface over the window background.
func TestShows(t *testing.T) {
	value := ""
	tt := ui.NewTester(view(Props{
		Value:        &value,
		Placeholder:  "0.00",
		Leading:      dollar,
		TrailingText: "USD",
	}), 320, 90)
	if !tt.HasText("USD") {
		t.Fatalf("missing the trailing text in %q", tt.Texts())
	}
}

// TestTypes edits the text in the field, which a click focuses.
func TestTypes(t *testing.T) {
	value := ""
	tt := ui.NewTester(view(Props{Value: &value, Placeholder: "Name"}), 320, 90)
	tt.ClickAt(60, 32)
	tt.Type("Alice")
	if value != "Alice" {
		t.Fatalf("Value = %q, want Alice", value)
	}
}

// TestDisabled keeps the field from being edited.
func TestDisabled(t *testing.T) {
	value := ""
	tt := ui.NewTester(view(Props{Value: &value, Placeholder: "Name", Disabled: true}), 320, 90)
	tt.ClickAt(60, 32)
	tt.Type("Alice")
	if value != "" {
		t.Fatalf("a disabled field edited to %q", value)
	}
}
