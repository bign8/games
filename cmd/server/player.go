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
			if a.Slug() == move {
				chosen = &a
				break
			}
		}
		if chosen == nil {
			a.write.Write([]byte("sInvalidMove... Try again!"))
		}
	}
	// TODO: handle scanner error!
	return *chosen
}

type gameMSG struct {
	SVG   string        `json:"SVG"`
	Moves []gameMoveMSG `json:"Moves,omitempty"`
}

type gameMoveMSG struct {
	Name string `json:"Name"`
	Type string `json:"Type"`
	Slug string `json:"Slug"`
}

func game4client(s games.State, done bool) []byte {
	moves := make([]gameMoveMSG, len(s.Actions()))

	for i, a := range s.Actions() {
		moves[i] = gameMoveMSG{
			Name: a.String(),
			Type: a.Type(),
			Slug: a.Slug(),
		}
	}

	data := gameMSG{
		SVG:   s.SVG(!done),
		Moves: moves,
	}
	js, err := json.Marshal(data)
	if err != nil {
		// TODO: handle error here
	}
	return js
}
