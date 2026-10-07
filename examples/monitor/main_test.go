package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/monitor/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.MonitorView, view.Width, view.Height)
}

// TestRenders draws the monitor headless: the title, the buttons, the
// meters with their readouts, the range slider, the stepper and the
// status line.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"System Monitor", "Refresh", "Scan",
		"CPU", "12%", "Memory", "64%", "Disk", "23%",
		"Load target", "20% – 80%",
		"Port", "8080",
		"Idle. Press Scan to scan for threats.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestRefresh cycles to the next reading, which moves the meters and
// names the reading in the status.
func TestRefresh(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Refresh"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("41%") || tt.HasText("12%") {
		t.Fatalf("Refresh did not move the meters: %q", tt.Texts())
	}
	if !tt.HasText("Showed reading 2 of 4.") {
		t.Fatal("Refresh did not name the reading")
	}
}

// TestScan runs the scan, which shows the spinner and its label, and
// stops it, which hides them.
func TestScan(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Scan"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Scanning || !tt.HasText("Scanning") {
		t.Fatalf("Scan did not start: %q", tt.Texts())
	}
	if err := tt.Click("Stop"); err != nil {
		t.Fatal(err)
	}
	if view.State.Scanning || tt.HasText("Scanning") {
		t.Fatalf("Stop did not end the scan: %q", tt.Texts())
	}
	if !tt.HasText("Stopped the scan.") {
		t.Fatal("Stop did not name the stop")
	}
}

// TestRange drags the track of the range slider, which moves the low
// knob toward the press and shows the new range in the readout.
func TestRange(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("20% – 80%")
	if !ok {
		t.Fatal("no load target readout")
	}
	// The track sits under the label row, its middle 10 below the top
	// of the Space(7) tall element. A press at the share 40% of the
	// track is nearer the low knob, which moves to 40.
	x := float32(24) + 0.4*(view.Width-48)
	y := r.Y + r.H + 2 + 10
	tt.Press(x, y)
	tt.Frame()
	tt.Release(x, y)
	tt.Frame()
	if view.State.TargetLow != 40 || view.State.TargetHigh != 80 {
		t.Fatalf("the drag set the target %v – %v, want 40 – 80", view.State.TargetLow, view.State.TargetHigh)
	}
	if !tt.HasText("40% – 80%") {
		t.Fatalf("the readout did not follow: %q", tt.Texts())
	}
}

// TestStepper clicks the up arrow of the stepper, which raises the port
// and shows it in the readout.
func TestStepper(t *testing.T) {
	tt := newTester(t)
	r, ok := tt.Find("8080")
	if !ok {
		t.Fatal("no port readout")
	}
	// The stepper sits under the label row; its up arrow is the top
	// half, 14 high, at the left of the padding 24.
	x := float32(24) + 10
	y := r.Y + r.H + 2 + 7
	tt.ClickAt(x, y)
	if view.State.Port != 8081 {
		t.Fatalf("the up arrow set the port %v, want 8081", view.State.Port)
	}
	if !tt.HasText("8081") {
		t.Fatalf("the readout did not follow: %q", tt.Texts())
	}
}

// TestDarkMode draws the monitor under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("System Monitor") {
		t.Fatal("dark monitor missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
