package chess

import (
	"fmt"
	"strconv"
	"strings"
)

//*

// // Out generats a minimal transmission of this data in []byte form
// func (s State) Out() string {
//
// 	// TODO: make this actually work
// 	// Consolidate 1's back down to 2-8 if possible
// 	// TODO: move all this logic to PARSE
// 	newBoard := s.board[:]
// 	length := 0
// 	for i := 63; i >= 0; i-- {
// 		if newBoard[i] == '1' {
// 			length++
// 			if length > 0 && i%8 == 0 { // line wrap
// 				newBoard = newBoard[:i] + strconv.Itoa(length) + newBoard[i+length:]
// 				length = 0
// 			}
// 		} else if length > 0 { // end of existing chain of numbers
// 			newBoard = newBoard[:i+1] + strconv.Itoa(length) + newBoard[i+length+1:]
// 			length = 0
// 		}
// 	}
//
// 	return "TODO"
// }

// Parse parses a state from []byte generated via Bytes()
func Parse(bits string) (*State, error) {
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
