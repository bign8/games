package chess

import (
	"fmt"
	"strconv"
	"strings"
)

var _lookup = []byte{'z', '1', '2', '3', '4', '5', '6', '7', '8'}

// FEN generats a minimal transmission of this data in []byte form
func (s State) FEN() string {
	pieces := make([]byte, 64+7+1+4+2) // squares + slashes + piece + castling + spaces
	index := 0

	length := 0
	for i := 0; i < 64; i++ {
		if i%8 == 0 && i != 0 {
			pieces[index] = '/'
			index++
		}
		if s.board[i] == '1' {
			length++
			if length > 0 && i%8 == 7 { // line wrap
				pieces[index] = _lookup[length]
				index++
				length = 0
			}
		} else if length > 0 { // end of existing chain of numbers
			pieces[index] = _lookup[length]
			index++
			length = 0
		} else {
			pieces[index] = s.board[i]
			index++
		}
	}

	// Space
	pieces[index] = ' '
	index++

	// Set active player
	if s.isBlack {
		pieces[index] = 'b'
	} else {
		pieces[index] = 'w'
	}
	index++

	// Space
	pieces[index] = ' '
	index++

	if s.castling == 0 {
		pieces[index] = '-'
		index++
	} else {
		if s.castling&8 == 8 {
			pieces[index] = 'K'
			index++
		}
		if s.castling&4 == 4 {
			pieces[index] = 'Q'
			index++
		}
		if s.castling&2 == 2 {
			pieces[index] = 'k'
			index++
		}
		if s.castling&1 == 1 {
			pieces[index] = 'q'
			index++
		}
	}

	// Set remaining parts
	enPassant := s.enPassant.String()
	halfmove := strconv.FormatUint(uint64(s.halfmove), 10)
	count := strconv.FormatUint(uint64(s.count), 10)
	return string(pieces[:index]) + " " + enPassant + " " + halfmove + " " + count
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
