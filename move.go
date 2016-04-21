package chess

import (
	"fmt"
	"strings"
)

// Move represents a single move in the game of chess
type Move struct {
	start Location
	stop  Location
}

// String prints a human readable move
func (m Move) String() string {
	return fmt.Sprintf("%s -> %s", m.start, m.stop)
}

// Moves gives the list of possible moves to take given a state of the game
func (s State) Moves() []Move {
	if s.moves == nil {
		s.moves = make([]Move, 0)

		var newMoves []Move
		idx := byte(0)
		normalBoard := strings.ToLower(s.board)
		for i := 0; i < len(s.board); i++ {
			if '0' < s.board[i] && s.board[i] < '9' {
				idx += s.board[i] - '0'
			} else {
				switch normalBoard[i] {
				case 'p':
					newMoves = s.pawnMoves(i)
				case 'r':
					newMoves = s.rookMoves(i)
				case 'n':
					newMoves = s.knightMoves(i)
				case 'b':
					newMoves = s.bishopMoves(i)
				case 'q':
					newMoves = s.queenMoves(i)
				case 'k':
					newMoves = s.kingMoves(i)
				}
				s.moves = append(s.moves, newMoves...)
				idx++
			}
			// switch  {
			//
			// }
		}
	}
	return s.moves
}

func (s State) pawnMoves(i int) []Move {
	return make([]Move, 0)
}

func (s State) rookMoves(i int) []Move {
	return make([]Move, 0)
}

func (s State) knightMoves(i int) []Move {
	return make([]Move, 0)
}

func (s State) bishopMoves(i int) []Move {
	return make([]Move, 0)
}

func (s State) queenMoves(i int) []Move {
	return make([]Move, 0)
}

func (s State) kingMoves(i int) []Move {
	return make([]Move, 1)
}
