package http

import (
	"net/http"

	"golang.org/x/net/websocket"
)

// New constructs a new http game handler
func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/{slug}", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.HandleFunc("/{slug}/{guid}", func(w http.ResponseWriter, r *http.Request) {

	})

	mux.Handle("/{slug}/{guid}/socket", websocket.Handler(nil))

	return mux
}
