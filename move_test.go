package chess

import (
	"errors"
	"testing"
)

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

func movesEqual(golden, moves []*Move) error {
	found := make([]bool, len(golden))
	missed := ""
_OUT:
	for _, move := range moves {
		for j, gold := range golden {
			if !found[j] && gold.Equals(move) {
				found[j] = true
				continue _OUT
			}
		}
		missed += ", m:" + move.String()
	}
	for i, bit := range found {
		if !bit {
			missed += ", g:" + golden[i].String()
		}
	}
	if len(missed) != 0 {
		return errors.New(missed[1:])
	}
	return nil
}

func BenchmarkMovesRook(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.rookMoves(start)
	}
}

func BenchmarkMovesPawn(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.pawnMoves(start)
	}
}

func BenchmarkMovesKnight(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.knightMoves(start)
	}
}

func BenchmarkMovesBishop(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.bishopMoves(start)
	}
}

func BenchmarkMovesQueen(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.queenMoves(start)
	}
}

func BenchmarkMovesKing(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.kingMoves(start)
	}
}

func TestRookMoves(t *testing.T) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	moves := board.rookMoves(start)
	golden := []*Move{
		&Move{start, Location(33), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(34), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(35), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(36), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(37), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(38), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(39), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(40), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(48), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(56), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(24), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(16), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(8), InvalidLocation, 0, nil, false, false},
		&Move{start, Location(0), InvalidLocation, 0, nil, false, false},
	}
	if err := movesEqual(golden, moves); err != nil {
		t.Errorf("Not all rook moves found: %s", err)
	}
}
