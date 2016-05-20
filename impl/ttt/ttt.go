package ttt

import (
	"fmt"

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
	switch m {
	case 0:
		return "Top Left Corner"
	case 1:
		return "Top Middle"
	case 2:
		return "Top Right Corner"
	case 3:
		return "Left Middle"
	case 4:
		return "Center Square"
	case 5:
		return "Right Middle"
	case 6:
		return "Bottom Left Corner"
	case 7:
		return "Bottom Middle"
	case 8:
		return "Bottom Right Corner"
	default:
		return fmt.Sprintf("Undefined Move: %d", m)
	}
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
	m := a.(tttMove)
	var board [9]byte
	copy(board[:], g.board[:])
	if g.ctr%2 == 0 {
		board[m] = 'X'
	} else {
		board[m] = 'O'
	}
	return &ttt{board, g.ctr + 1, g.players, nil}
}

func (g ttt) String() string {
	b := g.board
	return "╔═══╦═══╦═══╗\n║ " + string(b[0]) + " ║ " + string(b[1]) + " ║ " + string(b[2]) +
		" ║\n╠═══╬═══╬═══╣\n║ " + string(b[3]) + " ║ " + string(b[4]) + " ║ " + string(b[5]) +
		" ║\n╠═══╬═══╬═══╣\n║ " + string(b[6]) + " ║ " + string(b[7]) + " ║ " + string(b[8]) + " ║\n╚═══╩═══╩═══╝"
}

// Actions returns the next possible states given a particular state
func (g ttt) Actions() (moves []games.Action) {
	if !g.Player().Human() && g.ctr == 0 { // Starting move reducibility
		return []games.Action{tttMove(8), tttMove(7), tttMove(4)}
	}
	for j, bit := range g.board {
		if bit == ' ' {
			m := tttMove(j)
			moves = append(moves, m)
		}
	}
	return
}

// Terminal determines if we are currently in a winning state
// TODO: implement with bit masks
func (g ttt) Terminal() bool {
	if g.Error() != nil || g.ctr == 9 {
		return true
	}
	isWin, _ := g.isWin()
	return isWin
}

func (g ttt) isWin() (bool, byte) {
	// TODO: make this smarter
	chrs := g.board
	if chrs[0] != ' ' {
		if chrs[0] == chrs[1] && chrs[1] == chrs[2] { // top horiz
			return true, chrs[0]
		}
		if chrs[0] == chrs[3] && chrs[3] == chrs[6] { // left vert
			return true, chrs[0]
		}
		if chrs[0] == chrs[4] && chrs[4] == chrs[8] { // top-left to bot-right
			return true, chrs[0]
		}
	}
	if chrs[4] != ' ' {
		if chrs[3] == chrs[4] && chrs[4] == chrs[5] { // mid horiz
			return true, chrs[4]
		}
		if chrs[1] == chrs[4] && chrs[4] == chrs[7] { // mid vert
			return true, chrs[4]
		}
		if chrs[2] == chrs[4] && chrs[4] == chrs[6] { // top-right to bot-left
			return true, chrs[4]
		}
	}
	if chrs[8] != ' ' {
		if chrs[6] == chrs[7] && chrs[7] == chrs[8] { // bot horiz
			return true, chrs[8]
		}
		if chrs[2] == chrs[5] && chrs[5] == chrs[8] { // right vert
			return true, chrs[8]
		}
	}
	return false, ' '
}

func (g ttt) Utility() int {
	if isWin, chr := g.isWin(); !isWin {
		return 0
	} else if chr == 'X' {
		return 1
	}
	return -1
}

func (g ttt) SVG(active bool) string {
	// TODO: implement this function the right way
	return `
	<svg viewBox="0 0 100 100" xmlns="http://www.w3.org/2000/svg" xmlns:svg="http://www.w3.org/2000/svg">
		<path d="m62.193,11.333l0,24.785l-24,0l0,-24.785c0,-0.368 -0.112,-0.701 -0.293,-0.943c-0.181,-0.241 -0.431,-0.39 -0.707,-0.39c-0.552,0 -1,0.597 -1,1.333l0,24.785l-24.757,0c-0.367,0 -0.699,0.112 -0.94,0.293c-0.241,0.181 -0.39,0.431 -0.39,0.707c0,0.552 0.596,1 1.33,1l24.757,0l0,24l-24.757,0c-0.367,0 -0.699,0.112 -0.94,0.293c-0.241,0.181 -0.39,0.431 -0.39,0.707c0,0.552 0.596,1 1.33,1l24.757,0l0,24.549c0,0.368 0.112,0.701 0.293,0.943c0.181,0.241 0.431,0.39 0.707,0.39c0.552,0 1,-0.597 1,-1.333l0,-24.549l24,0l0,24.549c0,0.368 0.112,0.701 0.293,0.943c0.181,0.241 0.431,0.39 0.707,0.39c0.552,0 1,-0.597 1,-1.333l0,-24.549l24.372,0c0.367,0 0.699,-0.112 0.94,-0.293c0.24,-0.181 0.389,-0.431 0.389,-0.707c0,-0.552 -0.595,-1 -1.329,-1l-24.372,0l0,-24l24.372,0c0.367,0 0.699,-0.112 0.94,-0.293c0.24,-0.181 0.389,-0.431 0.389,-0.707c0,-0.552 -0.595,-1 -1.329,-1l-24.372,0l0,-24.785c0,-0.368 -0.112,-0.701 -0.293,-0.943c-0.181,-0.241 -0.431,-0.39 -0.707,-0.39c-0.552,0 -1,0.597 -1,1.333zm0,50.785l-24,0l0,-24l24,0l0,24z"/>
		<path fill="#ff0000" d="m38.775,76.269c0,6.306 5.112,11.418 11.418,11.418c6.306,0 11.418,-5.112 11.418,-11.418c0,-6.306 -5.112,-11.418 -11.418,-11.418c-6.306,0 -11.418,5.112 -11.418,11.418z"/>
		<path d="m14.476,66.598c0.823,-0.823 2.159,-0.824 2.983,0l6.687,6.687l6.687,-6.687c0.824,-0.824 2.16,-0.823 2.983,0c0.412,0.412 0.618,0.952 0.618,1.492c-0.001,0.539 -0.206,1.079 -0.618,1.492l-6.687,6.687l6.687,6.687c0.824,0.824 0.823,2.16 0,2.983c-0.412,0.412 -0.952,0.618 -1.492,0.619c-0.539,-0.001 -1.079,-0.206 -1.492,-0.619l-6.687,-6.687l-6.687,6.687c-0.824,0.824 -2.16,0.823 -2.983,0c-0.823,-0.823 -0.824,-2.159 0,-2.983l6.687,-6.687l-6.687,-6.687c-0.824,-0.825 -0.823,-2.16 0.001,-2.984z"/>
		<path d="m14.476,40.448c0.823,-0.824 2.159,-0.825 2.983,0l6.687,6.687l6.687,-6.687c0.824,-0.825 2.16,-0.823 2.983,0c0.412,0.412 0.618,0.952 0.618,1.492c-0.001,0.539 -0.206,1.079 -0.618,1.492l-6.687,6.687l6.687,6.687c0.824,0.824 0.823,2.16 0,2.983c-0.412,0.412 -0.952,0.618 -1.492,0.619c-0.539,-0.001 -1.079,-0.206 -1.492,-0.619l-6.687,-6.687l-6.687,6.687c-0.824,0.824 -2.16,0.823 -2.983,0c-0.823,-0.823 -0.824,-2.159 0,-2.983l6.687,-6.687l-6.687,-6.687c-0.823,-0.826 -0.823,-2.161 0.001,-2.984z"/>
		<path d="m40.523,14.395c0.823,-0.824 2.159,-0.825 2.983,0l6.687,6.687l6.687,-6.687c0.824,-0.825 2.16,-0.823 2.983,0c0.412,0.412 0.618,0.952 0.618,1.492c-0.001,0.539 -0.206,1.079 -0.618,1.492l-6.687,6.687l6.687,6.687c0.824,0.824 0.823,2.16 0,2.983c-0.412,0.412 -0.952,0.618 -1.492,0.619c-0.539,-0.001 -1.079,-0.206 -1.492,-0.619l-6.687,-6.687l-6.687,6.687c-0.824,0.824 -2.16,0.823 -2.983,0c-0.823,-0.823 -0.824,-2.159 0,-2.983l6.687,-6.687l-6.687,-6.687c-0.824,-0.826 -0.823,-2.161 0.001,-2.984z"/>
		<path d="m66.597,40.466c0.823,-0.824 2.159,-0.825 2.983,0l6.687,6.687l6.687,-6.687c0.824,-0.825 2.16,-0.823 2.983,0c0.412,0.412 0.618,0.952 0.618,1.492c-0.001,0.539 -0.206,1.079 -0.618,1.492l-6.687,6.687l6.687,6.687c0.824,0.824 0.823,2.16 0,2.983c-0.412,0.412 -0.952,0.618 -1.492,0.619c-0.539,-0.001 -1.079,-0.206 -1.492,-0.618l-6.687,-6.687l-6.687,6.687c-0.824,0.824 -2.16,0.823 -2.983,0c-0.823,-0.823 -0.824,-2.159 0,-2.983l6.687,-6.687l-6.687,-6.687c-0.824,-0.827 -0.823,-2.162 0.001,-2.985z"/>
		<path d="m66.474,66.598c0.823,-0.823 2.159,-0.824 2.983,0l6.687,6.687l6.687,-6.687c0.824,-0.824 2.16,-0.823 2.983,0c0.412,0.412 0.618,0.952 0.618,1.492c-0.001,0.539 -0.206,1.079 -0.618,1.492l-6.687,6.687l6.687,6.687c0.824,0.824 0.823,2.16 0,2.983c-0.412,0.412 -0.952,0.618 -1.492,0.619c-0.539,-0.001 -1.079,-0.206 -1.492,-0.619l-6.687,-6.687l-6.687,6.687c-0.824,0.824 -2.16,0.823 -2.983,0c-0.823,-0.823 -0.824,-2.159 0,-2.983l6.687,-6.687l-6.687,-6.687c-0.823,-0.825 -0.822,-2.16 0.001,-2.984z"/>
		<path d="m38.775,50.118c0,6.306 5.112,11.418 11.418,11.418c6.306,0 11.418,-5.112 11.418,-11.418c0,-6.306 -5.112,-11.418 -11.418,-11.418c-6.306,0 -11.418,5.112 -11.418,11.418z"/>
		<path d="m12.728,24.065c0,6.306 5.112,11.418 11.418,11.418c6.306,0 11.418,-5.112 11.418,-11.418c0,-6.306 -5.112,-11.418 -11.418,-11.418c-6.306,0 -11.418,5.112 -11.418,11.418z"/>
		<path d="m64.726,24.065c0,6.306 5.112,11.418 11.418,11.418s11.418,-5.112 11.418,-11.418c0,-6.306 -5.112,-11.418 -11.418,-11.418s-11.418,5.112 -11.418,11.418z"/>
	</svg>
	`
}

// Game is the fully described version of TTT
var Game = games.Game{
	Name:  "Tic-Tac-Toe",
	Slug:  "ttt",
	Start: New,
	Players: []games.PlayerConfig{
		games.PlayerConfig{
			Name: "X",
			Type: games.MaxPlayer,
		},
		games.PlayerConfig{
			Name: "O",
			Type: games.MinPlayer,
		},
	},
	Board: `<svg viewBox="0 0 100 100"><path d="m62.193,11.333l0,24.785l-24,0l0,-24.785c0,-0.368 -0.112,-0.701 -0.293,-0.943c-0.181,-0.241 -0.431,-0.39 -0.707,-0.39c-0.552,0 -1,0.597 -1,1.333l0,24.785l-24.757,0c-0.367,0 -0.699,0.112 -0.94,0.293c-0.241,0.181 -0.39,0.431 -0.39,0.707c0,0.552 0.596,1 1.33,1l24.757,0l0,24l-24.757,0c-0.367,0 -0.699,0.112 -0.94,0.293c-0.241,0.181 -0.39,0.431 -0.39,0.707c0,0.552 0.596,1 1.33,1l24.757,0l0,24.549c0,0.368 0.112,0.701 0.293,0.943c0.181,0.241 0.431,0.39 0.707,0.39c0.552,0 1,-0.597 1,-1.333l0,-24.549l24,0l0,24.549c0,0.368 0.112,0.701 0.293,0.943c0.181,0.241 0.431,0.39 0.707,0.39c0.552,0 1,-0.597 1,-1.333l0,-24.549l24.372,0c0.367,0 0.699,-0.112 0.94,-0.293c0.24,-0.181 0.389,-0.431 0.389,-0.707c0,-0.552 -0.595,-1 -1.329,-1l-24.372,0l0,-24l24.372,0c0.367,0 0.699,-0.112 0.94,-0.293c0.24,-0.181 0.389,-0.431 0.389,-0.707c0,-0.552 -0.595,-1 -1.329,-1l-24.372,0l0,-24.785c0,-0.368 -0.112,-0.701 -0.293,-0.943c-0.181,-0.241 -0.431,-0.39 -0.707,-0.39c-0.552,0 -1,0.597 -1,1.333zm0,50.785l-24,0l0,-24l24,0l0,24z"/></svg>`,
}

func init() {
	games.Register(Game)
}
