package ttt

import (
	"fmt"
	"strconv"

	"github.com/bign8/games"
)

type ttt struct {
	board   [9]byte
	ctr     uint8
	players []games.Player
	err     error
}

type tttMove uint8

// String does something
func (m tttMove) String() string {
	return "Position " + strconv.Itoa(int(m))
}

// Error tells if there is a problem with regular game play
func (g ttt) Error() error {
	return g.err
}

// Player returns the active player given a state
func (g ttt) Player() games.Player {
	return g.players[g.ctr%2]
}

// New takes creates a new game of ttt
func New(players ...games.Player) games.State {
	if len(players) != 2 {
		return &ttt{err: fmt.Errorf("invalid number of players: %d", len(players))}
	}
	var board [9]byte
	copy(board[:], "         ")
	return &ttt{board, 0, players, nil}
}

// Apply applies a given move to a state
func (g ttt) Apply(a games.Action) games.State {
	if g.Error() != nil {
		return g
	}
	// TODO: check for legal move
	m := a.(*tttMove)
	var board [9]byte
	copy(board[:], g.board[:])
	if g.ctr%2 == 0 {
		board[*m] = 'X'
	} else {
		board[*m] = 'O'
	}
	return &ttt{board, g.ctr + 1, g.players, nil}
}

func (g ttt) String() string {
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

// Actions returns the next possible states given a particular state
func (g ttt) Actions() (moves []games.Action) {
	for j, bit := range g.board {
		if bit == ' ' {
			m := tttMove(j)
			moves = append(moves, &m)
		}
	}
	return
}

// Terminal determines if we are currently in a winning state
// TODO: implement with bit masks
func (g ttt) Terminal() bool {
	if g.Error() != nil {
		return true
	}
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
