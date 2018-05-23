package main

// Game server (w/chat) based on the following article
// https://talks.golang.org/2012/chat.slide

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"

	"github.com/bign8/games"
	"github.com/bign8/games/impl"
)

// various HTML templates
var (
	p       = func(n string) string { return filepath.Join("cmd", "server", "tpl", n+".gohtml") }
	rootTpl = template.Must(template.ParseFiles(p("base"), p("root")))
	gameTpl = template.Must(template.ParseFiles(p("base"), p("game")))
	infoTpl = template.Must(template.ParseFiles(p("base"), p("info")))
)

// Environment parameters + defaults
var (
	defaults = map[string]string{
		"PORT": "4000",
	}
	def = func(s string) string {
		if v := os.Getenv(s); v != "" {
			return v
		}
		return defaults[s]
	}
	port = flag.String("port", def("PORT"), "port to serve on")
)

func main() {
	flag.Parse()

	// Setup routes
	r := mux.NewRouter()
	r.HandleFunc("/play/random", randomHandler)
	r.Handle("/play/{slug}/socket", websocket.Handler(socketHandler))
	r.HandleFunc("/play/{slug}", gameHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("cmd", "server", "www")))))
	r.HandleFunc("/about", aboutHandler)
	r.PathPrefix("/").HandlerFunc(rootHandler)

	// Spin up server
	fmt.Println("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, r); err != nil {
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
		Games map[string]showGame
		Year  int
	}{
		Games: process(impl.Map()),
		Year:  time.Now().Year(),
	})
}

type showGame struct {
	games.Game
	Board template.HTML
}

func process(in map[string]games.Game) map[string]showGame {
	new := make(map[string]showGame, len(in))
	for k, v := range in {
		new[k] = showGame{v, template.HTML(v.Board)}
	}
	return new
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Get(mux.Vars(r)["slug"])
	if !ok {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	gameTpl.Execute(w, struct {
		Game  games.Game
		Board template.HTML
		Year  int
	}{
		Game:  game,
		Board: template.HTML(game.Board),
		Year:  time.Now().Year(),
	})
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	infoTpl.Execute(w, nil)
}
