package chess

import (
	"fmt"
	"strconv"
	"strings"
)

//*

// FEN generats a minimal transmission of this data in []byte form
func (s State) FEN() string {

	// TODO: make this actually work
	// Consolidate 1's back down to 2-8 if possible
	// TODO: move all this logic to PARSE
	newBoard := ""
	parts := make([]string, 6)
	pieces := make([]string, 64+7)
	index := 0

	// Setup board text
	length := 0
	for i := 0; i < 64; i++ {
		if i%8 == 0 && i != 0 {
			pieces[index] = "/"
			index++
		}
		if s.board[i] == '1' {
			length++
			if length > 0 && i%8 == 7 { // line wrap
				pieces[index] = strconv.Itoa(length)
				length = 0
			}
		} else if length > 0 { // end of existing chain of numbers
			pieces[index] = strconv.Itoa(length)
			length = 0
		} else {
			pieces[index] = string(s.board[i])
		}
		index++
	}
	parts[0] = strings.Join(pieces, "")

	// Set active player
	if s.isBlack {
		parts[1] = "b"
	} else {
		parts[1] = "w"
	}

	// Set castling
	if s.castling == 0 {
		parts[2] = "-"
		newBoard += "-"
	} else {
		if s.castling&8 == 8 {
			parts[2] += "K"
		}
		if s.castling&4 == 4 {
			parts[2] += "Q"
		}
		if s.castling&2 == 2 {
			parts[2] += "k"
		}
		if s.castling&1 == 1 {
			parts[2] += "q"
		}
	}

	// Set enPassant
	parts[3] = s.enPassant.String()

	// Set Halfmove
	parts[4] = strconv.FormatUint(uint64(s.halfmove), 10)
	parts[5] = strconv.FormatUint(uint64(s.count), 10)
	return strings.Join(parts, " ")
}

// ParseFEN parses a state from []byte generated via Bytes()
func ParseFEN(bits string) (*State, error) {
	parts := strings.Split(bits, " ") // grid castle enPassant halfmove count

	// Performant replace
	var grid [64]byte
	ctr, board := 0, parts[0]
	for i := 0; i < len(board) && ctr < 64; i++ {
		if '0' < board[i] && board[i] < '9' {
			cap := ctr + int(board[i]-'0')
			for ctr < cap {
				grid[ctr] = '1'
				ctr++
			}
		} else if board[i] != '/' {
			grid[ctr] = board[i]
			ctr++
		}
	}
	if ctr != 64 {
		return nil, fmt.Errorf("invalid board dimensions: %d != 64", len(board))
	}

	// Parse castling
	castle := uint8(0)
	for i := 0; i < len(parts[2]); i++ {
		switch parts[2][i] {
		case 'K':
			castle |= 8
		case 'Q':
			castle |= 4
		case 'k':
			castle |= 2
		case 'q':
			castle |= 1
		}
	}

	// halfmove parsing
	halfmove, err := strconv.ParseUint(parts[4], 10, 8)
	if err != nil {
		return nil, err
	}

	// count parsing
	count, err := strconv.ParseUint(parts[5], 10, 32)
	if err != nil {
		return nil, err
	}

	return &State{
		board:     grid,
		isBlack:   parts[1] == "b",
		castling:  castle,
		enPassant: ParseLocation(parts[3]),
		halfmove:  uint8(halfmove),
		count:     uint32(count),
	}, nil
}
