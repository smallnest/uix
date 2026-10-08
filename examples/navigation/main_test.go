package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/fileupload"
	"github.com/smallnest/uix/examples/navigation/view"
)

// newTester draws the example headless in its initial state.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.NavigationView, view.Width, view.Height)
}

// pump runs frames as time passes, for the simulated upload to advance.
func pump(tt *ui.Tester, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(2 * time.Millisecond)
		tt.Frame()
	}
}

// TestRenders draws the example headless: the heading, the three
// components and their controls.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Navigation",
		"Carousel", "Pagination", "File upload",
		"Slide 1", "Previous slide", "Next slide", "Go to slide 1",
		"Page 1 of 12", "Previous", "Next", "Go to page 1",
		"Upload a file", "Upload sample", "PDF, JPG, PNG, XLSX (max 8 MB)",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestCarousel flips to the second slide with the Next button.
func TestCarousel(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Next slide"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if view.State.Carousel.Active != 1 {
		t.Fatalf("Active = %d, want 1", view.State.Carousel.Active)
	}
	if !tt.HasText("Slide 2") {
		t.Fatalf("missing the second slide in %q", tt.Texts())
	}
}

// TestPagination steps to a page through its cell.
func TestPagination(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Go to page 5"); err != nil {
		t.Fatal(err)
	}
	if view.State.Page != 5 {
		t.Fatalf("Page = %d, want 5", view.State.Page)
	}
	if !tt.HasText("Page 5 of 12") {
		t.Fatalf("missing the page line in %q", tt.Texts())
	}
}

// TestUpload starts a sample upload with the button: the zone turns
// busy with the file.
func TestUpload(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Upload sample"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if view.State.Upload.Phase != fileupload.Uploading {
		t.Fatalf("Phase = %v, want Uploading", view.State.Upload.Phase)
	}
	if !tt.HasText("sales-report.xlsx") {
		t.Fatalf("missing the file in %q", tt.Texts())
	}
}

// TestUploadCompletes runs the sample upload to the end and shows the
// name the completion reported.
func TestUploadCompletes(t *testing.T) {
	fileupload.UploadDuration, fileupload.CompleteHold = 8*time.Millisecond, 4*time.Millisecond
	t.Cleanup(func() {
		fileupload.UploadDuration, fileupload.CompleteHold = 2200*time.Millisecond, 1600*time.Millisecond
	})
	tt := newTester(t)
	if err := tt.Click("Upload sample"); err != nil {
		t.Fatal(err)
	}
	pump(tt, 12)
	if view.State.Uploaded != "sales-report.xlsx" {
		t.Fatalf("Uploaded = %q, want the file name", view.State.Uploaded)
	}
	if !tt.HasText("sales-report.xlsx uploaded") {
		t.Fatalf("missing the completion in %q", tt.Texts())
	}
}

// TestDark draws the example under the dark appearance.
func TestDark(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	for _, s := range []string{"Navigation", "Carousel", "Upload a file"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in the dark frame %q", s, tt.Texts())
		}
	}
}
