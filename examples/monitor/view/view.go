// Package view builds the monitor example: a window of MyGo native UI
// that shows the meter, spinner, rangeslider and stepper components
// together in a system monitor. The meters show the load of the CPU,
// the memory and the disk, Refresh cycles to another reading, Scan runs
// a fake scan with a spinner while it is on, the range slider sets the
// load target, and the stepper sets the port. The main package shows it
// in a window; the tests and the snapshot command draw it headless.
package view

import (
	"fmt"

	"github.com/egoist/mygo/ui"

	tokens "github.com/smallnest/uix/components/base"
	"github.com/smallnest/uix/components/meter"
	"github.com/smallnest/uix/components/rangeslider"
	"github.com/smallnest/uix/components/separator"
	"github.com/smallnest/uix/components/spinner"
	"github.com/smallnest/uix/components/stepper"
)

// Width and Height are the size of the example window.
const Width, Height = 640, 480

// State is the state of the example; the controls edit it in place.
var State = MonitorState{}

// MonitorState holds the monitor, as the controls edit it.
type MonitorState struct {
	// CPU, Memory and Disk are the load readings, in percent.
	CPU, Memory, Disk float64
	// Reading is which of the preset readings is shown.
	Reading int
	// Scanning runs the fake scan, which shows the spinner.
	Scanning bool
	// TargetLow and TargetHigh are the bounds of the load target, in
	// percent, which the range slider sets.
	TargetLow, TargetHigh float64
	// Port is the port to monitor, which the stepper sets.
	Port float64
	// Status is the last thing the controls did, shown at the bottom;
	// empty for none.
	Status string
}

// Reset puts the example in its initial state, for the tests and the
// snapshot command.
func Reset() {
	State = MonitorState{
		CPU: 12, Memory: 64, Disk: 23,
		Reading: 0,
		TargetLow: 20, TargetHigh: 80,
		Port:   8080,
		Status: "Idle. Press Scan to scan for threats.",
	}
}

// reading is one preset of load readings, which Refresh cycles through.
type reading struct {
	cpu, memory, disk float64
}

// readings are the load readings Refresh shows in turn. The first is
// also the initial state.
var readings = []reading{
	{cpu: 12, memory: 64, disk: 23},
	{cpu: 41, memory: 71, disk: 45},
	{cpu: 7, memory: 49, disk: 12},
	{cpu: 88, memory: 95, disk: 91},
}

// pct renders a load reading as a percentage.
func pct(v float64) string { return fmt.Sprintf("%.0f%%", v) }

// targetFormat renders the load target as a percentage range.
func targetFormat(lo, hi float64) string { return fmt.Sprintf("%.0f%% – %.0f%%", lo, hi) }

// MonitorView draws the example: a system monitor whose controls all
// work.
func MonitorView(c *ui.Context) {
	c.SetTheme(tokens.For(c.Theme().Dark))
	t := c.Theme()
	ui.Column(c).Fill().Children(func() {
		ui.Row(c).PaddingX(16).PaddingY(10).Gap(8).Children(func() {
			ui.Text(c, "System Monitor").Bold().Grow(1)
			if State.Scanning {
				spinner.Spinner(c, spinner.Props{Label: "Scanning"})
				ui.Text(c, "Scanning").TextColor(t.TextMuted)
			}
			if ui.Button(c, "Refresh").Clicked() {
				State.Reading = (State.Reading + 1) % len(readings)
				State.CPU, State.Memory, State.Disk = readings[State.Reading].cpu, readings[State.Reading].memory, readings[State.Reading].disk
				State.Status = fmt.Sprintf("Showed reading %d of %d.", State.Reading+1, len(readings))
			}
			label := "Scan"
			if State.Scanning {
				label = "Stop"
			}
			if ui.Button(c, label).Clicked() {
				if State.Scanning {
					State.Scanning = false
					State.Status = "Stopped the scan."
				} else {
					State.Scanning = true
					State.Status = "Scanning for threats…"
				}
			}
		})
		separator.Separator(c, separator.Props{})
		ui.Column(c).Padding(24).Gap(14).Children(func() {
			meter.Meter(c, meter.Props{Value: State.CPU, Max: 100, Label: "CPU", Format: pct})
			meter.Meter(c, meter.Props{Value: State.Memory, Max: 100, Label: "Memory", Format: pct})
			meter.Meter(c, meter.Props{Value: State.Disk, Max: 100, Label: "Disk", Format: pct,
				Levels: &ui.MeterLevels{Warning: 70, Critical: 90}})
			rangeslider.RangeSlider(c, rangeslider.Props{
				Low: &State.TargetLow, High: &State.TargetHigh,
				Min: 0, Max: 100, Step: 1,
				Label: "Load target", Format: targetFormat,
			})
			stepper.Stepper(c, stepper.Props{
				Value: &State.Port, Min: 1024, Max: 65535, Step: 1,
				Label: "Port",
			})
		})
		ui.Spacer(c)
		ui.Row(c).PaddingX(24).PaddingY(10).Children(func() {
			ui.Text(c, State.Status).TextColor(t.TextMuted)
		})
	})
}
