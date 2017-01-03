package connect4

import "github.com/bign8/games"

var (
	_    games.State  = (*c4)(nil)
	_    games.Action = (*c4move)(nil)
	Game              = games.Game{
		Name:    "Connect 4",
		Slug:    "connect4",
		Board:   "<!-- TODO -->",
		Players: []string{"Red", "Black"},
		Start:   nil,
		AI:      nil,
	}
)

type c4 struct {
	board   []byte
	ctr     uint8
	players []games.Actor
	err     error
}

func (s *c4) String() string                 { return "TODO" }
func (s *c4) Player() games.Actor            { return nil }
func (s *c4) Apply(games.Action) games.State { return nil }
func (s *c4) Actions() []games.Action        { return nil }
func (s *c4) Terminal() bool                 { return true }
func (s *c4) Utility(games.Actor) int        { return 0 }
func (s *c4) Error() error                   { return s.err }
func (s *c4) SVG(bool) string                { return "" }

type c4move uint8

func (move *c4move) String() string { return "TODO" }
func (move *c4move) Type() string   { return "" }
