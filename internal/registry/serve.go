package registry

import (
	"fmt"
	"net/http"
)

// Handler serves the embedded registry over HTTP, so uix on another
// machine can install from it:
//
//	GET /registry.json   the registry itself
//	GET /                a listing, for a browser
//
// Run uix registry build first, so the embedded registry is current. The
// registry JSON carries every file with its content, so one request
// installs a component. Why not serve the files individually: the client
// would have to know which files a component has and rewrite their paths,
// which the self-contained registry already does.
func Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/registry.json", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// A browser-based tool may fetch the registry from another origin.
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write(embedded)
	})
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		reg, err := Load()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "uix registry\n\ninstall with, from a Go module:\n  uix add <name> -registry %s/registry.json\n\ncomponents:\n", serverURL(r))
		for _, it := range reg.Items {
			fmt.Fprintf(w, "  %-12s %s (%d files)\n", it.Name, it.Type, len(it.Files))
		}
	})
	return mux
}

// serverURL derives the URL of this server from the request, for the
// listing page.
func serverURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return scheme + "://" + r.Host
}
