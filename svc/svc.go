package svc

import (
	"context"

	"github.com/bign8/games/impl"
)

// GameSlug is the slug of a game (defined in impl.Map()[*].Slug)
type GameSlug string

// PlayerID is a globally unique identifier for an ongoing match.
// Even if two players are playing the same game, their palyer ID is unique.
type PlayerID string

// GameService manages interfacing with various games
type GameService interface {
	NewPlayer(context.Context, GameSlug) (PlayerID, error)
	// Pair(slug string) (<-chan GameID, error)
	// Load(GameID) (games.Game, error)
	// Save(GameID, games.Game) error
}

// Random returns a random game slug
func Random() string {
	for key := range impl.Map() {
		return key
	}
	return "ttt" // you can always play tick-tac-toe!
}
