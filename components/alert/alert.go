// Package alert provides a shadcn/ui-style Alert: a colored notice with
// a bold title and an optional description, in the accent, success,
// warning or danger color.
//
//	alert.Alert(c, alert.Props{
//		Title:       "Heads up!",
//		Description: "You can add components to your app.",
//		Variant:     alert.Info,
//	})
package alert

import "github.com/egoist/mygo/ui"

// Variant is the color of an alert.
type Variant int

const (
	// Info is the default notice, in the accent color.
	Info Variant = iota
	// Success announces success, in the success color.
	Success
	// Warning announces a warning, in the warning color.
	Warning
	// Danger announces a failure, in the danger color.
	Danger
)

// Props describes the alert to draw.
type Props struct {
	// Title is the bold heading of the alert.
	Title string
	// Description is the text under the title.
	Description string
	// Variant is the color of the alert; Info unless set.
	Variant Variant
}

// Alert draws a rounded notice: a tinted face, a colored border and a
// colored bold title. It returns the notice, so a view can size it.
func Alert(c *ui.Context, p Props) ui.Element {
	t := c.Theme()
	col := variantColor(t, p.Variant)
	e := ui.Column(c).Radius(t.Radius).Clip().Padding(t.Space(3)).Gap(t.Space(1)).
		Background(ui.RGBA(col.R, col.G, col.B, 0.1)).
		Border(1, ui.RGBA(col.R, col.G, col.B, 0.4))
	e.Children(func() {
		if p.Title != "" {
			ui.Text(c, p.Title).Bold().TextColor(col)
		}
		if p.Description != "" {
			ui.Text(c, p.Description).FontSize(t.FontSize * 0.875).TextColor(t.TextMuted)
		}
	})
	return e
}

// variantColor returns the color of a variant: the accent for Info.
func variantColor(t *ui.Theme, v Variant) ui.Color {
	switch v {
	case Success:
		return t.Success
	case Warning:
		return t.Warning
	case Danger:
		return t.Danger
	}
	return t.Accent
}
