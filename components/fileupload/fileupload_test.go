package fileupload

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// fu is the harness state of the tests: the upload, the picker, the
// files the completion reported, and how often the zone picked.
type fu struct {
	st    State
	pick  func() (Picked, error)
	done  []string
	picks int
}

// view is the harness the tests draw the upload zone in, with its state.
func view(s *fu) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			FileUpload(c, Props{
				State: &s.st,
				Pick: func() (Picked, error) {
					s.picks++
					return s.pick()
				},
				OnComplete: func(name string, size int64) { s.done = append(s.done, name) },
			})
		})
	}
}

// pump runs frames as time passes, for the simulated upload to advance.
func pump(tt *ui.Tester, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(2 * time.Millisecond)
		tt.Frame()
	}
}

// TestIdle draws the invite: the zone, the tagline and the kinds with
// the size limit.
func TestIdle(t *testing.T) {
	s := &fu{}
	tt := ui.NewTester(view(s), 420, 240)
	for _, want := range []string{
		"Upload a file",
		"Drag and drop to upload or",
		"select",
		"PDF, JPG, PNG, XLSX (max 8 MB)",
	} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
	if s.st.Phase != Idle {
		t.Fatalf("Phase = %v, want Idle", s.st.Phase)
	}
}

// TestPick uploads the picked file: the zone turns busy with its name
// and the uploading line.
func TestPick(t *testing.T) {
	s := &fu{pick: func() (Picked, error) {
		return Picked{Name: "report.pdf", Size: 1258291}, nil
	}}
	tt := ui.NewTester(view(s), 420, 240)
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if s.st.Phase != Uploading {
		t.Fatalf("Phase = %v, want Uploading", s.st.Phase)
	}
	for _, want := range []string{"report.pdf", "Uploading 1.2 MB…"} {
		if !tt.HasText(want) {
			t.Fatalf("missing %q in %q", want, tt.Texts())
		}
	}
}

// TestCompletes runs the simulated upload to the end: the zone returns
// to idle and the completion reports the file.
func TestCompletes(t *testing.T) {
	UploadDuration, CompleteHold = 8*time.Millisecond, 4*time.Millisecond
	t.Cleanup(func() { UploadDuration, CompleteHold = 2200*time.Millisecond, 1600*time.Millisecond })
	s := &fu{pick: func() (Picked, error) {
		return Picked{Name: "report.pdf", Size: 1258291}, nil
	}}
	tt := ui.NewTester(view(s), 420, 240)
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	pump(tt, 12)
	if s.st.Phase != Idle {
		t.Fatalf("Phase = %v, want Idle after the upload", s.st.Phase)
	}
	if len(s.done) != 1 || s.done[0] != "report.pdf" {
		t.Fatalf("done = %v, want [report.pdf]", s.done)
	}
}

// TestRejects refuses a file of another kind: the zone stays idle with
// the reason.
func TestRejects(t *testing.T) {
	s := &fu{pick: func() (Picked, error) { return Picked{Name: "movie.mp4", Size: 100}, nil }}
	tt := ui.NewTester(view(s), 420, 240)
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	if s.st.Phase != Idle {
		t.Fatalf("Phase = %v, want Idle for a rejected file", s.st.Phase)
	}
	want := "Only PDF, JPG, JPEG, PNG, XLSX files are supported"
	if !strings.Contains(s.st.Rejection, "PDF, JPG, JPEG, PNG, XLSX") {
		t.Fatalf("Rejection = %q, want it to name the kinds", s.st.Rejection)
	}
	if !tt.HasText(want) {
		t.Fatalf("missing the rejection %q in %q", want, tt.Texts())
	}
}

// TestTooBig refuses a file over the limit with its size named.
func TestTooBig(t *testing.T) {
	s := &fu{pick: func() (Picked, error) { return Picked{Name: "big.pdf", Size: 9 << 20}, nil }}
	tt := ui.NewTester(view(s), 420, 240)
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	if s.st.Rejection != "That file is larger than 8 MB" {
		t.Fatalf("Rejection = %q, want the size limit", s.st.Rejection)
	}
}

// TestRejectionClears puts the rejection away after its hold, and the
// invite returns.
func TestRejectionClears(t *testing.T) {
	RejectionHold = 8 * time.Millisecond
	t.Cleanup(func() { RejectionHold = 2600 * time.Millisecond })
	s := &fu{pick: func() (Picked, error) { return Picked{Name: "movie.mp4", Size: 100}, nil }}
	tt := ui.NewTester(view(s), 420, 240)
	tt.Click("Upload a file")
	if s.st.Rejection == "" {
		t.Fatal("no rejection after the click")
	}
	pump(tt, 8)
	if s.st.Rejection != "" {
		t.Fatalf("Rejection = %q, want it cleared", s.st.Rejection)
	}
	if !tt.HasText("select") {
		t.Fatalf("the invite is missing from %q", tt.Texts())
	}
}

// TestBusyIgnoresClicks does not pick again while the zone uploads.
func TestBusyIgnoresClicks(t *testing.T) {
	s := &fu{pick: func() (Picked, error) { return Picked{Name: "a.pdf", Size: 1000}, nil }}
	tt := ui.NewTester(view(s), 420, 240)
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Upload a file"); err != nil {
		t.Fatal(err)
	}
	if s.picks != 1 {
		t.Fatalf("picks = %d, want 1 while uploading", s.picks)
	}
}

// TestCustomKindsAndLimit draws the app's own kinds and limit.
func TestCustomKindsAndLimit(t *testing.T) {
	tt := ui.NewTester(func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			FileUpload(c, Props{
				State:             &State{},
				AllowedExtensions: []string{"csv", "txt"},
				MaxBytes:          1024,
				Pick:              func() (Picked, error) { return Picked{Name: "a.csv", Size: 512}, nil },
			})
		})
	}, 420, 240)
	if !tt.HasText("CSV, TXT (max 1 KB)") {
		t.Fatalf("missing the kinds and limit in %q", tt.Texts())
	}
}

// TestStartShowsProgress presets an upload begun in the past, so the
// zone shows the progress it has made, without a dialog or a click.
func TestStartShowsProgress(t *testing.T) {
	s := &fu{}
	s.st.Start("report.pdf", 1258291, time.Now().Add(-time.Duration(0.4*float64(UploadDuration))))
	tt := ui.NewTester(view(s), 420, 240)
	if s.st.Phase != Uploading {
		t.Fatalf("Phase = %v, want Uploading", s.st.Phase)
	}
	if s.st.Progress < 35 || s.st.Progress > 45 {
		t.Fatalf("Progress = %d, want around 40", s.st.Progress)
	}
	if !tt.HasText("report.pdf") {
		t.Fatalf("missing the name in %q", tt.Texts())
	}
}

// TestDroppedStartsUpload begins the upload of a file dropped on the
// zone, the first of the drop, as a pick would.
func TestDroppedStartsUpload(t *testing.T) {
	s := &fu{}
	path := filepath.Join(t.TempDir(), "report.pdf")
	if err := os.WriteFile(path, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	startDropped(&s.st, time.Now(), path, defaultExtensions, defaultMaxBytes)
	if s.st.Phase != Uploading {
		t.Fatalf("Phase = %v, want Uploading", s.st.Phase)
	}
	if s.st.Name != "report.pdf" {
		t.Fatalf("Name = %q, want report.pdf", s.st.Name)
	}
	fi, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if s.st.Size != fi.Size() {
		t.Fatalf("Size = %d, want %d", s.st.Size, fi.Size())
	}
}

// TestDroppedRejects refuses a dropped file of another kind, as a pick.
func TestDroppedRejects(t *testing.T) {
	s := &fu{}
	path := filepath.Join(t.TempDir(), "movie.mp4")
	if err := os.WriteFile(path, []byte("x"), 0o600); err != nil {
		t.Fatal(err)
	}
	startDropped(&s.st, time.Now(), path, defaultExtensions, defaultMaxBytes)
	if s.st.Phase != Idle {
		t.Fatalf("Phase = %v, want Idle for a rejected drop", s.st.Phase)
	}
	if s.st.Rejection == "" {
		t.Fatal("no rejection after the drop")
	}
}

// TestDroppedGone ignores a path that vanished after the drop.
func TestDroppedGone(t *testing.T) {
	s := &fu{}
	startDropped(&s.st, time.Now(), filepath.Join(t.TempDir(), "gone.pdf"),
		defaultExtensions, defaultMaxBytes)
	if s.st.Phase != Idle || s.st.Rejection != "" {
		t.Fatalf("Phase = %v, Rejection = %q, want the vanished drop ignored",
			s.st.Phase, s.st.Rejection)
	}
}

// TestFormatSize names sizes the way BoardUI does.
func TestFormatSize(t *testing.T) {
	for _, tc := range []struct {
		bytes int64
		want  string
	}{
		{8 << 20, "8 MB"},
		{10 << 20, "10 MB"},
		{1048576, "1 MB"},
		{1258291, "1.2 MB"},
		{1572864, "1.5 MB"},
		{512 * 1024, "512 KB"},
		{100, "1 KB"},
	} {
		if got := formatFileSize(tc.bytes); got != tc.want {
			t.Errorf("formatFileSize(%d) = %q, want %q", tc.bytes, got, tc.want)
		}
	}
}

// TestAllowedLabel drops the duplicate jpeg from the tagline.
func TestAllowedLabel(t *testing.T) {
	if got := allowedLabel(defaultExtensions); got != "PDF, JPG, PNG, XLSX" {
		t.Fatalf("allowedLabel = %q, want PDF, JPG, PNG, XLSX", got)
	}
}
