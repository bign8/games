package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bign8/games"
)

func play(slug string, x, y io.ReadWriteCloser) {
	fmt.Fprintln(x, "sFound one! Say hi.")
	fmt.Fprintln(y, "sFound one! Say hi.")

	// Setup socket managers for sockets
	xMan := createManager(x)
	yMan := createManager(y)
	xChat := xMan.Room('u')
	yChat := yMan.Room('u')
	xGame := xMan.Room('g')
	yGame := yMan.Room('g')
	_, isBot := y.(bot)

	// Setup player chat-room // TODO: handle > 2 players
	errc := make(chan error, 1)
	go cp(xChat, yChat, errc)
	go cp(yChat, xChat, errc)

	// Convert actors to real players
	game := registry[slug]
	players := make([]games.Player, len(game.Players))
	players[0] = games.NewPlayer(newSocketActor(xGame, errc, false, game.AI), game.Players[0])
	players[1] = games.NewPlayer(newSocketActor(yGame, errc, isBot, game.AI), game.Players[1])

	// Play the game
	state := game.Start(players...)
	data := game4client(games.Run(state))
	xGame.Write(data) // Broadcast final game state
	yGame.Write(data)

	// Log errors if necessary
	if err := <-errc; err != nil {
		log.Println(err)
	}
	x.Close()
	y.Close()
}

func cp(w io.Writer, r io.Reader, errc chan<- error) {
	_, err := io.Copy(io.MultiWriter(w, chain), r) // copy chats to markov chain
	errc <- err
}

type actor struct {
	isBot bool
	ai    games.Actor
	s     *bufio.Scanner
	write io.Writer
}

func newSocketActor(s io.ReadWriteCloser, errc chan<- error, isBot bool, ai games.Actor) *actor {
	a := &actor{
		isBot: isBot,
		ai:    ai,
		s:     bufio.NewScanner(s),
		write: s,
	}
	return a
}

func (a *actor) Act(s games.State) games.Action {
	if a.isBot {
		time.Sleep(time.Second)
		return a.ai.Act(s)
	}

	actions := s.Actions()
	a.write.Write(game4client(s))
	var chosen *games.Action
	for chosen == nil && a.s.Scan() {
		move := a.s.Text()
		for _, a := range actions {
			if a.String() == move {
				chosen = &a
				break
			}
		}
		if chosen == nil {
			a.write.Write([]byte("sInvalidMove... Try again!"))
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

func game4client(s games.State) []byte {
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
	return js
}
