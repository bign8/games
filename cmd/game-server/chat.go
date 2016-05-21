package main

// https://talks.golang.org/2012/chat.slide#41
// https://talks.golang.org/2012/chat/both/chat.go

import (
	"fmt"
	"io"
	"log"
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
	fmt.Fprint(c, "sWaiting for a partner...")
	select {
	case partner <- c:
		// now handled by the other goroutine
	case p := <-partner:
		play(slug, p, c)
	case <-time.After(5 * time.Second):
		play(slug, Bot(), c)
	}
}
