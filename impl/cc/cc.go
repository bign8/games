package cc

import "github.com/bign8/games"

// Game is the fully described version of TTT
var Game = games.Game{
	Name:    "Chinese Checkers",
	Slug:    "cc",
	Board:   "<!-- TODO: board -->",
	Players: []string{"Black", "Blue", "White", "Red", "Green", "Yellow"},
	Counts:  []uint8{2, 3, 6}, // 2 players = 3 parts, 3 players = 2 parts, 6 players = 1 parts
	Start:   nil,
	AI:      nil,
}
