package chess

import "testing"

func BenchmarkMoves(b *testing.B) {
	game := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.moves = nil
		game.Moves()
	}
}
