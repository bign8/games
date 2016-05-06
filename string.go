package chess

import "strings"

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

	// Parse out player
	player := "White"
	if s.isBlack {
		player = "Black"
	}
	debug := ""
	if s.check {
		debug = "Player is in check!"
	}
	// debug = s.FEN() + "\n" + string(s.board[:])
	return top + strings.Join(rows, sep) + bot + "\n" + debug + "\n" + player + "'s Turn"
}
