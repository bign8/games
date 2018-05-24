package player

import (
	"bufio"
	"encoding/json"
	"io"

	"github.com/bign8/games"
)

// Socket creates a websocket based actor
func Socket(sock io.ReadWriteCloser, errc chan<- error) games.ActorBuilder {
	scanner := bufio.NewScanner(sock)
	return func(name string) games.Actor {
		return func(s games.State) games.Action {
			actions := s.Actions()
			sock.Write(ToJSON(s, false))
			var chosen games.Action
			for chosen == nil && scanner.Scan() {
				move := scanner.Text()
				for _, a := range actions {
					if a.Slug() == move {
						chosen = a
						break
					}
				}
				if chosen == nil {
					sock.Write([]byte("sInvalidMove... Try again!"))
				}
			}
			// TODO: handle scanner error!
			return chosen
		}

	}
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

func ToJSON(s games.State, done bool) []byte {
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
