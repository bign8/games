package minimax

import (
	"fmt"
	"time"

	"github.com/bign8/games"
)

type minimax struct {
	name string
	ctr  uint
}

// New creates a new player that interfaces with a human via Stdin/out/err
func New(name string) games.Actor {
	return &minimax{name: name, ctr: 0}
}

func (mm *minimax) Name() string {
	return mm.name
}

func (mm *minimax) Act(s games.State) games.Action {
	mm.ctr = 0
	start := time.Now()
	a, _ := mm.search(s)
	fmt.Printf(
		"%s chose %q after exploring %d games in %s\n",
		mm.name, a.String(), mm.ctr, time.Since(start),
	)
	return a
}

// MiniMax searches the full game tree until terminal nodes
func (mm *minimax) search(s games.State) (games.Action, int) {
	if s.Terminal() {
		mm.ctr++
		// fmt.Printf("%s - %d\n", s, s.Utility())
		return nil, s.Utility(s.Player())
	}

	actions := s.Actions()
	best := actions[0]
	_, min := mm.search(s.Apply(best))
	for _, a := range actions[1:] {
		if _, value := mm.search(s.Apply(a)); min > value {
			best, min = a, value
		}
	}
	return best, min
}
