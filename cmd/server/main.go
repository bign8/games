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

	"github.com/bign8/games"
	"github.com/bign8/games/cmd/server/store"
	"github.com/bign8/games/impl"
)

// various HTML templates (TODO: figure out how to use ParseGob or something with these guys)
var (
	p       = func(n string) string { return filepath.Join("cmd", "server", "tpl", n+".gohtml") }
	rootTpl = template.Must(template.ParseFiles(p("base"), p("root")))
	gameTpl = template.Must(template.ParseFiles(p("base"), p("game")))
	infoTpl = template.Must(template.ParseFiles(p("base"), p("info")))
	playTpl = template.Must(template.ParseFiles(p("base"), p("play")))
	// TODO: sexy error pages
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

type server struct {
	g map[string]games.Game
	r *mux.Router
	s store.Store
	// t *template.Template
}

func main() {
	flag.Parse()

	// Setup store
	s := &server{
		g: impl.Map(),
		r: mux.NewRouter(),
		s: store.NewMemoryStore(),
		// t: template.Must(template.ParseGlob("cmd/server/tpl/*.gohtml")),
	}

	// Setup routes
	s.r.HandleFunc("/play", s.random).Methods(http.MethodGet)                                         // Redirects to play/{slug}
	s.r.HandleFunc("/play/{slug}", s.lobby).Methods(http.MethodGet)                                   // Redirects to play/{slug}/{id}
	s.r.HandleFunc("/play/{slug}/board.svg", s.board).Methods(http.MethodGet).Name("board")           // Board is same for all players
	s.r.HandleFunc("/play/{slug}/{id}", s.game).Methods(http.MethodGet, http.MethodPost).Name("game") // Gives the ability to play game
	s.r.HandleFunc("/play/{slug}/{id}/state.svg", s.state).Methods(http.MethodGet)                    // Returns the given state of the game
	s.r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join("cmd", "server", "www")))))
	s.r.HandleFunc("/about", s.about)
	s.r.PathPrefix("/").HandlerFunc(s.root)

	// Spin up server
	fmt.Println("Serving on :" + *port)
	log.Fatal(http.ListenAndServe(":"+*port, s.r))
}

func (s *server) random(w http.ResponseWriter, r *http.Request) {
	urlStr := fmt.Sprintf("/play/%s", impl.Rand())
	http.Redirect(w, r, urlStr, http.StatusTemporaryRedirect)
}

func (s *server) root(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	list := impl.Map()

	// HTTP/2 push out the desired images
	if pusher, ok := w.(http.Pusher); ok {
		for slug := range list {
			url, err := s.r.Get("board").URL(slug)
			if err != nil {
				pusher.Push(url.String(), nil)
			}
		}
	}

	// Build the HTML to send to the client
	rootTpl = template.Must(template.ParseFiles(p("base"), p("root"))) // TODO: remove this line
	rootTpl.Execute(w, struct {
		Games map[string]games.Game
	}{Games: list})
}

func (s *server) lobby(w http.ResponseWriter, r *http.Request) {
	game, ok := s.g[mux.Vars(r)["slug"]]
	if !ok {
		http.Error(w, "Game Does Not Exist", http.StatusExpectationFailed)
		return
	}

	// get users cookie
	force := true
	cookie, err := r.Cookie("games-pid-" + game.Slug)
	if err == http.ErrNoCookie {
		force = false
		cookie = &http.Cookie{Name: "games-pid-" + game.Slug, Expires: time.Now().Add(time.Hour)} // * 24 * 30
		cookie.Value, err = s.s.NewPlayer(game)
		if err != nil {
			http.Error(w, "Could Not Create Player", http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, cookie)
	}

	// attempt to pair user
	gameID, err := s.s.Pair(game, cookie.Value, force)
	if err == nil {
		url, _ := s.r.Get("game").URL(cookie.Value, gameID) // TODO: handle error
		http.Redirect(w, r, url.String(), http.StatusSeeOther)
		return
	}
	playTpl = template.Must(template.ParseFiles(p("base"), p("play"))) // TODO: remove this line
	playTpl.Execute(w, game)
}

func (s *server) board(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Map()[mux.Vars(r)["slug"]]
	if !ok {
		http.Error(w, "Game Does Not Exist", http.StatusExpectationFailed)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprint(w, game.Board)
}

func (s *server) game(w http.ResponseWriter, r *http.Request) {
	game, ok := impl.Get(mux.Vars(r)["slug"])
	if !ok {
		http.Error(w, "Game Does Not Exist", http.StatusExpectationFailed)
		return
	}
	gameTpl = template.Must(template.ParseFiles(p("base"), p("game"))) // TODO: remove this line
	gameTpl.Execute(w, struct {
		Game  games.Game
		Board template.HTML
	}{
		Game:  game,
		Board: template.HTML(game.Board),
	})
}

func (s *server) state(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "TODO", http.StatusExpectationFailed)
}

func (s *server) about(w http.ResponseWriter, r *http.Request) {
	infoTpl = template.Must(template.ParseFiles(p("base"), p("info"))) // TODO: remove this line
	infoTpl.Execute(w, nil)
}
