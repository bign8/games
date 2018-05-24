package player

import (
	"math/rand"

	"github.com/bign8/games"
)

// Random creates a new player that interfaces with a human via Stdin/out/err
func Random(s games.State) games.Action {
	acts := s.Actions()
	idx := rand.Intn(len(acts))
	return acts[idx]
}
