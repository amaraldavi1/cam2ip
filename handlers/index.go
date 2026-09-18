// Package handlers provides HTTP handlers for the cam2ip application.
package handlers

import (
	_ "embed"
	"net/http"
)

//go:embed ui/index.html
var indexHTML []byte

// Index handler serves the web UI.
type Index struct {
}

// NewIndex returns new Index handler.
func NewIndex() *Index {
	return &Index{}
}

// ServeHTTP handles requests on incoming connections.
func (i *Index) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)

		return
	}

	if r.Method != "GET" && r.Method != "HEAD" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")

	_, _ = w.Write(indexHTML)
}
