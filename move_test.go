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

// castling board
// "r111k11r111111111111111111111111111111111111111111111111R111K11R"
