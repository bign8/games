package cc

import "github.com/bign8/games"

// Game is the fully described version of TTT
var Game = games.Game{
	Name:    "Chinese Checkers",
	Slug:    "cc",
	Board:   "<!-- TODO: board -->",
	Players: []string{"Black", "Blue", "White", "Red", "Green", "Yellow"},
	Start:   nil,
	AI:      nil,
}
