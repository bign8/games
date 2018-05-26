// Package games borrows terminology from "AI - A Modern Approach" Chapter 5
package games

import "errors"

// Starter is a function used to create a game's initial state
type Starter func(actors uint8) State

// Action is the base type for a game move
type Action interface {
	String() string // CLI representation of the action
	Type() string   // allows types of moves to be grouped
	Slug() string   // computer parsable ID of an action
}

// Actor is a method that choose an Action given a particular State
type Actor func(State) Action

// State is the base type for the state of a game
type State interface {
	String() string     // CLI represetation of the state
	Player() int        // index of the active player given a State (also index in Utility array)
	Apply(Action) State // Applying an action to a game
	Actions() []Action  // List of available actions in a State
	Utility() []int     // If the game is in a terminal state return the utility for each Actor, else nil
	SVG(bool) string    // Browser representation of a state (bool: editable)
	Terminal() bool     // Is the game complete

	// // For persistance
	// MarshalBinary() (data []byte, err error)
	// UnmarshalBinary(data []byte) error
}

// Game is contains all the meta-data surrounding a game so it can be played
type Game struct {
	Name   string  // Name of the game
	Slug   string  // Short name of game
	Board  string  // SVG of board state
	Counts []uint8 // Possible number of players to play a game
	Start  Starter `json:"-"` // Function to build the start state of the game // TODO: just have a state here
	AI     Actor   `json:"-"` // TODO: use smart enough ai that this can be removed
}

// Valid determines if a game configuration is valid.
func (g Game) Valid() error {
	if len(g.Counts) == 0 {
		return errors.New(g.Name + ": no Counts")
	}
	if g.AI == nil {
		return errors.New(g.Name + ": no AI")
	}
	if g.Start == nil {
		return errors.New(g.Name + ": no Start")
	}
	return nil
}

// Match is a game that is in progress
type Match struct {
	State          // active state of the game
	Actors []Actor // players engaged in the game
}

// Advance asks the current player to perform an action
func (m *Match) Advance() *Match {
	idx := m.Player()
	actor := m.Actors[idx]
	action := actor(m)
	m.State = m.Apply(action)
	return m
}

// Play constructs a match for a game given the set of actors.
// Build uses AI players to buffer insufficient numbers of players.
func (g Game) Play(actors ...Actor) *Match {
	length := uint8(len(actors))
	if err := g.Valid(); err != nil {
		panic(err)
	}

	// Find the first number of players larger than the current queue size
	count := g.Counts[0]
	for _, c := range g.Counts[1:] {
		count = c
		if c >= length {
			break
		}
	}

	// Make players or AIs based on the chosen count
	players := make([]Actor, count)
	for i := uint8(0); i < count; i++ {
		if i < length {
			players[i] = actors[i]
		} else {
			players[i] = g.AI
		}
	}
	return &Match{
		State:  g.Start(count),
		Actors: players,
	}
}

// MakeActors builds a series of actors given a method that can create them
func (g Game) MakeActors(builder func() Actor) []Actor {
	max := int(g.Counts[len(g.Counts)-1])
	actors := make([]Actor, max)
	for i := 0; i < max; i++ {
		actors[i] = builder()
	}
	return actors
}
