// Package engine borrows terminology from "AI - A Modern Approach" Chapter 5
package engine

import "fmt"

// Player is the active player of a game
type Player interface {
	fmt.Stringer
	Play(State) Action
}

// Action is the base type for a game move
type Action interface {
	fmt.Stringer
}

// State is the base type for the state of a game
type State interface {
	fmt.Stringer
	Player() Player              // Active player given a State
	Apply(Action) (State, error) // Applying an action to a game
	Actions() []Action           // List of available actions in a State
	Terminal() bool              // If the game is in a terminal state
}

// Utility gives a numeric value to a given state
type Utility func(State, Player) float64

// SearchAI generates a player that maximises the result of the Utilites provided
func SearchAI(ranks ...Utility) Player {
	return nil
}

// Run is the primary game runner
func Run(game State, players ...Player) (err error) {
	for ctr, all := 0, len(players); err == nil && !game.Terminal(); ctr++ {
		move := players[ctr%all].Play(game)
		game, err = game.Apply(move)
	}
	return err
}
