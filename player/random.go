package player

import (
	"math/rand"

	"github.com/bign8/games"
)

// Random creates a new player that interfaces with a human via Stdin/out/err
func Random(_ games.Game, name string) games.Actor {
	return func(s games.State) games.Action {
		acts := s.Actions()
		idx := rand.Intn(len(acts))
		return acts[idx]
	}
}
