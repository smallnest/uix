package registry

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// manifest is the optional uix.json in a component directory. It says how
// the component installs; every field has a default.
type manifest struct {
	Type                 string   `json:"type"`
	Install              string   `json:"install"`
	Dependencies         []string `json:"dependencies"`
	RegistryDependencies []string `json:"registryDependencies"`
}

// Build writes the registry for the components in dir to out.
//
// Each subdirectory of dir is one item, named after it. The build reads
// uix.json in each directory for how the item installs. The defaults are
// the base set at internal/ui and components at
// internal/ui/components/<name>. The files of an item are its non-test,
// non-manifest files, with their contents, so the registry is
// self-contained.
func Build(dir, out string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	var r Registry
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		item, err := buildItem(filepath.Join(dir, e.Name()), e.Name())
		if err != nil {
			return err
		}
		r.Items = append(r.Items, *item)
	}
	sort.Slice(r.Items, func(i, j int) bool { return r.Items[i].Name < r.Items[j].Name })
	data, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(out, append(data, '\n'), 0o644)
}

func buildItem(dir, name string) (*Item, error) {
	item := &Item{Name: name, Type: "components:ui"}
	if name == "base" {
		item.Type = "components:base"
	}
	var m manifest
	if data, err := os.ReadFile(filepath.Join(dir, "uix.json")); err == nil {
		if err := json.Unmarshal(data, &m); err != nil {
			return nil, fmt.Errorf("registry: %s: %v", filepath.Join(dir, "uix.json"), err)
		}
		if m.Type != "" {
			item.Type = m.Type
		}
		item.Dependencies = m.Dependencies
		item.RegistryDependencies = m.RegistryDependencies
	}
	install := m.Install
	if install == "" {
		install = "internal/ui/tokens"
		if item.Type != "components:base" {
			install = "internal/ui/components/" + name
		}
	}
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		if base == "uix.json" || strings.HasSuffix(base, "_test.go") {
			return nil
		}
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		item.Files = append(item.Files, File{
			Path:    filepath.ToSlash(filepath.Join(install, rel)),
			Content: string(content),
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return item, nil
}
