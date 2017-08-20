package connect4

import (
	"strconv"
	"strings"

	"github.com/bign8/games"
	"github.com/bign8/games/player/layer"
)

var (
	_ games.State  = (*c4)(nil)
	_ games.Action = (*c4move)(nil)

	// Game is the core game object that impl/impl.go works on stuff
	Game = games.Game{
		Name: "Connect 4",
		Slug: "c4",
		// TODO: https://upload.wikimedia.org/wikipedia/commons/d/dc/Puissance4_01.svg
		Board: `<svg viewBox="0 0 7 6" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
		<defs>
		  <circle id="o" cx=".5" cy=".5" r=".4" stroke="black" stroke-width=".05" fill="#d7d7d7" />
		</defs>
		<g>
		  <rect x="0" y="0" width="7" height="6" fill="#0070b9" />
		  <use xlink:href="#o" x="0" y="0"/>
		  <use xlink:href="#o" x="1" y="0"/>
		  <use xlink:href="#o" x="2" y="0"/>
		  <use xlink:href="#o" x="3" y="0"/>
		  <use xlink:href="#o" x="4" y="0"/>
		  <use xlink:href="#o" x="5" y="0"/>
		  <use xlink:href="#o" x="6" y="0"/>

		  <use xlink:href="#o" x="0" y="1"/>
		  <use xlink:href="#o" x="1" y="1"/>
		  <use xlink:href="#o" x="2" y="1"/>
		  <use xlink:href="#o" x="3" y="1"/>
		  <use xlink:href="#o" x="4" y="1"/>
		  <use xlink:href="#o" x="5" y="1"/>
		  <use xlink:href="#o" x="6" y="1"/>

		  <use xlink:href="#o" x="0" y="2"/>
		  <use xlink:href="#o" x="1" y="2"/>
		  <use xlink:href="#o" x="2" y="2"/>
		  <use xlink:href="#o" x="3" y="2"/>
		  <use xlink:href="#o" x="4" y="2"/>
		  <use xlink:href="#o" x="5" y="2"/>
		  <use xlink:href="#o" x="6" y="2"/>

		  <use xlink:href="#o" x="0" y="3"/>
		  <use xlink:href="#o" x="1" y="3"/>
		  <use xlink:href="#o" x="2" y="3"/>
		  <use xlink:href="#o" x="3" y="3"/>
		  <use xlink:href="#o" x="4" y="3"/>
		  <use xlink:href="#o" x="5" y="3"/>
		  <use xlink:href="#o" x="6" y="3"/>

		  <use xlink:href="#o" x="0" y="4"/>
		  <use xlink:href="#o" x="1" y="4"/>
		  <use xlink:href="#o" x="2" y="4"/>
		  <use xlink:href="#o" x="3" y="4"/>
		  <use xlink:href="#o" x="4" y="4"/>
		  <use xlink:href="#o" x="5" y="4"/>
		  <use xlink:href="#o" x="6" y="4"/>

		  <use xlink:href="#o" x="0" y="5"/>
		  <use xlink:href="#o" x="1" y="5"/>
		  <use xlink:href="#o" x="2" y="5"/>
		  <use xlink:href="#o" x="3" y="5"/>
		  <use xlink:href="#o" x="4" y="5"/>
		  <use xlink:href="#o" x="5" y="5"/>
		  <use xlink:href="#o" x="6" y="5"/>
		</g>
		</svg>
`,
		Players: []string{"Red", "Yellow"},
		Start:   New,
		AI:      layer.New,
	}
)

// New constructs a connect4 game
func New(actors ...games.Actor) games.State { return &c4{players: actors} }

type c4 struct {
	board   [7][]byte
	ctr     uint8
	players []games.Actor
}

func (s *c4) get(p int) byte {
	if int(master[p].row) < len(s.board[master[p].col]) {
		return s.board[master[p].col][master[p].row]
	}
	return ' '
}

func (s *c4) Actors() []games.Actor { return s.players }
func (s *c4) Player() int           { return int(s.ctr) % 2 }

func (s *c4) String() string {
	rows := make([]string, 0, 6)
	for i := 5; i >= 0; i-- {
		out := "|"
		for j := 0; j < 7; j++ {
			if len(s.board[j]) > i {
				switch s.board[j][i] {
				case 'R':
					out += "R"
				case 'Y':
					out += "Y"
				default:
					out += " "
				}
			} else {
				out += " "
			}
			out += "|"
		}
		rows = append(rows, out+"\n")
	}
	return "+-+-+-+-+-+-+-+\n" + strings.Join(rows, "+-+-+-+-+-+-+-+\n") + "+-+-+-+-+-+-+-+"
}

func (s *c4) Apply(action games.Action) games.State {
	a := action.(c4move)
	idx := int(uint8(a))
	next := &c4{
		ctr:     s.ctr + 1,
		players: s.players,
	}
	copy(next.board[:], s.board[:])
	next.board[idx] = append(next.board[idx], s.Actors()[s.Player()].Name()[0])
	return next
}

func (s *c4) Actions() []games.Action {
	acts := make([]games.Action, 0, 7)
	for i := 0; i < 7; i++ {
		if len(s.board[i]) < 6 {
			acts = append(acts, c4move(uint8(i)))
		}
	}
	return acts
}

type point struct{ col, row int8 }

// returns the value of the inARow array found (-1 if not found)
func isInARow(s *c4) int {
	for i := 0; i < len(master); i += 4 {
		a := s.get(i)
		b := s.get(i + 1)
		c := s.get(i + 2)
		d := s.get(i + 3)
		if a != ' ' && a == b && a == c && a == d {
			return i
		}
	}
	return -1
}

// Terminal checks if there exists a 4-in-a-row
func (s *c4) Terminal() bool {
	if len(s.Actions()) == 0 {
		return true
	}
	return isInARow(s) >= 0
}
func (s *c4) Utility() []int {
	val := isInARow(s)
	res := make([]int, 2)
	if val >= 0 {
		for i, a := range s.players {
			if s.get(val) == a.Name()[0] {
				res[i] = 1
			} else {
				res[i] = -1
			}
		}
	}
	return res
}

func (s *c4) SVG(bool) string {
	// TODO
	return ""
}

// Which column to drop the players token
type c4move uint8

func (move c4move) String() string { return "Column " + strconv.Itoa(int(move)+1) }
func (move c4move) Type() string   { return "" }
