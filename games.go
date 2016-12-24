// Package games borrows terminology from "AI - A Modern Approach" Chapter 5
package games

import "fmt"

// SVGChooseMove is the JS function that needs to be called within a SVG for a move to be chosen
const SVGChooseMove = `parent.N8.games.chooseMove`

// Starter is a function used to create a game's initial state
type Starter func(...Actor) State

// ActorBuilder is a builder of actors
type ActorBuilder func(g Game, name string) Actor

// Action is the base type for a game move
type Action fmt.Stringer

// Actor is a method that choose an Action given a particular State
type Actor interface {
	Name() string
	Act(State) Action
}

// State is the base type for the state of a game
type State interface {
	fmt.Stringer
	Player() Actor      // Active player given a State
	Apply(Action) State // Applying an action to a game
	Actions() []Action  // List of available actions in a State
	Terminal() bool     // If the game is in a terminal state
	Utility(Actor) int  // Each game defines their own utility
	Error() error       // If any problem exists in regular game-play
	SVG(bool) string    // Browser representation of a state (editable)
}

// Game is contains all the meta-data surrounding a game so it can be played
type Game struct {
	Name    string   // Name of the game
	Board   string   // SVG of board state
	Players []string // List of Player names
	Start   Starter  `json:"-"`
	AI      Actor    `json:"-"`
}

// Run is the primary game runner
func Run(g Game, ab ActorBuilder) (final State) {
	actors := make([]Actor, len(g.Players))
	for i, name := range g.Players {
		actors[i] = ab(g, name)
	}
	game := g.Start(actors...)
	for !game.Terminal() {
		game = game.Apply(game.Player().Act(game))
	}
	return game
}
