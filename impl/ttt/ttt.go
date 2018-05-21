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

// String does something
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

// https://thenounproject.com/term/tic-tac-toe/25029/

var svgHead = `<svg viewBox="0 0 100 100">`
var svgTail = `</svg>`
var svgPrefix = `<path d="m`
var svgOSuffix = `c0,6.306 5.112,11.418 11.418,11.418c6.306,0 11.418,-5.112 11.418,-11.418c0,-6.306 -5.112,-11.418 -11.418,-11.418c-6.306,0 -11.418,5.112 -11.418,11.418z"/>`
var svgXSuffix = `c0.823,-0.823 2.159,-0.824 2.983,0l6.687,6.687l6.687,-6.687c0.824,-0.824 2.16,-0.823 2.983,0c0.412,0.412 0.618,0.952 0.618,1.492c-0.001,0.539 -0.206,1.079 -0.618,1.492l-6.687,6.687l6.687,6.687c0.824,0.824 0.823,2.16 0,2.983c-0.412,0.412 -0.952,0.618 -1.492,0.619c-0.539,-0.001 -1.079,-0.206 -1.492,-0.619l-6.687,-6.687l-6.687,6.687c-0.824,0.824 -2.16,0.823 -2.983,0c-0.823,-0.823 -0.824,-2.159 0,-2.983l6.687,-6.687l-6.687,-6.687c-0.824,-0.825 -0.823,-2.161 0.001,-2.984z"/>`
var svgOPos = []string{
	`12.728,24.065`, `38.775,24.065`, `64.726,24.065`,
	`12.728,50.118`, `38.775,50.118`, `64.726,50.118`,
	`12.728,76.269`, `38.775,76.269`, `64.726,76.269`,
}
var svgXPos = []string{
	`14.476,14.395`, `40.523,14.395`, `66.597,14.395`,
	`14.476,40.466`, `40.523,40.466`, `66.597,40.466`,
	`14.476,66.598`, `40.523,66.598`, `66.597,66.598`,
}
var svgTargetPos = []string{
	`x="11" y="11"`, `x="38" y="11"`, `x="64" y="11"`,
	`x="11" y="38"`, `x="38" y="38"`, `x="64" y="38"`,
	`x="11" y="64"`, `x="38" y="64"`, `x="64" y="64"`,
}
var svgTargetID = []string{`p1`, `p2`, `p3`, `p4`, `p5`, `p6`, `p7`, `p8`, `p9`}
var svgBoard = `<svg viewBox="0 0 100 100">
<path d="m62.193,11.333l0,24.785l-24,0l0,-24.785c0,-0.368 -0.112,-0.701 -0.293,-0.943c-0.181,-0.241 -0.431,-0.39 -0.707,-0.39c-0.552,0 -1,0.597 -1,1.333l0,24.785l-24.757,0c-0.367,0 -0.699,0.112 -0.94,0.293c-0.241,0.181 -0.39,0.431 -0.39,0.707c0,0.552 0.596,1 1.33,1l24.757,0l0,24l-24.757,0c-0.367,0 -0.699,0.112 -0.94,0.293c-0.241,0.181 -0.39,0.431 -0.39,0.707c0,0.552 0.596,1 1.33,1l24.757,0l0,24.549c0,0.368 0.112,0.701 0.293,0.943c0.181,0.241 0.431,0.39 0.707,0.39c0.552,0 1,-0.597 1,-1.333l0,-24.549l24,0l0,24.549c0,0.368 0.112,0.701 0.293,0.943c0.181,0.241 0.431,0.39 0.707,0.39c0.552,0 1,-0.597 1,-1.333l0,-24.549l24.372,0c0.367,0 0.699,-0.112 0.94,-0.293c0.24,-0.181 0.389,-0.431 0.389,-0.707c0,-0.552 -0.595,-1 -1.329,-1l-24.372,0l0,-24l24.372,0c0.367,0 0.699,-0.112 0.94,-0.293c0.24,-0.181 0.389,-0.431 0.389,-0.707c0,-0.552 -0.595,-1 -1.329,-1l-24.372,0l0,-24.785c0,-0.368 -0.112,-0.701 -0.293,-0.943c-0.181,-0.241 -0.431,-0.39 -0.707,-0.39c-0.552,0 -1,0.597 -1,1.333zm0,50.785l-24,0l0,-24l24,0l0,24z"/>
<!--<text font-family="&#x27;Helvetica Neue', Helvetica, Arial-Unicode, Arial, Sans-serif" font-weight="bold" font-size="5px" fill="#000000" y="115" x="0">Created by TNS</text>
<text font-family="&#x27;Helvetica Neue', Helvetica, Arial-Unicode, Arial, Sans-serif" font-weight="bold" font-size="5px" fill="#000000" y="120" x="0">from the Noun Project</text>-->
</svg>`

func (g ttt) SVG(active bool) string {
	ctr := 0
	pieces := make([]string, 9)
	for i, bit := range g.board {
		if bit == 'X' {
			pieces[ctr] = svgPrefix + svgXPos[i] + svgXSuffix
			ctr++
		} else if bit == 'O' {
			pieces[ctr] = svgPrefix + svgOPos[i] + svgOSuffix
			ctr++
		}
	}
	pieces = pieces[:ctr]

	// Clickable targets
	var groups string
	if active {
		suffix := svgXSuffix
		pos := svgXPos
		if g.Actors()[g.Player()].Name() == "O" {
			suffix = svgOSuffix
			pos = svgOPos
		}
		ctr = 0
		hover, target := make([]string, 9), make([]string, 9)
		for i, bit := range g.board {
			if bit == ' ' {
				hover[ctr] = `<rect height="25" width="25" ` + svgTargetPos[i] + ` fill="transparent" ontouchend="` + games.SVGChooseMove + `('` + moveNames[i] + `')" onclick="` + games.SVGChooseMove + `('` + moveNames[i] + `')" onmouseover="` + svgTargetID[i] + `.setAttribute('opacity', '0.5')" onmouseout="` + svgTargetID[i] + `.setAttribute('opacity', '0')"/>`
				target[ctr] = `<path id="` + svgTargetID[i] + `" opacity="0" d="m` + pos[i] + suffix
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
	Board:   svgBoard,
	Players: []string{"X", "O"},
	Start:   New,
	AI:      minimax.New,
}
