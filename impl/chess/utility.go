package chess

import "github.com/bign8/games"

func (s State) Utility(a games.Actor) int {
	val := ValueUtility(s)
	if "Black" == a.Name() {
		val *= -1
	}
	return val
}

// ValueUtility is a uses the common standard value of pieces to rate a state.
// https://en.wikipedia.org/wiki/Chess_piece_relative_value
func ValueUtility(s State) int { // TODO: convert to games.State
	ctr := 0
	for _, square := range s.board {
		switch square {
		case 'P':
			ctr++
		case 'p':
			ctr--
		case 'N':
			fallthrough
		case 'B':
			ctr += 3
		case 'n':
			fallthrough
		case 'b':
			ctr -= 3
		case 'R':
			ctr += 5
		case 'r':
			ctr -= 5
		case 'Q':
			ctr += 9
		case 'q':
			ctr -= 9
		}
	}
	return ctr
}
