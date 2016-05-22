// Package games borrows terminology from "AI - A Modern Approach" Chapter 5
package games

import "fmt"

// SVGChooseMove is the JS function that needs to be called within a SVG for a move to be chosen
const SVGChooseMove = `parent.N8.games.chooseMove`

// Starter is a function used to create a game's initial state
type Starter func(...Player) State

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
	SVG(bool) string    // Browser representation of a state (editable)
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
	Board   string
	Start   Starter        `json:"-"`
	Players []PlayerConfig `json:"-"`
	AI      Actor          `json:"-"`
}
