package main

import (
	"fmt"
	"os"

	"github.com/bign8/games"
	"github.com/bign8/games/impl/ttt"
	"github.com/bign8/games/player/cli"
)

type implConfig struct {
	name  string
	start games.Starter
	names []string
}

var impl = map[string]implConfig{
	// "Chess":        chess.New,
	"ttt": implConfig{
		name:  "Tick-Tac-Toe",
		start: ttt.New,
		names: []string{"X", "O"},
	},
}

type playerConfig struct {
	name   string
	create games.Gamer
}

var player = map[string]playerConfig{
	"cli": playerConfig{
		name:   "Human via Command Line",
		create: cli.New,
	},
}

func main() {
	// TODO: pick game
	config := impl["ttt"]

	// TODO: setup players
	p1 := player["cli"].create(config.names[0])
	p2 := player["cli"].create(config.names[1])

	// Play Game
	game := config.start(p1, p2)
	game = games.Run(game)

	// Print terminal message
	if game.Terminal() && (p1.Human() || p2.Human()) {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}

	// Print error message
	if game.Error() != nil {
		fmt.Fprintf(os.Stderr, "Error executing game: %s", game.Error())
		os.Exit(1)
	}
}
