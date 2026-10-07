package tooltip

import (
	"testing"
	"time"

	"github.com/egoist/mygo/ui"
)

// frame draws a Save button with the tooltip of the props on it.
func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			b := ui.Button(c, "Save")
			p.Anchor = b
			Tooltip(c, p)
		})
	}
}

// hover rests the pointer on the anchor for longer than the tooltip
// delay, and draws the frame the tip would show in.
func hover(t *testing.T, tt *ui.Tester, s string) {
	t.Helper()
	r, ok := tt.Find(s)
	if !ok {
		t.Fatalf("no %q in %q", s, tt.Texts())
	}
	tt.Move(r.X+r.W/2, r.Y+r.H/2)
	tt.Frame()
	time.Sleep(650 * time.Millisecond)
	tt.Frame()
}

// TestShows draws the tip once the pointer has rested on the anchor.
func TestShows(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Saves the file"}), 300, 100)
	hover(t, tt, "Save")
	if !tt.HasText("Saves the file") {
		t.Fatal("the tip did not show")
	}
}

// TestHidden keeps the tip away while the pointer has not rested.
func TestHidden(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Saves the file"}), 300, 100)
	if tt.HasText("Saves the file") {
		t.Fatal("the tip shows without a rest")
	}
}

// TestDisabled keeps the tip away even as the pointer rests.
func TestDisabled(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Saves the file", Disabled: true}), 300, 100)
	hover(t, tt, "Save")
	if tt.HasText("Saves the file") {
		t.Fatal("a disabled tip shows")
	}
}

// TestDarkMode draws the tip under the dark appearance too.
func TestDarkMode(t *testing.T) {
	tt := ui.NewTester(frame(Props{Text: "Saves the file"}), 300, 100)
	tt.SetDark(true)
	hover(t, tt, "Save")
	if !tt.HasText("Saves the file") {
		t.Fatal("tip missing under dark mode")
	}
}
