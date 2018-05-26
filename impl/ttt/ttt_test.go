package ttt

import (
	"testing"

	"github.com/bign8/games/util/assert"
)

func Test(t *testing.T) {
	game := New(2)
	assert.Equal(t, game.String(), "╔═══╦═══╦═══╗\n║   ║   ║   ║\n╠═══╬═══╬═══╣\n"+
		"║   ║   ║   ║\n╠═══╬═══╬═══╣\n║   ║   ║   ║\n╚═══╩═══╩═══╝", "default string")
	assert.Equal(t, game.SVG(false), `<svg viewBox="0 0 90 90" stroke="black" stroke-linecap="round"><g></g></svg>`, "svg")
	game.Utility()

	// First move
	game = game.Apply(game.Actions()[0])
	game.SVG(true)
	game.Utility()

	// Second move
	game = game.Apply(game.Actions()[0])
	game.SVG(true)
	game.Utility()
}

func TestIsWin(t *testing.T) {
	check := func(board string, done bool, winner byte, msg string) {
		if len(board) != 9 {
			t.Errorf("Invalid board: %q %q", board, msg)
		}
		var b [9]byte
		copy(b[:], board)
		win, er := isWin(b)
		assert.Equal(t, done, win, msg)
		assert.Equal(t, winner, er, msg)
	}
	check("         ", false, ' ', "empty")
	check("xxx      ", true, 'x', "top horz")
	check("x  x  x  ", true, 'x', "tl to br")
	check("x   x   x", true, 'x', "ltr diag")
	check("   xxx   ", true, 'x', "mid horz")
	check(" x  x  x ", true, 'x', "mid vert")
	check("  x x x  ", true, 'x', "bl to tr")
	check("      xxx", true, 'x', "bot horz")
	check("  x  x  x", true, 'x', "rit vert")
}

func BenchmarkStateString(b *testing.B) {
	game := New(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = game.String()
	}
}

func BenchmarkNewState(b *testing.B) {
	for i := 0; i < b.N; i++ {
		New(2)
	}
}

func BenchmarkApply(b *testing.B) {
	game := New(2)
	move := game.Actions()[0]
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Apply(move)
	}
}

func BenchmarkUtility(b *testing.B) {
	game := New(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.Utility()
	}
}

func BenchmarkSVG(b *testing.B) {
	game := New(2)
	game = game.Apply(game.Actions()[4])
	game = game.Apply(game.Actions()[4])
	game = game.Apply(game.Actions()[3])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.SVG(false)
	}
}
