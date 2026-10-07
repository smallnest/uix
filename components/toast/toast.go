// Package toast provides shadcn/ui-style toasts for the MyGo native
// toolkit: a viewport to draw them in, and a way to push them.
//
//	toast.Viewport(c)   // once in the view
//	if b.Clicked() {
//		toast.Push(c, toast.Props{Title: "Saved", Type: toast.Success})
//	}
package toast

import (
	"time"

	"github.com/egoist/mygo/ui"
)

// Type names the kind of a toast, which the viewport colors.
type Type string

const (
	// Info is a neutral note, in the accent color.
	Info Type = "info"
	// Success is a good outcome, in the success color.
	Success Type = "success"
	// Warning is a caution, in the warning color.
	Warning Type = "warning"
	// Error is a failure, in the danger color.
	Error Type = "error"
)

// Props describes a toast to push.
type Props struct {
	// Title is the one-line summary.
	Title string
	// Description is the detail under it.
	Description string
	// Type colors the toast; Info when empty.
	Type Type
	// Action labels a button that runs OnAction, as "Undo".
	Action string
	// OnAction runs for the action button, which also closes the toast.
	OnAction func()
	// Timeout is how long the toast shows: 4 seconds when zero.
	Timeout time.Duration
}

// Push shows a toast of the given props, at the bottom right of the
// window.
func Push(c *ui.Context, p Props) {
	typ := p.Type
	if typ == "" {
		typ = Info
	}
	c.AddToast(ui.Toast{
		Title:       p.Title,
		Description: p.Description,
		Type:        string(typ),
		Action:      p.Action,
		OnAction:    p.OnAction,
		Timeout:     p.Timeout,
	})
}

// Viewport draws the toasts of the window, bottom right, styled by the
// theme. Call it once in the view, wherever; MyGo shows one viewport.
func Viewport(c *ui.Context) *ui.Element {
	t := c.Theme()
	return ui.ToastViewportBase(c, func(viewport *ui.Element, toasts []ui.Toast) {
		viewport.Padding(16).AlignItems(ui.End).Gap(t.Space(2))
		for _, ts := range toasts {
			tp := ui.ToastBase(c, ts)
			root := tp.Root
			root.Row().Gap(t.Space(3)).Padding(t.Space(2), t.Space(3)).Radius(t.Radius).
				Background(t.Background).Border(1, t.Border).MaxWidth(t.Space(72))
			root.Shadow(0, 6, 20, 0, ui.RGBA(0, 0, 0, 0.18))
			root.Children(func() {
				ui.Box(c).Size(t.Space(2.5), t.Space(2.5)).Radius(t.Space(1.25)).Background(toastColor(t, ts.Type)).Shrink(0)
				ui.Column(c).Gap(t.Space(0.5)).Grow(1).Children(func() {
					ui.Text(c, ts.Title).Bold().SingleLine()
					if ts.Description != "" {
						ui.Text(c, ts.Description).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
					}
				})
				if ts.Action != "" {
					tp.ActionButton().Label(ts.Action).Children(func() { ui.Text(c, ts.Action).TextColor(t.Accent).Bold() })
				}
				tp.CloseButton().Label("Close").Children(func() { ui.Text(c, "✕").TextColor(t.TextMuted) })
			})
		}
	})
}

// toastColor returns the dot color of a toast type.
func toastColor(t *ui.Theme, typ string) ui.Color {
	switch typ {
	case "success":
		return t.Success
	case "warning":
		return t.Warning
	case "error":
		return t.Danger
	}
	return t.Accent
}
