// Package theme holds the palettes a uix app can pick for its design
// tokens, and regenerates the tokens package from them, as the uix theme
// command does. A palette names only the colors the tokens override; the
// neutrals, fonts and metrics still come from the MyGo theme, so a
// palette keeps the system-tuned look.
package theme

// Palette is one color scheme for the uix tokens.
type Palette struct {
	// Name is the argument of uix theme set.
	Name string
	// Description is what the palette is for, shown by uix theme list.
	Description string
	// Radius is the corner radius the tokens set, in points.
	Radius float32
	// Light and Dark are the colors of the two appearances.
	Light Scale
	Dark  Scale
}

// Scale is the color scale of one appearance. Every value is a hex color
// like "#6366f1".
type Scale struct {
	Accent        string
	AccentHover   string
	AccentPressed string
	AccentText    string
	Danger        string
}

// The tokens derive Selection and Focus from the accent at a fixed
// alpha, so a palette does not name them. The alphas follow the
// hand-written tokens: a soft selection and a stronger focus ring,
// stronger in the dark appearance where the accent is lighter.
const (
	lightSelectionAlpha float32 = 0.25
	lightFocusAlpha     float32 = 0.55
	darkSelectionAlpha  float32 = 0.4
	darkFocusAlpha      float32 = 0.6
)
