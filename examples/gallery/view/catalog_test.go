package view

import "testing"

// TestCatalogCompleteness makes sure every entry carries an icon and a
// demo, so the sidebar draws every item with one and every page has
// something to show.
func TestCatalogCompleteness(t *testing.T) {
	for _, s := range catalog {
		for _, e := range s.Entries {
			if e.Icon == nil {
				t.Errorf("%s (%s) has no icon", e.Name, e.ID)
			}
			if e.Build == nil {
				t.Errorf("%s (%s) has no demo", e.Name, e.ID)
			}
			if e.Name == "" || e.ID == "" {
				t.Errorf("an entry of %q lacks a name or an ID", s.Title)
			}
		}
	}
}
