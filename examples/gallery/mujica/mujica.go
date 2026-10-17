// Package mujica ports the MujicaUI gallery's demos into the uix
// gallery. The demos live in MujicaUI's examples as package main, so
// this package carries the files over and exposes them as a catalog.
// Every demo is self-contained: its state sits in ui.Local under the
// page ID, and it draws only through its (c, log) arguments, so the
// uix shell can render any page by calling core.Use and the demo.
package mujica

import (
	"github.com/egoist/mygo/ui"

	"git.zacharyzhang.com/ZacharyZhang-NY/MujicaUI/core"
)

// Locale is the built-in copy the MujicaUI demos draw with. The uix
// gallery speaks Chinese, so the demos do too; the hardcoded demo
// labels are Chinese either way.
const Locale = core.ZhCN

// Use applies the MujicaUI theme for the rest of the frame, as
// MujicaUI's own gallery does before it draws a page. MujicaUI
// components panic without it, so every ported page calls it first.
func Use(c *ui.Context) {
	core.Use(c, core.Settings{Locale: Locale})
}

// Pages returns every page of MujicaUI's gallery, in the order the
// gallery lists them: the catalog's 30 categories, IDs "001" to "500".
func Pages() []Page {
	return catalog()
}
