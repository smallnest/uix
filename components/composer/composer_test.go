package composer

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func str(s string) *string { return &s }

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Composer(c, p)
		})
	}
}

// TestRenders shows the provider and the message count, and the pill
// draws a frame (the placeholder paints inside it, not as a label).
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Value:        str(""),
		Provider:     "Anthropic",
		MessageCount: 2,
	}), 420, 120)
	for _, s := range []string{"Anthropic", "2 messages"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestNewChat reads "New chat" for an empty thread.
func TestNewChat(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Value:    str(""),
		Provider: "Anthropic",
	}), 420, 120)
	if !tt.HasText("New chat") {
		t.Fatalf("missing New chat in %q", tt.Texts())
	}
}

// TestShortensModel shows the name after the last slash in the chip.
func TestShortensModel(t *testing.T) {
	tt := ui.NewTester(frame(Props{
		Value:    str(""),
		Model:    "openai/gpt-5-nano",
		Provider: "Anthropic",
	}), 420, 120)
	if !tt.HasText("gpt-5-nano") {
		t.Fatalf("missing the short model in %q", tt.Texts())
	}
	if tt.HasText("openai") {
		t.Fatalf("the chip shows the provider prefix in %q", tt.Texts())
	}
}

// TestSubmits runs OnSubmit with the typed text and clears the field.
func TestSubmits(t *testing.T) {
	value := ""
	var sent string
	tt := ui.NewTester(frame(Props{
		Value:    &value,
		OnSubmit: func(s string) { sent = s },
		Provider: "Anthropic",
	}), 420, 120)
	if err := tt.Click("Message"); err != nil {
		t.Fatal(err)
	}
	tt.Type("hello")
	tt.Key(0, ui.KeyEnter)
	if sent != "hello" {
		t.Fatalf("Enter sent %q, want hello", sent)
	}
	if value != "" {
		t.Fatalf("the field still holds %q after the send", value)
	}
}

// TestSendButton submits with a click, as the button is there for.
func TestSendButton(t *testing.T) {
	value := "click me"
	sent := 0
	tt := ui.NewTester(frame(Props{
		Value:    &value,
		OnSubmit: func(s string) { sent++ },
		Provider: "Anthropic",
	}), 420, 120)
	if err := tt.Click("Send message"); err != nil {
		t.Fatal(err)
	}
	if sent != 1 {
		t.Fatalf("the send button sent %d times, want 1", sent)
	}
}

// TestStop stops the run with a click while busy.
func TestStop(t *testing.T) {
	value := ""
	stopped := 0
	tt := ui.NewTester(frame(Props{
		Value:    &value,
		Busy:     true,
		OnStop:   func() { stopped++ },
		Provider: "Anthropic",
	}), 420, 120)
	if err := tt.Click("Stop generating"); err != nil {
		t.Fatal(err)
	}
	if stopped != 1 {
		t.Fatalf("the stop button stopped %d times, want 1", stopped)
	}
}

// TestBusyDraws renders the lit pill without a panic.
func TestBusyDraws(t *testing.T) {
	value := "thinking"
	tt := ui.NewTester(frame(Props{
		Value:    &value,
		Busy:     true,
		Model:    "openai/gpt-5-nano",
		Provider: "Anthropic",
	}), 420, 120)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
