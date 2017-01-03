package main

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/bign8/games"
)

type actor struct {
	name  string
	s     *bufio.Scanner
	write io.Writer
}

func newSocketActor(name string, s io.ReadWriteCloser, errc chan<- error) *actor {
	a := &actor{
		name:  name,
		s:     bufio.NewScanner(s),
		write: s,
	}
	return a
}

func (a *actor) Name() string {
	return a.name
}

func (a *actor) Act(s games.State) games.Action {
	actions := s.Actions()
	a.write.Write(game4client(s, false))
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
	Type string
	SVG  string
}

func game4client(s games.State, done bool) []byte {
	moves := make([]gameMoveMSG, len(s.Actions()))

	for i, a := range s.Actions() {
		moves[i] = gameMoveMSG{
			Name: a.String(),
			Type: a.Type(),
			SVG:  s.Apply(a).SVG(false),
		}
	}

	data := gameMSG{
		SVG:   s.SVG(!done),
		Moves: moves,
	}
	js, _ := json.Marshal(data)
	return js
}
