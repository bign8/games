// Package games borrows terminology from "AI - A Modern Approach" Chapter 5
package games

import "fmt"

// Starter is a function used to create a game's initial state
type Starter func(...Player) State

// Gamer is a function used to create a game's player
type Gamer func(string) Player

// Player is the active player of a game
type Player interface {
	fmt.Stringer
	Play(State) Action
	Human() bool
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
