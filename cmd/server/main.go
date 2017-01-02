package main

// Game server (w/chat) based on the following article
// https://talks.golang.org/2012/chat.slide

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"

	"github.com/bign8/games/impl"
)

func main() {
	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/play/random", randomHandler)
	r.Handle("/play/{slug}/socket", websocket.Handler(socketHandler))
	r.HandleFunc("/play/{slug}", gameHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("www"))))
	r.HandleFunc("/about", aboutHandler)
	r.PathPrefix("/").HandlerFunc(rootHandler)

	// Spin up server
	err := http.ListenAndServe("localhost:4000", r)
	if err != nil {
		log.Fatal(err)
	}
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	urlStr := fmt.Sprintf("/play/%s", impl.Rand())
	http.Redirect(w, r, urlStr, http.StatusTemporaryRedirect)
}
