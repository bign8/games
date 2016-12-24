package mancala

import "github.com/bign8/games"

var Game = games.Game{
	Name:  "Mancala",
	Slug:  "mancala",
	Board: "<!-- TODO -->",
	Start: nil,
	Players: []games.PlayerConfig{
		{Name: "Left", Type: games.MaxPlayer},
		{Name: "Right", Type: games.MinPlayer},
	},
	AI: nil,
}
