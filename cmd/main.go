package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/bign8/games"
	"github.com/bign8/games/impl"
	"github.com/bign8/games/player"
)

func getImpl(in *bufio.Reader) games.Game {
	fmt.Println("Choosing game to play")
	slugs := make([]string, 0, impl.Len())
	for slug, config := range impl.Map() {
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
		if config, ok := impl.Get(str); !ok {
			fmt.Fprintf(os.Stderr, "Input slug not found: %s\n", str)
		} else {
			return config
		}
	}
}

type playerConfig struct {
	name   string
	create games.Actor
}

func getPlayer(in *bufio.Reader) playerConfig {
	var player = map[string]playerConfig{
		"cli": playerConfig{
			name:   "Human via Command Line",
			create: player.CLI(in),
		},
		"mm": playerConfig{
			name:   "MiniMax Search",
			create: player.Minimax,
		},
		"lay": playerConfig{
			name:   "Layer Search",
			create: player.Layer,
		},
	}

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

func playerBuilder(in *bufio.Reader) func() games.Actor {
	return func() games.Actor {
		fmt.Print("=================================================================\nChoose player \n")
		return getPlayer(in).create
	}
}

func main() {
	// Choose Game
	in := bufio.NewReader(os.Stdin)

	// Play Game
	game := games.Run(getImpl(in), playerBuilder(in))

	// Print terminal message
	if nil != game.Utility() {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}
}
