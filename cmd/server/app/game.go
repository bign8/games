package app

import (
	"flag"
	"fmt"
	"io"
	"log"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"

	"github.com/bign8/games"
	"github.com/bign8/games/impl"
	"github.com/bign8/games/util/socket"
)

var pool = &poolManager{games: make(map[string]chan<- *socket.Socket)}
var tout = flag.Duration("tout", 5*time.Second, "matcher timeout")

// Socket is called when a client wants to start a game
func Socket(ws *websocket.Conn) {
	slug := mux.Vars(ws.Request())["slug"] // TODO: verify valid slug
	s := socket.New(ws)
	pool.Match(s, slug)
	<-s.Done
}

// poolManager is designed to setup games as game requests come in
type poolManager struct {
	games map[string]chan<- *socket.Socket
	mu    sync.Mutex
}

// Match puts a socket request into a pool of incomming start game requests
func (pm *poolManager) Match(s *socket.Socket, slug string) {
	fmt.Fprint(s.Room('s'), "Waiting for a partner...")
	pm.mu.Lock()
	ch, ok := pm.games[slug]
	if !ok {
		cha := make(chan *socket.Socket) // new variable because chan types
		pm.games[slug] = cha
		ch = cha
		go pm.pair(slug, cha)
	}
	pm.mu.Unlock()
	ch <- s
}

// pair attempt to create games based on incomming start game requests
// TODO: terminate this go-routine if it runs for a long amount of time
func (pm *poolManager) pair(slug string, ch <-chan *socket.Socket) {
	game, _ := impl.Get(slug) // TODO: panic if game hasn't been found
	queue := make([]*socket.Socket, 0, len(game.Players))
	var wait <-chan time.Time

	// Fix Case where Counts is not set on the registered game object
	if len(game.Counts) == 0 {
		game.Counts = append(game.Counts, uint8(len(game.Players)))
	}
	max := game.Counts[len(game.Counts)-1]

	// Start the game and reset queue
	start := func() {
		go play(game, queue...)
		queue = make([]*socket.Socket, 0, len(game.Players))
		wait = nil
	}

	// Actual pairing loop
	for {
		select {
		case p := <-ch: // Incomming game start request
			queue = append(queue, p)
			if uint8(len(queue)) == max {
				start()
			} else if wait == nil {
				wait = time.After(*tout) // Max wait of first to enter a room
			}
		case <-wait: // first person in queue has been waiting for a while
			start()
		}
	}
}

func play(game games.Game, players ...*socket.Socket) {
	isBot := padPlayers(game, &players)
	chats := make([]io.ReadWriteCloser, len(players))
	gamez := make([]io.ReadWriteCloser, len(players))
	systm := make([]io.ReadWriteCloser, len(players))
	for i, p := range players {
		chats[i] = p.Room('u')
		gamez[i] = p.Room('g')
		systm[i] = p.Room('s')
	}

	// Setup all gamers to communicate with each other
	// TODO: make chats much less go-routine heavy
	// TODO: use multi-writers here
	errc := make(chan error, 1)
	for i := 0; i < len(players); i++ {
		for j := i + 1; j < len(players); j++ {
			go cp(chats[i], chats[j], errc)
			go cp(chats[j], chats[i], errc)
		}
		fmt.Fprintln(systm[i], "Found one! Say hi.")
	}

	// Setup the player builder
	i := -1
	builder := func(g games.Game, name string) games.Actor {
		i++
		if isBot[i] {
			return g.AI(g, name)
		}
		return newSocketActor(name, gamez[i], errc)
	}

	// Play the game (and broadcast final state)
	data := game4client(games.Run(game, builder), true)
	for _, g := range gamez {
		g.Write(data)
	}

	// Log errors if necessary
	if err := <-errc; err != nil {
		log.Println("play", err)
	}

	// TODO: close down sockets
}

// padPlayers adds bots to the array of players
// returns an array of bools telling if players are bots or not
// TODO: make this really work
func padPlayers(game games.Game, players *[]*socket.Socket) []bool {
	// Find the first number of players larger than the current queue size
	cnt := game.Counts[0]
	for _, c := range game.Counts[1:] {
		cnt = c
		if c >= uint8(len(*players)) {
			break
		}
	}
	isBot := make([]bool, cnt)

	// Make bots to fill the void between desired size and number in queue
	length := uint8(len(*players))
	if length < cnt {
		for i := length; i < cnt; i++ {
			isBot[i] = true
			*players = append(*players, socket.New(Bot()))
		}
	}
	return isBot
}
