package view

import (
	"github.com/egoist/mygo/plugins/loading"
	"github.com/egoist/mygo/ui"
)

// The loading page shows every indicator of MyGo's loading plugin: the
// 29 that loading.dev defines, each drawn by MyGo itself and named under
// it. The page reads the plugin's own All() rather than a list of its
// own, so an indicator the plugin adds appears here without an edit,
// and one it renames renames here too.

// The cell's geometry: the indicator's side, and the width of the box
// that holds it and its name, which keeps the wrapped rows in columns.
const (
	ldSize = 40
	ldCell = 96
)

// demoLoading lists the indicators.
func demoLoading(c *ui.Context, t *ui.Theme) {
	ui.Row(c).FillWidth().Wrap().Gap(t.Space(3)).AlignItems(ui.Center).Children(func() {
		for _, s := range loading.All() {
			loadingCell(c, t, s)
		}
	})
}

// loadingCell is one indicator with its name under it. The name is drawn
// as well as labelled because a glyph that moves tells a sighted reader
// nothing about what it is called, and because the page's own test finds
// each indicator by that text.
func loadingCell(c *ui.Context, t *ui.Theme, s loading.Spinner) ui.Element {
	return ui.Column(c).Width(ldCell).Center().Gap(t.Space(2)).Children(func() {
		// Each indicator is a progress indicator of unknown length to the
		// accessibility tree already; the label is what names it there.
		s.New(c, loading.Options{Size: ldSize, Color: t.Text}).Label(s.Name)
		ui.Text(c, s.Name).FontSize(t.FontSize * 0.8).TextColor(t.TextMuted)
	})
}
