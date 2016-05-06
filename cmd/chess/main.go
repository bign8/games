package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/bign8/chess"
)

var (
	version = "0.0.0.0"
	hash    = "zzzzz"
	slug    = "2016-05-01 UTC"
)

func getLocation(reader *bufio.Reader, prompt string) chess.Location {
	for {
		fmt.Printf(prompt)
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Problem reading input:", err)
			continue
		}

		out := chess.ParseLocation(strings.Trim(str, "\r\n\t "))
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

func getNumberLessThan(reader *bufio.Reader, prompt string, max int) int {
	for {
		fmt.Printf("%s [0-%d] > ", prompt, max-1)
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
		return idx
	}
}

func getMove(reader *bufio.Reader, state *chess.State) *chess.Move {

	// Generate input tree
	tree := make(map[chess.Location]map[chess.Location][]*chess.Move)
	for _, move := range state.Moves() {
		tier, ok := tree[move.Start]
		if !ok {
			tier = make(map[chess.Location][]*chess.Move)
			tree[move.Start] = tier
		}
		tier[move.Stop] = append(tier[move.Stop], move)
	}

	// TODO: auto-pick if only one source

	// Request Source
	set := make(map[chess.Location]struct{})
	for key := range tree {
		set[key] = struct{}{}
	}
	start := getLocationFromSet(reader, "Enter Starting Piece", set)

	// TODO: auto-pick if only one destination

	// Request Destination
	set = make(map[chess.Location]struct{})
	for key := range tree[start] {
		set[key] = struct{}{}
	}
	stop := getLocationFromSet(reader, "Enter Piece Destination", set)

	// Auto-pick if only one move
	moves := tree[start][stop]
	if len(moves) == 1 {
		fmt.Println("User Chose:", moves[0])
		return moves[0]
	}

	// Custom UI for choosing a duplicate ending move
	fmt.Println("ID\tMove")
	for i, m := range moves {
		fmt.Printf("%d\t%s\n", i, m)
	}
	idx := getNumberLessThan(reader, "Enter Move ID", len(moves))
	return moves[idx]
}

func main() {
	fmt.Printf("Chess Version: %s\nCommit Hash: %s\nBuild Stamp: %s\n\n", version, hash, slug)

	// Start printing input info
	game := chess.New()
	reader := bufio.NewReader(os.Stdin)

	// actual game runtime
	var err error
	for err == nil && !game.Terminal() {
		fmt.Printf("%s\n\n", game)
		fmt.Printf("\nMoves: %+q\n\n", game.Moves())
		move := getMove(reader, game)
		fmt.Printf("Used Move: %s\n", move)
		game, err = game.Apply(move)
	}

	// Print terminal message
	if game.Terminal() {
		fmt.Printf("Game Complete\n\n%s\n", game)
	}

	// Error handling
	if err != nil {
		fmt.Fprintln(os.Stderr, "Chess Error:", err)
	}
}
