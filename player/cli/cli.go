package cli

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bign8/games"
)

type cliPlayer struct {
	name   string
	reader *bufio.Reader
}

// New creates a new player that interfaces with a human via Stdin/out/err
func New(name string) games.Player {
	return &cliPlayer{
		name:   name,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (cli cliPlayer) Human() bool {
	return true
}

func (cli cliPlayer) String() string {
	return cli.name
}

func (cli cliPlayer) Play(s games.State) games.Action {
	moves := s.Actions()
	max := len(moves)

	// TODO: print state + moves side by side

	// Print state of the union
	fmt.Println("=================================================================\n" + s.String() + "\n" + cli.String() + "'s available moves:")
	for i, move := range moves {
		fmt.Printf("  %d: %s\n", i, move)
	}

	// Get Human move
	for {
		fmt.Printf("Choose your move [0-%d] > ", max-1)
		str, err := cli.reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Problem reading input:", err)
			continue
		}
		str = strings.Trim(str, "\r\n\t ")
		idx, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Problem processing number:", err)
			continue
		}
		if idx >= max {
			fmt.Fprintln(os.Stderr, "Value not permitted. Try again...")
			continue
		}
		return moves[idx]
	}
}
