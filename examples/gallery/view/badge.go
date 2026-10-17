package view

import "github.com/egoist/mygo/ui"

// libraryColor returns the color that marks a library in the gallery:
// the accent for uix's own components, the success for mygo's
// primitives, the warning for onevra's, the danger for MujicaUI's.
func libraryColor(t *ui.Theme, lib Library) ui.Color {
	switch lib {
	case Mygo:
		return t.Success
	case Onevra:
		return t.Warning
	case Mujica:
		return t.Danger
	}
	return t.Accent
}

// libraryName is the label of a library, shown in its badge.
func libraryName(lib Library) string {
	switch lib {
	case Mygo:
		return "mygo"
	case Onevra:
		return "onevra"
	case Mujica:
		return "mujica"
	}
	return "uix"
}

// libraryBadge is the small pill that marks which library a component
// comes from: its name in the library's color on a tinted face.
func libraryBadge(c *ui.Context, t *ui.Theme, lib Library) {
	col := libraryColor(t, lib)
	ui.Box(c).Padding(t.Space(0.75), t.Space(2)).Radius(t.Space(1)).
		Background(col.Alpha(0.14)).Border(1, col.Alpha(0.35)).Children(func() {
		ui.Text(c, libraryName(lib)).FontSize(t.FontSize * 0.75).FontWeight(600).TextColor(col).SingleLine()
	})
}
