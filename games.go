// Package games borrows terminology from "AI - A Modern Approach" Chapter 5
package games

import (
	"fmt"
	"sync"
)

// Starter is a function used to create a game's initial state
type Starter func(...Player) State

// Gamer is a function used to create a game's player
type Gamer func(string, PlayerType) Player

// PlayerType allows for direct player comparison during tree search
type PlayerType int

const (
	// UnknownPlayer is the default PlayerType
	UnknownPlayer PlayerType = iota
	// MaxPlayer is a player attempting to maximize the utility function
	MaxPlayer
	// MinPlayer is a player attempting to minimize the unility function
	MinPlayer
)

// Player is the active player of a game
type Player interface {
	fmt.Stringer
	Play(State) Action
	Human() bool
	Type() PlayerType
}

// Action is the base type for a game move
type Action interface {
	fmt.Stringer
}

// State is the base type for the state of a game
type State interface {
	fmt.Stringer
	Player() Player     // Active player given a State
	Apply(Action) State // Applying an action to a game
	Actions() []Action  // List of available actions in a State
	Terminal() bool     // If the game is in a terminal state
	Utility() int       // Each game can define their own utility
	Error() error       // If any problem exists in regular game-play
}

// AI generates a player that maximises the result of the Utilites provided
func AI(ranks ...Utility) Player {
	return nil
}

// Run is the primary game runner
func Run(game State) State {
	for !game.Terminal() {
		game = game.Apply(game.Player().Play(game))
	}
	return game
}

// Game is contains all the meta-data surrounding a game so it can be played
type Game struct {
	Name    string
	Slug    string
	Start   Starter        `json:"-"`
	Players []PlayerConfig `json:"-"`
	AI      Gamer          `json:"-"`
}

// PlayerConfig is the starting configuration for a player
type PlayerConfig struct {
	Name string
	Type PlayerType
}

var mu sync.RWMutex
var registry = make(map[string]Game)

// Register allows implementations to register their specific version of a game
func Register(game Game) error {
	mu.Lock()
	defer mu.Unlock()
	if _, ok := registry[game.Slug]; ok {
		return fmt.Errorf("games: Game already registered with slug: %s", game.Slug)
	}
	registry[game.Slug] = game
	return nil
}

// List returns all the currently registerd games in the system
func List() []Game {
	mu.RLock()
	defer mu.RUnlock()
	res := make([]Game, len(registry))
	i := 0
	for _, game := range registry {
		res[i] = game
		i++
	}
	return res
}

// Get returns the full registered game by slug
func Get(slug string) *Game {
	mu.RLock()
	defer mu.RUnlock()
	g, ok := registry[slug]
	if !ok {
		return nil
	}
	return &g
}
