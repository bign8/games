package main

import (
	"fmt"
	"io/ioutil"

	"github.com/bign8/chess"
)

func main() {
	version, err := ioutil.ReadFile("VERSION")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Chess %s", version)
	game := chess.New()
	fmt.Printf("%s\n", game)
}
