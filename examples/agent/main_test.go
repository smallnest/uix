package main

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/chat"
	"github.com/smallnest/uix/examples/agent/view"
)

// newTester resets the example and returns a tester of its window.
func newTester(t *testing.T) *ui.Tester {
	t.Helper()
	view.Reset()
	return ui.NewTester(view.AgentView, view.Width, view.Height)
}

// drive advances the demo's paces by running frames as time passes, and
// returns after it has run out of steps or frames.
func drive(tt *ui.Tester, done func() bool) {
	for i := 0; i < 80 && !done(); i++ {
		time.Sleep(3 * time.Millisecond)
		tt.Frame()
	}
}

// TestRenders draws the agent headless: the shell with the sidebar, the
// trail, the heading, and the empty chat with its suggestions.
func TestRenders(t *testing.T) {
	tt := newTester(t)
	for _, s := range []string{
		"Board team", "Mertcan", "Chat",
		"Agent", "Log", "Notifications",
		"What can I help with?",
		"Explain this starter", "Suggest a name for the app",
		"Write the product update",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestSwitchPages walks the sidebar: the log shows its first step at
// work, and the notifications page its items and tabs.
func TestSwitchPages(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Log"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("Planning the task") {
		t.Fatalf("the log does not work on its first step: %q", tt.Texts())
	}
	if err := tt.Click("Notifications"); err != nil {
		t.Fatal(err)
	}
	for _, s := range []string{"All 3", "Mentions 1", "System 1", "Ada mentioned you"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q on the notifications page in %q", s, tt.Texts())
		}
	}
}

// TestChatSubmits sends a suggestion: the chat shows the words of the
// user, and the demo model starts to answer.
func TestChatSubmits(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Explain this starter"); err != nil {
		t.Fatal(err)
	}
	if len(view.State.Messages) != 2 {
		t.Fatalf("the chat holds %d messages, want 2", len(view.State.Messages))
	}
	if view.State.Messages[0].Role != chat.User ||
		view.State.Messages[1].Role != chat.Assistant {
		t.Fatalf("the chat holds wrong roles: %+v", view.State.Messages)
	}
	if !view.State.Busy || !view.State.Thinking {
		t.Fatalf("the model is busy=%v thinking=%v, want true true",
			view.State.Busy, view.State.Thinking)
	}
	if !tt.HasText("Explain this starter") {
		t.Fatalf("the user's words are missing from %q", tt.Texts())
	}
}

// TestChatStreams delivers the demo model's reply word by word, until
// the answer is done.
func TestChatStreams(t *testing.T) {
	view.ThinkPace, view.StreamPace = time.Millisecond, time.Millisecond
	tt := newTester(t)
	if err := tt.Click("Write the product update"); err != nil {
		t.Fatal(err)
	}
	drive(tt, func() bool { return !view.State.Busy })
	if view.State.Busy {
		t.Fatal("the model still writes after the frames ran out")
	}
	last := view.State.Messages[len(view.State.Messages)-1]
	if len(last.Text) == 0 {
		t.Fatal("the model answered nothing")
	}
	if !tt.HasText("shipped three fixes") {
		t.Fatalf("the answer is missing from %q", tt.Texts())
	}
}

// TestChatStops cuts the model's answer off and keeps the words it
// wrote.
func TestChatStops(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Explain this starter"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Stop generating"); err != nil {
		t.Fatal(err)
	}
	if view.State.Busy || view.State.Thinking {
		t.Fatalf("the model is busy=%v thinking=%v after the stop, want false false",
			view.State.Busy, view.State.Thinking)
	}
	if len(view.State.Messages) != 2 {
		t.Fatalf("the stop dropped a message: %d", len(view.State.Messages))
	}
}

// TestLogReveals the log's steps one by one, until it finishes.
func TestLogReveals(t *testing.T) {
	view.LogPace = time.Millisecond
	tt := newTester(t)
	if err := tt.Click("Log"); err != nil {
		t.Fatal(err)
	}
	drive(tt, func() bool { return view.State.Revealed == 4 })
	if view.State.Revealed != 4 {
		t.Fatalf("the log revealed %d steps, want 4", view.State.Revealed)
	}
	tt.Frame()
	if !tt.HasText("Writing the summary") {
		t.Fatalf("the last step is missing from %q", tt.Texts())
	}
}

// TestNotifications reads all and filters by tab, as the page's
// controls work.
func TestNotifications(t *testing.T) {
	tt := newTester(t)
	if err := tt.Click("Notifications"); err != nil {
		t.Fatal(err)
	}
	if err := tt.Click("Mark all read"); err != nil {
		t.Fatal(err)
	}
	if !tt.HasText("No unread notifications") {
		t.Fatalf("the header still counts unread in %q", tt.Texts())
	}
	if err := tt.Click("Mentions 1"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if !tt.HasText("Ada mentioned you") || tt.HasText("Backup complete") {
		t.Fatalf("the Mentions tab shows %q", tt.Texts())
	}
}

// TestDarkMode draws the agent under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := newTester(t)
	tt.SetDark(true)
	if !tt.HasText("What can I help with?") {
		t.Fatal("dark agent missing the chat prompt")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
