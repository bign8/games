package main

import (
	"bufio"
	"io"
	"log"
	"sync"
)

type manager struct {
	source io.ReadWriteCloser
	rooms  map[byte]io.Writer
	mu     sync.RWMutex
	outMU  sync.Mutex
}

func createManager(source io.ReadWriteCloser) *manager {
	man := &manager{
		source: source,
		rooms:  make(map[byte]io.Writer),
	}
	go func() {
		scanner := bufio.NewScanner(source) // Reader
		for scanner.Scan() {
			msg := scanner.Bytes()
			man.mu.RLock()
			obj, ok := man.rooms[msg[0]]
			man.mu.RUnlock()
			if ok {
				obj.Write(msg[1:]) // TODO: error check here
			} else {
				log.Printf("Invalid room message (%x): %s", msg[0], msg[1:])
			}
		}
	}()
	return man
}

func (m *manager) Room(name byte) io.ReadWriteCloser {
	// TODO: verify no duplicate room names
	read, write := io.Pipe()
	r := &room{
		read: read,
		name: name,
		man:  m,
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.rooms[name] = write
	return r
}

type room struct {
	read io.ReadCloser
	name byte
	man  *manager
}

func (r *room) Read(b []byte) (int, error) {
	return r.read.Read(b)
}

func (r *room) Write(b []byte) (int, error) {
	r.man.outMU.Lock()
	defer r.man.outMU.Unlock()
	bits := make([]byte, len(b)+1)
	bits[0] = r.name
	copy(bits[1:], b)
	num, err := r.man.source.Write(bits)
	if err != nil {
		return num, err
	}
	return num - 1, nil
}

func (r *room) Close() error {
	r.man.mu.Lock()
	defer r.man.mu.Unlock()
	delete(r.man.rooms, r.name)
	return r.read.Close()
}
