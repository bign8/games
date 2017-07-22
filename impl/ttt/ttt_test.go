package ttt

import (
	"testing"

	"github.com/bign8/games"
)

var (
	_  games.Actor = (*badActor)(nil)
	bs string
)

type badActor string

func (ba badActor) Name() string                   { return string(ba) }
func (ba badActor) Act(s games.State) games.Action { return s.Actions()[0] }

func BenchmarkStateString(b *testing.B) {
	p1 := badActor("asdf")
	p2 := badActor("qwer")
	game := New(&p1, &p2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bs = game.String()
	}

}

func BenchmarkNewState(b *testing.B) {
	p1 := badActor("asdf")
	p2 := badActor("qwer")
	for i := 0; i < b.N; i++ {
		New(&p1, &p2)
	}
}

func BenchmarkApply(b *testing.B) {
	game := newGame()
	move := game.Actions()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Apply(move)
	}
}

func BenchmarkUtility(b *testing.B) {
	game := newGame()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Utility()
	}
}

func BenchmarkSVG(b *testing.B) {
	game := newGame()
	game = game.Apply(game.Actions()[4])
	game = game.Apply(game.Actions()[4])
	game = game.Apply(game.Actions()[3])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.SVG(false)
	}
}

func newGame() games.State {
	p1 := badActor("asdf")
	p2 := badActor("qwer")
	return New(&p1, &p2)
}

type dumb struct{}

func (p dumb) Act(s games.State) games.Action {
	return s.Actions()[0]
}
