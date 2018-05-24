package player

import (
	"math/rand"

	"github.com/bign8/games"
)

// Layer constructs a new layer player
func Layer(s games.State) games.Action {
	actions := s.Actions()
	moves := []games.Action{actions[0]}
	val := value(s, actions[0])
	for _, a := range actions {
		test := value(s, a)
		if test > val {
			moves = []games.Action{a}
			val = test
		} else if test == val {
			moves = append(moves, a)
		}
	}
	return moves[rand.Intn(len(moves))]
}

func value(s games.State, a games.Action) int { return s.Apply(a).Utility()[s.Player()] }
