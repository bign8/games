package chess

import (
	"errors"
	"testing"
)

func BenchmarkMoves(b *testing.B) {
	game := New().(*State)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.moves = nil
		game.Moves()
	}
}

func BenchmarkMovesClip(b *testing.B) {
	game := New().(*State)
	moves := game.Moves()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		game.clipCheckMoves(moves)
	}
}

// castling board
// "r111k11r111111111111111111111111111111111111111111111111R111K11R"

func movesEqual(golden, moves []Move) error {
	found := make([]bool, len(golden))
	missed := ""
_OUT:
	for _, move := range moves {
		for j, gold := range golden {
			if !found[j] && gold.Equals(&move) {
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
	buffer := make([]Move, 0, 14)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.rookMoves(start, buffer)
	}
}

func BenchmarkMovesPawn(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	buffer := make([]Move, 0, 14)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.pawnMoves(start, buffer)
	}
}

func BenchmarkMovesKnight(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	buffer := make([]Move, 0, 8)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.knightMoves(start, buffer)
	}
}

func BenchmarkMovesBishop(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	buffer := make([]Move, 0, 14)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.bishopMoves(start, buffer)
	}
}

func BenchmarkMovesQueen(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	buffer := make([]Move, 0, 28)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.queenMoves(start, buffer)
	}
}

func BenchmarkMovesKing(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	buffer := make([]Move, 0, 10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		board.kingMoves(start, buffer)
	}
}

func TestRookMoves(t *testing.T) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	moves := make([]Move, 0, 14)
	moves = board.rookMoves(start, moves)
	golden := []Move{
		Move{start, Location(33), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(34), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(35), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(36), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(37), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(38), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(39), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(40), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(48), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(56), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(24), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(16), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(8), InvalidLocation, 0, nil, false, ""},
		Move{start, Location(0), InvalidLocation, 0, nil, false, ""},
	}
	if err := movesEqual(golden, moves); err != nil {
		t.Errorf("Not all rook moves found: %s", err)
	}
}

func TestStateIsCheck(t *testing.T) {
	start := Location(0)

	// bishop
	if board, _ := ParseFEN("8/1B6/8/8/8/8/8/8 w KQkq - 0 1"); !board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for bishop check")
	}

	// queen (bishop)
	if board, _ := ParseFEN("8/1Q6/8/8/8/8/8/8 w KQkq - 0 1"); !board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for queen (bishop) check")
	}

	// rook
	if board, _ := ParseFEN("8/R7/8/8/8/8/8/8 w KQkq - 0 1"); !board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for rook check")
	}

	// queen (rook)
	if board, _ := ParseFEN("8/Q7/8/8/8/8/8/8 w KQkq - 0 1"); !board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for queen (rook) check")
	}

	// knight
	if board, _ := ParseFEN("8/8/1N6/8/8/8/8/8 w KQkq - 0 1"); !board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for knight check")
	}

	// pawn
	if board, _ := ParseFEN("8/1P6/8/8/8/8/8/8 w KQkq - 0 1"); !board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for pawn check")
	}

	// pawn 2
	if board, _ := ParseFEN("8/8/8/8/8/8/6p1/8 b KQkq - 0 1"); !board.isCheck(Location(63), false) {
		t.Error("isCheck: returned wrong value for pawn 2 check")
	}

	// none
	if board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1"); board.isCheck(start, true) {
		t.Error("isCheck: returned wrong value for no check")
	}
}

func BenchmarkStateIsCheck(b *testing.B) {
	board, _ := ParseFEN("8/8/8/8/8/8/8/8 w KQkq - 0 1")
	start := Location(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if board.isCheck(start, false) {
			b.Error("Should never be in check here")
		}
	}
}

func TestBlackKingCheckRegression(t *testing.T) {
	board, _ := ParseFEN("R4k2/6p1/4N3/8/8/8/8/8 b kq - 0 1")
	// t.Log("Regression Board State:\n" + board.String())
	moves := board.Moves()
	if len(moves) != 2 {
		t.Errorf("Invalid number of moves for given state: %d != 2", len(moves))
	}
	if !t.Failed() {
		if moves[0].Type() != "King" || moves[1].Type() != "King" {
			t.Errorf("Invalid type for moves: %s != %s != King", moves[0].Type(), moves[1].Type())
		}
	}
}
