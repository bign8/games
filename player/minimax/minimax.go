package minimax

import "github.com/bign8/games"

type minimax struct {
	name string
	ctr  uint
}

// New creates a new player that interfaces with a human via Stdin/out/err
func New(_ games.Game, name string) games.Actor {
	return &minimax{name: name, ctr: 0}
}

func (mm *minimax) Name() string {
	return mm.name
}

func (mm *minimax) Act(s games.State) games.Action {
	a, _ := mm.search(s)
	return a
}

func (mm *minimax) search(s games.State) (games.Action, int) {
	if s.Terminal() {
		return nil, s.Utility(s.Player())
	}
	a := s.Actions()
	score := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		_, score[i] = mm.search(s.Apply(a[i]))
		score[i] *= -1 // allows us to always maximize (only works for 2 players)
	}
	pos := 0
	for i := 1; i < len(score); i++ {
		if score[i] > score[pos] {
			pos = i
		}
	}
	return a[pos], score[pos]
}
