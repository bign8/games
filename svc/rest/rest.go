package rest

import (
	"net/http"

	"github.com/bign8/games"
	"github.com/bign8/games/impl"
	"github.com/bign8/games/svc"
)

type rest struct {
	http.Handler
	svc.GameService
	prefix string
}

// New constructs a new http game handler
// GET  /<prefix>/ -> list available endpoints
// POST /<prefix>/random -> redirects to random game
// POST /<prefix>/{slug} -> creates a game and redirects to /<prefix>/{slug}/{guid}
// GET  /<prefix>/{slug}/{guid} -> gets an active game state and moves if authenticated
// POST /<prefix>/{slug}/{guid} -> post the available move ID and advance the game
func New(s svc.GameService, prefix string) http.Handler {
	mux := http.NewServeMux()
	o := &rest{
		Handler:     mux,
		GameService: s,
		prefix:      prefix,
	}
	mux.HandleFunc("/", o.list)
	mux.HandleFunc("/rand", o.rand)
	for slug, game := range impl.Map() {
		mux.Handle("/"+slug+"/", http.StripPrefix("/"+slug+"/", o.manage(game)))
		mux.Handle("/"+slug, o.create(game))
	}
	return o
}

// redirect is a wrapper function for http.Redirect that properly prepends the prefix
func (o *rest) redirect(w http.ResponseWriter, r *http.Request, url string, code int) {
	http.Redirect(w, r, o.prefix+url, code)
}

// list returns a json list of all available games
func (o *rest) list(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "list available games "+r.URL.Path, http.StatusNotImplemented)
}

// rand (POST-only) forwards the create game request to a random games slug
func (o *rest) rand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "POST only endpoint", http.StatusMethodNotAllowed)
		return
	}
	o.redirect(w, r, "/"+svc.Random(), http.StatusTemporaryRedirect)
}

func (o *rest) create(game games.Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST only endpoint", http.StatusMethodNotAllowed)
			return
		}

		// Returns a token that represents the given player per game
		// this is in the form of the guid, guids for different opponents (even if playing the same game) are different
		http.Error(w, "start a game "+r.URL.Path, http.StatusNotImplemented)
	}
}

func (o *rest) manage(game games.Game) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "manage a game "+r.URL.Path, http.StatusNotImplemented)
	}
}
