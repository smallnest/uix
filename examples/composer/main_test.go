package main

import (
	"slices"
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/composer/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.ComposerView, view.Width, view.Height)
}

// TestRenders draws the composer headless: the title, the fields with
// their captions, the calendar, the reminder group, and the status line.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"New appointment", "Task", "Tags", "Due", "Remind",
		"October 2026", "Remind me by", "Mail", "Calendar", "Bounce",
		"Type a title, or pick a due day.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestCompletes completes the title from the tasks known: typing the
// start shows the task, and the arrows and Enter take it.
func TestCompletes(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Task title"); err != nil {
		t.Fatal(err)
	}
	tt.Type("Wa")
	if !tt.HasText("Water the plants") {
		t.Fatalf("Wa suggests %q", tt.Texts())
	}
	tt.Key(0, ui.KeyDown)
	tt.Key(0, ui.KeyEnter)
	if view.State.Title != "Water the plants" {
		t.Fatalf("the completion took %q, want Water the plants", view.State.Title)
	}
}

// TestAdds adds tags typed, on a comma and on Enter, which the status
// names.
func TestAdds(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Add a tag"); err != nil {
		t.Fatal(err)
	}
	tt.Type("urgent,")
	tt.Type("later")
	tt.Key(0, ui.KeyEnter)
	if !slices.Equal(view.State.Tags, []string{"urgent", "later"}) {
		t.Fatalf("the tags are %q, want [urgent later]", view.State.Tags)
	}
	if !tt.HasText("Tags: urgent, later.") {
		t.Fatalf("the status does not name the tags: %q", tt.Texts())
	}
}

// TestCalendar picks a due day with a click, which the status names.
func TestCalendar(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("October 20, 2026"); err != nil {
		t.Fatal(err)
	}
	if view.State.Due.Day() != 20 || view.State.Due.Month() != 10 {
		t.Fatalf("the calendar chose %v, want Oct 20", view.State.Due)
	}
	want := "Due “" + view.State.Due.Format("Mon, Jan 2, 2006") + "”."
	if !tt.HasText(want) {
		t.Fatalf("the status does not name the due: %q", tt.Texts())
	}
}

// TestRemind checks one reminder, which the status confirms.
func TestRemind(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Mail"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Mail || view.State.Schedule {
		t.Fatalf("the box did not check only Mail: %v %v", view.State.Mail, view.State.Schedule)
	}
	if !tt.HasText("Reminders updated.") {
		t.Fatalf("the status does not name the reminders: %q", tt.Texts())
	}
}

// TestCreate confirms the appointment with the button in the footer.
func TestCreate(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Create"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("created") {
		t.Fatalf("the status does not confirm: %q", tt.Texts())
	}
}

// TestDarkMode draws the composer under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("New appointment") {
		t.Fatal("dark composer missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
