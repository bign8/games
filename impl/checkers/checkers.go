package checkers

import "github.com/bign8/games"

var Game = games.Game{
	Name:  "Checkers",
	Slug:  "checkers",
	Board: "<!-- TODO -->",
	Start: nil,
	Players: []games.PlayerConfig{
		{Name: "Red", Type: games.MaxPlayer},
		{Name: "Black", Type: games.MinPlayer},
	},
	AI: nil,
}
