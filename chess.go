package chess

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

	moves  []*Move        // cache of available moves
	pieces map[uint8]byte // cache of occupied locations
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

// IsBlack returns if current player is black
func (s State) IsBlack() bool {
	return s.isBlack
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
// TODO: benchmark this a ton
func (s State) Apply(m *Move) (*State, error) {
	var found bool
	for _, move := range s.Moves() {
		if move.Equals(m) {
			found = true
			break
		}
	}
	if !found {
		return nil, errors.New("chess: move not permitted")
	}

	// migrating to mutable state
	board := s.board
	board = strings.Replace(board, "8", "11111111", -1)
	board = strings.Replace(board, "7", "1111111", -1)
	board = strings.Replace(board, "6", "111111", -1)
	board = strings.Replace(board, "5", "11111", -1)
	board = strings.Replace(board, "4", "1111", -1)
	board = strings.Replace(board, "3", "111", -1)
	board = strings.Replace(board, "2", "11", -1)

	// Make Move
	board = board[:m.Stop] + board[m.Start:m.Start+1] + board[m.Stop+1:]
	board = board[:m.Start] + "1" + board[m.Start+1:]

	// Consolidate 1's back down to 2-8 if possible
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

	fmt.Printf("New Board\n%q\n%q\n", board, newBoard)

	state := &State{
		board:     newBoard,
		isBlack:   !s.isBlack,
		castling:  s.castling,
		enPassant: s.enPassant,
		halfmove:  s.halfmove,
		count:     s.count + 1,
	}
	return state, nil
}
