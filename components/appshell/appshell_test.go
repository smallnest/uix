package appshell

import (
	"testing"

	"github.com/egoist/mygo/ui"

	"github.com/smallnest/uix/components/sidebar"
)

var sections = []sidebar.Section{
	{Title: "General", Items: []sidebar.Item{
		{ID: "chat", Label: "Chat"},
		{ID: "dashboard", Label: "Dashboard"},
	}},
	{Title: "Account", Items: []sidebar.Item{
		{ID: "settings", Label: "Settings"},
	}},
}

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			AppShell(c, p)
		})
	}
}

// TestRenders shows the trail, the heading and the sidebar.
func TestRenders(t *testing.T) {
	selected := "chat"
	tt := ui.NewTester(frame(Props{
		Title:    "Dashboard",
		Team:     "Board team",
		Member:   "Mertcan",
		Selected: &selected,
		Sections: sections,
		Content:  func(c *ui.Context) { ui.Text(c, "Here is the page.") },
	}), 700, 420)
	for _, s := range []string{"Board team", "Mertcan", "Dashboard", "Chat", "Settings", "Here is the page."} {
		if !tt.HasText(s) {
			t.Fatalf("missing %q in %q", s, tt.Texts())
		}
	}
}

// TestChoosesSidebar picks a sidebar item, as the shell passes through.
func TestChoosesSidebar(t *testing.T) {
	selected := "chat"
	tt := ui.NewTester(frame(Props{
		Title:    "Dashboard",
		Selected: &selected,
		Sections: sections,
	}), 700, 420)
	if err := tt.Click("Settings"); err != nil {
		t.Fatal(err)
	}
	if selected != "settings" {
		t.Fatalf("the sidebar chose %q, want settings", selected)
	}
}

// TestActions draws the header actions on the right.
func TestActions(t *testing.T) {
	selected := "chat"
	tt := ui.NewTester(frame(Props{
		Title:    "Dashboard",
		Selected: &selected,
		Sections: sections,
		Actions: func(c *ui.Context) {
			ui.Text(c, "Create ticket")
		},
	}), 700, 420)
	if !tt.HasText("Create ticket") {
		t.Fatalf("missing the action in %q", tt.Texts())
	}
}

// TestEmptyTrail leaves out the ancestors when none are given, and the
// title heads the page alone.
func TestEmptyTrail(t *testing.T) {
	selected := "chat"
	tt := ui.NewTester(frame(Props{
		Title:    "Chat",
		Selected: &selected,
		Sections: sections,
	}), 700, 420)
	if !tt.HasText("Chat") {
		t.Fatalf("missing the page title in %q", tt.Texts())
	}
}
