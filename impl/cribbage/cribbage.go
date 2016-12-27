package cribbage

import "github.com/bign8/games"

// Game is the fully described version of TTT
var Game = games.Game{
	Name:    "Cribbage",
	Slug:    "cribbage",
	Board:   "<!-- TODO: board -->",
	Players: []string{"Red", "Green", "Blue"},
	Counts:  []uint8{2, 3},
	Start:   nil,
	AI:      nil,
}
