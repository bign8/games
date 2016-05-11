package search

import "github.com/bign8/games"

// MiniMax searches the full game tree until terminal nodes
func MiniMax(s games.State, u games.Utility) (games.Action, int) {
	if s.Terminal() {
		return nil, u(s)
	}
	compare := func(a, b int) bool { return a < b }
	if s.Player().Type() == games.MinPlayer {
		compare = func(a, b int) bool { return a > b }
	}

	actions := s.Actions()
	best := actions[0]
	_, cap := MiniMax(s.Apply(best), u)
	for _, a := range actions[1:] {
		if _, value := MiniMax(s.Apply(a), u); compare(cap, value) {
			best, cap = a, value
		}
	}
	return best, cap
}
