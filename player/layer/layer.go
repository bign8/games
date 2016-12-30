package layer

import (
	"math/rand"

	"github.com/bign8/games"
)

type actor string

func New(_ games.Game, name string) games.Actor {
	return actor(name)
}

func (a actor) Name() string { return string(a) }
func (a actor) Act(s games.State) games.Action {
	player := s.Player()
	actions := s.Actions()
	moves := []games.Action{actions[0]}
	value := s.Apply(actions[0]).Utility(player)
	for _, a := range actions {
		test := s.Apply(a).Utility(player)
		if test > value {
			moves = []games.Action{a}
			value = test
		} else if test == value {
			moves = append(moves, a)
		}
	}
	return moves[rand.Intn(len(moves))]
}
