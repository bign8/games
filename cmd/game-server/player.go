package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/bign8/games"
)

func play(slug string, x, y io.ReadWriteCloser) {
	fmt.Fprintln(x, "sFound one! Say hi.")
	fmt.Fprintln(y, "sFound one! Say hi.")

	// Convert actors to real players
	game := registry[slug]
	players := make([]games.Player, len(game.Players))

	// Generate Players
	// TODO: support more than 2-way communication (more players)
	ltr := make(chan string)
	rtl := make(chan string)
	errc := make(chan error, 1)
	a := newChatActor(x, ltr, rtl, errc, game.AI)
	b := newChatActor(y, rtl, ltr, errc, game.AI)
	players[0] = games.NewPlayer(a, game.Players[0])
	players[1] = games.NewPlayer(b, game.Players[1])

	// Play the game
	state := game.Start(players...)
	data := game4client(games.Run(state))
	a.out <- data
	b.out <- data

	// Log errors if necessary
	if err := <-errc; err != nil {
		log.Println(err)
	}
	x.Close()
	y.Close()
}

// func cp(w io.Writer, r io.Reader, errc chan<- error) {
// 	_, err := io.Copy(w, r)
// 	errc <- err
// }

type actor struct {
	moves chan string
	in    chan string
	out   chan<- string
	isBot bool
	ai    games.Actor
}

func newChatActor(s io.ReadWriteCloser, in chan string, out chan<- string, errc chan<- error, ai games.Actor) actor {
	_, ok := s.(bot)
	a := actor{make(chan string, 5), in, out, ok, ai}

	// TODO: combine these go-routines
	go func() {
		var msg string
		scanner := bufio.NewScanner(s)
		for scanner.Scan() {
			msg = scanner.Text()
			if msg[0] == 'g' {
				a.moves <- msg[1:]
			} else {
				out <- msg
			}
		}
		errc <- scanner.Err()
	}()

	go func() {
		for msg := range in {
			fmt.Fprint(s, msg)
		}
	}()

	return a
}

func (p actor) Act(state games.State) games.Action {
	if p.isBot {
		return p.ai.Act(state)
	}

	actions := state.Actions()
	p.in <- game4client(state)
	var move string
	var chosen *games.Action
	for chosen == nil {
		move = <-p.moves
		for _, a := range actions {
			if a.String() == move {
				chosen = &a
				break
			}
		}
		if chosen == nil {
			p.in <- "sInvalidMove... Try again!"
		}
	}
	return *chosen
}

type gameMSG struct {
	SVG   string
	Moves []gameMoveMSG
}

type gameMoveMSG struct {
	Name string
	SVG  string
}

func game4client(s games.State) string {
	moves := make([]gameMoveMSG, len(s.Actions()))

	for i, a := range s.Actions() {
		moves[i] = gameMoveMSG{
			Name: a.String(),
			SVG:  s.Apply(a).SVG(false),
		}
	}

	data := gameMSG{
		SVG:   s.SVG(true),
		Moves: moves,
	}
	js, _ := json.Marshal(data)
	return "g" + string(js)
}
