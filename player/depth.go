package player

import (
	"math/rand"

	"github.com/bign8/games"
)

// Depth walks `max` layers and chooses the utility that maximizes each outcome
func Depth(max int) games.Actor {
	return func(s games.State) games.Action {
		act, _ := depth(max, s)
		return act
	}
}

func depth(steps int, s games.State) (games.Action, int) {
	if steps <= 0 || s.Terminal() {
		return nil, s.Utility()[s.Player()]
	}
	a := s.Actions()
	score := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		_, score[i] = depth(steps-1, s.Apply(a[i]))
		score[i] *= -1 // allows us to always maximize (only works for 2 players)
	}

	// Find all the best moves possible
	idxs := []int{0}
	for i := 1; i < len(score); i++ {
		if score[i] > score[idxs[0]] {
			idxs = []int{i}
		} else if score[i] == score[idxs[0]] {
			idxs = append(idxs, i)
		}
	}

	// Choose randomly between them
	pos := idxs[rand.Intn(len(idxs))]
	return a[pos], score[pos]
}
