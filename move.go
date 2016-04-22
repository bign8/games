package chess

import (
	"fmt"
	"strings"
)

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
func (s State) Moves() []*Move {
	if s.moves == nil {
		idx := byte(0)
		s.moves = make([]*Move, 0)

		// map all pieces
		if s.pieces == nil {
			s.pieces = make(map[uint8]byte)
			for i := 0; i < len(s.board); i++ {
				if '0' < s.board[i] && s.board[i] < '9' {
					idx += s.board[i] - '0'
				} else {
					s.pieces[idx] = s.board[i]
					idx++
				}
			}
			idx = 0
		}

		var newMoves []*Move
		normalBoard := strings.ToLower(s.board)
		for i := 0; i < len(s.board); i++ {
			if '0' < s.board[i] && s.board[i] < '9' {
				idx += s.board[i] - '0'
			} else {
				if s.black(idx) == s.isBlack {
					next := Location(idx)
					switch normalBoard[i] {
					case 'p':
						newMoves = s.pawnMoves(next)
					case 'r':
						newMoves = s.rookMoves(next)
					case 'n':
						newMoves = s.knightMoves(next)
					case 'b':
						newMoves = s.bishopMoves(next)
					case 'q':
						newMoves = s.queenMoves(next)
					case 'k':
						newMoves = s.kingMoves(next)
					}
					s.moves = append(s.moves, newMoves...)
				}
				idx++
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
	if _, ok := s.pieces[move.toInt()]; !ok {
		res = append(res, NewMove(loc, move))
	}
	if _, ok := s.pieces[start.toInt()]; !ok && isStarting {
		res = append(res, NewMove(loc, start))
	}
	if _, ok := s.pieces[left.toInt()]; ok {
		res = append(res, NewMove(loc, left))
	}
	if _, ok := s.pieces[right.toInt()]; ok {
		res = append(res, NewMove(loc, right))
	}
	return res
}

func (s State) rookMoves(start Location) []*Move {
	return make([]*Move, 0)
}

func (s State) knightMoves(start Location) []*Move {
	return make([]*Move, 0)
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
	return 'a' < s.pieces[idx] && s.pieces[idx] < 'z'
}
