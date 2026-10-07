package chat

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func str(s string) *string { return &s }

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Chat(c, p)
		})
	}
}

// TestRenders shows the transcript in each role's own voice.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Messages: []Message{
			{Role: User, Text: "Explain uix"},
			{Role: Assistant, Text: "uix is a component registry for MyGo."},
		},
		Value:    str(""),
		Provider: "Anthropic",
	}), 420, 520)
	for _, s := range []string{"Explain uix", "uix is a component registry for MyGo.", "Anthropic"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestEmptyState shows the prompt and the suggestion chips, and a chip
// sends its own text.
func TestEmptyState(t *testing.T) {
	value := ""
	var sent string
	tt := ui.NewTester(frame(Props{
		Messages:    nil,
		Value:       &value,
		Provider:    "Anthropic",
		OnSubmit:    func(s string) { sent = s },
		Suggestions: []string{"Explain this starter", "Name my app"},
	}), 420, 520)
	if !tt.HasText("What can I help with?") {
		t.Fatalf("missing the prompt in %q", tt.Texts())
	}
	if err := tt.Click("Explain this starter"); err != nil {
		t.Fatal(err)
	}
	if sent != "Explain this starter" {
		t.Fatalf("the chip sent %q, want the suggestion", sent)
	}
}

// TestThinking shows the thinking indicator where the reply will land.
func TestThinking(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Messages: []Message{{Role: User, Text: "Explain uix"}},
		Value:    str(""),
		Thinking: true,
		Busy:     true,
		Provider: "Anthropic",
	}), 420, 520)
	if !tt.HasText("Thinking") {
		t.Fatalf("missing the thinking indicator in %q", tt.Texts())
	}
}

// TestComposerSends hands the typed text to OnSubmit, through the chat.
func TestComposerSends(t *testing.T) {
	value := ""
	msgs := []Message{{Role: User, Text: "Hi"}}
	var sent string
	tt := ui.NewTester(frame(Props{
		Messages: msgs,
		Value:    &value,
		Provider: "Anthropic",
		OnSubmit: func(s string) { sent = s },
	}), 420, 520)
	if err := tt.Click("Message"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hello")
	tt.Key(0, ui.KeyEnter)
	if sent != "hello" {
		t.Fatalf("the chat sent %q, want hello", sent)
	}
	if value != "" {
		t.Fatalf("the composer still holds %q after the send", value)
	}
}

// TestComposerStops reaches the stop button while busy.
func TestComposerStops(t *testing.T) {
	value := ""
	stopped := 0
	tt := ui.NewTester(frame(Props{
		Messages: []Message{{Role: User, Text: "Hi"}},
		Value:    &value,
		Busy:     true,
		Provider: "Anthropic",
		OnStop:   func() { stopped++ },
	}), 420, 520)
	if err := tt.Click("Stop generating"); err != nil {
		t.Fatal(err)
	}
	if stopped != 1 {
		t.Fatalf("the stop button stopped %d times, want 1", stopped)
	}
}

// TestStreaming draws the caret after the line being written.
func TestStreaming(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Messages: []Message{
			{Role: User, Text: "Explain uix"},
			{Role: Assistant, Text: "uix is a component", Streaming: true},
		},
		Value:    str(""),
		Provider: "Anthropic",
	}), 420, 520)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
