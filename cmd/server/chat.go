package main

// https://talks.golang.org/2012/chat.slide
// https://talks.golang.org/2012/chat/both/chat.go

import (
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
)

type socket struct {
	io.ReadWriteCloser
	done chan bool
}

func (s socket) Close() error {
	s.ReadWriteCloser.Close()
	s.done <- true
	return nil
}

var chain = NewChain(2) // 2-word prefixes

func socketHandler(ws *websocket.Conn) {
	slug := mux.Vars(ws.Request())["slug"] // TODO: verify valid slug
	s := socket{ws, make(chan bool)}
	go match(s, slug)
	<-s.done
}

type roomManager struct {
	rooms map[string]chan io.ReadWriteCloser
	mu    sync.Mutex
}

func (rm *roomManager) Get(slug string) chan io.ReadWriteCloser {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	if rm.rooms == nil {
		rm.rooms = make(map[string]chan io.ReadWriteCloser)
	}
	cha, ok := rm.rooms[slug]
	if !ok {
		cha = make(chan io.ReadWriteCloser)
		rm.rooms[slug] = cha
	}
	return cha
}

var partner = &roomManager{}

// TODO: allow this to handle games with more than 2 players
func match(c io.ReadWriteCloser, slug string) {
	fmt.Fprint(c, "sWaiting for a partner...")
	cha := partner.Get(slug)
	select {
	case cha <- c:
		// now handled by the other goroutine
	case p := <-cha:
		play(slug, p, c)
	case <-time.After(5 * time.Second):
		play(slug, c, Bot())
	}
}
