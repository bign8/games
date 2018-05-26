package svc

import "github.com/bign8/games"

// GameID is a globally unique identifier for an ongoing match.
type GameID string

// GameService manages interfacing with various games
type GameService interface {
	List() ([]games.Game, error)
	Pair(slug string) (<-chan GameID, error)
	Load(GameID) (games.Game, error)
	Save(GameID, games.Game) error
}

// Random returns a random game
func Random(svc GameService) games.Game {
	return games.Game{ /* todo */ }
}
