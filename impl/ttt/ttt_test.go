package ttt

import (
	"testing"

	"github.com/bign8/games"
)

func BenchmarkStateString(b *testing.B) {
	game := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.String()
	}
}

func BenchmarkNewState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New()
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

func BenchmarkTerminal(b *testing.B) {
	game := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Terminal()
	}
}

func BenchmarkUtility(b *testing.B) {
	game := New()
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
	p1 := games.PlayerConfig{Name: "asdf", Type: games.MinPlayer}
	p2 := games.PlayerConfig{Name: "qwer", Type: games.MaxPlayer}
	return New(
		games.NewPlayer(&dumb{}, p1),
		games.NewPlayer(&dumb{}, p2),
	)
}

type dumb struct{}

func (p dumb) Act(s games.State) games.Action {
	return s.Actions()[0]
}
