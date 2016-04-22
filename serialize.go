package chess

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
