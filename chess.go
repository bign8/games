package chess

import "errors"

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
	moves     []Move
}

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

// Bytes generats a minimal transmission of this data in []byte form
func (s State) Bytes() []byte {
	return []byte("TODO")
}

// Parse parses a state from []byte generated via Bytes()
func Parse(bits []byte) (*State, error) {
	// TODO: parse state from bytes
	return &State{}, errors.New("TODO: not implemented")
}

// Apply executes a move on a given state of the board
func (s State) Apply(m Move) (*State, error) {
	return nil, errors.New("TODO: Not implemented")
}
