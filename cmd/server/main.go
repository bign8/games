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

	"github.com/gorilla/mux"

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
	rout *mux.Router // TODO: inject this into handlers
)

func main() {
	flag.Parse()

	// Setup routes
	rout = mux.NewRouter()
	rout.HandleFunc("/play", randomHandler).Methods(http.MethodGet)                            // Redirects to play/{slug}
	rout.HandleFunc("/play/{slug}", lobbyHandler).Methods(http.MethodGet)                      // Redirects to play/{slug}/{id}
	rout.HandleFunc("/play/{slug}/board.svg", boardHandler).Methods(http.MethodGet)            // Board is same for all players
	rout.HandleFunc("/play/{slug}/{id}", gameHandler).Methods(http.MethodGet, http.MethodPost) // Gives the ability to play game
	rout.HandleFunc("/play/{slug}/{id}/state.svg", stateHandler).Methods(http.MethodGet)       // Returns the given state of the game

	// TODO: depricate
	// r.Handle("/play/{slug}/socket", websocket.Handler(socketHandler))

	rout.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("cmd", "server", "www")))))
	rout.HandleFunc("/about", aboutHandler)
	rout.PathPrefix("/").HandlerFunc(rootHandler)

	// Spin up server
	fmt.Println("Serving on :" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, rout))
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
	list := impl.Map()

	// HTTP/2 push out the desired images
	if pusher, ok := w.(http.Pusher); ok {
		for slug := range list {
			url, err := rout.Get("board").URL(slug)
			if err != nil {
				pusher.Push(url.String(), nil)
			}
		}
	}

	// Build the HTML to send to the client
	rootTpl.Execute(w, struct {
		Games map[string]games.Game
	}{Games: list})
}

func lobbyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "TODO", http.StatusExpectationFailed)
}

func boardHandler(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Map()[mux.Vars(r)["slug"]]
	if !ok {
		http.Error(w, "Game Does Not Exist", http.StatusExpectationFailed)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, game.Board)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Get(mux.Vars(r)["slug"])
	if !ok {
		http.Error(w, "Game Does Not Exist", http.StatusExpectationFailed)
		return
	}
	gameTpl.Execute(w, struct {
		Game  games.Game
		Board template.HTML
	}{
		Game:  game,
		Board: template.HTML(game.Board),
	})
}

func stateHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "TODO", http.StatusExpectationFailed)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	infoTpl.Execute(w, nil)
}
