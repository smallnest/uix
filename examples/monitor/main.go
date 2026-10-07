// Example monitor opens a window of MyGo native UI that shows the
// meter, spinner, rangeslider and stepper components together in a
// system monitor. The meters show the load of the CPU, the memory and
// the disk, Refresh cycles to another reading, Scan runs a fake scan
// with a spinner while it is on, the range slider sets the load target,
// and the stepper sets the port.
//
//	go run ./examples/monitor                open the window
//	go test ./examples/monitor/...           render and check it headless
//	go run ./examples/monitor/cmd/snapshot   write monitor.png, monitor-scanning.png and monitor-dark.png
package main

import (
	"log"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/monitor/view"
)

func main() {
	mygo.App.WhenReady(func() {
		mygo.NewWindow(mygo.WindowOptions{
			Title:   "uix monitor",
			Width:   view.Width,
			Height:  view.Height,
			Content: ui.View(view.MonitorView),
		})
	})
	if err := mygo.App.Run(); err != nil {
		log.Fatal(err)
	}
}
