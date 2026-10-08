// Package registry stores the components a uix app installs, as one
// self-contained JSON file, and installs them into a Go module.
//
// The model follows shadcn/ui: a component is source code that you copy
// into your project and own, not a library you import. The registry
// carries each file with its content, so the uix binary is standalone
// and the registry can later be served over HTTP.
//
// A component imports only github.com/egoist/mygo and the sibling
// components it composes, such as the tooltip a stat card shows. The
// copy model has no path rewriting in the consumer, so Install rewrites
// the import of a sibling to the path the copy has in the consumer's
// module, as shadcn/ui rewrites its imports to the consumer's alias.
package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	_ "embed" // for the embedded registry, below
)

// File is one file of a component, with its content.
type File struct {
	Path    string `json:"path"`    // path to write in the consumer module
	Content string `json:"content"` // the file's source
}

// Item is one entry of the registry: the base set or a component.
type Item struct {
	Name                 string   `json:"name"`
	Type                 string   `json:"type"` // components:base or components:ui
	Install              string   `json:"install,omitempty"` // where it installs in the consumer
	Import               string   `json:"import,omitempty"`  // its import path in this module
	RegistryDependencies []string `json:"registryDependencies,omitempty"`
	Dependencies         []string `json:"dependencies,omitempty"` // Go modules to fetch
	Files                []File   `json:"files"`
}

// Registry is the whole registry.
type Registry struct {
	Items []Item `json:"items"`
}

//go:embed components.json
var embedded []byte

//go:generate go run ./cmd/uix registry build

// Load reads the registry embedded in the binary.
func Load() (*Registry, error) {
	var r Registry
	if err := json.Unmarshal(embedded, &r); err != nil {
		return nil, fmt.Errorf("registry: %v", err)
	}
	return &r, nil
}

// Fetch reads the registry at url over HTTP, as uix add does when you
// point it at a remote registry. Why not cache the fetched registry: it
// is small, and a fresh fetch keeps every add up to date, as shadcn/ui
// fetches its registry per command.
func Fetch(url string) (*Registry, error) {
	client := &http.Client{Timeout: 20 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("fetch %s: %v", url, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("fetch %s: %s", url, resp.Status)
	}
	var r Registry
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("fetch %s: %v", url, err)
	}
	return &r, nil
}

// Get finds the item with the given name.
func (r *Registry) Get(name string) (*Item, error) {
	for i := range r.Items {
		if r.Items[i].Name == name {
			return &r.Items[i], nil
		}
	}
	return nil, fmt.Errorf("registry: no component %q; available: %s", name, r.Names())
}

// Names lists the item names, for errors and uix list.
func (r *Registry) Names() string {
	var names []string
	for _, it := range r.Items {
		names = append(names, it.Name)
	}
	return strings.Join(names, ", ")
}

// Install writes item and its registry dependencies into dir, a Go
// module, and fetches the Go modules they need. seen guards against
// dependency cycles.
func (r *Registry) Install(dir, name string, seen map[string]bool) error {
	if seen[name] {
		return nil
	}
	seen[name] = true
	item, err := r.Get(name)
	if err != nil {
		return err
	}
	for _, dep := range item.RegistryDependencies {
		if err := r.Install(dir, dep, seen); err != nil {
			return err
		}
	}
	// Rewrite the imports of the copies so the components that compose
	// one another point at the copy in the consumer's module, not at the
	// registry's module. Each copy keeps the path it had in the registry,
	// under the consumer's module path, as shadcn/ui rewrites its
	// imports to the consumer's alias.
	rewrites := map[string]string{}
	if mod := consumerModule(dir); mod != "" {
		for i := range r.Items {
			it := &r.Items[i]
			if it.Import != "" && it.Install != "" {
				rewrites[it.Import] = mod + "/" + it.Install
			}
		}
	}
	for _, f := range item.Files {
		target := filepath.Join(dir, filepath.FromSlash(f.Path))
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return err
		}
		content := f.Content
		for from, to := range rewrites {
			content = strings.ReplaceAll(content, `"`+from+`"`, `"`+to+`"`)
		}
		if err := os.WriteFile(target, []byte(content), 0o644); err != nil {
			return err
		}
		fmt.Printf("uix: wrote %s\n", f.Path)
	}
	return GoGet(dir, item.Dependencies...)
}

// consumerModule returns the module path of the Go module in dir, or ""
// if dir has no go.mod. Install rewrites the imports of the copies to
// this path, so they compose one another in the consumer's module.
func consumerModule(dir string) string {
	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return ""
	}
	for _, line := range strings.Split(string(data), "\n") {
		if f := strings.Fields(line); len(f) == 2 && f[0] == "module" {
			return f[1]
		}
	}
	return ""
}

// GoGet fetches the Go modules a component needs, from dir.
func GoGet(dir string, deps ...string) error {
	if len(deps) == 0 {
		return nil
	}
	args := append([]string{"get"}, deps...)
	cmd := exec.Command("go", args...)
	cmd.Dir = dir
	// Components need a recent Go; let it fetch the toolchain itself.
	cmd.Env = append(os.Environ(), "GOTOOLCHAIN=auto")
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}
	if err != nil {
		return err
	}
	// go get records the modules in go.mod but not always the sums of
	// their transitive packages; tidy fills them, so the consumer builds
	// at once instead of asking for a go.sum entry per missing package.
	return run(dir, "go", "mod", "tidy")
}

// run runs a command from dir, printing its output.
func run(dir, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOTOOLCHAIN=auto")
	out, err := cmd.CombinedOutput()
	if len(out) > 0 {
		fmt.Print(string(out))
	}
	return err
}
