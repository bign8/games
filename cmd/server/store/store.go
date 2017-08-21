package store

import (
	"crypto/rand"
	"errors"

	"github.com/bign8/games"
)

// Store holds active games
type Store interface {
	NewPlayer(game games.Game) (playerID string, err error)

	// Pairs a player with opponents. Iff force is set, assert AI players
	Pair(game games.Game, playerID string, force bool) (gameID string, err error)
}

// NewMemoryStore creates an in-memory store
func NewMemoryStore() Store {
	return &memoryStore{
		playerz: make(map[string]bool),
	}
}

type memoryStore struct {
	playerz  map[string]bool // playerID
	gamePool map[string][]string
}

func (mem *memoryStore) NewPlayer(game games.Game) (playerID string, err error) {
	pid, err := uid()
	if err != nil {
		return "", err
	}
	mem.playerz[pid] = true
	return pid, nil
}

const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func uid() (string, error) {
	bytes := make([]byte, 40)
	_, err := rand.Read(bytes) // TODO: fallback to non-crypto
	if err == nil {
		for i, b := range bytes {
			bytes[i] = letters[b%byte(len(letters))]
		}
	}
	return string(bytes), err
}

func (mem *memoryStore) Pair(game games.Game, playerID string, force bool) (gameID string, err error) {
	return "", errors.New("TODO")
}
