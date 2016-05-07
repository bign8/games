package chess

import "testing"

func TestParse(t *testing.T) {
	board := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	state, err := ParseFEN(board)
	if err != nil {
		t.Errorf("Failed with: %s", err)
	}

	var golden [64]byte
	copy(golden[:], "rnbqkbnrpppppppp11111111111111111111111111111111PPPPPPPPRNBQKBNR")
	if state.board != golden {
		t.Errorf("Failed with board missmatch\n%q\n%q", golden, state.board)
	}

	out := state.FEN()
	if out != board {
		t.Errorf("Board and State don't match\n%q\n%q", board, out)
	}

}

func TestFENFullCycle(t *testing.T) {
	start := "rnbqkbnr/1ppp1Qpp/8/p3p3/2B1P3/8/PPPP1PPP/RNB1K1NR b KQkq - 0 1"
	board, err := ParseFEN(start)
	if err != nil {
		t.Fatalf("Failed with: %s", err)
	}
	end := board.FEN()
	if end != start {
		t.Fatalf("Boards don't match: %q != %q", start, end)
	}
}

func BenchmarkParseFEN(b *testing.B) {
	board := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ParseFEN(board)
	}
}

func BenchmarkStringFEN(b *testing.B) {
	state, err := ParseFEN("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		b.Errorf("Cannot parse FEN: %s", err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		state.FEN()
	}
}
