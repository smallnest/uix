// Package tokens provides the palette a uix app installs, the way the CSS
// variables of shadcn/ui give a web app its design tokens.
//
// The MyGo theme already names its colors semantically: Background,
// Surface, Text, Accent, Danger, Border and so on. This package curates
// their values into one palette. Apply it at the top of a view:
//
//	c.SetTheme(tokens.For(c.Theme().Dark))
//
// The mapping to the shadcn/ui tokens is:
//
//	background   -> Background    primary     -> Accent
//	foreground   -> Text          primary-fg  -> AccentText
//	secondary    -> Surface       destructive -> Danger
//	muted        -> SurfaceHover  border      -> Border
//	muted-fg     -> TextMuted     ring        -> Focus
//
// uix theme set regenerates this file from the palette it names.
package tokens

import "github.com/egoist/mygo/ui"

// Light returns the light appearance with the indigo palette. It
// starts from the MyGo light theme, which keeps the system-tuned
// neutrals, fonts and metrics, and overrides the values the palette
// names. Change a value here and every component restyles at once.
func Light() *ui.Theme {
	t := ui.LightTheme()
	t.Radius = 8
	t.Accent = ui.Hex("#6366f1")
	t.AccentHover = ui.Hex("#4f46e5")
	t.AccentPressed = ui.Hex("#4338ca")
	t.AccentText = ui.Hex("#ffffff")
	t.Danger = ui.Hex("#e11d48")
	t.Selection = ui.RGBA(99, 102, 241, 0.25)
	t.Focus = ui.RGBA(99, 102, 241, 0.55)
	return t
}

// Dark returns the dark appearance with the indigo palette.
func Dark() *ui.Theme {
	t := ui.DarkTheme()
	t.Radius = 8
	t.Accent = ui.Hex("#818cf8")
	t.AccentHover = ui.Hex("#6366f1")
	t.AccentPressed = ui.Hex("#4f46e5")
	t.AccentText = ui.Hex("#ffffff")
	t.Danger = ui.Hex("#f43f5e")
	t.Selection = ui.RGBA(129, 140, 248, 0.4)
	t.Focus = ui.RGBA(129, 140, 248, 0.6)
	return t
}

// For returns Light or Dark after the given appearance.
func For(dark bool) *ui.Theme {
	if dark {
		return Dark()
	}
	return Light()
}
