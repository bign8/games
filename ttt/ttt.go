package ttt

import "strconv"

// Player is the player for a game of tic-tac-toe
type Player bool

// Move is the move performed on a tic-tac-toe game
type Move uint8

// String does something
func (m Move) String() string {
	return "Position " + strconv.Itoa(int(m))
}

// State is the state of a tic-tac-toe game
type State struct {
	board  [9]byte
	player Player
}

// New takes creates a new game of ttt
func New() State {
	var board [9]byte
	copy(board[:], "         ")
	return State{board, false}
}

// Apply applies a given move to a state
func (g State) Apply(m Move) (State, error) {
	var board [9]byte
	copy(board[:], g.board[:])
	if g.player {
		board[m] = 'X'
	} else {
		board[m] = 'O'
	}
	res := State{board, !g.player}
	return res, nil
}

func (g State) String() string {
	b := g.board
	for i, x := range g.board {
		if x == ' ' {
			b[i] = strconv.Itoa(i)[0]
		}
	}
	return "╔═══╦═══╦═══╗\n║ " + string(b[0]) + " ║ " + string(b[1]) + " ║ " + string(b[2]) +
		" ║\n╠═══╬═══╬═══╣\n║ " + string(b[3]) + " ║ " + string(b[4]) + " ║ " + string(b[5]) +
		" ║\n╠═══╬═══╬═══╣\n║ " + string(b[6]) + " ║ " + string(b[7]) + " ║ " + string(b[8]) + " ║\n╚═══╩═══╩═══╝"
}

// Moves returns the next possible states given a particular state
func (g State) Moves() (moves []Move) {
	for j, bit := range g.board {
		if bit == ' ' {
			moves = append(moves, Move(j))
		}
	}
	return
}

// Terminal determines if we are currently in a winning state
// TODO: implement with bit masks
func (g State) Terminal() bool {
	// TODO: make this smarter
	// chrs := iToT(uint16(g.board))
	chrs := g.board
	if chrs[0] != ' ' {
		// p := sToPlayer(chrs[0])
		if chrs[0] == chrs[1] && chrs[1] == chrs[2] { // top horiz
			return true
		}
		if chrs[0] == chrs[3] && chrs[3] == chrs[6] { // left vert
			return true
		}
		if chrs[0] == chrs[4] && chrs[4] == chrs[8] { // top-left to bot-right
			return true
		}
	}
	if chrs[4] != ' ' {
		// p := sToPlayer(chrs[4])
		if chrs[3] == chrs[4] && chrs[4] == chrs[5] { // mid horiz
			return true
		}
		if chrs[1] == chrs[4] && chrs[4] == chrs[7] { // mid vert
			return true
		}
		if chrs[2] == chrs[4] && chrs[4] == chrs[6] { // top-right to bot-left
			return true
		}
	}
	if chrs[8] != ' ' {
		// p := sToPlayer(chrs[8])
		if chrs[6] == chrs[7] && chrs[7] == chrs[8] { // bot horiz
			return true
		}
		if chrs[2] == chrs[5] && chrs[5] == chrs[8] { // right vert
			return true
		}
	}
	return false
}
