// Package fileupload provides the file upload of BoardUI for the MyGo
// native toolkit: a drop zone you click to pick a file, which then
// uploads with a progress ring and a percentage until it is done.
//
//	fileupload.FileUpload(c, fileupload.Props{
//		State: &state.Upload,
//	})
//
// The zone names the accepted kinds and the size limit, picks a file
// through the native open dialog, rejects files outside them with the
// reason, and simulates the upload, so the component needs no storage
// layer: connect OnComplete to your upload/storage layer. The app owns
// the State, as it owns the scroll of a chat. Drag and drop of BoardUI
// has no native file counterpart in MyGo, so the zone selects with a
// click instead; Pick can replace the dialog with any picker.
package fileupload

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/egoist/mygo"
	"github.com/egoist/mygo/ui"
)

// Phase is where the upload is.
type Phase int

// The phases of the zone.
const (
	// Idle invites a file.
	Idle Phase = iota
	// Uploading shows the file and its progress.
	Uploading
	// Complete announces the finished upload before returning to Idle.
	Complete
)

// The timings of the simulated upload. Tests shorten them to run the
// simulation in a few frames.
var (
	// UploadDuration is how long a simulated upload takes.
	UploadDuration = 2200 * time.Millisecond
	// CompleteHold is how long the complete state shows before it resets.
	CompleteHold = 1600 * time.Millisecond
	// RejectionHold is how long a rejection message shows.
	RejectionHold = 2600 * time.Millisecond
	// progressTick is how often the ring repaints while uploading.
	progressTick = 100 * time.Millisecond
)

// State is where the upload is, kept by the app.
type State struct {
	// Phase of the zone.
	Phase Phase
	// Name and Size of the file uploading or uploaded.
	Name string
	Size int64
	// Progress is the upload in percent, 0 to 100.
	Progress int
	// Rejection is why the last file was not accepted, "" when none.
	Rejection string

	// started is when the upload began, next when it moves on, and
	// rejectionUntil when the rejection clears; the simulation reads
	// them as time passes.
	started        time.Time
	next           time.Time
	rejectionUntil time.Time
}

// Start begins the simulated upload of a valid file at the time now, as
// picking it would: pass time.Now to start now, or an earlier time to
// show progress already made. Examples and tests use it without a
// dialog.
func (s *State) Start(name string, size int64, now time.Time) {
	s.Phase = Uploading
	s.Name = name
	s.Size = size
	s.Progress = 0
	s.Rejection = ""
	s.started = now
	s.next = time.Time{}
}

// Picked is a file the picker chose.
type Picked struct {
	Name string
	Size int64
}

// Props describes the upload zone to draw.
type Props struct {
	// State is where the upload is, for the app to keep.
	State *State
	// Pick chooses a file when the zone is clicked; nil opens the native
	// dialog filtered to the allowed kinds.
	Pick func() (Picked, error)
	// AllowedExtensions are the accepted file extensions without dots;
	// pdf, jpg, jpeg, png and xlsx when nil.
	AllowedExtensions []string
	// MaxBytes is the largest accepted file; 8 MB when zero.
	MaxBytes int64
	// Height is the height of the zone; 164 when zero.
	Height float32
	// OnComplete reports the file when the simulated upload finishes.
	OnComplete func(name string, size int64)
}

// The kinds and the limit BoardUI accepts.
var defaultExtensions = []string{"pdf", "jpg", "jpeg", "png", "xlsx"}

const defaultMaxBytes int64 = 8 << 20

// FileUpload draws the upload zone and returns it, so a view can chain
// more calls on it.
func FileUpload(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	st := p.State
	if st == nil {
		st = &State{}
	}
	exts := p.AllowedExtensions
	if len(exts) == 0 {
		exts = defaultExtensions
	}
	maxBytes := p.MaxBytes
	if maxBytes <= 0 {
		maxBytes = defaultMaxBytes
	}
	h := p.Height
	if h <= 0 {
		h = 164
	}
	radius := t.Radius * 2

	// The simulation advances with the clock, so the ring moves whether
	// or not the view changes otherwise.
	tick(c, st, p)

	z := ui.ButtonBase(c).FillWidth().Height(h).Radius(radius).Background(t.Surface)
	z.Label("Upload a file")
	if st.Phase != Idle {
		z.Disabled(true)
		// The ring around the zone and the percent badge above its middle,
		// drawn each frame as the progress advances.
		z.Draw(func(pt *ui.Painter, r ui.Rect) {
			ring := ui.Rect{X: r.X + 1, Y: r.Y + 1, W: r.W - 2, H: r.H - 2}
			pt.Stroke(ring, t.Border, radius, 2)
			frac := float32(st.Progress) / 100
			if frac > 0 {
				pt.StrokePath(ringPath(ring, radius, frac), 2, t.Accent)
			}
			badge := fmt.Sprintf("%d%%", st.Progress)
			size := t.FontSize * 0.8
			w := float32(len(badge))*size*0.55 + 12
			br := ui.Rect{X: r.X + r.W/2 - w/2, Y: r.Y + 8, W: w, H: 18}
			pt.Fill(br, t.Accent, 9)
			tw := float32(len(badge)) * size * 0.55
			pt.Text(br.X+(br.W-tw)/2, br.Y+br.H/2+size*0.35, badge, size, t.AccentText)
		})
	} else {
		// The dashed outline of the invite, darker while the pointer
		// rests on the zone.
		z.Draw(func(pt *ui.Painter, r ui.Rect) {
			dash := t.Border
			if z.Hovered() {
				dash = t.Text.Alpha(0.35)
			}
			pt.StrokeDashed(ui.Rect{X: r.X + 1, Y: r.Y + 1, W: r.W - 2, H: r.H - 2}, dash, radius, 2)
		})
	}
	z.Children(func() {
		ui.Column(c).FillWidth().Children(func() {
			if st.Phase == Idle {
				idle(c, t, st, exts, maxBytes)
			} else {
				busy(c, t, st)
			}
		})
	})
	if st.Phase == Idle && z.Clicked() {
		pick := p.Pick
		if pick == nil {
			pick = defaultPick(exts)
		}
		if f, err := pick(); err == nil && f.Name != "" {
			startUpload(st, c.Now(), f, exts, maxBytes)
		}
	}
	return z
}

// idle is the invite: the cloud icon, the tagline, and the kinds with
// the size limit, or the rejection instead of the tagline.
func idle(c *ui.Context, t *ui.Theme, st *State, exts []string, maxBytes int64) {
	ui.Column(c).FillWidth().Grow(1).Center().Gap(t.Space(1.5)).Children(func() {
		ui.Box(c).Size(40, 40).Background(t.SurfaceHover).Radius(20).Center().Children(func() {
			cloudIcon(c, t)
		})
		if st.Rejection == "" {
			ui.Row(c).Center().Gap(2).Children(func() {
				ui.Text(c, "Drag and drop to upload or ").TextColor(t.TextMuted).SingleLine()
				ui.Text(c, "select").TextColor(t.Accent).SingleLine()
			})
		} else {
			ui.Text(c, st.Rejection).TextColor(t.Danger).SingleLine()
		}
		ui.Text(c, fmt.Sprintf("%s (max %s)", allowedLabel(exts), formatFileSize(maxBytes))).
			FontSize(t.FontSize * 0.85).TextColor(t.TextMuted.Alpha(0.7)).SingleLine()
	})
}

// busy is the upload: the file icon, its name, and the line that moves
// from uploading to done.
func busy(c *ui.Context, t *ui.Theme, st *State) {
	ui.Column(c).FillWidth().Grow(1).Center().Gap(t.Space(1)).Children(func() {
		ui.Box(c).Size(40, 40).Radius(20).Border(1, t.Border).Center().Children(func() {
			fileIcon(c, t, t.Text)
		})
		ui.Text(c, st.Name).TextColor(t.Text).SingleLine()
		if st.Phase == Uploading {
			ui.Text(c, fmt.Sprintf("Uploading %s…", formatFileSize(st.Size))).TextColor(t.TextMuted).SingleLine()
		} else {
			ui.Text(c, "Uploaded successfully!").TextColor(t.Success).SingleLine()
		}
	})
}

// tick advances the simulation: the progress with the clock while
// uploading, the return to idle after the complete hold, and the
// clearing of a rejection. It asks for the frames those take.
func tick(c *ui.Context, st *State, p Props) {
	now := c.Now()
	switch st.Phase {
	case Uploading:
		elapsed := now.Sub(st.started)
		st.Progress = int(float64(elapsed) * 100 / float64(UploadDuration))
		if st.Progress >= 100 {
			st.Progress = 100
			st.Phase = Complete
			st.next = now.Add(CompleteHold)
			c.After(CompleteHold)
			return
		}
		// A frame in a tick keeps the ring moving; the completion timer
		// wakes it at the end.
		c.After(progressTick)
		c.After(UploadDuration - elapsed)
	case Complete:
		if st.next.IsZero() {
			// Set only now, so a preset complete state holds from here.
			st.next = now.Add(CompleteHold)
			c.After(CompleteHold)
			return
		}
		if now.Before(st.next) {
			c.After(st.next.Sub(now))
			return
		}
		if p.OnComplete != nil {
			p.OnComplete(st.Name, st.Size)
		}
		st.Phase = Idle
		st.Name, st.Size, st.Progress = "", 0, 0
	}
	if st.Rejection != "" {
		if now.Before(st.rejectionUntil) {
			c.After(st.rejectionUntil.Sub(now))
		} else {
			st.Rejection = ""
		}
	}
}

// startUpload validates the picked file and begins its upload, or
// rejects it with the reason, as BoardUI does.
func startUpload(st *State, now time.Time, f Picked, exts []string, maxBytes int64) {
	if !accepted(f.Name, exts) {
		st.Rejection = fmt.Sprintf("Only %s files are supported", upperAll(exts))
		st.rejectionUntil = now.Add(RejectionHold)
		return
	}
	if f.Size > maxBytes {
		st.Rejection = fmt.Sprintf("That file is larger than %s", formatFileSize(maxBytes))
		st.rejectionUntil = now.Add(RejectionHold)
		return
	}
	st.Start(f.Name, f.Size, now)
}

// defaultPick is the native open dialog, filtered to the allowed kinds.
func defaultPick(exts []string) func() (Picked, error) {
	return func() (Picked, error) {
		paths, err := mygo.Dialog.Open(mygo.OpenDialogOptions{
			Title:       "Upload a file",
			ButtonLabel: "Upload",
			Filters:     []mygo.FileFilter{{Name: "Supported files", Extensions: exts}},
		})
		if err != nil || len(paths) == 0 {
			return Picked{}, err
		}
		st, err := os.Stat(paths[0])
		if err != nil {
			return Picked{}, err
		}
		return Picked{Name: filepath.Base(paths[0]), Size: st.Size()}, nil
	}
}

// accepted reports whether the file's extension is among the kinds.
func accepted(name string, exts []string) bool {
	e := strings.ToLower(extensionFor(name))
	for _, x := range exts {
		if strings.ToLower(x) == e {
			return true
		}
	}
	return false
}

// extensionFor is the suffix after the last dot, "" without one.
func extensionFor(name string) string {
	if i := strings.LastIndexByte(name, '.'); i >= 0 {
		return name[i+1:]
	}
	return ""
}

// allowedLabel is the kinds the tagline names, without the duplicate
// jpeg, as BoardUI names them.
func allowedLabel(exts []string) string {
	var kept []string
	for _, e := range exts {
		if e != "jpeg" {
			kept = append(kept, strings.ToUpper(e))
		}
	}
	return strings.Join(kept, ", ")
}

// upperAll is the kinds a rejection names.
func upperAll(exts []string) string {
	up := make([]string, len(exts))
	for i, e := range exts {
		up[i] = strings.ToUpper(e)
	}
	return strings.Join(up, ", ")
}

// formatFileSize names a size in bytes the way BoardUI does: whole MB
// from ten up, one decimal below, whole KB below a MB and at least one.
func formatFileSize(bytes int64) string {
	if bytes >= 1024*1024 {
		mb := float64(bytes) / (1024 * 1024)
		if mb >= 10 {
			return fmt.Sprintf("%d MB", int64(math.Round(mb)))
		}
		// One decimal, and a whole number drops it, as BoardUI shows it.
		v := math.Round(mb*10) / 10
		if v == float64(int64(v)) {
			return fmt.Sprintf("%d MB", int64(v))
		}
		return fmt.Sprintf("%.1f MB", v)
	}
	return fmt.Sprintf("%d KB", max(1, bytes/1024))
}

// cloudIcon is the cloud with the up arrow of the invite, punched out
// of the cloud with the circle's color.
func cloudIcon(c *ui.Context, t *ui.Theme) {
	ui.Box(c).Size(24, 24).Draw(func(pt *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		fg := t.TextMuted
		hole := t.SurfaceHover
		// The cloud: a band under three bumps.
		pt.Fill(ui.Rect{X: cx - 11, Y: cy + 1, W: 22, H: 7}, fg, 3.5)
		for _, b := range []struct{ x, y, rad float32 }{
			{cx - 6.5, cy - 1, 4.5},
			{cx, cy - 3.5, 6},
			{cx + 6.5, cy - 1, 4.5},
		} {
			pt.Fill(ui.Rect{X: b.x - b.rad, Y: b.y - b.rad, W: 2 * b.rad, H: 2 * b.rad}, fg, b.rad)
		}
		// The up arrow, cut out of the cloud.
		pt.Fill(ui.Rect{X: cx - 1.5, Y: cy - 2, W: 3, H: 8}, hole, 1.5)
		var head ui.Path
		head.MoveTo(cx-5, cy-1).LineTo(cx, cy-7).LineTo(cx+5, cy-1).LineTo(cx+2.5, cy-1).
			LineTo(cx+2.5, cy+2).LineTo(cx-2.5, cy+2).LineTo(cx-2.5, cy-1).Close()
		pt.FillPath(&head, hole)
	})
}

// fileIcon is the document with the folded corner and two lines of the
// busy view, as a line icon.
func fileIcon(c *ui.Context, t *ui.Theme, fg ui.Color) {
	ui.Box(c).Size(24, 24).Draw(func(pt *ui.Painter, r ui.Rect) {
		cx, cy := r.X+r.W/2, r.Y+r.H/2
		var doc ui.Path
		doc.MoveTo(cx-9, cy-12).LineTo(cx+3, cy-12).LineTo(cx+10, cy-5).LineTo(cx+10, cy+12).
			LineTo(cx-9, cy+12).Close()
		pt.StrokePath(&doc, 1.8, fg)
		pt.Fill(ui.Rect{X: cx - 6, Y: cy - 1, W: 9, H: 2}, fg, 1)
		pt.Fill(ui.Rect{X: cx - 6, Y: cy + 4, W: 12, H: 2}, fg, 1)
	})
}

// ringPath is the perimeter of the rounded rectangle from the top
// center clockwise, up to the given fraction of it: the progress arc of
// the ring.
func ringPath(r ui.Rect, radius float32, frac float32) *ui.Path {
	radius = min(radius, r.W/2, r.H/2)
	x0, y0, x1, y1 := r.X, r.Y, r.X+r.W, r.Y+r.H
	type piece struct {
		length float32
		at     func(a float32) xy
	}
	arc := func(cx, cy, from, to float32) piece {
		return piece{math.Pi / 2 * radius, func(a float32) xy {
			ang := from + (to-from)*a
			return xy{cx + radius*float32(math.Cos(float64(ang))), cy + radius*float32(math.Sin(float64(ang)))}
		}}
	}
	// The eight pieces of the perimeter: the four straight edges and the
	// four quarter arcs at the corners.
	pieces := []piece{
		{r.W/2 - radius, func(a float32) xy { return xy{x0 + r.W/2 + (r.W/2-radius)*a, y0} }},
		arc(x1-radius, y0+radius, -math.Pi/2, 0),
		{r.H - 2*radius, func(a float32) xy { return xy{x1, y0 + radius + (r.H-2*radius)*a} }},
		arc(x1-radius, y1-radius, 0, math.Pi/2),
		{r.W - 2*radius, func(a float32) xy { return xy{x1 - (r.W-2*radius)*a, y1} }},
		arc(x0+radius, y1-radius, math.Pi/2, math.Pi),
		{r.H - 2*radius, func(a float32) xy { return xy{x0, y1 - (r.H-2*radius)*a} }},
		arc(x0+radius, y0+radius, math.Pi, 3*math.Pi/2),
	}
	total := float32(0)
	for i := range pieces {
		total += pieces[i].length
	}
	// Walk the pieces up to the target, a point about every 1.5 points.
	target := frac * total
	var path ui.Path
	started := false
	cum := float32(0)
	for _, pc := range pieces {
		if target <= cum {
			break
		}
		avail := target - cum
		partial := avail < pc.length
		n := int(min(avail, pc.length)/1.5) + 1
		to := float32(1)
		if partial {
			to = avail / pc.length
		}
		for k := 0; k <= n; k++ {
			q := pc.at(float32(k) / float32(n) * to)
			if !started {
				path.MoveTo(q.X, q.Y)
				started = true
			} else {
				path.LineTo(q.X, q.Y)
			}
		}
		cum += pc.length
	}
	return &path
}

// xy is a point on the ring.
type xy struct{ X, Y float32 }
