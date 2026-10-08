package carousel

import (
	"fmt"
	"testing"

	"github.com/egoist/mygo/ui"
)

// cs is the harness state of the tests.
type cs struct {
	Scroll ui.ScrollState
	State  State
}

// view is the harness the tests draw the carousel in: three slides that
// each show their number, so the visible one is readable.
func view(s *cs) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			Carousel(c, Props{
				Count: 3,
				Slide: func(c *ui.Context, i int) {
					ui.Box(c).FillWidth().Height(120).Background(c.Theme().Surface).
						Radius(c.Theme().Radius).Center().Children(func() {
						ui.Text(c, fmt.Sprintf("Slide %d", i+1)).TextColor(c.Theme().Text)
					})
				},
				Scroll: &s.Scroll,
				State:  &s.State,
			})
		})
	}
}

// TestShows draws the carousel with its buttons, dots and first slide.
func TestShows(t *testing.T) {
	s := &cs{}
	tt := ui.NewTester(view(s), 360, 220)
	for _, label := range []string{"Previous slide", "Next slide", "Go to slide 1",
		"Go to slide 2", "Go to slide 3", "Slide 1"} {
		if !tt.HasText(label) {
			t.Fatalf("missing %q in %q", label, tt.Texts())
		}
	}
	if s.State.Active != 0 || !s.State.AtStart || s.State.AtEnd {
		t.Fatalf("state = %+v, want Active 0 AtStart true AtEnd false", s.State)
	}
}

// TestNext scrolls to the second slide with the Next button: the state
// moves with it and the second slide is shown.
func TestNext(t *testing.T) {
	s := &cs{}
	tt := ui.NewTester(view(s), 360, 220)
	if err := tt.Click("Next slide"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if s.State.Active != 1 {
		t.Fatalf("Active = %d, want 1", s.State.Active)
	}
	if !tt.HasText("Slide 2") {
		t.Fatalf("the second slide is missing from %q", tt.Texts())
	}
}

// TestPrev scrolls back with the Previous button.
func TestPrev(t *testing.T) {
	s := &cs{}
	tt := ui.NewTester(view(s), 360, 220)
	tt.Click("Next slide")
	tt.Click("Previous slide")
	tt.Frame()
	if s.State.Active != 0 {
		t.Fatalf("Active = %d, want 0", s.State.Active)
	}
	if !tt.HasText("Slide 1") {
		t.Fatalf("the first slide is missing from %q", tt.Texts())
	}
}

// TestDot jumps straight to a slide by its dot.
func TestDot(t *testing.T) {
	s := &cs{}
	tt := ui.NewTester(view(s), 360, 220)
	tt.Click("Go to slide 3")
	tt.Frame()
	if s.State.Active != 2 {
		t.Fatalf("Active = %d, want 2", s.State.Active)
	}
	if !tt.HasText("Slide 3") {
		t.Fatalf("the third slide is missing from %q", tt.Texts())
	}
}

// TestEnds disables the Previous button at the start and the Next at the
// end, and the state marks them.
func TestEnds(t *testing.T) {
	s := &cs{}
	tt := ui.NewTester(view(s), 360, 220)
	tt.Click("Go to slide 3")
	tt.Frame()
	if !s.State.AtEnd {
		t.Fatalf("AtEnd = false at the last slide")
	}
	if s.State.AtStart {
		t.Fatalf("AtStart = true at the last slide")
	}
	// The Next button is disabled at the end and does not move it.
	tt.Click("Next slide")
	tt.Frame()
	if s.State.Active != 2 {
		t.Fatalf("Active = %d, want 2 after a disabled Next", s.State.Active)
	}
}
