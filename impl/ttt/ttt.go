package ttt

import (
	"fmt"
	"strings"

	"github.com/bign8/games"
	"github.com/bign8/games/player/minimax"
)

type ttt struct {
	board   [9]byte
	ctr     uint8
	players []games.Actor
}

type tttMove uint8

var moveNames = []string{
	"Top Left Corner", "Top Middle", "Top Right Corner",
	"Left Middle", "Center Square", "Right Middle",
	"Bottom Left Corner", "Bottom Middle", "Bottom Right Corner",
}

func (m tttMove) Type() string { return "" }
func (m tttMove) Slug() string { return svgName[m] }
func (m tttMove) String() string {
	if m > 8 {
		return fmt.Sprintf("Undefined Move: %d", m)
	}
	return moveNames[m]
}
func (g ttt) Actors() []games.Actor { return g.players }
func (g ttt) Player() int           { return int(g.ctr) % 2 }

// New takes creates a new game of ttt
func New(players ...games.Actor) games.State {
	if len(players) != 2 {
		panic(fmt.Sprintf("invalid number of players: %d", len(players)))
	}
	var board [9]byte
	copy(board[:], "         ")
	return &ttt{board, 0, players}
}

// Apply applies a given move to a state
func (g ttt) Apply(a games.Action) games.State {
	// TODO: check for legal move
	m := a.(tttMove)
	var board [9]byte
	copy(board[:], g.board[:])
	if g.ctr%2 == 0 {
		board[m] = 'X'
	} else {
		board[m] = 'O'
	}
	return &ttt{board, g.ctr + 1, g.players}
}

func (g ttt) String() string {
	b := g.board
	return "╔═══╦═══╦═══╗\n║ " + string(b[0]) + " ║ " + string(b[1]) + " ║ " + string(b[2]) +
		" ║\n╠═══╬═══╬═══╣\n║ " + string(b[3]) + " ║ " + string(b[4]) + " ║ " + string(b[5]) +
		" ║\n╠═══╬═══╬═══╣\n║ " + string(b[6]) + " ║ " + string(b[7]) + " ║ " + string(b[8]) + " ║\n╚═══╩═══╩═══╝"
}

// Actions returns the next possible states given a particular state
func (g ttt) Actions() (moves []games.Action) {
	if g.Terminal() {
		return nil
	}
	// if !g.Player().Human() && g.ctr == 0 { // Starting move reducibility
	// 	return []games.Action{tttMove(8), tttMove(7), tttMove(4)}
	// }
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
	if g.ctr == 9 {
		return true
	}
	isWin, _ := isWin(g.board)
	return isWin
}

func isWin(chrs [9]byte) (bool, byte) {
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

// Utility for TTT is simple: 1 for a win, -1 for a loss, 0 if game is in progress
func (g ttt) Utility() []int {
	res := make([]int, 2)
	if isWin, chr := isWin(g.board); isWin {
		for i, a := range g.players {
			if chr == a.Name()[0] {
				res[i] = 1
			} else {
				res[i] = -1
			}
		}
	}
	return res
}

const (
	svgHead = `<svg viewBox="0 0 90 90" stroke="black" stroke-linecap="round">`
	svgTail = `</svg>`
	svgPath = `<path d="M`
	svgXend = ` m10,10 l-20,-20 m20,0 l-20,20" stroke-width="4" />`
	svgOend = ` m-12,0 a12,12 0 1,0 24,0 a12,12 0 1,0 -24,0" />`
	svgRect = ` m-12,-12 h24 v24 h-24 z" `
	svgGame = `<svg viewBox="-1 -1 92 92" stroke-width="2" stroke="black" stroke-linecap="round">
	<line x1="30" y1="00" x2="30" y2="90" />
	<line x1="60" y1="00" x2="60" y2="90" />
	<line x1="00" y1="30" x2="90" y2="30" />
	<line x1="00" y1="60" x2="90" y2="60" />
</svg>`
)

var (
	svgName = []string{`p1`, `p2`, `p3`, `p4`, `p5`, `p6`, `p7`, `p8`, `p9`}
	svgSpot = []string{
		`15,15`, `45,15`, `75,15`,
		`15,45`, `45,45`, `75,45`,
		`15,75`, `45,75`, `75,75`,
	}
)

func (g ttt) SVG(active bool) string {
	ctr := 0
	pieces := make([]string, 9)
	for i, bit := range g.board {
		if bit == 'X' {
			pieces[ctr] = svgPath + svgSpot[i] + svgXend
			ctr++
		} else if bit == 'O' {
			pieces[ctr] = svgPath + svgSpot[i] + svgOend
			ctr++
		}
	}
	pieces = pieces[:ctr]

	// Clickable targets
	var groups string
	if active {
		suffix := svgOend
		if int(g.ctr)%2 == 0 {
			suffix = svgXend
		}
		ctr = 0
		hover, target := make([]string, 9), make([]string, 9)
		for i, bit := range g.board {
			if bit == ' ' {
				hover[ctr] = svgPath + svgSpot[i] + svgRect + ` data-show="` + svgName[i] + `" />`
				target[ctr] = `<path data-slug="` + svgName[i] + `" d="m` + svgSpot[i] + suffix
				ctr++
			}
		}
		groups = "<g>" + strings.Join(target[:ctr], "") + "</g><g>" + strings.Join(hover[:ctr], "") + "</g>"
	}
	return svgHead + "<g>" + strings.Join(pieces, "") + "</g>" + groups + svgTail
}

// Game is the fully described version of TTT
var Game = games.Game{
	Name:    "Tic-Tac-Toe",
	Slug:    "ttt",
	Board:   svgGame,
	Players: []string{"X", "O"},
	Start:   New,
	AI:      minimax.New,
}
