package minimax

import (
	"fmt"
	"time"

	"github.com/bign8/games"
)

type minimax struct {
	ctr uint
}

// New creates a new player that interfaces with a human via Stdin/out/err
func New() games.Actor {
	return &minimax{ctr: 0}
}

func (mm *minimax) Act(s games.State) games.Action {
	mm.ctr = 0
	start := time.Now()
	a, _ := mm.search(s)
	fmt.Printf(
		"%s chose %q after exploring %d games in %s\n",
		s.Player().Name, a.String(), mm.ctr, time.Since(start),
	)
	return a
}

// MiniMax searches the full game tree until terminal nodes
func (mm *minimax) search(s games.State) (games.Action, int) {
	if s.Terminal() {
		mm.ctr++
		// fmt.Printf("%s - %d\n", s, s.Utility())
		return nil, s.Utility()
	}
	compare := func(a, b int) bool { return a < b }
	if s.Player().Type == games.MinPlayer {
		compare = func(a, b int) bool { return a > b }
	}

	actions := s.Actions()
	best := actions[0]
	_, cap := mm.search(s.Apply(best))
	for _, a := range actions[1:] {
		if _, value := mm.search(s.Apply(a)); compare(cap, value) {
			best, cap = a, value
		}
	}
	return best, cap
}
