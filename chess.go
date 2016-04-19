package chess

import "strings"

// State is an internal representation of a chess game.
// - FEN notation: https://en.wikipedia.org/wiki/FEN
// - See: http://golang-sizeof.tips/?t=Ly8gTm90ZXM6Ci8vIC0gc2l6ZXM6IGh0dHBzOi8vZ29sYW5nLm9yZy9wa2cvYnVpbHRpbi8KLy8gLSBGRU4gbm90YXRpb246IGh0dHBzOi8vZW4ud2lraXBlZGlhLm9yZy93aWtpL0ZvcnN5dGglRTIlODAlOTNFZHdhcmRzX05vdGF0aW9uCgpzdHJ1Y3QgewoJYm9hcmQgc3RyaW5nIC8vIGFyYml0cmFyeSBzaXplCglpc0JsYWNrIGJvb2wgLy8gZGV0ZXJtaW5lcyBhY3RpdmUgY29sb3IgKGNvdWxkIG1hcCBvbiBjYXN0bGluZykKCWNhc3RsaW5nIHVpbnQ4IC8vIGJpdCBtYXNrZWQgbnVtYmVyIGluIEtRa3Egb3JkZXIKCWVuUGFzc2FudCB1aW50OCAvLyBib2FyZCBwb3NpdGlvbiAwIChuL2EpICsgMSAtIDY0CgloYWxmbW92ZSB1aW50OCAvLyBtYXggNTAgKGxpbWl0ZWQgYnkgcnVsZSkgW3R5cGU6MjU1XQoJY291bnQgdWludDMyIC8vIG1heCBvZiA0Mjk0OTY3Mjk1IChsaW1pdGVkIGJ5IHR5cGUpCn0KCi8vIENhbiBpZ25vcmUgdGhlIGxhc3QgMiBmb3IgcG9zc2libGUgbW92ZXMgbWFw
// - sizes: https://golang.org/pkg/builtin/
// - Can ignore the last 2 for possible moves map
type State struct {
	board     string // arbitrary size
	isBlack   bool   // determines active color (could map on castling)
	castling  uint8  // bit masked number in KQkq order
	enPassant uint8  // board position 0 (n/a) + 1 - 64
	halfmove  uint8  // max 50 (limited by rule) [type:255]
	count     uint32 // max of 4294967295 (limited by type)
}

var chrLookup = map[uint8]string{
	'p': "♟", 'r': "♜", 'n': "♞", 'b': "♝", 'q': "♛", 'k': "♚",
	'P': "♙", 'R': "♖", 'N': "♘", 'B': "♗", 'Q': "♕", 'K': "♔",
}

var numLookup = map[uint8]int{
	'1': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7,
}

const col = " ║ "
const top = "╔═══╦═══╦═══╦═══╦═══╦═══╦═══╦═══╗\n"
const sep = "\n╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣\n"
const bot = "\n╚═══╩═══╩═══╩═══╩═══╩═══╩═══╩═══╝\n  A   B   C   D   E   F   G   H"

// New begins a brand new game
func New() *State {
	return &State{
		board:     "rnbqkbnrpppppppp8888PPPPPPPPRNBQKBNR", // '/'s removed
		isBlack:   false,
		castling:  15,
		enPassant: 0,
		halfmove:  0,
		count:     1,
	}
}

// String is to implement the fmt.Stringer interface
func (s State) String() string {
	bits := make([]string, 64)
	for i := 0; i < 64; i++ {
		bits[i] = " "
	}
	for i, scanner := 0, 0; i < 64; i++ {
		in := s.board[scanner]
		scanner++
		if chr, ok := chrLookup[in]; ok {
			bits[i] = chr
			continue
		}
		i += numLookup[in]
	}

	rows := []string{
		"║ " + strings.Join(bits[0:8], col) + " ║  8",
		"║ " + strings.Join(bits[8:16], col) + " ║  7",
		"║ " + strings.Join(bits[16:24], col) + " ║  6",
		"║ " + strings.Join(bits[24:32], col) + " ║  5",
		"║ " + strings.Join(bits[32:40], col) + " ║  4",
		"║ " + strings.Join(bits[40:48], col) + " ║  3",
		"║ " + strings.Join(bits[48:56], col) + " ║  2",
		"║ " + strings.Join(bits[56:64], col) + " ║  1",
	}
	return top + strings.Join(rows, sep) + bot
}

// Bytes generats a minimal transmission of this data in []byte form
func (s State) Bytes() []byte {
	return []byte("TODO")
}

// Parse parses a state from []byte generated via Bytes()
func Parse(bits []byte) *State {
	// TODO: parse state from bytes
	return &State{}
}
