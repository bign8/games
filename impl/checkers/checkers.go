package checkers

import "github.com/bign8/games"

// Game is a fully configured checkers game.
var Game = games.Game{
	Name:    "Checkers",
	Slug:    "checkers",
	Board:   `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1"><rect x="-.1" y="-.1" width="8.2" height="8.2" fill="#999"/><path fill="#FFF" d="M0,0H8v1H0zm0,2H8v1H0zm0 2H8v1H0zm0,2H8v1H0zM1,0V8h1V0zm2,0V8h1V0zm2 0V8h1V0zm2,0V8h1V0z"/></svg>`,
	Players: []string{"Red", "Black"},
	Start:   nil,
	AI:      nil,
	Counts:  []uint8{2},
}
