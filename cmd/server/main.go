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

	"golang.org/x/net/websocket"

	"github.com/bign8/games"
	"github.com/bign8/games/cmd/server/app"
	"github.com/bign8/games/impl"
	"github.com/bign8/games/player"
	"github.com/bign8/games/svc/memory"
	"github.com/bign8/games/svc/rest"
)

// various HTML templates
var (
	rootTPL = load("root")
	gameTPL = load("game")
	infoTPL = load("info")
)

// Environment parameters + defaults
var (
	host = flag.String("host", def("HOST", ""), "host to serve on")
	port = flag.String("port", def("PORT", "4000"), "port to serve on")
)

func main() {
	flag.Parse()

	// Setup API server
	http.Handle("/api/", http.StripPrefix("/api", rest.New(memory.New(), "/api")))

	// Setup standard gorilla router
	fs := http.FileServer(http.Dir(filepath.Join("cmd", "server", "www")))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// Setup programatic routes
	http.Handle("/play/random", http.HandlerFunc(randomHandler))
	http.Handle("/about", wrap(infoTPL, aboutHandler))
	http.Handle("/", wrap(rootTPL, rootHandler))

	// Bind specific game implementations
	for slug, game := range impl.Map() {
		http.Handle("/play/"+slug+"/socket", websocket.Handler(app.Socket(game)))
		http.Handle("/play/"+slug, wrap(gameTPL, gameHandler(game)))
	}

	// Spin up server
	log.Print("Serving on :" + *port)
	if err := http.ListenAndServe(*host+":"+*port, nil); err != nil {
		log.Fatal(err)
	}
}

// load hydrates a template from persistence
// TODO: bundle static files in go binary
func load(name string) *template.Template {
	p := func(n string) string {
		return filepath.Join("cmd", "server", "tpl", n+".gohtml")
	}
	return template.Must(template.ParseFiles(p("base"), p(name))).Option("missingkey=error")
}

// def gives the environment variable if present, otherwise it returns def
func def(env, str string) string {
	if v := os.Getenv(env); v != "" {
		return v
	}
	return str
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

func gameHandler(game games.Game) webpage {
	return func(r *http.Request) interface{} {
		return showGame{
			Game:  game,
			Board: template.HTML(game.Board),
		}
	}
}

func aboutHandler(r *http.Request) interface{} {
	// TODO: generate CSRF token for contact form (or something)
	return nil
}
