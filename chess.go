package chess

import "errors"

// State is an internal representation of a chess game.
// - FEN notation: https://en.wikipedia.org/wiki/FEN
// - See: http://golang-sizeof.tips/?t=Ly8gTm90ZXM6Ci8vIC0gc2l6ZXM6IGh0dHBzOi8vZ29sYW5nLm9yZy9wa2cvYnVpbHRpbi8KLy8gLSBGRU4gbm90YXRpb246IGh0dHBzOi8vZW4ud2lraXBlZGlhLm9yZy93aWtpL0ZvcnN5dGglRTIlODAlOTNFZHdhcmRzX05vdGF0aW9uCgpzdHJ1Y3QgewoJYm9hcmQgWzY0XWJ5dGUgLy8gYXJiaXRyYXJ5IHNpemUKCWlzQmxhY2sgYm9vbCAvLyBkZXRlcm1pbmVzIGFjdGl2ZSBjb2xvciAoY291bGQgbWFwIG9uIGNhc3RsaW5nKQoJY2FzdGxpbmcgdWludDggLy8gYml0IG1hc2tlZCBudW1iZXIgaW4gS1FrcSBvcmRlcgoJZW5QYXNzYW50IHVpbnQ4IC8vIGJvYXJkIHBvc2l0aW9uIDAgKG4vYSkgKyAxIC0gNjQKCWhhbGZtb3ZlIHVpbnQ4IC8vIG1heCA1MCAobGltaXRlZCBieSBydWxlKSBbdHlwZToyNTVdCgljb3VudCB1aW50MzIgLy8gbWF4IG9mIDQyOTQ5NjcyOTUgKGxpbWl0ZWQgYnkgdHlwZSkKfQoKLy8gQ2FuIGlnbm9yZSB0aGUgbGFzdCAyIGZvciBwb3NzaWJsZSBtb3ZlcyBtYXA=
// - sizes: https://golang.org/pkg/builtin/
// - Can ignore the last 2 for possible moves map
type State struct {
	board     [64]byte // fixed size
	isBlack   bool     // determines active color (could map on castling)
	castling  uint8    // bit masked number in KQkq order
	enPassant uint8    // board position 0 (n/a) + 1 - 64
	halfmove  uint8    // max 50 (limited by rule) [type:255]
	count     uint32   // max of 4294967295 (limited by type)
	moves     []*Move  // cache of available moves
}

// New begins a brand new game
func New() *State {
	var board [64]byte
	copy(board[:], "rnbqkbnrpppppppp11111111111111111111111111111111PPPPPPPPRNBQKBNR")
	return &State{
		board:     board,
		isBlack:   false,
		castling:  15,
		enPassant: 0,
		halfmove:  0,
		count:     1,
	}
}

// Apply executes a move on a given state of the board
func (s State) Apply(m *Move) (*State, error) {
	var found bool
	for _, move := range s.Moves() {
		found = move.Equals(m)
		if found {
			break
		}
	}
	if !found {
		return nil, errors.New("chess: move not permitted")
	}

	// Make Move
	var board [64]byte
	copy(board[:], s.board[:])
	board[m.Stop] = s.board[m.Start]
	board[m.Start] = '1'

	// Generate new board... TODO: fix count + halfmove + nePassant + castling
	state := &State{
		board:     board,
		isBlack:   !s.isBlack,
		castling:  s.castling,
		enPassant: s.enPassant,
		halfmove:  s.halfmove,
		count:     s.count + 1,
	}
	return state, nil
}
