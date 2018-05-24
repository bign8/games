package random

import (
	"math/rand"

	"github.com/bign8/games"
)

type random struct {
	name string
}

// New creates a new player that interfaces with a human via Stdin/out/err
func New(_ games.Game, name string) games.Actor {
	return &random{name: name}
}

func (r *random) Name() string {
	return r.name
}

func (r *random) Act(s games.State) games.Action {
	acts := s.Actions()
	idx := rand.Intn(len(acts))
	return acts[idx]
}
