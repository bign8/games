package chess

import (
	"github.com/bign8/games"
	"github.com/bign8/games/player"
)

var _ games.State = (*State)(nil)

// State is an internal representation of a chess game.
// - FEN notation: https://en.wikipedia.org/wiki/FEN
// - See: http://golang-sizeof.tips/?t=Ly8gTm90ZXM6Ci8vIC0gc2l6ZXM6IGh0dHBzOi8vZ29sYW5nLm9yZy9wa2cvYnVpbHRpbi8KLy8gLSBGRU4gbm90YXRpb246IGh0dHBzOi8vZW4ud2lraXBlZGlhLm9yZy93aWtpL0ZvcnN5dGglRTIlODAlOTNFZHdhcmRzX05vdGF0aW9uCgpzdHJ1Y3QgewoJYm9hcmQgWzY0XWJ5dGUgLy8gYXJiaXRyYXJ5IHNpemUKCWlzQmxhY2sgYm9vbCAvLyBkZXRlcm1pbmVzIGFjdGl2ZSBjb2xvciAoY291bGQgbWFwIG9uIGNhc3RsaW5nKQoJY2FzdGxpbmcgdWludDggLy8gYml0IG1hc2tlZCBudW1iZXIgaW4gS1FrcSBvcmRlcgoJZW5QYXNzYW50IHVpbnQ4IC8vIGJvYXJkIHBvc2l0aW9uIDAgKG4vYSkgKyAxIC0gNjQKCWhhbGZtb3ZlIHVpbnQ4IC8vIG1heCA1MCAobGltaXRlZCBieSBydWxlKSBbdHlwZToyNTVdCgljb3VudCB1aW50MzIgLy8gbWF4IG9mIDQyOTQ5NjcyOTUgKGxpbWl0ZWQgYnkgdHlwZSkKfQoKLy8gQ2FuIGlnbm9yZSB0aGUgbGFzdCAyIGZvciBwb3NzaWJsZSBtb3ZlcyBtYXA=
// - sizes: https://golang.org/pkg/builtin/
// - Can ignore the last 2 for possible moves map
type State struct {
	board     [64]byte // fixed size
	isBlack   bool     // determines active color (could map on castling)
	castling  uint8    // bit masked number in KQkq order
	enPassant Location // location behind last double pawn move
	halfmove  uint8    // max 50 (limited by rule) [type:255]
	count     uint32   // max of 4294967295 (limited by type)
	moves     []*Move  // cache of available moves
	check     bool
}

// New begins a brand new game
func New(actors uint8) games.State {
	var board [64]byte
	copy(board[:], "rnbqkbnrpppppppp11111111111111111111111111111111PPPPPPPPRNBQKBNR")
	// copy(board[:], "rnbqkbnr1ppp1ppp11111111p111p11Q11B1P11111111111PPPP1PPPRNB1K1NR")
	return &State{
		board:     board,
		isBlack:   false,
		castling:  15,
		enPassant: InvalidLocation,
		halfmove:  0,
		count:     1,
	}
}

// Player return the index of the active player
func (s State) Player() int {
	if s.isBlack {
		return 1
	}
	return 0
}

// Apply executes a move on a given state of the board
func (s State) Apply(mo games.Action) games.State {
	m := mo.(*Move)
	var found bool
	for _, move := range s.Moves() {
		found = move.Equals(m)
		if found {
			break
		}
	}
	if !found {
		panic("chess: move not permitted")
	}

	// Should reset halfmove count https://en.wikipedia.org/wiki/Fifty-move_rule
	halfmove := s.halfmove + 1
	if s.piece(m.Stop) || s.board[m.Start] == 'p' || s.board[m.Start] == 'P' {
		halfmove = 0
	}

	// enPassant - remove pawn
	isPawn := s.board[m.Start] == 'p' || s.board[m.Start] == 'P'
	if isPawn && m.Stop == s.enPassant && s.board[m.Stop] == '1' {
		_, col := m.Stop.rowCol()
		row, _ := m.Start.rowCol()
		s.board[locFromRowCol(row, col)] = '1'
	}

	// castling - track if king or rook moves - KQkq - TODO: clean up this ligic
	var kq uint8
	var home int8
	if s.isBlack {
		kq = (s.castling >> 2) & 3
		home = 0
	} else {
		kq = (s.castling >> 0) & 3
		home = 7
	}
	if kq > 0 {
		if s.board[m.Start] == 'K' || s.board[m.Start] == 'k' {
			kq = 0 // If king is moved, castling is no longer possible
		} else if row, col := m.Start.rowCol(); (s.board[m.Start] == 'r' || s.board[m.Start] == 'R') && row == home {
			if col == 0 {
				kq = kq & 2
			} else {
				kq = kq & 1
			}
		}
	}
	var castling uint8
	if s.isBlack {
		castling = kq<<2 | (s.castling & 3)
	} else {
		castling = kq<<0 | (s.castling & 12)
	}

	// Make Move
	var board [64]byte
	copy(board[:], s.board[:])
	board[m.Stop] = s.board[m.Start]
	board[m.Start] = '1'
	if m.castling != nil {
		board[m.castling.Stop] = s.board[m.castling.Start]
		board[m.castling.Start] = '1'
	}

	// Promotion
	if m.promotion > 0 {
		list := "_RNBQ"
		if s.isBlack {
			list = "_rnbq"
		}
		board[m.Stop] = list[m.promotion]
	}

	// fullmove count is only incremented after black's move
	count := s.count
	if s.isBlack {
		count++
	}

	// Generate new board... TODO: fix castling
	return &State{
		board:     board,
		isBlack:   !s.isBlack,
		castling:  castling,
		enPassant: m.passing,
		halfmove:  halfmove,
		count:     count,
		check:     m.check,
	}
}

// Terminal determines if the active game state is a complete move
func (s State) Terminal() bool {
	// https://en.wikipedia.org/wiki/Chess#End_of_the_game
	// TODO: checkmate: https://en.wikipedia.org/wiki/Checkmate
	// TODO: stalemate: https://en.wikipedia.org/wiki/Stalemate
	// TODO: halfmove: https://en.wikipedia.org/wiki/Fifty-move_rule
	// TODO: more win/draw cases ...
	return len(s.Moves()) == 0
}

// Game is the fully configured chess game
var Game = games.Game{
	Name:   "Chess",
	Slug:   "chess",
	Board:  `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1"><rect x="-.1" y="-.1" width="8.2" height="8.2" fill="#999"/><path fill="#FFF" d="M0,0H8v1H0zm0,2H8v1H0zm0,2H8v1H0zm0,2H8v1H0zM1,0V8h1V0zm2,0V8h1V0zm2,0V8h1V0zm2,0V8h1V0z"/></svg>`,
	Start:  New,
	AI:     player.Layer,
	Counts: []uint8{2},
}
