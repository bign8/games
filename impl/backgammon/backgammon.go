package backgammon

import "github.com/bign8/games"

// Game is the fully described version of TTT
var Game = games.Game{
	Name:    "Backgammon",
	Slug:    "backgammon",
	Board:   "<!-- TODO: board -->",
	Players: []string{"Black", "White"},
	Start:   nil,
	AI:      nil,
}
