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
	a, _ := mm.search(s, s.Player())
	fmt.Printf(
		"%s chose %q after exploring %d games in %s\n",
		mm.name, a.String(), mm.ctr, time.Since(start),
	)
	return a
}

// MiniMax searches the full game tree until terminal nodes
func (mm *minimax) search(s games.State, p games.Actor) (games.Action, int) {
	if s.Terminal() {
		mm.ctr++
		u := s.Utility(p)
		// fmt.Printf("%s - %d for %s\n", s, u, p.Name())
		return nil, u
	}

	actions := s.Actions()
	myBest := actions[0]
	_, myScore := mm.search(s.Apply(myBest), p)
	bestScore := myScore
	for _, a := range actions[1:] {
		_, value := mm.search(s.Apply(a), p)
		myScore += value
		if value > bestScore {
			myBest = a
			bestScore = value
		}
	}
	return myBest, myScore
}
