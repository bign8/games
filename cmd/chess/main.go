package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/bign8/chess"
)

func getLocation(reader *bufio.Reader, prompt string) chess.Location {
	for {
		fmt.Printf(prompt)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Problem reading input:", err)
			continue
		}
		out := chess.ParseLocation(str)
		if out == chess.InvalidLocation {
			fmt.Fprintln(os.Stderr, "Invalid input. Try again...")
			continue
		}
		return out
	}
}

func getLocationFromSet(reader *bufio.Reader, prompt string, set map[chess.Location]struct{}) chess.Location {
	list := make([]string, len(set))
	i := 0
	for loc := range set {
		list[i] = loc.String()
		i++
	}
	sort.Strings(list)
	prompt = fmt.Sprintf("%s [%s] > ", prompt, strings.Join(list, ", "))

	loc := getLocation(reader, prompt)
	for _, ok := set[loc]; !ok; _, ok = set[loc] {
		fmt.Fprintln(os.Stderr, "Location not permitted. Try again...")
		loc = getLocation(reader, prompt)
	}
	return loc
}

func getMove(reader *bufio.Reader, state *chess.State) *chess.Move {
	// TODO: use better data structures here to improve lookup performance
	moves := state.Moves()

	// Convert starts of moves to a printable string
	set := make(map[chess.Location]struct{})
	for _, move := range moves {
		set[move.Start] = struct{}{}
	}
	start := getLocationFromSet(reader, "Enter Starting Piece", set)

	// Clip the list of moves based on starting input
	set = make(map[chess.Location]struct{})
	for _, move := range moves {
		if move.Start == start {
			set[move.Stop] = struct{}{}
		}
	}

	// Auto-pick move if only one available
	if len(set) == 1 {
		for _, move := range moves {
			if move.Start == start {
				return move
			}
		}
	}

	// Get stop of intended move
	stop := getLocationFromSet(reader, "Enter Piece Destination", set)
	for _, move := range moves {
		if move.Start == start && move.Stop == stop {
			return move
		}
	}
	return nil
}

func main() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}

	// Start printing input info
	fmt.Printf("Chess %s", version)
	game := chess.New()
	reader := bufio.NewReader(os.Stdin)

	// actual game runtime
	for err == nil {
		fmt.Printf("%s\n\n", game)
		fmt.Printf("\nMoves: %+q\n\n", game.Moves())
		move := getMove(reader, game)
		fmt.Printf("Used Move: %s\n", move)
		game, err = game.Apply(move)
	}

	// Error handling
	if err != nil {
		fmt.Fprintln(os.Stderr, "Chess Error:", err)
	}
}
