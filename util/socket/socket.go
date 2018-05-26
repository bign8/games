package socket

import (
	"bufio"
	"io"
	"log"
	"sync"
)

var _ io.ReadWriteCloser = (*room)(nil)

func New(socket io.ReadWriteCloser) *Socket {
	s := &Socket{
		Done:  make(chan struct{}),
		core:  socket,
		rooms: make(map[byte]io.Writer, 3),
	}
	go s.recv()
	return s
}

type Socket struct {
	Done chan struct{}

	// Communication
	core    io.ReadWriteCloser
	writeMu sync.Mutex

	// Rooms
	rooms   map[byte]io.Writer
	roomsMu sync.RWMutex
}

func (s *Socket) recv() {
	scanner := bufio.NewScanner(s.core)
	for scanner.Scan() {
		msg := scanner.Bytes()
		if len(msg) == 0 {
			log.Println("Zero length message from somebody")
			continue
		}
		s.roomsMu.RLock()
		obj, ok := s.rooms[msg[0]]
		s.roomsMu.RUnlock()
		if ok {
			obj.Write(append(msg[1:], '\n')) // TODO: error check
		} else {
			log.Printf("Invalid room message (%x): %s", msg[0], msg[1:])
		}
	}
	close(s.Done)
}

// Close down a given socket
func (s *Socket) Close() error { return s.core.Close() }

// Room starts up a new single character prefixed room
func (s *Socket) Room(name byte) io.ReadWriteCloser {
	// TODO: detect duplicate room names
	read, write := io.Pipe()
	s.roomsMu.Lock()
	s.rooms[name] = write
	s.roomsMu.Unlock()
	return &room{
		read: read,
		name: name,
		conn: s,
	}
}

// Methods Below are called from child Room objects

func (s *Socket) roomWrite(room byte, msg []byte) (int, error) {
	s.writeMu.Lock()
	defer s.writeMu.Unlock()
	return s.core.Write(append([]byte{room}, msg...))
}

func (s *Socket) roomClose(room byte) error {
	s.roomsMu.Lock()
	defer s.roomsMu.Unlock()
	delete(s.rooms, room)
	return nil
}

// Room object

type room struct {
	read io.Reader
	name byte
	conn *Socket
}

func (r *room) Read(b []byte) (int, error) { return r.read.Read(b) }
func (r *room) Close() error               { return r.conn.roomClose(r.name) }
func (r *room) Write(b []byte) (int, error) {
	i, err := r.conn.roomWrite(r.name, b)
	return i - 1, err // we are clipping the room character
}
