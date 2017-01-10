package connect4

import (
	"strconv"

	"github.com/bign8/games"
)

var (
	_    games.State  = (*c4)(nil)
	_    games.Action = (*c4move)(nil)
	Game              = games.Game{
		Name: "Connect 4",
		Slug: "c4",
		// TODO: https://upload.wikimedia.org/wikipedia/commons/d/dc/Puissance4_01.svg
		Board: `<svg width="350" height="300" viewBox="1 1 7 6" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
		<defs>
		  <circle id="o" cx=".5" cy=".5" r=".4" stroke="black" stroke-width=".05" fill="#d7d7d7" />
		</defs>
		<g>
		  <rect x="0" y="0" width="8" height="7" fill="#0070b9" />
		  <use xlink:href="#o" x="1" y="1"/>
		  <use xlink:href="#o" x="2" y="1"/>
		  <use xlink:href="#o" x="3" y="1"/>
		  <use xlink:href="#o" x="4" y="1"/>
		  <use xlink:href="#o" x="5" y="1"/>
		  <use xlink:href="#o" x="6" y="1"/>
		  <use xlink:href="#o" x="7" y="1"/>

		  <use xlink:href="#o" x="1" y="2"/>
		  <use xlink:href="#o" x="2" y="2"/>
		  <use xlink:href="#o" x="3" y="2"/>
		  <use xlink:href="#o" x="4" y="2"/>
		  <use xlink:href="#o" x="5" y="2"/>
		  <use xlink:href="#o" x="6" y="2"/>
		  <use xlink:href="#o" x="7" y="2"/>

		  <use xlink:href="#o" x="1" y="3"/>
		  <use xlink:href="#o" x="2" y="3"/>
		  <use xlink:href="#o" x="3" y="3"/>
		  <use xlink:href="#o" x="4" y="3"/>
		  <use xlink:href="#o" x="5" y="3"/>
		  <use xlink:href="#o" x="6" y="3"/>
		  <use xlink:href="#o" x="7" y="3"/>

		  <use xlink:href="#o" x="1" y="4"/>
		  <use xlink:href="#o" x="2" y="4"/>
		  <use xlink:href="#o" x="3" y="4"/>
		  <use xlink:href="#o" x="4" y="4"/>
		  <use xlink:href="#o" x="5" y="4"/>
		  <use xlink:href="#o" x="6" y="4"/>
		  <use xlink:href="#o" x="7" y="4"/>

		  <use xlink:href="#o" x="1" y="5"/>
		  <use xlink:href="#o" x="2" y="5"/>
		  <use xlink:href="#o" x="3" y="5"/>
		  <use xlink:href="#o" x="4" y="5"/>
		  <use xlink:href="#o" x="5" y="5"/>
		  <use xlink:href="#o" x="6" y="5"/>
		  <use xlink:href="#o" x="7" y="5"/>

		  <use xlink:href="#o" x="1" y="6"/>
		  <use xlink:href="#o" x="2" y="6"/>
		  <use xlink:href="#o" x="3" y="6"/>
		  <use xlink:href="#o" x="4" y="6"/>
		  <use xlink:href="#o" x="5" y="6"/>
		  <use xlink:href="#o" x="6" y="6"/>
		  <use xlink:href="#o" x="7" y="6"/>
		</g>
		</svg>
`,
		Players: []string{"Red", "Yellow"},
		Start:   New,
		AI:      nil,
	}
)

func New(actors ...games.Actor) games.State {
	// FUTURE: can we assume this?
	if len(actors) != 2 {
		return games.StateErrInvalidNumberOfActors
	}
	return &c4{players: actors}
}

type c4 struct {
	board   [7][]byte
	ctr     uint8
	players []games.Actor
}

func (s *c4) get(p point) byte {
	if int(p.row) < len(s.board[p.col]) {
		return s.board[p.col][p.row]
	}
	return ' '
}

func (s *c4) Error() error        { return nil }
func (s *c4) Player() games.Actor { return s.players[s.ctr%2] }

func (s *c4) String() string {
	// TODO: improve the performance here
	out := "+-+-+-+-+-+-+-+\n"
	for i := 5; i >= 0; i-- {
		out += "|"
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
		out += "\n"
		if i != 0 {
			out += "+-+-+-+-+-+-+-+\n"
		}
	}
	return out + "+-+-+-+-+-+-+-+"
}

func (s *c4) Apply(action games.Action) games.State {
	a, ok := action.(c4move)
	if !ok {
		return games.StateErrInvalidMove
	}
	idx := int(uint8(a))
	next := &c4{
		ctr:     s.ctr + 1,
		players: s.players,
	}
	copy(next.board[:], s.board[:])
	next.board[idx] = append(next.board[idx], s.Player().Name()[0])
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

type point struct {
	col int8
	row int8
}

// returns the value of the inARow array found (-1 if not found)
func isInARow(s *c4) int {
outer:
	for i := 0; i < len(inARow); i++ {
		start := s.get(inARow[i][0])
		if start == ' ' {
			continue outer
		}
		for j := 1; j < len(inARow[i]); j++ {
			if start != s.get(inARow[i][j]) {
				continue outer
			}
		}
		return i
	}
	return -1
}

// Terminal checks if there exists a 4-in-a-row
func (s *c4) Terminal() bool {
	return isInARow(s) >= 0
}

func (s *c4) Utility(a games.Actor) int {
	val := isInARow(s)
	if val >= 0 {
		pt := inARow[val][0]
		if s.board[pt.col][pt.row] == a.Name()[0] {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func (s *c4) SVG(bool) string {
	// TODO
	return ""
}

// Which column to drop the players token
type c4move uint8

func (move c4move) String() string { return "Column " + strconv.Itoa(int(move)+1) }
func (move c4move) Type() string   { return "" }
