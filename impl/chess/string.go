package chess

var chrLookup = map[uint8]rune{
	'p': '♟', 'r': '♜', 'n': '♞', 'b': '♝', 'q': '♛', 'k': '♚',
	'P': '♙', 'R': '♖', 'N': '♘', 'B': '♗', 'Q': '♕', 'K': '♔',
}

var numLookup = map[uint8]int{
	'1': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7,
}

const col = " ║ "
const top = "╔═══╦═══╦═══╦═══╦═══╦═══╦═══╦═══╗\n"
const sep = "\n╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣\n"
const bot = "\n╚═══╩═══╩═══╩═══╩═══╩═══╩═══╩═══╝\n  A   B   C   D   E   F   G   H"

// toGrid converts the current state to 64 length string that represents a board
func (s State) toGrid() []rune {
	bits := make([]rune, 64)
	for i := 0; i < 64; i++ {
		bits[i] = ' '
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
	return bits
}

// String is to implement the fmt.Stringer interface
func (s State) String() string {
	bits := s.toGrid()

	join := func(i int) string {
		return string(bits[i]) + col +
			string(bits[i+1]) + col +
			string(bits[i+2]) + col +
			string(bits[i+3]) + col +
			string(bits[i+4]) + col +
			string(bits[i+5]) + col +
			string(bits[i+6]) + col +
			string(bits[i+7])
	}

	rows := "║ " + join(0) + " ║  8" + sep +
		"║ " + join(8) + " ║  7" + sep +
		"║ " + join(16) + " ║  6" + sep +
		"║ " + join(24) + " ║  5" + sep +
		"║ " + join(32) + " ║  4" + sep +
		"║ " + join(40) + " ║  3" + sep +
		"║ " + join(48) + " ║  2" + sep +
		"║ " + join(56) + " ║  1"

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
	return top + rows + bot + "\n" + debug + "\n" + player + "'s Turn"
}
