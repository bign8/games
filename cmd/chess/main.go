package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

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

func main() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Chess %s", version)
	game := chess.New()
	fmt.Printf("%s\n", game)

	// // Start reading user input
	// reader := bufio.NewReader(os.Stdin)
	// loc := getLocation(reader, "Enter Piece Position > ")
	// fmt.Printf("You entered: %d %s\n", loc, loc.String())
	fmt.Printf("Moves: %+q\n", game.Moves())
}
