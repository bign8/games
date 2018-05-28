// Package impl provides a registry for all loaded game types.
package impl

import (
	"sync"

	"github.com/bign8/games"
	"github.com/bign8/games/impl/backgammon"
	"github.com/bign8/games/impl/cc"
	"github.com/bign8/games/impl/checkers"
	"github.com/bign8/games/impl/chess"
	"github.com/bign8/games/impl/connect4"
	"github.com/bign8/games/impl/cribbage"
	gos "github.com/bign8/games/impl/go"
	"github.com/bign8/games/impl/mancala"
	"github.com/bign8/games/impl/ttt"
)

var (
	reg = map[string]games.Game{}
	mux sync.RWMutex
)

// Get returns a registered name via it's short slug.
func Get(slug string) (g games.Game, ok bool) {
	mux.RLock()
	defer mux.RUnlock()
	g, ok = reg[slug]
	return
}

// Len gives the number of registered games.
func Len() int {
	mux.RLock()
	defer mux.RUnlock()
	return len(reg)
}

// Rand provides a random registerd game from the registry.
func Rand() string {
	mux.RLock()
	defer mux.RUnlock()
	for slug := range reg {
		return slug
	}
	return ""
}

// Map provides a duplicate of the internal game registry.
func Map() map[string]games.Game {
	mux.RLock()
	defer mux.RUnlock()
	res := make(map[string]games.Game, len(reg))
	for key, value := range reg {
		res[key] = value
	}
	return res
}

func init() {
	mux.Lock()
	defer mux.Unlock()
	register := func(g games.Game) {
		if _, ok := reg[g.Slug]; ok {
			panic(g.Slug + " already registered")
		}
		if err := g.Valid(); err == nil {
			reg[g.Slug] = g
		}
		// TODO: log invalid cases
	}
	register(cc.Game)
	register(gos.Game)
	register(ttt.Game)
	register(chess.Game)
	register(mancala.Game)
	register(checkers.Game)
	register(connect4.Game)
	register(cribbage.Game)
	register(backgammon.Game)
}
