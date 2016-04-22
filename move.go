package chess

import "fmt"

// NewMove takes 2 chess locations and builds a move.
func NewMove(src, dst Location) *Move {
	return &Move{
		Start: src,
		Stop:  dst,
	}
}

// Move represents a single move in the game of chess
type Move struct {
	Start Location
	Stop  Location
}

// Equals checks if Moves are equal
func (m Move) Equals(n *Move) bool {
	return m.Start == n.Start && m.Stop == n.Stop
}

// String prints a human readable move
func (m Move) String() string {
	return fmt.Sprintf("%s -> %s", m.Start, m.Stop)
}

// Moves gives the list of possible moves to take given a state of the game
func (s *State) Moves() []*Move {
	if s.moves == nil {
		s.moves = make([]*Move, 0)
		var newMoves []*Move
		for idx := byte(0); idx < 64; idx++ {
			if s.piece(idx) && s.black(idx) == s.isBlack {
				next := Location(idx)
				switch s.board[idx] {
				case 'p':
					fallthrough
				case 'P':
					newMoves = s.pawnMoves(next)
				case 'r':
					fallthrough
				case 'R':
					newMoves = s.rookMoves(next)
				case 'n':
					fallthrough
				case 'N':
					newMoves = s.knightMoves(next)
				case 'b':
					fallthrough
				case 'B':
					newMoves = s.bishopMoves(next)
				case 'q':
					fallthrough
				case 'Q':
					newMoves = s.queenMoves(next)
				case 'k':
					fallthrough
				case 'K':
					newMoves = s.kingMoves(next)
				}
				s.moves = append(s.moves, newMoves...)
			}
		}
	}
	return s.moves
}

func (s State) pawnMoves(loc Location) (res []*Move) {
	// https://en.wikipedia.org/wiki/Pawn_(chess)
	var isStarting bool
	var move, start, left, right Location
	row, _ := loc.rowCol()
	if s.isBlack {
		move = loc.offset(1, 0)
		start = loc.offset(2, 0)
		left = loc.offset(1, -1)
		right = loc.offset(1, 1)
		isStarting = row == 1
	} else {
		move = loc.offset(-1, 0)
		start = loc.offset(-2, 0)
		left = loc.offset(-1, -1)
		right = loc.offset(-1, 1)
		isStarting = row == 6
	}

	if !s.piece(move.toInt()) {
		res = append(res, NewMove(loc, move))
	}
	if !s.piece(start.toInt()) && isStarting {
		res = append(res, NewMove(loc, start))
	}
	idx := left.toInt()
	if left != InvalidLocation && s.piece(idx) && s.black(idx) != s.isBlack {
		res = append(res, NewMove(loc, left))
	}
	idx = right.toInt()
	if right != InvalidLocation && s.piece(idx) && s.black(idx) != s.isBlack {
		res = append(res, NewMove(loc, right))
	}
	return res
}

func (s State) rookMoves(start Location) (res []*Move) {
	return res
}

func (s State) knightMoves(loc Location) (res []*Move) {
	x := []int8{1, 1, 2, 2, -1, -1, -2, -2}
	y := []int8{2, -2, 1, -1, 2, -2, 1, -1}
	for i := 0; i < len(x); i++ {
		if m := loc.offset(x[i], y[i]); m != InvalidLocation {
			idx := m.toInt()
			if s.piece(idx) && s.black(idx) == s.isBlack {
				continue
			}
			res = append(res, NewMove(loc, m))
		}
	}
	return res
}

func (s State) bishopMoves(start Location) []*Move {
	return make([]*Move, 0)
}

func (s State) queenMoves(start Location) []*Move {
	return make([]*Move, 0)
}

func (s State) kingMoves(start Location) []*Move {
	return make([]*Move, 0)
}

func (s State) black(idx uint8) bool {
	return 'a' < s.board[idx] && s.board[idx] < 'z'
}

func (s State) piece(idx uint8) bool {
	return s.board[idx] != '1'
}
