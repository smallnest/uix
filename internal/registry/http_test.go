package registry

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestFetch reads a registry served over HTTP: the served JSON decodes
// into the items it carries.
func TestFetch(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"items":[{"name":"base","type":"components:base","files":[{"path":"a.go","content":"package a"}]}]}`))
	}))
	defer srv.Close()
	r, err := Fetch(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(r.Items) != 1 || r.Items[0].Name != "base" {
		t.Fatalf("items = %+v, want one named base", r.Items)
	}
	if r.Items[0].Files[0].Content != "package a" {
		t.Fatalf("file content = %q, want package a", r.Items[0].Files[0].Content)
	}
}

// TestFetchError reports a non-200 response as an error, so a wrong URL
// does not look like an empty registry.
func TestFetchError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "boom", http.StatusInternalServerError)
	}))
	defer srv.Close()
	if _, err := Fetch(srv.URL); err == nil {
		t.Fatal("Fetch succeeded, want an error")
	}
}

// TestHandler serves the embedded registry: the JSON decodes and has
// items, and the listing names the components.
func TestHandler(t *testing.T) {
	srv := httptest.NewServer(Handler())
	defer srv.Close()

	resp, err := http.Get(srv.URL + "/registry.json")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("GET /registry.json: %s", resp.Status)
	}
	var r Registry
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		t.Fatal(err)
	}
	if len(r.Items) == 0 {
		t.Fatal("served registry has no items")
	}

	list, err := http.Get(srv.URL + "/")
	if err != nil {
		t.Fatal(err)
	}
	defer list.Body.Close()
	body, err := io.ReadAll(list.Body)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(body), r.Items[0].Name) {
		t.Fatal("listing does not name the components")
	}
}
