// Package log provides the AgentLog of BoardUI for the MyGo native
// toolkit: a transcript of an agent's work, each line joined to the next
// by a curved tree guide, with a working row of stars while the agent is
// still going. It is the machinery of a streaming transcript; what a row
// says is the app's own business.
//
//	log.Log(c, log.Props{
//		Items:    steps,
//		Revealed: revealed,
//		Running:  working,
//		Working:  "Searching the docs",
//	})
//
// The reveal is the app's own pacing, not the component's: pass how many
// items are shown, and the app advances it with a timer or with the real
// events of the work, so the log never pretends steps take equal time.
package log

import (
	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/thinking"
)

// Props describes the log to draw.
type Props struct {
	// Items are the lines of the work, in order.
	Items []string
	// Revealed is how many items are shown, as the app has paced them;
	// it is clamped to the length of Items. 0 shows none.
	Revealed int
	// Running shows the working row below the revealed items, with the
	// stars indicator, while the log is still going.
	Running bool
	// Working is the label of the working row, such as "Searching the
	// docs".
	Working string
}

// Log draws the transcript and returns it, so a view can chain more
// calls on it.
func Log(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	n := min(p.Revealed, len(p.Items))
	return ui.Column(c).FillWidth().Gap(t.Space(1)).Children(func() {
		for i := 0; i < n; i++ {
			row(c, t, p.Items[i], i == n-1)
		}
		if p.Running {
			workingRow(c, t, p.Working)
		}
	})
}

// row is one completed line of work, with the stretch of guide that
// belongs to it: a branch at the top and a trunk running to the next
// row, which the last row omits so the guide finishes on the corner.
func row(c *ui.Context, t *ui.Theme, label string, last bool) {
	guide := t.TextMuted.Alpha(0.45)
	ui.Row(c).FillWidth().AlignItems(ui.Stretch).Children(func() {
		// The branch geometry, as the docs sidebar draws it: down the
		// trunk, round a corner at 14px, out to the label.
		ui.Box(c).Width(12).Shrink(0).Draw(func(pt *ui.Painter, r ui.Rect) {
			var branch ui.Path
			branch.MoveTo(r.X+0.5, r.Y)
			branch.LineTo(r.X+0.5, r.Y+8)
			branch.QuadTo(r.X+0.5, r.Y+14, r.X+6.5, r.Y+14)
			branch.LineTo(r.X+11.5, r.Y+14)
			pt.StrokePath(&branch, 1, guide)
			if !last {
				pt.Line(r.X+0.5, r.Y+14, r.X+0.5, r.Y+r.H, 1, guide)
			}
		})
		ui.Text(c, label).TextColor(t.Text).Padding(0, t.Space(1))
	})
}

// workingRow is the line the agent is on now, with the stars indicator.
// It deliberately sits off the guide: the tree records what the agent
// did, and this line has not happened yet, so hanging it on the tree
// would file it alongside the findings.
func workingRow(c *ui.Context, t *ui.Theme, label string) {
	ui.Row(c).Padding(0, t.Space(1)).Children(func() {
		thinking.Thinking(c, thinking.Props{
			Variant: thinking.Stars,
			Label:   label,
			Tone:    thinking.Subtle,
		})
	})
}
