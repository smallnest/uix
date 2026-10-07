package tokenfield

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			TokenField(c, p)
		})
	}
}

// TestRenders shows the tokens as chips and the label.
func TestRenders(t *testing.T) {
	tags := []string{"go"}
	tt := ui.NewTester(frame(Props{
		Tokens: &tags, Suggestions: []string{"rust", "python"}, Label: "Tags",
	}), 400, 200)
	for _, s := range []string{"Tags", "go"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestAdds adds tokens typed, on a comma and on Enter.
func TestAdds(t *testing.T) {
	tags := []string{"go"}
	tt := ui.NewTester(frame(Props{
		Tokens: &tags, Suggestions: []string{"rust", "python"}, Label: "Tags",
	}), 400, 240)
	if err := tt.Click("Tags"); err != nil {
		t.Fatal(err)
	}
	tt.Type("zig,")
	tt.Type("lua")
	tt.Key(0, ui.KeyEnter)
	if !slices.Equal(tags, []string{"go", "zig", "lua"}) {
		t.Fatalf("typed tokens: %q", tags)
	}
}

// TestRemoves takes out a token with its button, and the last with
// Backspace in the empty field.
func TestRemoves(t *testing.T) {
	tags := []string{"go", "zig"}
	tt := ui.NewTester(frame(Props{
		Tokens: &tags, Suggestions: []string{"rust"}, Label: "Tags",
	}), 400, 240)
	if err := tt.Click("Remove go"); err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(tags, []string{"zig"}) {
		t.Fatalf("the button left %q, want [zig]", tags)
	}
	if err := tt.Click("Tags"); err != nil {
		t.Fatal(err)
	}
	tt.Key(0, ui.KeyBackspace)
	if len(tags) != 0 {
		t.Fatalf("Backspace left %q, want none", tags)
	}
}

// TestSuggests adds a suggestion with the arrows and Enter.
func TestSuggests(t *testing.T) {
	tags := []string{"go"}
	tt := ui.NewTester(frame(Props{
		Tokens: &tags, Suggestions: []string{"typescript", "python", "rust"}, Label: "Tags",
	}), 400, 240)
	if err := tt.Click("Tags"); err != nil {
		t.Fatal(err)
	}
	tt.Type("ty")
	if !tt.HasText("typescript") {
		t.Fatalf("ty suggests %q", tt.Texts())
	}
	tt.Key(0, ui.KeyDown)
	tt.Key(0, ui.KeyEnter)
	if !slices.Equal(tags, []string{"go", "typescript"}) {
		t.Fatalf("the suggestion added %q", tags)
	}
}
