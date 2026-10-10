package inputotp

import (
	"testing"

	"github.com/egoist/mygo/ui"
)

// otp draws the code field in a padded column and returns the state the
// tests keep.
func otp(st *State, onComplete func(string)) func(c *ui.Context) {
	return func(c *ui.Context) {
		ui.Column(c).Fill().Padding(16).Children(func() {
			OTP(c, Props{State: st, Length: 6, OnComplete: onComplete})
		})
	}
}

// focus clicks the first cell, which gives the field the focus.
func focus(tt *ui.Tester) {
	tt.ClickAt(40, 40)
}

// TestTypes fills the cells with what is typed, one character each, and
// runs OnComplete when the last fills.
func TestTypes(t *testing.T) {
	st := &State{}
	done := ""
	tt := ui.NewTester(otp(st, func(code string) { done = code }), 420, 120)
	focus(tt)
	tt.Type("123456")
	if st.Code != "123456" {
		t.Fatalf("Code = %q, want 123456", st.Code)
	}
	if done != "123456" {
		t.Fatalf("OnComplete = %q, want 123456", done)
	}
}

// TestStopsAtLength types past the last cell without growing the code.
func TestStopsAtLength(t *testing.T) {
	st := &State{}
	tt := ui.NewTester(otp(st, nil), 420, 120)
	focus(tt)
	tt.Type("1234567890")
	if st.Code != "123456" {
		t.Fatalf("Code = %q, want 123456", st.Code)
	}
}

// TestBackspace takes the character before the caret: after typing two,
// one Backspace leaves the first.
func TestBackspace(t *testing.T) {
	st := &State{}
	tt := ui.NewTester(otp(st, nil), 420, 120)
	focus(tt)
	tt.Type("12")
	tt.Key(0, ui.KeyBackspace)
	if st.Code != "1" {
		t.Fatalf("Code = %q, want 1", st.Code)
	}
}

// TestArrows move the caret: Left past the typed code, then typing
// replaces the character there.
func TestArrows(t *testing.T) {
	st := &State{}
	tt := ui.NewTester(otp(st, nil), 420, 120)
	focus(tt)
	tt.Type("12")
	tt.Key(0, ui.KeyLeft)
	tt.Type("9")
	if st.Code != "19" {
		t.Fatalf("Code = %q, want 19", st.Code)
	}
}
