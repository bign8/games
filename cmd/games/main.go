package main

import (
	"fmt"
	"os"

	"github.com/bign8/games"
	"github.com/bign8/games/ttt"
)

var all = map[string]games.Starter{
	// "Chess":        chess.New(),
	"Tick-Tac-Toe": ttt.New,
}

func main() {
	// TODO: pick game
	fn := all["Tick-Tac-Toe"]

	// TODO: setup players
	p1 := games.NewCLIPlayer("Player 1")
	p2 := games.NewCLIPlayer("Player 2")

	// Play Game
	game := fn(p1, p2)
	game = games.Run(game)

	// Print terminal message
	if game.Terminal() {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}

	// Print error message
	if game.Error() != nil {
		fmt.Fprintf(os.Stderr, "Error executing game: %s", game.Error())
	}
}
