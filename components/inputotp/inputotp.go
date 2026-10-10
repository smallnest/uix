// Package inputotp provides a shadcn/ui-style input for a one-time
// password, or any short code: a row of cells, one character each, that
// typing fills in turn, as a verification code a message brings. The
// arrows move the caret among the cells, Backspace takes the character
// before it, and pasting a longer code fills the cells it reaches.
//
//	inputotp.OTP(c, inputotp.Props{
//		State:      &code,
//		Length:     6,
//		OnComplete: func(code string) { app.verify(code) },
//	})
package inputotp

import "github.com/egoist/mygo/ui"

// State is where the code is, kept by the app.
type State struct {
	// Code is the code typed so far, one character a cell.
	Code string
	// cursor is the cell the typing fills next, -1 for none.
	cursor int
}

// Props describes the code field to draw.
type Props struct {
	// State is where the code is, for the app to keep.
	State *State
	// Length is the number of cells; 6 when zero.
	Length int
	// Disabled keeps the cells from being edited.
	Disabled bool
	// OnComplete runs when the last cell fills, with the code; Enter in
	// a full field runs it too.
	OnComplete func(code string)
}

// OTP draws the cells of the code and returns them, so a view can chain
// more calls on it. The cells are one stop of Tab; the caret cell reads
// on the accent while the field has the focus.
func OTP(c *ui.Context, p Props) *ui.Element {
	t := c.Theme()
	st := p.State
	if st == nil {
		st = &State{}
	}
	n := p.Length
	if n <= 0 {
		n = 6
	}
	// The code and the caret within the cells, as the last frame left
	// them or the app preset them.
	if len(st.Code) > n {
		st.Code = st.Code[:n]
	}
	if st.cursor < 0 || st.cursor > len(st.Code) {
		st.cursor = len(st.Code)
	}
	if st.cursor > n {
		st.cursor = n
	}

	cell := t.Space(10)
	gap := t.Space(1.5)
	row := ui.Row(c).Gap(gap).Focusable()
	if p.Disabled {
		row.Disabled(true)
	}
	row.HandleInput(func(ev ui.InputEvent) bool {
		if p.Disabled {
			return false
		}
		switch ev.Kind {
		case ui.InputText:
			took := false
			for _, r := range ev.Text {
				if st.cursor >= n {
					break
				}
				if st.cursor == len(st.Code) {
					st.Code += string(r)
				} else {
					st.Code = st.Code[:st.cursor] + string(r) + st.Code[st.cursor+1:]
				}
				st.cursor++
				took = true
			}
			if took && len(st.Code) == n && p.OnComplete != nil {
				p.OnComplete(st.Code)
			}
			return took
		case ui.InputKeyDown:
			switch ev.Key {
			case ui.KeyBackspace:
				if st.cursor > 0 {
					st.cursor--
					if st.cursor < len(st.Code) {
						st.Code = st.Code[:st.cursor] + st.Code[st.cursor+1:]
					}
				}
				return true
			case ui.KeyLeft:
				if st.cursor > 0 {
					st.cursor--
				}
				return true
			case ui.KeyRight:
				if st.cursor < len(st.Code) {
					st.cursor++
				}
				return true
			case ui.KeyEnter:
				if len(st.Code) == n && p.OnComplete != nil {
					p.OnComplete(st.Code)
				}
				return true
			}
		case ui.InputPointerDown:
			st.cursor = int(ev.X) / int(cell+gap)
			if st.cursor > n-1 {
				st.cursor = n - 1
			}
			if st.cursor > len(st.Code) {
				st.cursor = len(st.Code)
			}
			row.Focus()
			return true
		}
		return false
	})
	// The caret of the focused cell, where the input methods show their
	// candidates.
	row.TextCaret(ui.Rect{X: float32(st.cursor) * (cell + gap), Y: 0, W: cell, H: cell})
	row.Children(func() {
		for i := 0; i < n; i++ {
			ch := ""
			if i < len(st.Code) {
				ch = string(st.Code[i])
			}
			face := ui.Box(c).Size(cell, cell).Radius(t.Radius).Background(t.Surface).
				Border(1, t.Border).Center()
			if i == st.cursor && row.Focused() {
				face.Border(2, t.Accent)
			}
			face.Children(func() {
				if ch != "" {
					ui.Text(c, ch).FontSize(t.FontSize * 1.25)
				}
			})
		}
	})
	return row
}
