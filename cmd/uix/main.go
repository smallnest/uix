// Command uix installs shadcn/ui-style components into a MyGo native UI
// app. It copies source into your project, as shadcn/ui does: you own the
// code, and the components import only github.com/egoist/mygo.
//
//	uix init          install the base set (design tokens)
//	uix add button    install a component and what it depends on
//	uix list          list the components in the registry
//	uix theme set     write a palette into the tokens
//
// Run the commands in the Go module you are building, or pass -dir.
// Point -registry at an HTTP registry to install from another machine;
// the URL can also sit in the uix.config.json of the module.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/smallnest/uix/internal/registry"
	"github.com/smallnest/uix/internal/theme"
)

const usage = `uix installs shadcn/ui-style components into a MyGo native UI app.

Commands:
  uix init              install the base set (design tokens) into a Go module
  uix add <name>        install a component and its dependencies into a Go module
  uix list              list the components in the registry
  uix theme list        list the palettes a module can pick
  uix theme set <name>  write a palette into the tokens of a module
  uix theme set custom <hex>  write a palette from your own accent color
  uix registry build    rebuild the registry JSON from ./components
  uix registry serve    serve the embedded registry over HTTP

Options:
  -dir <dir>            the Go module to install into (default ".")
  -registry <url>       an HTTP registry to install from, instead of the
                        embedded one; the URL also applies when a module's
                        uix.config.json records one

Run init and add in the Go module you are building, or pass -dir to point
at it. Each component is copied into the module as source you own.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
	}
	var err error
	switch os.Args[1] {
	case "init":
		err = cmdInstall("base", os.Args[2:])
	case "add":
		err = cmdAdd(os.Args[2:])
	case "list":
		err = cmdList(os.Args[2:])
	case "theme":
		err = cmdTheme(os.Args[2:])
	case "registry":
		err = cmdRegistry(os.Args[2:])
	case "help", "-h", "--help":
		fmt.Print(usage)
		return
	default:
		fmt.Fprintf(os.Stderr, "uix: unknown command %q\n\n%s", os.Args[1], usage)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "uix:", err)
		os.Exit(1)
	}
}

// options are the options uix reads from the command line.
type options struct {
	dir      string // the Go module to install into
	registry string // an HTTP registry to install from; empty means config or embedded
}

// splitArgs pulls -dir and -registry out of args, wherever they sit, and
// returns the options and the remaining positional arguments. Go's flag
// package stops at the first positional argument, so an option that
// follows a component name would otherwise be swallowed.
func splitArgs(args []string) (o options, rest []string) {
	o.dir = "."
	for i := 0; i < len(args); i++ {
		a := args[i]
		switch {
		case a == "-dir" || a == "--dir":
			if i+1 < len(args) {
				o.dir = args[i+1]
				i++
			}
		case a == "-registry" || a == "--registry":
			if i+1 < len(args) {
				o.registry = args[i+1]
				i++
			}
		case strings.HasPrefix(a, "-dir=") || strings.HasPrefix(a, "--dir="):
			o.dir = a[strings.IndexByte(a, '=')+1:]
		case strings.HasPrefix(a, "-registry=") || strings.HasPrefix(a, "--registry="):
			o.registry = a[strings.IndexByte(a, '=')+1:]
		default:
			rest = append(rest, a)
		}
	}
	return o, rest
}

// loadRegistry returns the registry to install from: the -registry URL,
// the URL in the module's uix.config.json, or the embedded registry.
func loadRegistry(o options) (*registry.Registry, error) {
	url := o.registry
	if url == "" {
		url = configRegistry(o.dir)
	}
	if url == "" {
		return registry.Load()
	}
	r, err := registry.Fetch(url)
	if err != nil {
		return nil, err
	}
	fmt.Printf("uix: registry %s\n", url)
	return r, nil
}

// configRegistry reads the registry URL of the module at dir, from
// uix.config.json, as the components.json of shadcn/ui records its
// registry. A missing or empty file means the embedded registry.
func configRegistry(dir string) string {
	data, err := os.ReadFile(filepath.Join(dir, "uix.config.json"))
	if err != nil {
		return ""
	}
	var c struct {
		Registry string `json:"registry"`
	}
	if err := json.Unmarshal(data, &c); err != nil {
		return ""
	}
	return c.Registry
}

func cmdInstall(name string, args []string) error {
	o, _ := splitArgs(args)
	r, err := loadRegistry(o)
	if err != nil {
		return err
	}
	return r.Install(o.dir, name, map[string]bool{})
}

func cmdAdd(args []string) error {
	o, names := splitArgs(args)
	if len(names) == 0 {
		return fmt.Errorf("uix add: give a component name, e.g. uix add button")
	}
	r, err := loadRegistry(o)
	if err != nil {
		return err
	}
	seen := map[string]bool{}
	for _, name := range names {
		if err := r.Install(o.dir, name, seen); err != nil {
			return err
		}
	}
	return nil
}

func cmdList(args []string) error {
	o, _ := splitArgs(args)
	r, err := loadRegistry(o)
	if err != nil {
		return err
	}
	for _, it := range r.Items {
		fmt.Printf("%s\t%s (%d files)\n", it.Name, it.Type, len(it.Files))
	}
	return nil
}

// cmdTheme dispatches the theme subcommands.
func cmdTheme(args []string) error {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		fmt.Fprint(os.Stderr, "uix theme: pick a palette for the tokens.\n\n  uix theme list          list the palettes\n  uix theme set <name>    write a palette into the tokens of the module\n  uix theme set custom <hex>  write a palette from your own accent color\n")
		return nil
	}
	switch args[0] {
	case "list":
		return cmdThemeList(args[1:])
	case "set":
		return cmdThemeSet(args[1:])
	}
	return fmt.Errorf("uix theme: unknown command %q", args[0])
}

func cmdThemeList(args []string) error {
	for _, p := range theme.Palettes() {
		fmt.Printf("%s\t%s\n", p.Name, p.Description)
	}
	fmt.Printf("%s\t%s\n", theme.CustomPalette.Name, theme.CustomPalette.Description)
	return nil
}

func cmdThemeSet(args []string) error {
	o, rest := splitArgs(args)
	if len(rest) == 0 {
		return fmt.Errorf("uix theme set: give a palette name, e.g. uix theme set emerald")
	}
	var (
		p   *theme.Palette
		err error
	)
	if rest[0] == theme.CustomPalette.Name {
		// A custom accent needs the hex that derives it; the palette is
		// built from it, not looked up.
		if len(rest) < 2 {
			return fmt.Errorf("uix theme set custom: give a hex color, e.g. uix theme set custom #f97316")
		}
		p, err = theme.Custom(rest[1])
	} else {
		p, err = theme.Find(rest[0])
	}
	if err != nil {
		return err
	}
	r, err := loadRegistry(o)
	if err != nil {
		return err
	}
	base, err := r.Get("base")
	if err != nil {
		return err
	}
	if len(base.Files) == 0 {
		return fmt.Errorf("uix theme set: the base set of the registry carries no files")
	}
	// The tokens install where the base set says, so a registry fork that
	// moves them keeps the theme command in step.
	return theme.Apply(o.dir, filepath.Dir(base.Files[0].Path), p)
}

// cmdRegistry dispatches the registry subcommands.
func cmdRegistry(args []string) error {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		fmt.Fprint(os.Stderr, "uix registry: build or serve the registry.\n\n  uix registry build [-root dir] [-out path]   write the registry JSON\n  uix registry serve [-addr host:port]          serve it over HTTP\n")
		return nil
	}
	switch args[0] {
	case "build":
		return cmdBuild(args[1:])
	case "serve":
		return cmdServe(args[1:])
	}
	return fmt.Errorf("uix registry: unknown command %q", args[0])
}

func cmdBuild(args []string) error {
	fs := flag.NewFlagSet("uix registry build", flag.ExitOnError)
	root := fs.String("root", ".", "the uix repository root")
	out := fs.String("out", "internal/registry/components.json", "where to write the registry")
	fs.Parse(args)
	componentsDir := filepath.Join(*root, "components")
	if !filepath.IsAbs(*out) {
		*out = filepath.Join(*root, *out)
	}
	return registry.Build(componentsDir, *out)
}

func cmdServe(args []string) error {
	fs := flag.NewFlagSet("uix registry serve", flag.ExitOnError)
	addr := fs.String("addr", ":7373", "the address to listen on")
	fs.Parse(args)
	fmt.Printf("uix: serving the registry at %s/registry.json\n", serveURL(*addr))
	return http.ListenAndServe(*addr, registry.Handler())
}

// serveURL turns an address into a URL for printing: a bare port becomes
// localhost, because http://:7373 is not something to click.
func serveURL(addr string) string {
	if strings.HasPrefix(addr, ":") {
		return "http://localhost" + addr
	}
	return "http://" + addr
}
