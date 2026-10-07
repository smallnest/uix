package textarea

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

func frame(p Props) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Gap(8).Children(func() {
			Textarea(c, p)
		})
	}
}

// The value of a text area is drawn by its editor, not by a text node,
// so these tests check that a frame renders and that a field of many
// lines comes through.
func TestRenders(t *testing.T) {
	value := "line one\nline two\nline three"
	tt := ui.NewTester(frame(Props{Value: &value, Placeholder: "notes"}), 240, 200)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}

// TestRows draws the field at two heights and at the default.
func TestRows(t *testing.T) {
	for _, rows := range []int{0, 2, 8} {
		value := ""
		tt := ui.NewTester(frame(Props{Value: &value, Rows: rows}), 240, 240)
		img := tt.Image()
		if img == nil || img.Bounds().Dy() == 0 {
			t.Fatalf("rows %d: no frame rendered", rows)
		}
	}
}

// TestTypes focuses the field with a press and types into it: the value
// the view shares with the editor changes, as it does in a window.
func TestTypes(t *testing.T) {
	value := ""
	tt := ui.NewTester(frame(Props{Value: &value, Placeholder: "notes"}), 240, 200)
	tt.Press(20, 25)
	tt.Release(20, 25)
	tt.Type("hello")
	if value != "hello" {
		t.Fatalf("value %q, want hello", value)
	}
}

// TestDisabled keeps the value of a field the user cannot edit.
func TestDisabled(t *testing.T) {
	value := "locked"
	tt := ui.NewTester(frame(Props{Value: &value, Disabled: true}), 240, 200)
	tt.Press(20, 25)
	tt.Release(20, 25)
	tt.Type("x")
	if value != "locked" {
		t.Fatalf("disabled field became %q, want locked", value)
	}
}

// TestDarkMode draws the field under the dark appearance too.
func TestDarkMode(t *testing.T) {
	value := "line one\nline two"
	tt := ui.NewTester(frame(Props{Value: &value}), 240, 200)
	tt.SetDark(true)
	img := tt.Image()
	if img == nil || img.Bounds().Dx() == 0 {
		t.Fatal("no frame rendered")
	}
}
