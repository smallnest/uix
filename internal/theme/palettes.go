package theme

import (
	"fmt"
	"strings"
)

// palettes are the color schemes uix theme offers. The colors are the
// values the Tailwind scales name, so a palette looks familiar; the
// tokens curate them into the MyGo theme. Add an entry here and uix
// theme list shows it.
var palettes = []Palette{
	{
		Name:        "indigo",
		Description: "the default uix accent",
		Radius:      8,
		Light: Scale{Accent: "#6366f1", AccentHover: "#4f46e5", AccentPressed: "#4338ca", AccentText: "#ffffff", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#818cf8", AccentHover: "#6366f1", AccentPressed: "#4f46e5", AccentText: "#ffffff", Danger: "#f43f5e"},
	},
	{
		Name:        "violet",
		Description: "a violet accent",
		Radius:      8,
		Light: Scale{Accent: "#8b5cf6", AccentHover: "#7c3aed", AccentPressed: "#6d28d9", AccentText: "#ffffff", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#a78bfa", AccentHover: "#8b5cf6", AccentPressed: "#7c3aed", AccentText: "#ffffff", Danger: "#f43f5e"},
	},
	{
		Name:        "sky",
		Description: "a sky-blue accent",
		Radius:      8,
		Light: Scale{Accent: "#0ea5e9", AccentHover: "#0284c7", AccentPressed: "#0369a1", AccentText: "#ffffff", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#38bdf8", AccentHover: "#0ea5e9", AccentPressed: "#0284c7", AccentText: "#ffffff", Danger: "#f43f5e"},
	},
	{
		Name:        "emerald",
		Description: "a green accent",
		Radius:      8,
		Light: Scale{Accent: "#10b981", AccentHover: "#059669", AccentPressed: "#047857", AccentText: "#ffffff", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#34d399", AccentHover: "#10b981", AccentPressed: "#059669", AccentText: "#ffffff", Danger: "#f43f5e"},
	},
	{
		Name:        "rose",
		Description: "a rose accent",
		Radius:      8,
		Light: Scale{Accent: "#f43f5e", AccentHover: "#e11d48", AccentPressed: "#be123c", AccentText: "#ffffff", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#fb7185", AccentHover: "#f43f5e", AccentPressed: "#e11d48", AccentText: "#ffffff", Danger: "#f43f5e"},
	},
	{
		Name:        "amber",
		Description: "a warm amber accent",
		Radius:      8,
		Light: Scale{Accent: "#f59e0b", AccentHover: "#d97706", AccentPressed: "#b45309", AccentText: "#422006", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#fbbf24", AccentHover: "#f59e0b", AccentPressed: "#d97706", AccentText: "#422006", Danger: "#f43f5e"},
	},
	{
		Name:        "slate",
		Description: "a restrained slate accent",
		Radius:      8,
		Light: Scale{Accent: "#334155", AccentHover: "#1e293b", AccentPressed: "#0f172a", AccentText: "#ffffff", Danger: "#e11d48"},
		Dark:  Scale{Accent: "#cbd5e1", AccentHover: "#94a3b8", AccentPressed: "#64748b", AccentText: "#0f172a", Danger: "#f43f5e"},
	},
}

// Palettes lists the palettes, in the order uix theme list shows them.
func Palettes() []Palette { return palettes }

// Find returns the palette with the given name.
func Find(name string) (*Palette, error) {
	for i := range palettes {
		if palettes[i].Name == name {
			return &palettes[i], nil
		}
	}
	return nil, fmt.Errorf("uix theme: no palette %q; available: %s", name, names())
}

// names lists the palette names, for errors and uix theme list.
func names() string {
	var ns []string
	for _, p := range palettes {
		ns = append(ns, p.Name)
	}
	return strings.Join(ns, ", ")
}
