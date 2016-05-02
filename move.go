package chess

import "fmt"

// NewMove takes 2 chess locations and builds a move.
func NewMove(src, dst Location) Move {
	return Move{
		Start:     src,
		Stop:      dst,
		passing:   InvalidLocation,
		promotion: 0,
	}
}

/*
Chr | upper | lower |  u bin  |  l bin  |
 r  |   82  |  114  | 1010010 | 1110010 |
 n  |   78  |  110  | 1001110 | 1101110 |
 b  |   66  |   98  | 1000010 | 1100010 |
 q  |   81  |  113  | 1010001 | 1110001 |
*/

// Move represents a single move in the game of chess
type Move struct {
	Start     Location
	Stop      Location
	passing   Location
	promotion uint8 // 0:n/a, 1:rook, 2:knight, 3:bishop, 4:queen
	castling  *Move // nil unless castling
	check     bool
}

var promotionLookup = []string{"n/a", "Rook", "Knight", "Bishop", "Queen"}

// Equals checks if Moves are equal
func (m Move) Equals(n *Move) bool {
	pos := m.Start == n.Start && m.Stop == n.Stop
	special := n.passing == n.passing && n.promotion == n.promotion
	var castling bool
	if m.castling == nil && n.castling == nil {
		castling = true
	} else if m.castling == nil || n.castling == nil {
		castling = false
	} else {
		castling = m.castling.Start == n.castling.Start && m.castling.Stop == n.castling.Stop
	}
	return pos && special && castling
}

// String prints a human readable move
func (m Move) String() string {
	template := "%s -> %s"

	if m.promotion > 0 {
		template += " (" + promotionLookup[m.promotion] + ")"
	}
	if m.castling != nil {
		template += " (castling)"
	}
	if m.check {
		template += " (check)"
	}
	// template += " " + strconv.Itoa(int(m.Start.toInt())) + " " + strconv.Itoa(int(m.Stop.toInt()))
	return fmt.Sprintf(template, m.Start, m.Stop)
}

// Moves gives the list of possible moves to take given a state of the game
func (s *State) Moves() []*Move {
	if s.moves == nil {
		newMoves := make([]Move, 0, 100) // 206 theory
		for idx := Location(0); idx < 64; idx++ {
			if s.board[idx] != '1' {
				// TODO: remove this and figure out a better way
				if s.black(idx) != s.isBlack {
					continue
				}
				switch s.board[idx] {
				case 'p':
					fallthrough
				case 'P':
					newMoves = s.pawnMoves(idx, newMoves) // 12 * 8
				case 'r':
					fallthrough
				case 'R':
					newMoves = s.rookMoves(idx, newMoves) // 14 * 2
				case 'n':
					fallthrough
				case 'N':
					newMoves = s.knightMoves(idx, newMoves) // 8 * 2
				case 'b':
					fallthrough
				case 'B':
					newMoves = s.bishopMoves(idx, newMoves) // 14 * 2
				case 'q':
					fallthrough
				case 'Q':
					newMoves = s.queenMoves(idx, newMoves) // 28 * 1
				case 'k':
					fallthrough
				case 'K':
					newMoves = s.kingMoves(idx, newMoves) // 10 * 1
				}
			}
		}

		// Generate pointer based move array
		s.moves = make([]*Move, 0, len(newMoves))
		for i := range newMoves {
			s.moves = append(s.moves, &(newMoves[i]))
		}

		// Remove moves that place king in check
		var mine, yours Location
		tail := len(s.moves)
		for i := 0; i < tail; i++ {
			m := s.moves[i]

			// apply move to state
			orig := s.board[m.Stop]
			s.board[m.Stop] = s.board[m.Start]
			s.board[m.Start] = '1'

			// find each king
			// TODO: only do this if the king has been moved
			numFound := 0
			for i := Location(0); i < 64 && numFound <= 2; i++ {
				if s.board[i] == 'k' || s.board[i] == 'K' {
					numFound++
					if s.black(i) == s.isBlack {
						mine = i
					} else {
						yours = i
					}
				}
			}

			// is my or their king in check?
			if s.isCheck(mine, s.isBlack) {
				tail--
				s.moves[i], s.moves[tail] = s.moves[tail], s.moves[i]
				i--
			} else if s.isCheck(yours, !s.isBlack) {
				s.moves[i].check = true
			}

			// revert move on state
			s.board[m.Start] = s.board[m.Stop]
			s.board[m.Stop] = orig

		}
		s.moves = s.moves[:tail]
		// fmt.Printf("Length of moves: %d, tail: %d\n", len(s.moves), tail)
	}
	return s.moves
}

func (s State) isCheck(loc Location, isBlack bool) bool {
	var temp, temp2 Location

	// Checking Pawns
	if isBlack {
		temp = loc.offset(1, 1)
		temp2 = loc.offset(1, -1)
	} else {
		temp = loc.offset(-1, 1)
		temp2 = loc.offset(-1, -1)
	}
	if temp != InvalidLocation && (s.board[temp] == 'p' || s.board[temp] == 'P') && s.black(temp) != isBlack {
		return true
	}
	if temp2 != InvalidLocation && (s.board[temp2] == 'p' || s.board[temp2] == 'P') && s.black(temp2) != isBlack {
		return true
	}

	// Checking Bishop/Queen
	for i := 0; i < len(bishopX); i++ {
		temp = loc.offset(bishopX[i], bishopY[i])
		for temp != InvalidLocation && s.board[temp] == '1' {
			temp = temp.offset(bishopX[i], bishopY[i])
		}
		if temp != InvalidLocation && (s.board[temp] == 'b' || s.board[temp] == 'B' || s.board[temp] == 'q' || s.board[temp] == 'Q') && s.black(temp) != isBlack {
			return true
		}
	}

	// Checking Rook/Queen
	for i := 0; i < len(rookX); i++ {
		temp = loc.offset(rookX[i], rookY[i])
		for temp != InvalidLocation && s.board[temp] == '1' {
			temp = temp.offset(rookX[i], rookY[i])
		}
		if temp != InvalidLocation && (s.board[temp] == 'r' || s.board[temp] == 'R' || s.board[temp] == 'q' || s.board[temp] == 'Q') && s.black(temp) != isBlack {
			return true
		}
	}

	// Checking Knights
	for i := 0; i < len(knightX); i++ {
		temp = loc.offset(knightX[i], knightY[i])
		if temp != InvalidLocation && (s.board[temp] == 'n' || s.board[temp] == 'N') && s.black(temp) != isBlack {
			return true
		}
	}
	return false
}

func (s State) pawnMoves(loc Location, res []Move) []Move {
	// https://en.wikipedia.org/wiki/Pawn_(chess)
	// TODO: cleanup duplicate logic here
	var isStarting bool
	var move, start, left, right Location
	row, _ := loc.rowCol()
	if s.isBlack {
		move = loc.offset(1, 0)
		start = loc.offset(2, 0)
		left = loc.offset(1, -1)
		right = loc.offset(1, 1)
		isStarting = row == 1
	} else {
		move = loc.offset(-1, 0)
		start = loc.offset(-2, 0)
		left = loc.offset(-1, -1)
		right = loc.offset(-1, 1)
		isStarting = row == 6
	}

	if move != InvalidLocation && !s.piece(move) {
		if row, _ := move.rowCol(); row == 0 || row == 7 {
			// Promotion (rook, knight, bishop, queen)
			for i := uint8(1); i < 5; i++ {
				m := NewMove(loc, move)
				m.promotion = i
				res = append(res, m)
			}
		} else {
			// Regular move
			res = append(res, NewMove(loc, move))
		}

		// Double start (set enPassant variable)
		if start != InvalidLocation && !s.piece(start) && isStarting {
			m := NewMove(loc, start)
			m.passing = move
			res = append(res, m)
		}
	}
	if left != InvalidLocation && ((s.piece(left) && s.black(left) != s.isBlack) || left == s.enPassant) {
		if row, _ := left.rowCol(); row == 0 || row == 7 {
			// Promotion (rook, knight, bishop, queen)
			for i := uint8(1); i < 5; i++ {
				m := NewMove(loc, left)
				m.promotion = i
				res = append(res, m)
			}
		} else {
			// Regular move
			res = append(res, NewMove(loc, left))
		}
	}
	if right != InvalidLocation && ((s.piece(right) && s.black(right) != s.isBlack) || right == s.enPassant) {
		if row, _ := right.rowCol(); row == 0 || row == 7 {
			// Promotion (rook, knight, bishop, queen)
			for i := uint8(1); i < 5; i++ {
				m := NewMove(loc, right)
				m.promotion = i
				res = append(res, m)
			}
		} else {
			// Regular move
			res = append(res, NewMove(loc, right))
		}
	}
	return res
}

var rookX, rookY = []int8{0, 1, -1, 0}, []int8{1, 0, 0, -1}

func (s State) rookMoves(loc Location, res []Move) []Move {
	// https://en.wikipedia.org/wiki/Rook_(chess)
	for i := 0; i < len(rookX); i++ {
		next := loc.offset(rookX[i], rookY[i])
		idx := next.toInt()
		for next != InvalidLocation && s.board[idx] == '1' {
			res = append(res, NewMove(loc, next))
			next = next.offset(rookX[i], rookY[i])
			idx = next.toInt()
		}
		if next != InvalidLocation && s.board[idx] != '1' && s.black(next) != s.isBlack {
			res = append(res, NewMove(loc, next))
		}
	}
	return res
}

var knightX, knightY = []int8{1, 1, 2, 2, -1, -1, -2, -2}, []int8{2, -2, 1, -1, 2, -2, 1, -1}

func (s State) knightMoves(loc Location, res []Move) []Move {
	// https://en.wikipedia.org/wiki/Knight_(chess)
	for i := 0; i < len(knightX); i++ {
		if m := loc.offset(knightX[i], knightY[i]); m != InvalidLocation {
			if s.piece(m) && s.black(m) == s.isBlack {
				continue
			}
			res = append(res, NewMove(loc, m))
		}
	}
	return res
}

var bishopX, bishopY = []int8{1, 1, -1, -1}, []int8{1, -1, 1, -1}

func (s State) bishopMoves(loc Location, res []Move) []Move {
	// https://en.wikipedia.org/wiki/Bishop_(chess)
	for i := 0; i < len(bishopX); i++ {
		next := loc.offset(bishopX[i], bishopY[i])
		for next != InvalidLocation && s.board[next] == '1' {
			res = append(res, NewMove(loc, next))
			next = next.offset(bishopX[i], bishopY[i])
		}
		if next != InvalidLocation && s.board[next] != '1' && s.black(next) != s.isBlack {
			res = append(res, NewMove(loc, next))
		}
	}
	return res
}

var allX, allY = []int8{0, 1, 1, 1, -1, -1, 0, -1}, []int8{1, 1, 0, -1, 1, 0, -1, -1}

func (s State) queenMoves(loc Location, res []Move) []Move {
	// https://en.wikipedia.org/wiki/Queen_(chess)
	for i := 0; i < 8; i++ {
		next := loc.offset(allX[i], allY[i])
		for next != InvalidLocation && !s.piece(next) {
			res = append(res, NewMove(loc, next))
			next = next.offset(allX[i], allY[i])
		}
		if next != InvalidLocation && s.board[next] != '1' && s.black(next) != s.isBlack {
			res = append(res, NewMove(loc, next))
		}
	}
	return res
}

func (s State) kingMoves(loc Location, res []Move) []Move {
	// https://en.wikipedia.org/wiki/King_(chess)
	for i := 0; i < 8; i++ {
		if m := loc.offset(allX[i], allY[i]); m != InvalidLocation {
			if s.board[m] != '1' && s.black(m) == s.isBlack {
				continue
			}
			res = append(res, NewMove(loc, m))
		}
	}

	// Castling is only permitted if we are not in check
	if !s.check {
		var kq uint8
		var home int8
		if s.isBlack {
			kq = (s.castling >> 2) & 3
			home = 0
		} else {
			kq = (s.castling >> 0) & 3
			home = 7
		}
		if kq&2 == 2 {
			// check for kingside castle
			rook := locFromRowCol(home, 7)
			knight := locFromRowCol(home, 6)
			bishop := locFromRowCol(home, 5)
			if s.board[knight] == '1' && s.board[bishop] == '1' {
				m := NewMove(loc, knight)
				m.castling = new(Move) // malloc
				*m.castling = NewMove(rook, bishop)
				res = append(res, m)
			}
		}
		if kq&1 == 1 {
			// check for queenside castle
			rook := locFromRowCol(home, 0)
			knight := locFromRowCol(home, 1)
			bishop := locFromRowCol(home, 2)
			queen := locFromRowCol(home, 3)
			if s.board[knight] == '1' && s.board[bishop] == '1' && s.board[queen] == '1' {
				m := NewMove(loc, bishop)
				m.castling = new(Move) // malloc
				*m.castling = NewMove(rook, queen)
				res = append(res, m)
			}
		}
	}
	return res
}

func (s State) black(idx Location) bool {
	return 'a' < s.board[idx] && s.board[idx] < 'z'
}

func (s State) piece(idx Location) bool {
	return s.board[idx] != '1'
}
