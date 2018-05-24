package random

import (
	"math/rand"

	"github.com/bign8/games"
)

// New creates a new player that interfaces with a human via Stdin/out/err
func New(_ games.Game, name string) games.Actor {
	return func(s games.State) games.Action {
		acts := s.Actions()
		idx := rand.Intn(len(acts))
		return acts[idx]
	}
}
