package gos

import "github.com/bign8/games"

var Game = games.Game{
	Name:  "Go",
	Slug:  "go",
	Board: "<!-- TODO -->",
	Start: nil,
	Players: []games.PlayerConfig{
		{Name: "Black", Type: games.MaxPlayer},
		{Name: "White", Type: games.MinPlayer},
	},
	AI: nil,
}
