package main

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/examples/profile/view"
)

// reset puts the example in its initial state, so the tests are
// independent of each other.
func reset() { view.State.Following = false }

// TestRenders draws the profile page headless and checks that the card,
// the avatar, the divider and the alerts came through: MyGo renders
// frames in memory, no window needed.
func TestRenders(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ProfileView, view.Width, view.Height)
	for _, s := range []string{
		"Profile", "Ada Lovelace", "@ada", "Follow", "Edit",
		"Heads up!", "Beta", "Danger", "This action cannot be undone.",
	} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q", s)
		}
	}
	if tt.HasText("Followed!") {
		t.Fatal("the success alert shows before the follow click")
	}
}

// TestFollows clicks the follow button, which swaps its label and shows
// a success alert; a second click undoes both.
func TestFollows(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ProfileView, view.Width, view.Height)
	if err := tt.Click("Follow"); err != nil {
		t.Fatal(err)
	}
	if !view.State.Following {
		t.Fatal("click did not set the following state")
	}
	if !tt.HasText("Followed!") || !tt.HasText("You now follow Ada Lovelace.") {
		t.Fatal("the success alert is missing after the follow click")
	}
	if err := tt.Click("Following"); err != nil {
		t.Fatal(err)
	}
	if view.State.Following {
		t.Fatal("second click did not clear the following state")
	}
	if tt.HasText("Followed!") {
		t.Fatal("the success alert still shows after the unfollow click")
	}
}

// TestDarkMode draws the profile page under the dark appearance too.
func TestDarkMode(t *testing.T) {
	reset()
	tt := ui.NewTester(view.ProfileView, view.Width, view.Height)
	tt.SetDark(true)
	if !tt.HasText("Profile") {
		t.Fatal("dark profile page missing the title")
	}
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
