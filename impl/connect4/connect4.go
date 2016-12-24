package connect4

import "github.com/bign8/games"

var Game = games.Game{
	Name:  "Connect 4",
	Slug:  "connect-4",
	Board: "<!-- TODO -->",
	Start: nil,
	Players: []games.PlayerConfig{
		{Name: "Red", Type: games.MaxPlayer},
		{Name: "Black", Type: games.MinPlayer},
	},
	AI: nil,
}
