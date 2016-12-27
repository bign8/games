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
	a, _ := mm.search(s, s.Player(), 1)
	fmt.Printf(
		"%s chose %q after exploring %d games in %s\n",
		mm.name, a.String(), mm.ctr, time.Since(start),
	)
	return a
}

// MiniMax searches the full game tree until terminal nodes
func (mm *minimax) search(s games.State, p games.Actor, dist int) (games.Action, float32) {
	if s.Terminal() {
		mm.ctr++
		u := float32(s.Utility(p)) / float32(dist)
		fmt.Printf("%s - %f for %s\n", s, u, p.Name())
		return nil, u
	}

	actions := s.Actions()
	myBest := actions[0]
	_, myScore := mm.search(s.Apply(myBest), p, dist+1)
	bestScore := myScore
	for _, a := range actions[1:] {
		_, value := mm.search(s.Apply(a), p, dist+1)
		myScore += value
		if value > bestScore {
			myBest = a
			bestScore = value
			fmt.Printf("%d - %s updating to %f for %s\n", dist, a.String(), bestScore, p.Name())
		}
	}
	return myBest, myScore
}
