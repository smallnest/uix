package theme

import (
	"fmt"
	"os"
	"path/filepath"
)

// Apply writes the palette into the tokens package of the module at dir.
// tokensDir is the install path of the base set, as the registry records
// it. The tokens.go must already exist: the module owns its tokens, and
// the command only restyles them, it does not create them.
func Apply(dir, tokensDir string, p *Palette) error {
	content, err := Generate(p)
	if err != nil {
		return err
	}
	rel := filepath.ToSlash(filepath.Join(tokensDir, "tokens.go"))
	target := filepath.Join(dir, filepath.FromSlash(rel))
	if _, err := os.Stat(target); err != nil {
		return fmt.Errorf("no tokens package at %s: run uix init first", rel)
	}
	if err := os.WriteFile(target, content, 0o644); err != nil {
		return err
	}
	fmt.Printf("uix: wrote %s with the %s palette\n", rel, p.Name)
	return nil
}
