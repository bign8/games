package main

// https://talks.golang.org/2012/chat.slide#41
// https://talks.golang.org/2012/chat/both/chat.go

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/bign8/games"
	"github.com/bign8/games/player/cli"
	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type socket struct {
	io.Reader
	io.Writer
	done chan bool
}

func (s socket) Close() error {
	s.done <- true
	return nil
}

var chain = NewChain(2) // 2-word prefixes

func socketHandler(ws *websocket.Conn) {
	slug := mux.Vars(ws.Request())["slug"]
	// TODO: verify valid slug
	log.Printf("Socket connected of type: %s", slug)
	r, w := io.Pipe()
	go func() {
		_, err := io.Copy(io.MultiWriter(w, chain), ws)
		w.CloseWithError(err)
	}()
	s := socket{r, ws, make(chan bool)}
	go match(s, slug)
	<-s.done
}

var partner = make(chan io.ReadWriteCloser)

func match(c io.ReadWriteCloser, slug string) {
	fmt.Fprint(c, "Waiting for a partner...")
	select {
	case partner <- c:
		// now handled by the other goroutine
	case p := <-partner:
		play(slug, p, c)
	case <-time.After(5 * time.Second):
		play(slug, Bot(), c)
	}
}

func play(slug string, a, b io.ReadWriteCloser) {
	fmt.Fprintln(a, "Found one! Say hi.")
	fmt.Fprintln(b, "Found one! Say hi.")

	// TODO: actually initialize and run the game
	game := registry[slug]
	players := make([]games.Player, len(game.Players))
	for i, config := range game.Players {
		players[i] = cli.New(config.Name, config.Type)
	}
	state := game.Start(players[0], players[1])
	state = state.Apply(state.Actions()[5])
	state = state.Apply(state.Actions()[4])
	state = state.Apply(state.Actions()[3])
	svg := state.SVG(false)
	fmt.Fprintln(a, svg)
	fmt.Fprintln(b, svg)

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

// Bot returns an io.ReadWriteCloser that responds to
// each incoming write with a generated sentence.
func Bot() io.ReadWriteCloser {
	r, out := io.Pipe() // for outgoing data
	return bot{r, out}
}

type bot struct {
	io.ReadCloser
	out io.Writer
}

func (b bot) Write(buf []byte) (int, error) {
	go b.speak()
	return len(buf), nil
}

func (b bot) speak() {
	time.Sleep(time.Second)
	msg := chain.Generate(10) // at most 10 words
	b.out.Write([]byte(msg))
}
