package player

import "github.com/bign8/games"

// Minimax creates a new player that interfaces with a human via Stdin/out/err
func Minimax(s games.State) games.Action {
	a, _ := search(s)
	return a
}

// TODO: fix this to work for more than 2 players
func search(s games.State) (games.Action, int) {
	if s.Terminal() {
		return nil, s.Utility()[s.Player()]
	}
	a := s.Actions()
	score := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		_, score[i] = search(s.Apply(a[i]))
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
