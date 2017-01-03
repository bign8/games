package main

// Game server (w/chat) based on the following article
// https://talks.golang.org/2012/chat.slide

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"

	"github.com/bign8/games/impl"
)

var (
	p       = func(n string) string { return filepath.Join("cmd", "server", "tpl", n+".gohtml") }
	rootTpl = template.Must(template.ParseFiles(p("base"), p("root")))
	gameTpl = template.Must(template.ParseFiles(p("base"), p("game")))
	infoTpl = template.Must(template.ParseFiles(p("base"), p("info")))
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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	rootTpl.Execute(w, struct {
		Games interface{}
	}{
		Games: impl.Map(),
	})
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Get(mux.Vars(r)["slug"])
	if !ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	gameTpl.Execute(w, struct {
		Game  interface{}
		Board template.HTML
	}{
		Game:  game,
		Board: template.HTML(game.Board),
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	infoTpl.Execute(w, nil)
}
