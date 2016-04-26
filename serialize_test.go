package chess

import "testing"

func TestParse(t *testing.T) {
	board := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	state, err := Parse(board)
	if err != nil {
		t.Errorf("Failed with: %s", err)
	}

	var golden [64]byte
	copy(golden[:], "rnbqkbnrpppppppp11111111111111111111111111111111PPPPPPPPRNBQKBNR")
	if state.board != golden {
		t.Errorf("Failed with board missmatch\n%q\n%q", golden, state.board)
	}
}

func BenchmarkParse(b *testing.B) {
	board := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(board)
	}
}
