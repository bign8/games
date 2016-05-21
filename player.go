package games

// PlayerType allows for direct player comparison during tree search
type PlayerType int

// Useful PlayerTypes to help AI's move through a search space
const (
	UnknownPlayer PlayerType = iota
	MaxPlayer
	MinPlayer
)

// PlayerConfig is the starting configuration for a player
type PlayerConfig struct {
	Name string
	Type PlayerType
}

// Actor is a method that choose an Action given a particular State
type Actor interface {
	Act(State) Action
}

// Player is a base object that holds data pertaining to a particular person
type Player struct {
	Name  string
	Type  PlayerType
	Human bool // TODO: actually set this somewhere
	actor Actor
}

// Play on a Player calls the actors Act method
func (p Player) Play(s State) Action {
	return p.actor.Act(s)
}

// NewPlayer constructs a Player out of an actor and a PlayerConfig
func NewPlayer(a Actor, c PlayerConfig) Player {
	return Player{c.Name, c.Type, false, a}
}
