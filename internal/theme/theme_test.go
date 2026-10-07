package theme

import (
	"bytes"
	"go/format"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// TestPalettesWellFormed checks that every palette has a name, a
// description and valid hex colors, and that the two appearances differ.
func TestPalettesWellFormed(t *testing.T) {
	seen := map[string]bool{}
	for _, p := range Palettes() {
		if p.Name == "" {
			t.Fatal("a palette has no name")
		}
		if seen[p.Name] {
			t.Fatalf("duplicate palette %q", p.Name)
		}
		seen[p.Name] = true
		if p.Description == "" {
			t.Fatalf("%s: no description", p.Name)
		}
		for _, s := range []Scale{p.Light, p.Dark} {
			for _, hex := range []string{s.Accent, s.AccentHover, s.AccentPressed, s.AccentText, s.Danger} {
				if _, _, _, err := parseHex(hex); err != nil {
					t.Fatalf("%s: %v", p.Name, err)
				}
			}
		}
		if p.Light.Accent == p.Dark.Accent {
			t.Fatalf("%s: same accent in both appearances", p.Name)
		}
	}
}

// TestFind finds a palette by name and rejects a missing one.
func TestFind(t *testing.T) {
	p, err := Find("emerald")
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != "emerald" {
		t.Fatalf("Find(emerald) = %q", p.Name)
	}
	if _, err := Find("nope"); err == nil {
		t.Fatal("Find accepted a missing palette")
	}
}

// TestGenerate renders a palette into the tokens package: the file
// carries the palette colors, the selection and focus derived from the
// accent, and it stays gofmt-clean.
func TestGenerate(t *testing.T) {
	p, err := Find("indigo")
	if err != nil {
		t.Fatal(err)
	}
	got, err := Generate(p)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{
		"the indigo palette",
		`t.Accent = ui.Hex("#6366f1")`,
		`t.Accent = ui.Hex("#818cf8")`,
		`t.Selection = ui.RGBA(99, 102, 241, 0.25)`,
		`t.Focus = ui.RGBA(99, 102, 241, 0.55)`,
		`t.Selection = ui.RGBA(129, 140, 248, 0.4)`,
		`t.Focus = ui.RGBA(129, 140, 248, 0.6)`,
	} {
		if !strings.Contains(string(got), want) {
			t.Fatalf("the generated tokens lack %q", want)
		}
	}
	if _, err := format.Source(got); err != nil {
		t.Fatalf("the generated tokens are not gofmt-clean: %v", err)
	}
}

// TestIndigoMatchesBase keeps the base set that uix init installs and
// the file uix theme set indigo writes the same, so the two never drift:
// regenerate the base with uix theme set indigo after editing the
// template or a palette.
func TestIndigoMatchesBase(t *testing.T) {
	p, err := Find("indigo")
	if err != nil {
		t.Fatal(err)
	}
	got, err := Generate(p)
	if err != nil {
		t.Fatal(err)
	}
	_, file, _, _ := runtime.Caller(0)
	base := filepath.Join(filepath.Dir(file), "..", "..", "components", "base", "tokens.go")
	want, err := os.ReadFile(base)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("components/base/tokens.go differs from uix theme set indigo; run uix theme set indigo to regenerate it")
	}
}

// TestCustom derives a palette from one accent color: the light
// appearance keeps the given accent and darkens it for hover and press,
// the dark appearance lightens it, the text color reads white or
// near-black by the luminance, and a very light accent keeps its color
// in the dark appearance.
func TestCustom(t *testing.T) {
	p, err := Custom("#6366f1")
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != CustomPalette.Name {
		t.Fatalf("Name = %q, want %q", p.Name, CustomPalette.Name)
	}
	if p.Light.Accent != "#6366f1" {
		t.Fatalf("Light.Accent = %q, want the given color", p.Light.Accent)
	}
	if p.Light.AccentHover == p.Light.Accent || p.Light.AccentPressed == p.Light.Accent {
		t.Fatal("a light step did not darken the accent")
	}
	if p.Light.AccentHover == p.Light.AccentPressed {
		t.Fatal("the light hover and press steps are equal")
	}
	if p.Dark.Accent == p.Light.Accent {
		t.Fatal("the dark accent did not lighten")
	}
	if p.Dark.AccentHover != p.Light.Accent {
		t.Fatalf("Dark.AccentHover = %q, want the given color", p.Dark.AccentHover)
	}
	if p.Dark.AccentPressed == p.Dark.Accent || p.Dark.AccentPressed == p.Dark.AccentHover {
		t.Fatal("a dark step did not darken the accent")
	}
	// A mid-tone accent takes white text; a light one takes near-black.
	if p.Light.AccentText != "#ffffff" {
		t.Fatalf("text on #6366f1 = %q, want white", p.Light.AccentText)
	}
	pAmber, err := Custom("#f59e0b")
	if err != nil {
		t.Fatal(err)
	}
	if pAmber.Light.AccentText != "#0f172a" {
		t.Fatalf("text on #f59e0b = %q, want near-black", pAmber.Light.AccentText)
	}
	// A very light accent keeps its color in the dark appearance, so its
	// hue does not wash out.
	pLight, err := Custom("#fdf4ff")
	if err != nil {
		t.Fatal(err)
	}
	if pLight.Dark.Accent != "#fdf4ff" {
		t.Fatalf("dark accent of a light color = %q, want it kept", pLight.Dark.Accent)
	}
	// The derived palette renders gofmt-clean tokens.
	got, err := Generate(p)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := format.Source(got); err != nil {
		t.Fatalf("the generated tokens are not gofmt-clean: %v", err)
	}
	// A malformed hex is an error.
	for _, bad := range []string{"nope", "#12", "#1234567", "6366f1"} {
		if _, err := Custom(bad); err == nil {
			t.Fatalf("Custom(%q) accepted a bad hex", bad)
		}
	}
}

// TestApply writes the palette into a module, and refuses a module that
// has no tokens yet.
func TestApply(t *testing.T) {
	p, err := Find("emerald")
	if err != nil {
		t.Fatal(err)
	}
	dir := t.TempDir()
	tokensDir := filepath.ToSlash(filepath.Join("internal", "ui", "tokens"))
	target := filepath.Join(dir, tokensDir, "tokens.go")
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(target, []byte("package tokens\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Apply(dir, tokensDir, p); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(target)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `ui.Hex("#10b981")`) {
		t.Fatal("the applied file lacks the emerald accent")
	}
	if err := Apply(t.TempDir(), tokensDir, p); err == nil {
		t.Fatal("Apply accepted a module without tokens")
	}
}
