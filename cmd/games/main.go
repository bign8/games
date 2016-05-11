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

type implConfig struct {
	name  string
	start games.Starter
	names []string
	types []games.PlayerType
}

var impl = map[string]implConfig{
	// "Chess":        chess.New,
	"ttt": implConfig{
		name:  "Tick-Tac-Toe",
		start: ttt.New,
		names: []string{"X", "O"},
		types: []games.PlayerType{games.MaxPlayer, games.MinPlayer},
	},
}

func getImpl(in *bufio.Reader) implConfig {
	slugs := make([]string, 0, len(impl))
	for slug, config := range impl {
		fmt.Printf("\t%s: %s\n", slug, config.name)
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
	create games.Gamer
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
	human := false
	players := make([]games.Player, len(config.names))
	for i, name := range config.names {
		fmt.Printf("=================================================================\nChoosing player %s (%d/%d)\n", name, i+1, len(config.names))
		players[i] = getPlayer(in).create(name, config.types[i])
		human = human || players[i].Human()
	}

	// Play Game
	game := config.start(players...)
	game = games.Run(game)

	// Print terminal message
	if game.Terminal() && human {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}

	// Print error message
	if game.Error() != nil {
		fmt.Fprintf(os.Stderr, "Error executing game: %s", game.Error())
		os.Exit(1)
	}
}
