package app

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/bign8/games"
)

type actor struct {
	s     *bufio.Scanner
	write io.Writer
}

func newSocketActor(s io.ReadWriteCloser, errc chan<- error) games.Actor {
	a := &actor{
		s:     bufio.NewScanner(s),
		write: s,
	}
	return a.Act
}

func (a *actor) Act(s games.State) games.Action {
	actions := s.Actions()
	a.write.Write(game4client(s, false))
	var chosen games.Action
	for chosen == nil && a.s.Scan() {
		move := a.s.Text()
		for _, a := range actions {
			if a.Slug() == move {
				chosen = a
				break
			}
		}
		if chosen == nil {
			a.write.Write([]byte("sInvalidMove... Try again!"))
		}
	}
	// TODO: handle scanner error!
	return chosen
}

type gameMSG struct {
	SVG   string        `json:"svg"`
	Moves []gameMoveMSG `json:"moves,omitempty"`
}

type gameMoveMSG struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Slug string `json:"slug"`
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
