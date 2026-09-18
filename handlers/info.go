package handlers

import (
	"encoding/json"
	"net/http"
)

// Info describes the running server for the web UI.
type Info struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Device  string `json:"device"`
	Format  string `json:"format"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Quality int    `json:"quality"`
	Delay   int    `json:"delay"`
	Lazy    bool   `json:"lazy"`
	Auth    bool   `json:"auth"`
	Viewers int    `json:"viewers"`
}

// InfoHandler serves server information as JSON.
type InfoHandler struct {
	info   Info
	stream *Stream
}

// NewInfo returns new info handler.
func NewInfo(info Info, stream *Stream) *InfoHandler {
	return &InfoHandler{info, stream}
}

// ServeHTTP handles requests on incoming connections.
func (i *InfoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "HEAD" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	info := i.info
	info.Viewers = i.stream.Viewers()

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Cache-Control", "no-store")

	_ = json.NewEncoder(w).Encode(info)
}
