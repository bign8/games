package ttt

import (
	"testing"

	"github.com/bign8/games"
	"github.com/bign8/games/util/assert"
)

var (
	_  games.Actor = (*badActor)(nil)
	bs string
)

type badActor string

func (ba badActor) Name() string                   { return string(ba) }
func (ba badActor) Act(s games.State) games.Action { return s.Actions()[0] }

func Test(t *testing.T) {
	p1 := badActor("asdf")
	p2 := badActor("qwer")
	game := New(&p1, &p2)
	assert.String(t, game.String(), "╔═══╦═══╦═══╗\n║   ║   ║   ║\n╠═══╬═══╬═══╣\n"+
		"║   ║   ║   ║\n╠═══╬═══╬═══╣\n║   ║   ║   ║\n╚═══╩═══╩═══╝", "default string")
	game.Apply(game.Actions()[0])
}

func TestIsWin(t *testing.T) {
	check := func(board string, done bool, winner byte, msg string) {
		if len(board) != 9 {
			t.Errorf("Invalid board: %q %q", board, msg)
		}
		var b [9]byte
		copy(b[:], board)
		win, er := isWin(b)
		assert.Bool(t, done, win, msg)
		assert.Byte(t, winner, er, msg)
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
