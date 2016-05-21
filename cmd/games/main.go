package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bign8/games"
	"github.com/bign8/games/impl/ttt"
	"github.com/bign8/games/player/cli"
	"github.com/bign8/games/player/minimax"
)

var impl = map[string]games.Game{
	ttt.Game.Slug: ttt.Game,
}

func getImpl(in *bufio.Reader) games.Game {
	slugs := make([]string, 0, len(impl))
	for slug, config := range impl {
		fmt.Printf("\t%s: %s\n", slug, config.Name)
		slugs = append(slugs, slug)
	}
	for {
		fmt.Printf("Choose game (%s) > ", strings.Join(slugs, "|"))
		str, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Problem reading input:", err)
			continue
		}
		str = strings.Trim(str, "\r\n\t ")
		if config, ok := impl[str]; !ok {
			fmt.Fprintf(os.Stderr, "Input slug not found: %s\n", str)
		} else {
			return config
		}
	}
}

type playerConfig struct {
	name   string
	create func() games.Actor
}

var player = map[string]playerConfig{
	"cli": playerConfig{
		name:   "Human via Command Line",
		create: cli.New,
	},
	"mm": playerConfig{
		name:   "MiniMax Search",
		create: minimax.New,
	},
}

func getPlayer(in *bufio.Reader) playerConfig {
	slugs := make([]string, 0, len(player))
	for slug, config := range player {
		fmt.Printf("\t%s: %s\n", slug, config.name)
		slugs = append(slugs, slug)
	}
	for {
		fmt.Printf("Choose player (%s) > ", strings.Join(slugs, "|"))
		str, err := in.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Problem reading input:", err)
			continue
		}
		str = strings.Trim(str, "\r\n\t ")
		if config, ok := player[str]; !ok {
			fmt.Fprintf(os.Stderr, "Input slug not found: %s\n", str)
		} else {
			return config
		}
	}
}

func main() {
	// Choose Game
	in := bufio.NewReader(os.Stdin)
	fmt.Println("Choosing game to play")
	config := getImpl(in)

	// Setup Players
	// human := false
	count := len(config.Players)
	players := make([]games.Player, count)
	for i, p := range config.Players {
		fmt.Printf("=================================================================\nChoosing player %s (%d/%d)\n", p.Name, i+1, count)
		players[i] = games.NewPlayer(getPlayer(in).create(), p)
		// human = human || players[i].Human
	}

	// Play Game
	game := config.Start(players...)
	game = games.Run(game)

	// Print terminal message
	if game.Terminal() {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}

	// Print error message
	if game.Error() != nil {
		fmt.Fprintf(os.Stderr, "Error executing game: %s", game.Error())
		os.Exit(1)
	}
}
