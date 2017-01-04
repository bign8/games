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
		Slug: "connect4",
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
		Start:   nil,
		AI:      nil,
	}
)

type c4 struct {
	board   [7][]byte
	ctr     uint8
	players []games.Actor
	err     error
}

func (s *c4) String() string                 { return "TODO" }
func (s *c4) Player() games.Actor            { return nil }
func (s *c4) Apply(games.Action) games.State { return nil }
func (s *c4) Actions() []games.Action {
	acts := make([]games.Action, 0, 7)
	for i := 0; i < 7; i++ {
		if len(s.board[i]) < 6 {
			acts = append(acts, c4move(uint8(i)))
		}
	}
	return acts
}
func (s *c4) Terminal() bool          { return true }
func (s *c4) Utility(games.Actor) int { return 0 }
func (s *c4) Error() error            { return s.err }
func (s *c4) SVG(bool) string         { return "" }

type c4move uint8

func (move c4move) String() string { return "Column " + strconv.Itoa(int(move)+1) }
func (move c4move) Type() string   { return "" }
