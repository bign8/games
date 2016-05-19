package main

// Game server (w/chat) based on the following article
// https://talks.golang.org/2012/chat.slide

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"

	"github.com/bign8/games"
	"github.com/bign8/games/impl/ttt"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

// This stupid line of code allows all the implementation to run init
var preloader = []games.Starter{
	ttt.New,
}

func main() {
	// Setup routes
	r := mux.NewRouter()
	r.Handle("/api/v0.0.0/socket", websocket.Handler(socketHandler))
	r.HandleFunc("/api/v0.0.0/games", gamesHandler)
	r.HandleFunc("/play/random", randomHandler)
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
		Data: games.List(),
	}

	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Numer of games: %d", len(games.List()))
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	rootTemplate.Execute(w, games.List())
}

var rootTemplate = template.Must(template.ParseFiles("www/index.html"))

func randomHandler(w http.ResponseWriter, r *http.Request) {
	list := games.List()
	urlStr := fmt.Sprintf("/play/%s", list[rand.Intn(len(list))].Slug)
	http.Redirect(w, r, urlStr, http.StatusTemporaryRedirect)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	game := games.Get(mux.Vars(r)["slug"])
	if game == nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Fprintf(w, "Playing %s baby!", game.Name)
}
