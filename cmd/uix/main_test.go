package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestSplitArgs pulls -dir and -registry out of any position, so an
// option after a component name is not swallowed.
func TestSplitArgs(t *testing.T) {
	o, rest := splitArgs([]string{"button", "-registry", "https://r.example/registry.json", "-dir", "app"})
	if o.dir != "app" {
		t.Fatalf("dir = %q, want app", o.dir)
	}
	if o.registry != "https://r.example/registry.json" {
		t.Fatalf("registry = %q", o.registry)
	}
	if len(rest) != 1 || rest[0] != "button" {
		t.Fatalf("rest = %v, want [button]", rest)
	}

	o, rest = splitArgs([]string{"--registry=https://r.example/r.json"})
	if o.registry != "https://r.example/r.json" {
		t.Fatalf("registry = %q", o.registry)
	}
	if len(rest) != 0 {
		t.Fatalf("rest = %v, want none", rest)
	}
}

// TestConfigRegistry reads the registry URL of a module from its
// uix.config.json, and treats a missing file as no URL.
func TestConfigRegistry(t *testing.T) {
	dir := t.TempDir()
	if got := configRegistry(dir); got != "" {
		t.Fatalf("configRegistry of an empty module = %q, want none", got)
	}
	config := filepath.Join(dir, "uix.config.json")
	if err := os.WriteFile(config, []byte(`{"registry":"https://r.example/r.json"}`), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := configRegistry(dir); got != "https://r.example/r.json" {
		t.Fatalf("configRegistry = %q, want the URL", got)
	}
}

// TestServeURL prints a clickable URL for an address: a bare port names
// localhost.
func TestServeURL(t *testing.T) {
	if got := serveURL(":7373"); got != "http://localhost:7373" {
		t.Fatalf("serveURL(:7373) = %q", got)
	}
	if got := serveURL("0.0.0.0:8080"); got != "http://0.0.0.0:8080" {
		t.Fatalf("serveURL(0.0.0.0:8080) = %q", got)
	}
}

// TestThemeSet writes a palette into the tokens of a module, and errors
// when the module has no tokens or the palette is unknown.
func TestThemeSet(t *testing.T) {
	dir := t.TempDir()
	if err := cmdTheme([]string{"set", "emerald", "-dir", dir}); err == nil {
		t.Fatal("theme set accepted a module without tokens")
	}
	tokens := filepath.Join(dir, "internal", "ui", "tokens")
	if err := os.MkdirAll(tokens, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tokens, "tokens.go"), []byte("package tokens\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := cmdTheme([]string{"set", "emerald", "-dir", dir}); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(tokens, "tokens.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `ui.Hex("#10b981")`) {
		t.Fatal("the emerald accent is missing from the applied tokens")
	}
	if err := cmdTheme([]string{"set", "nope", "-dir", dir}); err == nil {
		t.Fatal("theme set accepted an unknown palette")
	}
}

// TestThemeSetCustom derives a palette from a hex color and writes it
// into the tokens, and errors when the hex is missing or malformed.
func TestThemeSetCustom(t *testing.T) {
	dir := t.TempDir()
	tokens := filepath.Join(dir, "internal", "ui", "tokens")
	if err := os.MkdirAll(tokens, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(tokens, "tokens.go"), []byte("package tokens\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := cmdTheme([]string{"set", "custom", "#f97316", "-dir", dir}); err != nil {
		t.Fatal(err)
	}
	got, err := os.ReadFile(filepath.Join(tokens, "tokens.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(got), `ui.Hex("#f97316")`) {
		t.Fatal("the custom accent is missing from the applied tokens")
	}
	if !strings.Contains(string(got), "with the custom palette") {
		t.Fatal("the applied tokens do not name the custom palette")
	}
	if err := cmdTheme([]string{"set", "custom", "-dir", dir}); err == nil {
		t.Fatal("theme set custom accepted a missing hex")
	}
	if err := cmdTheme([]string{"set", "custom", "nope", "-dir", dir}); err == nil {
		t.Fatal("theme set custom accepted a malformed hex")
	}
}
