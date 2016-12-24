package impl

import (
	"sync"

	"github.com/bign8/games"
	"github.com/bign8/games/impl/checkers"
	"github.com/bign8/games/impl/chess"
	"github.com/bign8/games/impl/connect4"
	gos "github.com/bign8/games/impl/go"
	"github.com/bign8/games/impl/mancala"
	"github.com/bign8/games/impl/ttt"
)

var (
	reg = map[string]games.Game{}
	mux sync.RWMutex
)

func Get(slug string) (g games.Game, ok bool) {
	mux.RLock()
	defer mux.RUnlock()
	g, ok = reg[slug]
	return
}

func Len() int {
	mux.RLock()
	defer mux.RUnlock()
	return len(reg)
}

func Rand() string {
	mux.RLock()
	defer mux.RUnlock()
	for slug := range reg {
		return slug
	}
	return ""
}

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
	reg["go"] = gos.Game
	reg["ttt"] = ttt.Game
	reg["chess"] = chess.Game
	reg["mancala"] = mancala.Game
	reg["checkers"] = checkers.Game
	reg["connect4"] = connect4.Game
}
