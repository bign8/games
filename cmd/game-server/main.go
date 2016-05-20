package main

// Game server (w/chat) based on the following article
// https://talks.golang.org/2012/chat.slide

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/bign8/games"
	"github.com/bign8/games/impl/ttt"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// This stupid line of code allows all the implementation to run init
var registry = map[string]games.Game{
	ttt.Game.Slug: ttt.Game,
}

func main() {
	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/api/v0.0.0/games", gamesHandler)
	r.HandleFunc("/play/random", randomHandler)
	r.Handle("/play/{slug}/socket", websocket.Handler(socketHandler))
	r.HandleFunc("/play/{slug}", gameHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("www"))))
	r.PathPrefix("/").HandlerFunc(rootHandler)

	// Spin up server
	err := http.ListenAndServe("localhost:4000", r)
	if err != nil {
		log.Fatal(err)
	}
}

type response struct {
	Code int
	Msg  string
	Data interface{}
}

func gamesHandler(w http.ResponseWriter, r *http.Request) {
	data := response{
		Code: 200,
		Msg:  "Success",
		Data: registry,
	}
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Numer of games: %d", len(registry))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	rootTemplate.Execute(w, struct {
		Games interface{}
	}{
		Games: registry,
	})
}

var rootTemplate = template.Must(template.ParseFiles("www/index.html"))

func randomHandler(w http.ResponseWriter, r *http.Request) {
	for _, game := range registry { // TODO: verify this is random iteration
		urlStr := fmt.Sprintf("/play/%s", game.Slug)
		http.Redirect(w, r, urlStr, http.StatusTemporaryRedirect)
		return
	}
}

var gameTemplate = template.Must(template.ParseFiles("www/game.html"))

func gameHandler(w http.ResponseWriter, r *http.Request) {
	game, ok := registry[mux.Vars(r)["slug"]]
	if !ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	gameTemplate.Execute(w, struct {
		Game interface{}
	}{
		Game: game,
	})
	// fmt.Fprintf(w, "Playing %s baby!", game.Name)
}
