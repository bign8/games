package ttt

import (
	"testing"

	"github.com/bign8/games"
	"github.com/bign8/games/player/cli"
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
	p := cli.New("asdf", games.MinPlayer)
	game := New(p, p)
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
	p := cli.New("asdf", games.MinPlayer)
	game := New(p, p)
	game = game.Apply(game.Actions()[4])
	game = game.Apply(game.Actions()[4])
	game = game.Apply(game.Actions()[3])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.SVG(false)
	}
}
