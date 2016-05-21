package main

import (
	"fmt"
	"io"
	"log"

	"github.com/bign8/games"
)

func play(slug string, a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "sFound one! Say hi.")
	fmt.Fprintln(b, "sFound one! Say hi.")

	// TODO: actually initialize and run the game
	game := registry[slug]
	players := make([]games.Player, len(game.Players))
	for i, config := range game.Players {
		players[i] = games.NewPlayer(newChatActor(), config)
	}
	state := game.Start(players[0], players[1])
	state = state.Apply(state.Actions()[5])
	state = state.Apply(state.Actions()[4])
	state = state.Apply(state.Actions()[3])
	svg := state.SVG(false)
	fmt.Fprintln(a, "g"+svg)
	fmt.Fprintln(b, "g"+svg)

	// Start conversation between the players
	errc := make(chan error, 1)
	go cp(a, b, errc)
	go cp(b, a, errc)
	if err := <-errc; err != nil {
		log.Println(err)
	}
	a.Close()
	b.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(w, r)
	errc <- err
}

type actor struct{}

func newChatActor() games.Actor {
	return &actor{}
}

func (p actor) Act(state games.State) games.Action {
	return state.Actions()[0]
}
