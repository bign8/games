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

func getMove(reader *bufio.Reader, state *chess.State) *chess.Move {
	var start, stop chess.Location
	var found bool
	for {
		moves := state.Moves()

		// Convert starts of moves to a printable string
		set := make(map[string]struct{})
		for _, move := range moves {
			set[move.Start.String()] = struct{}{}
		}
		strMoves := make([]string, len(set))
		i := 0
		for key := range set {
			strMoves[i] = key
			i++
		}
		sort.Strings(strMoves)
		moveStarts := strings.Join(strMoves, ", ")

		// Get the first location from the user
		var short []*chess.Move
		start = getLocation(reader, fmt.Sprintf("Enter Starting Piece [%s] > ", moveStarts))
		for _, move := range state.Moves() {
			if move.Start == start {
				found = true
				short = append(short, move)
			}
		}
		if !found {
			fmt.Fprintln(os.Stderr, "Piece not found. Try again...")
			continue
		}

		// Convert stops of short list to a printable string
		set = make(map[string]struct{})
		for _, move := range short {
			set[move.Stop.String()] = struct{}{}
		}
		i = 0
		strMoves = make([]string, len(set))
		for key := range set {
			strMoves[i] = key
			i++
		}
		sort.Strings(strMoves)
		moveStops := strings.Join(strMoves, ", ")

		// Get the second location from the user
		stop = getLocation(reader, fmt.Sprintf("Enter Piece Destination [%s] > ", moveStops))
		for _, move := range short {
			if move.Stop == stop {
				return move
			}
		}
		fmt.Fprintln(os.Stderr, "Move not allowed. Try again...")
	}
}

func main() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Chess %s", version)
	game := chess.New()
	fmt.Printf("%s\n", game)

	// // Start reading user input
	reader := bufio.NewReader(os.Stdin)
	// loc := getLocation(reader, "Enter Piece Position > ")
	// fmt.Printf("You entered: %d %s\n", loc, loc.String())
	fmt.Printf("Moves: %+q\n", game.Moves())
	move := getMove(reader, game)
	fmt.Printf("Used Move: %s\n", move)
}
