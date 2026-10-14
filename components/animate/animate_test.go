// Package animate tests the animations: the transition presets build
// the motion they promise, and the decorators drive their property from
// a Keyframes that reduce motion completes at once.
package animate

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// TestPresets builds every transition preset and checks the motion it
// promises, with the house defaults and with custom options.
func TestPresets(t *testing.T) {
	p := Panel(Options{})
	if p.Duration != 200*time.Millisecond || p.Ease == nil || p.Enter == nil {
		t.Fatalf("panel defaults %+v", p)
	}
	if p.Enter.Opacity != 0 || p.Enter.Scale != 0.95 || p.Enter.Y != 8 || p.Enter.X != 0 {
		t.Fatalf("panel motion %+v", p.Enter)
	}
	if p.Exit != nil {
		t.Fatal("Exit off by default")
	}
	if tr := Rise(Options{}); tr.Enter.Y != 8 || tr.Enter.Scale != 0 || tr.Enter.Opacity != 0 {
		t.Fatalf("rise motion %+v", tr.Enter)
	}
	if tr := Fade(Options{}); tr.Enter.Opacity != 0 || tr.Enter.Y != 0 || tr.Enter.Scale != 0 {
		t.Fatalf("fade motion %+v", tr.Enter)
	}
	for from, want := range map[Direction][2]float32{
		Left: {-8, 0}, Right: {8, 0}, Up: {0, -8}, Down: {0, 8},
	} {
		tr := FromEdge(from, Options{})
		if tr.Enter.X != want[0] || tr.Enter.Y != want[1] {
			t.Fatalf("%v slides %v, want %v", from, tr.Enter, want)
		}
	}
	p = Panel(Options{Duration: 300 * time.Millisecond, Distance: 12, Scale: 0.9, Exit: true})
	if p.Duration != 300*time.Millisecond || p.Enter.Y != 12 || p.Enter.Scale != 0.9 {
		t.Fatalf("custom panel %+v", p)
	}
	if p.Exit == nil || p.Exit.Opacity != 0 || p.Exit.Y != 12 || p.Exit.Scale != 0.9 {
		t.Fatalf("custom exit %+v", p.Exit)
	}
}

// boxFrame draws a full-fill box animated by apply, so tests sample the
// pixel the box paints.
func boxFrame(apply func(e ui.Element) ui.Element) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Children(func() {
			apply(ui.Box(c).Fill().Background(ui.RGBA(200, 60, 60, 1)))
		})
	}
}

// center returns the color of the middle pixel of tt's frame.
func center(tt *ui.Tester) (r, g, b uint32) {
	img := tt.Image()
	c := img.RGBAAt(img.Bounds().Dx()/2, img.Bounds().Dy()/2)
	return uint32(c.R), uint32(c.G), uint32(c.B)
}

// TestReduceMotionCompletes finishes each one-shot entrance at its final
// value when the desktop asks for less motion, so nothing is hidden.
func TestReduceMotionCompletes(t *testing.T) {
	for _, tc := range []struct {
		name  string
		apply func(e ui.Element) ui.Element
	}{
		{"FadeIn", func(e ui.Element) ui.Element { return FadeIn(e, Options{}) }},
		{"RiseIn", func(e ui.Element) ui.Element { return RiseIn(e, Options{}) }},
		{"ScaleIn", func(e ui.Element) ui.Element { return ScaleIn(e, Options{}) }},
		{"SlideIn", func(e ui.Element) ui.Element { return SlideIn(e, Options{}, Up) }},
	} {
		tt := ui.NewTester(boxFrame(tc.apply), 100, 100)
		tt.SetPreferences(ui.Preferences{ReduceMotion: true})
		r, g, b := center(tt)
		if r < 180 || g > 90 || b > 90 {
			t.Fatalf("%s: center %d,%d,%d, want the opaque box", tc.name, r, g, b)
		}
	}
}

// TestEntranceStartsHidden draws each one-shot entrance without reduce
// motion, whose first frame is at the animation's start: the box has
// not yet faded in, so the center shows the window, not the box.
func TestEntranceStartsHidden(t *testing.T) {
	for _, tc := range []struct {
		name  string
		apply func(e ui.Element) ui.Element
	}{
		{"FadeIn", func(e ui.Element) ui.Element { return FadeIn(e, Options{}) }},
		{"ScaleIn", func(e ui.Element) ui.Element { return ScaleIn(e, Options{}) }},
	} {
		tt := ui.NewTester(boxFrame(tc.apply), 100, 100)
		r, g, b := center(tt)
		if r < 180 {
			t.Fatalf("%s: center %d,%d,%d, want the window behind the start value", tc.name, r, g, b)
		}
	}
}

// TestEnterTransitionClick opens a Panel surface on a later frame, so its
// Enter fires, and clicks a button inside it while it is mid-entrance:
// the click lands because testing hits the drawn, transformed position.
func TestEnterTransitionClick(t *testing.T) {
	var clicks int
	open := false
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Children(func() {
			if open {
				ui.Box(c).Fill().Center().Transition(Panel(Options{})).TransformOrigin(0.5, 0.5).Children(func() {
					ui.ButtonBase(c).Label("Go").Children(func() { ui.Text(c, "Go") }).
						OnClick(func() { clicks++ })
				})
			}
		})
	}, 200, 150)
	open = true
	tt.Frame()
	if err := tt.Click("Go"); err != nil {
		t.Fatal(err)
	}
	if clicks != 1 {
		t.Fatalf("clicks %d, want 1", clicks)
	}
}

// TestDecoratorsRender draws each decorator and the loops in a layout
// without panicking, and keeps the text of the element around them.
func TestDecoratorsRender(t *testing.T) {
	for _, tc := range []struct {
		name  string
		apply func(e ui.Element) ui.Element
	}{
		{"Pulse", func(e ui.Element) ui.Element { return Pulse(e, Options{}) }},
		{"Spin", func(e ui.Element) ui.Element { return Spin(e, Options{}) }},
		{"SlideRight", func(e ui.Element) ui.Element { return SlideIn(e, Options{}, Right) }},
		{"SlideDown", func(e ui.Element) ui.Element { return SlideIn(e, Options{}, Down) }},
	} {
		tt := ui.NewTester(func(c *ui.Context) {
			ui.Column(c).Fill().Padding(16).Children(func() {
				ui.Text(c, "Before")
				tc.apply(ui.Box(c).Size(30, 30).Background(ui.RGBA(200, 60, 60, 1)))
				ui.Text(c, "After")
			})
		}, 160, 120)
		if !tt.HasText("Before") || !tt.HasText("After") {
			t.Fatalf("%s: texts %q", tc.name, tt.Texts())
		}
		if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
			t.Fatalf("%s: no frame", tc.name)
		}
	}
}
