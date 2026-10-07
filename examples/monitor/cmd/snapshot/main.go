// Command snapshot renders the monitor example headless into PNG files,
// so the components can be seen without opening a window:
//
//	go run ./examples/monitor/cmd/snapshot
package main

import (
	"image/png"
	"os"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/monitor/view"
)

func main() {
	// The plain shot draws the monitor in its initial state.
	view.Reset()
	save(ui.NewTester(view.MonitorView, view.Width, view.Height), "monitor.png")

	// The scanning shot runs the scan, raises the readings — the disk
	// past its warning level, which turns its meter yellow — narrows
	// the load target and raises the port, as the user would with the
	// controls.
	view.Reset()
	view.State.CPU, view.State.Memory, view.State.Disk = 41, 71, 87
	view.State.Scanning = true
	view.State.TargetLow, view.State.TargetHigh = 40, 70
	view.State.Port = 8443
	view.State.Status = "Scanning for threats…"
	save(ui.NewTester(view.MonitorView, view.Width, view.Height), "monitor-scanning.png")

	// The dark shot draws the same monitor under the dark appearance.
	view.Reset()
	tt := ui.NewTester(view.MonitorView, view.Width, view.Height)
	tt.SetDark(true)
	save(tt, "monitor-dark.png")
}

// save writes the frame of the tester into a PNG file.
func save(tt *ui.Tester, name string) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if err := png.Encode(f, tt.Image()); err != nil {
		panic(err)
	}
}
