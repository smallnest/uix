package richtext

import (
	"bytes"
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			RichText(c, p)
		})
	}
}

// TestRenders shows the text of the spans, joined.
func TestRenders(t *testing.T) {
	tt := ui.NewTester(frame(Props{Spans: []ui.Span{
		{Text: "Report "},
		{Text: "draft", Weight: 600},
		{Text: " — final"},
	}}), 400, 100)
	for _, s := range []string{"Report draft — final", "draft"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestHighlight draws the same text with and without a highlighted span,
// which differ, so the styles of the spans reach the render.
func TestHighlight(t *testing.T) {
	plain := ui.NewTester(frame(Props{Spans: []ui.Span{{Text: "Find this word"}}}), 400, 100)
	marked := ui.NewTester(frame(Props{Spans: []ui.Span{
		{Text: "Find "},
		{Text: "this", Background: ui.RGB(255, 235, 59)},
		{Text: " word"},
	}}), 400, 100)
	if bytes.Equal(plain.Image().Pix, marked.Image().Pix) {
		t.Fatal("the span styles do not reach the render")
	}
}
