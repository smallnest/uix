// Package animate gives any element a motion of its own, over the
// animation engine of MyGo: named entrances (fade, rise, slide, scale),
// attention loops (pulse, spin), and transition presets for surfaces
// that appear and go. It is the one uix package that draws nothing: a
// decorator returns the element it animated, and a preset returns an
// ElementTransition, so a view chains either into what it already builds.
//
//	card := ui.Box(c).Padding(12).Children(...)
//	card = animate.ScaleIn(card, animate.Options{})
//
//	ui.DialogBase(c, &open, func(back, panel ui.Element) {
//		back.Transition(animate.Panel(animate.Options{}))
//	})
//
// Every animation honors the desktop's reduce-motion preference: MyGo
// completes a Keyframes at its last value and skips an Enter transition
// that moves, so nothing of the content is lost.
package animate

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Options configures one animation. The zero value is the house motion:
// 200 ms, eased out, moving 8 DIPs, starting at 95% scale.
type Options struct {
	// Duration is how long one play takes, or one period of a Pulse or
	// Spin. Zero uses 200 ms, or 1400 ms for Pulse and 800 ms for Spin.
	Duration time.Duration
	// Delay holds the entrance at its first value before it plays, so
	// rows enter one after another with a growing Delay.
	Delay time.Duration
	// Ease is how the value goes; EaseOut when nil.
	Ease ui.Easing
	// Distance is how many DIPs a rise or slide moves; 8 when zero.
	Distance float32
	// Scale is the starting scale of a ScaleIn or a Panel; 0.95 when
	// zero, and 1 for an entrance that does not scale.
	Scale float32
	// Run replays a one-shot entrance when it changes, as a button
	// replays a card entering.
	Run uint64
	// Exit also animates the surface going away, back where it came
	// from. Off by default: MyGo keeps a fading copy of the element
	// while it exits, which delays its closing and keeps its text in
	// the tree; hosts that want that motion set it.
	Exit bool
}

// Direction names the side a slide comes from, or a surface's edge.
type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

// A key prefix scopes every animation to its element: MyGo keeps an
// element's animations apart by key, and the kind keeps two animations
// of one element apart. One entrance per element — the entrance drives
// the opacity, so a second would fight it.
const keyPrefix = "uix.animate."

// entrance plays the 0→1 sequence of an entrance and returns its value,
// which the callers drive their property with. The element and key must
// stay stable across frames for MyGo to continue the animation.
func entrance(e ui.Element, key string, o Options) float32 {
	return e.Keyframes(keyPrefix+key, []ui.Keyframe{
		{At: 0, Value: 0},
		{At: o.duration(), Value: 1, Ease: o.ease()},
	}, ui.KeyframeOptions{Run: o.Run, Delay: o.Delay})
}

// FadeIn fades e from transparent to opaque, holding the result.
func FadeIn(e ui.Element, o Options) ui.Element {
	return e.Opacity(entrance(e, "fade", o))
}

// RiseIn fades e in while it rises Distance, as a toast or a menu comes
// up from below its anchor.
func RiseIn(e ui.Element, o Options) ui.Element {
	p := entrance(e, "rise", o)
	return e.Opacity(p).Translate(0, o.distance()*(1-p))
}

// SlideIn fades e in while it slides from the side, as a panel sweeps
// over from an edge.
func SlideIn(e ui.Element, o Options, from Direction) ui.Element {
	p := entrance(e, "slide", o)
	x, y := offset(o.distance(), from)
	return e.Opacity(p).Translate(x*(1-p), y*(1-p))
}

// ScaleIn fades e in while it grows from Scale around its center, as a
// card or a dialog takes its place. It sets the element's transform and
// origin, so apply it after any transform the caller keeps.
func ScaleIn(e ui.Element, o Options) ui.Element {
	p := entrance(e, "scale", o)
	s := o.scale() + (1-o.scale())*p
	return e.Opacity(p).Transform(ui.Scaling(s, s)).TransformOrigin(0.5, 0.5)
}

// Pulse breathes e's opacity between half and full, as a placeholder or
// a live dot. It keeps requesting frames while e is built.
func Pulse(e ui.Element, o Options) ui.Element {
	phase := e.Loop(keyPrefix+"pulse", o.period(1400*time.Millisecond), ui.Bounce(ui.EaseInOut))
	return e.Opacity(0.5 + 0.5*phase)
}

// Spin turns e around its center, once per period, as a spinner. It
// keeps requesting frames while e is built.
func Spin(e ui.Element, o Options) ui.Element {
	phase := e.Loop(keyPrefix+"spin", o.period(800*time.Millisecond), ui.Linear)
	return e.Rotate(phase*360).TransformOrigin(0.5, 0.5)
}

// Fade is the Enter/Exit of a surface that comes and goes in place:
// pure opacity, nothing moves. A toast is a Fade with a Rise.
func Fade(o Options) ui.ElementTransition {
	return transition(o, func(m *ui.Motion) { m.Opacity = 0 })
}

// Rise is the Enter/Exit of a surface that rises from below, as a toast
// floats up from the corner.
func Rise(o Options) ui.ElementTransition {
	return transition(o, func(m *ui.Motion) {
		m.Opacity = 0
		m.Y = o.distance()
	})
}

// Panel is the Enter/Exit of a centered surface, as a dialog or an
// empty state: it fades in, grows from Scale, and rises a little.
func Panel(o Options) ui.ElementTransition {
	return transition(o, func(m *ui.Motion) {
		m.Opacity = 0
		m.Scale = o.scale()
		m.Y = o.distance()
	})
}

// FromEdge is the Enter/Exit of a surface that lives on a window edge,
// as a drawer or a sheet: it fades in while it slides over from that
// side.
func FromEdge(from Direction, o Options) ui.ElementTransition {
	return transition(o, func(m *ui.Motion) {
		m.Opacity = 0
		m.X, m.Y = offset(o.distance(), from)
	})
}

// transition builds the Enter of a surface, and, when o.Exit asks, the
// Exit that sends it back where it came from.
func transition(o Options, set func(*ui.Motion)) ui.ElementTransition {
	in := &ui.Motion{}
	set(in)
	t := ui.ElementTransition{Duration: o.duration(), Ease: o.ease(), Enter: in}
	if o.Exit {
		e := *in
		t.Exit = &e
	}
	return t
}

// offset is the translate that slides in from the side: from the left
// means it comes from further left, negative on the x axis.
func offset(d float32, from Direction) (x, y float32) {
	switch from {
	case Left:
		return -d, 0
	case Right:
		return d, 0
	case Up:
		return 0, -d
	case Down:
		return 0, d
	}
	return 0, 0
}

func (o Options) duration() time.Duration {
	if o.Duration > 0 {
		return o.Duration
	}
	return 200 * time.Millisecond
}

func (o Options) period(fallback time.Duration) time.Duration {
	if o.Duration > 0 {
		return o.Duration
	}
	return fallback
}

func (o Options) ease() ui.Easing {
	if o.Ease != nil {
		return o.Ease
	}
	return ui.EaseOut
}

func (o Options) distance() float32 {
	if o.Distance > 0 {
		return o.Distance
	}
	return 8
}

func (o Options) scale() float32 {
	if o.Scale > 0 {
		return o.Scale
	}
	return 0.95
}
