package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bign8/games"
)

// CLI creates a new player that interfaces with a human via Stdin/out/err
func CLI(buf *bufio.Reader) games.Actor {
	return func(s games.State) games.Action {
		moves := s.Actions()
		max := len(moves)

		// TODO: print state + moves side by side

		// Print state of the union
		fmt.Println("=================================================================\n" + s.String() + "\nAvailable moves:")
		for i, move := range moves {
			fmt.Printf("  %d: %s\n", i, move)
		}

		// Get Human move
		for {
			fmt.Printf("Choose your move [0-%d] > ", max-1)
			str, err := buf.ReadString('\n')
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
}
