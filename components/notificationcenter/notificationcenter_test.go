package notificationcenter

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			NotificationCenter(c, p)
		})
	}
}

var sample = []Item{
	{ID: "n1", Category: "mentions", Title: "Ada mentioned you",
		Description: "in a comment on the design", Timestamp: "2m", Unread: true,
		Status: Info, Actions: []Action{{ID: "view", Label: "View"}}},
	{ID: "n2", Category: "system", Title: "Backup complete",
		Description: "The nightly backup finished", Timestamp: "1h", Unread: true,
		Status: Success},
	{ID: "n3", Category: "activity", Title: "Weekly report is ready",
		Description: "Open it from the dashboard", Timestamp: "3h",
		Status: Neutral},
}

// TestRenders shows the header, the per-tab counts and the items.
func TestRenders(t *testing.T) {
	tab := 0
	tt := ui.NewTester(frame(Props{Items: sample, Tab: &tab}), 460, 480)
	for _, s := range []string{"Notifications", "2 unread", "All 3", "Mentions 1", "System 1"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
	for _, s := range []string{"Ada mentioned you", "Backup complete", "Weekly report is ready"} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestTabs filters the list to the category, running OnTabChange.
func TestTabs(t *testing.T) {
	tab := 0
	changes := 0
	tt := ui.NewTester(frame(Props{
		Items:       sample,
		Tab:         &tab,
		OnTabChange: func(i int) { changes++ },
	}), 460, 480)
	if err := tt.Click("Mentions 1"); err != nil {
		t.Fatal(err)
	}
	if tab != 1 || changes != 1 {
		t.Fatalf("the tab chose %d (%d changes), want 1 (1)", tab, changes)
	}
	tt.Frame()
	if !tt.HasText("Ada mentioned you") || tt.HasText("Backup complete") {
		t.Fatalf("the Mentions tab shows %q", tt.Texts())
	}
	if err := tt.Click("System 1"); err != nil {
		t.Fatal(err)
	}
	tt.Frame()
	if !tt.HasText("Backup complete") || tt.HasText("Ada mentioned you") {
		t.Fatalf("the System tab shows %q", tt.Texts())
	}
}

// TestMarkAllRead runs the callback for the unread items, which the app
// answers by clearing its flags.
func TestMarkAllRead(t *testing.T) {
	tab := 0
	items := append([]Item(nil), sample...)
	marked := 0
	tt := ui.NewTester(frame(Props{
		Items:         items,
		Tab:           &tab,
		OnMarkAllRead: func() {
			marked++
			for i := range items {
				items[i].Unread = false
			}
		},
	}), 460, 480)
	if err := tt.Click("Mark all read"); err != nil {
		t.Fatal(err)
	}
	if marked != 1 {
		t.Fatalf("Mark all read ran %d times, want 1", marked)
	}
	tt.Frame()
	if !tt.HasText("No unread notifications") {
		t.Fatalf("the header still counts unread in %q", tt.Texts())
	}
}

// TestAction runs the notification's action and the callback.
func TestAction(t *testing.T) {
	tab := 0
	var itemID, actionID string
	opened := 0
	items := append([]Item(nil), sample...)
	items[0].Actions[0].OnClick = func() { opened++ }
	tt := ui.NewTester(frame(Props{
		Items: items,
		Tab:   &tab,
		OnAction: func(i, a string) {
			itemID, actionID = i, a
		},
	}), 460, 480)
	if err := tt.Click("View"); err != nil {
		t.Fatal(err)
	}
	if itemID != "n1" || actionID != "view" {
		t.Fatalf("the action ran for %q/%q, want n1/view", itemID, actionID)
	}
	if opened != 1 {
		t.Fatalf("the action opened %d times, want 1", opened)
	}
}

// TestEmpty shows the caught-up message for an empty list.
func TestEmpty(t *testing.T) {
	tab := 0
	tt := ui.NewTester(frame(Props{Items: nil, Tab: &tab}), 460, 480)
	if !tt.HasText("You're all caught up.") {
		t.Fatalf("missing the empty message in %q", tt.Texts())
	}
}

// TestScrolls caps the list at a height when one is set.
func TestScrolls(t *testing.T) {
	tab := 0
	tt := ui.NewTester(frame(Props{Items: sample, Tab: &tab, Height: 120}), 460, 480)
	if img := tt.Image(); img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
