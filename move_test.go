package chess

import "testing"

func BenchmarkMoves(b *testing.B) {
	game := New()
	for i := 0; i < b.N; i++ {
		game.Moves()
	}
}
