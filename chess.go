package chess

import "errors"

const newGame = "rnbqkbnrpppppppp11111111111111111111111111111111PPPPPPPPRNBQKBNR"

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

	moves  []*Move        // cache of available moves
	pieces map[uint8]byte // cache of occupied locations
}

// New begins a brand new game
func New() *State {
	var board [64]byte
	copy(board[:], newGame)
	return &State{
		board:     board,
		isBlack:   false,
		castling:  15,
		enPassant: 0,
		halfmove:  0,
		count:     1,
	}
}

// IsBlack returns if current player is black
func (s State) IsBlack() bool {
	return s.isBlack
}

/*
// Out generats a minimal transmission of this data in []byte form
func (s State) Out() string {

	// TODO: make this actually work
	// Consolidate 1's back down to 2-8 if possible
	// TODO: move all this logic to PARSE
	newBoard := board[:]
	length := 0
	for i := 63; i >= 0; i-- {
		if newBoard[i] == '1' {
			length++
			if length > 0 && i%8 == 0 { // line wrap
				newBoard = newBoard[:i] + strconv.Itoa(length) + newBoard[i+length:]
				length = 0
			}
		} else if length > 0 { // end of existing chain of numbers
			newBoard = newBoard[:i+1] + strconv.Itoa(length) + newBoard[i+length+1:]
			length = 0
		}
	}

	return "TODO"
}

// Parse parses a state from []byte generated via Bytes()
func Parse(bits string) (*State, error) {

	// migrating to mutable state
	// TODO: move all this logic to PARSE
	board := "parse_board from full set of bits"
	board = strings.Replace(board, "8", "11111111", -1)
	board = strings.Replace(board, "7", "1111111", -1)
	board = strings.Replace(board, "6", "111111", -1)
	board = strings.Replace(board, "5", "11111", -1)
	board = strings.Replace(board, "4", "1111", -1)
	board = strings.Replace(board, "3", "111", -1)
	board = strings.Replace(board, "2", "11", -1)
	// TODO: parse state from bytes
	return &State{}, errors.New("TODO: not implemented")
}
//*/

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
