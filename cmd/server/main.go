package main

// Game server (w/chat) based on the following article
// https://talks.golang.org/2012/chat.slide

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"

	"github.com/bign8/games"
	"github.com/bign8/games/cmd/server/app"
	"github.com/bign8/games/impl"
	"github.com/bign8/games/player"
)

// various HTML templates
var (
	p       = func(n string) string { return filepath.Join("cmd", "server", "tpl", n+".gohtml") }
	rootTPL = template.Must(template.ParseFiles(p("base"), p("root"))).Option("missingkey=error")
	gameTPL = template.Must(template.ParseFiles(p("base"), p("game"))).Option("missingkey=error")
	infoTPL = template.Must(template.ParseFiles(p("base"), p("info"))).Option("missingkey=error")
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

	// Setup standard gorilla router
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir(filepath.Join("cmd", "server", "www")))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", fs))

	// Setup programatic routes
	r.Handle("/play/random", http.HandlerFunc(randomHandler))
	r.Handle("/play/{slug}/socket", websocket.Handler(app.Socket))
	r.Handle("/play/{slug}", wrap(gameTPL, gameHandler))
	r.Handle("/about", wrap(infoTPL, aboutHandler))
	r.Handle("/", wrap(rootTPL, rootHandler))

	// Spin up server
	log.Print("Serving on :" + *port)
	if err := http.ListenAndServe(":"+*port, r); err != nil {
		log.Fatal(err)
	}
}

// webpage is a simple handler for dealing with standard web-page assets
// iff response is an error, internal server error will be returned
type webpage func(r *http.Request) interface{}

type redirect string // do a temporary redirect to this location

func wrap(tpl *template.Template, fn webpage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := fn(r)
		if url, ok := data.(redirect); ok {
			http.Redirect(w, r, string(url), http.StatusTemporaryRedirect)
			return
		}
		err := tpl.Execute(w, struct {
			Data interface{}
			Year int
			Path string
		}{
			Data: data,
			Year: time.Now().Year(),
			Path: r.URL.Path,
		})
		if err != nil {
			log.Println("Failure to render", r.Method, r.URL.Path, err)
			http.Error(w, "Unable to render page.", http.StatusInternalServerError)
		}
	}
}

func randomHandler(w http.ResponseWriter, r *http.Request) {
	urlStr := "/play/" + impl.Rand()
	http.Redirect(w, r, urlStr, http.StatusTemporaryRedirect)
}

func rootHandler(r *http.Request) interface{} {
	if r.URL.Path != "/" {
		log.Printf("Invalid root path: %q", r.URL.Path)
		return redirect("/")
	}

	// Convert games for rendering
	list := impl.Map()
	output := make(map[string]showGame, len(list))
	for slug, game := range list {
		if err := game.Valid(); err != nil {
			log.Print("skipping" + err.Error())
			continue
		}

		// show the first 3 moves of a game
		match := game.Play(player.Random).Advance().Advance().Advance()
		output[slug] = showGame{
			Game:  game,
			Board: template.HTML(game.Board),
			First: template.HTML(match.SVG(false)),
		}
	}
	return output
}

type showGame struct {
	games.Game
	Board template.HTML
	First template.HTML
}

func gameHandler(r *http.Request) interface{} {
	slug := mux.Vars(r)["slug"]
	game, ok := impl.Get(slug)
	if !ok {
		log.Printf("Unable to find game: %q", slug)
		return redirect("/")
	}
	return struct {
		Game  games.Game
		Board template.HTML
	}{
		Game:  game,
		Board: template.HTML(game.Board),
	}
}

func aboutHandler(r *http.Request) interface{} {
	return nil
}
