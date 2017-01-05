package games

import "errors"

// Varous error states a game can commonly enter
var (
	StateErrInvalidNumberOfActors = Error("Invalid Number of Actors")
	StateErrInvalidMove           = Error("Invalid Move")
)

// Error creates an erronious state that only returns errors
func Error(s string) State { return err(s) }

type err string

func (e err) String() string     { return "Error: " + string(e) }
func (e err) Player() Actor      { return nil }
func (e err) Apply(Action) State { return e }
func (e err) Actions() []Action  { return nil }
func (e err) Terminal() bool     { return true }
func (e err) Utility(Actor) int  { return 0 }
func (e err) Error() error       { return errors.New(string(e)) }
func (e err) SVG(bool) string    { return e.String() }
