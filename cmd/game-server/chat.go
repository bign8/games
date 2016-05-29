package main

// https://talks.golang.org/2012/chat.slide
// https://talks.golang.org/2012/chat/both/chat.go

import (
	"fmt"
	"io"
	"time"

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
	slug := mux.Vars(ws.Request())["slug"] // TODO: verify valid slug
	s := socket{ws, ws, make(chan bool)}
	go match(s, slug)
	<-s.done
}

var partner = make(chan io.ReadWriteCloser)

// TODO: allow this to handle games with more than 2 players
func match(c io.ReadWriteCloser, slug string) {
	fmt.Fprint(c, "sWaiting for a partner...")
	select {
	case partner <- c:
		// now handled by the other goroutine
	case p := <-partner:
		play(slug, p, c)
	case <-time.After(5 * time.Second):
		play(slug, c, Bot())
	}
}
