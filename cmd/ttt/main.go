package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/bign8/games/ttt"
)

var (
	version = "0.0.0.0"
	hash    = "zzzzz"
	slug    = "2016-05-01 UTC"
)

func getMove(reader *bufio.Reader, state ttt.State) ttt.Move {
	moves := state.Moves()
	max := len(moves)

	fmt.Println("Moves:")
	for i, move := range moves {
		fmt.Printf("  %d: %s\n", i, move)
	}

	for {
		fmt.Printf("Choose your move [0-%d] > ", max-1)
		str, err := reader.ReadString('\n')
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

func main() {
	fmt.Printf("Tick-Tac-Toe Version: %s\nCommit Hash: %s\nBuild Stamp: %s\n\n", version, hash, slug)

	// Start printing input info
	game := ttt.New()
	reader := bufio.NewReader(os.Stdin)

	// actual game runtime
	for !game.Terminal() {
		fmt.Printf("%s\n\n", game)
		move := getMove(reader, game)
		fmt.Printf("Used Move: %s\n", move)
		game = game.Apply(move)
	}

	// Print terminal message
	if game.Terminal() {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}

	// Error handling
	err := game.Error()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Chess Error:", err)
	}
}
