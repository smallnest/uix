package theme

import (
	"bytes"
	_ "embed"
	"fmt"
	"strconv"
	"text/template"
)

//go:embed tokens.tmpl
var tokensTmplText string

var tokensTmpl = template.Must(template.New("tokens").Parse(tokensTmplText))

// view is the data the template fills in. Selection and Focus arrive
// already rendered as ui.RGBA expressions, so the template stays a
// straightforward copy of the tokens package.
type view struct {
	Name   string
	Radius float32
	Light  scaleView
	Dark   scaleView
}

// scaleView is the rendered scale of one appearance.
type scaleView struct {
	Accent, AccentHover, AccentPressed, AccentText, Danger string
	Selection, Focus                                        string
}

// Generate renders the tokens package for the palette.
func Generate(p *Palette) ([]byte, error) {
	v, err := p.view()
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := tokensTmpl.Execute(&buf, v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// view builds the template data of the palette.
func (p *Palette) view() (view, error) {
	ls, err := rgbaExpr(p.Light.Accent, lightSelectionAlpha)
	if err != nil {
		return view{}, err
	}
	lf, err := rgbaExpr(p.Light.Accent, lightFocusAlpha)
	if err != nil {
		return view{}, err
	}
	ds, err := rgbaExpr(p.Dark.Accent, darkSelectionAlpha)
	if err != nil {
		return view{}, err
	}
	df, err := rgbaExpr(p.Dark.Accent, darkFocusAlpha)
	if err != nil {
		return view{}, err
	}
	return view{
		Name:   p.Name,
		Radius: p.Radius,
		Light: scaleView{
			Accent: p.Light.Accent, AccentHover: p.Light.AccentHover,
			AccentPressed: p.Light.AccentPressed, AccentText: p.Light.AccentText,
			Danger: p.Light.Danger, Selection: ls, Focus: lf,
		},
		Dark: scaleView{
			Accent: p.Dark.Accent, AccentHover: p.Dark.AccentHover,
			AccentPressed: p.Dark.AccentPressed, AccentText: p.Dark.AccentText,
			Danger: p.Dark.Danger, Selection: ds, Focus: df,
		},
	}, nil
}

// rgbaExpr renders a hex color at an alpha as the ui.RGBA expression the
// tokens write, so the template needs no color math of its own.
func rgbaExpr(hex string, alpha float32) (string, error) {
	r, g, b, err := parseHex(hex)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("ui.RGBA(%d, %d, %d, %v)", r, g, b, alpha), nil
}

// parseHex reads a hex color like "#6366f1" into its channels. Why not
// a color library: the standard library has no hex parser, and eight
// lines beat a new dependency for the binary.
func parseHex(s string) (r, g, b uint8, err error) {
	if len(s) != 7 || s[0] != '#' {
		return 0, 0, 0, fmt.Errorf("hex color %q: want #rrggbb", s)
	}
	n, err := strconv.ParseUint(s[1:], 16, 24)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("hex color %q: %v", s, err)
	}
	return uint8(n >> 16), uint8(n >> 8), uint8(n), nil
}
