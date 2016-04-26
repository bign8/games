package chess

import "fmt"

// NewMove takes 2 chess locations and builds a move.
func NewMove(src, dst Location) *Move {
	return &Move{
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
	mate      bool
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
		if m.mate {
			template += " (check)"
		} else {
			template += " (checkmate)"
		}
	}
	return fmt.Sprintf(template, m.Start, m.Stop)
}

// Moves gives the list of possible moves to take given a state of the game
func (s *State) Moves() []*Move {
	if s.mate {
		return nil
	}
	if s.moves == nil {
		s.moves = make([]*Move, 0)
		var newMoves []*Move
		var attacks []*Move
		for idx := byte(0); idx < 64; idx++ {
			if s.piece(idx) {
				next := Location(idx)
				switch s.board[idx] {
				case 'p':
					fallthrough
				case 'P':
					newMoves = s.pawnMoves(next)
				case 'r':
					fallthrough
				case 'R':
					newMoves = s.rookMoves(next)
				case 'n':
					fallthrough
				case 'N':
					newMoves = s.knightMoves(next)
				case 'b':
					fallthrough
				case 'B':
					newMoves = s.bishopMoves(next)
				case 'q':
					fallthrough
				case 'Q':
					newMoves = s.queenMoves(next)
				case 'k':
					fallthrough
				case 'K':
					newMoves = s.kingMoves(next)
				}
				if s.black(idx) == s.isBlack {
					s.moves = append(s.moves, newMoves...)
				} else {
					attacks = append(attacks, newMoves...)
				}
			}
		}
		// TODO: figure out if move is check-mate
		// TODO: other types of check
		// - Discovered Check: https://en.wikipedia.org/wiki/Discovered_check
		// - Double Check: https://en.wikipedia.org/wiki/Double_check
		// - Cross Check: https://en.wikipedia.org/wiki/Cross-check
		// TODO: VERIFY WE ARE NOT PUTTING THE KING INTO CHECK
		badSquares := make(map[Location]struct{})
		for _, attack := range attacks {
			badSquares[attack.Stop] = struct{}{}
		}
		var intoCheck []int
		for i, move := range s.moves {
			isKingStop := s.board[move.Stop] == 'k' || s.board[move.Stop] == 'K'
			s.moves[i].check = s.black(move.Stop.toInt()) != s.isBlack && isKingStop
			if _, ok := badSquares[move.Stop]; isKingStop && ok {
				intoCheck = append(intoCheck, i)
			}
		}
		// TODO: remove intoChecks from moves
		for i := len(intoCheck) - 1; i >= 0; i-- {

		}
	}
	return s.moves
}

func (s State) pawnMoves(loc Location) (res []*Move) {
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

	if move != InvalidLocation && !s.piece(move.toInt()) {
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
		if start != InvalidLocation && !s.piece(start.toInt()) && isStarting {
			m := NewMove(loc, start)
			m.passing = move
			res = append(res, m)
		}
	}
	idx := left.toInt()
	if left != InvalidLocation && ((s.piece(idx) && s.black(idx) != s.isBlack) || left == s.enPassant) {
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
	idx = right.toInt()
	if right != InvalidLocation && ((s.piece(idx) && s.black(idx) != s.isBlack) || right == s.enPassant) {
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

func (s State) rookMoves(loc Location) (res []*Move) {
	// https://en.wikipedia.org/wiki/Rook_(chess)
	x := []int8{0, 1, -1, 0}
	y := []int8{1, 0, 0, -1}
	for i := 0; i < len(x); i++ {
		next := loc.offset(x[i], y[i])
		idx := next.toInt()
		for next != InvalidLocation && !s.piece(idx) {
			res = append(res, NewMove(loc, next))
			next = next.offset(x[i], y[i])
			idx = next.toInt()
		}
		if next != InvalidLocation && s.piece(idx) && s.black(idx) != s.isBlack {
			res = append(res, NewMove(loc, next))
		}
	}
	return res
}

func (s State) knightMoves(loc Location) (res []*Move) {
	// https://en.wikipedia.org/wiki/Knight_(chess)
	x := []int8{1, 1, 2, 2, -1, -1, -2, -2}
	y := []int8{2, -2, 1, -1, 2, -2, 1, -1}
	for i := 0; i < len(x); i++ {
		if m := loc.offset(x[i], y[i]); m != InvalidLocation {
			idx := m.toInt()
			if s.piece(idx) && s.black(idx) == s.isBlack {
				continue
			}
			res = append(res, NewMove(loc, m))
		}
	}
	return res
}

func (s State) bishopMoves(loc Location) (res []*Move) {
	// https://en.wikipedia.org/wiki/Bishop_(chess)
	x := []int8{1, 1, -1, -1}
	y := []int8{1, -1, 1, -1}
	for i := 0; i < len(x); i++ {
		next := loc.offset(x[i], y[i])
		idx := next.toInt()
		for next != InvalidLocation && !s.piece(idx) {
			res = append(res, NewMove(loc, next))
			next = next.offset(x[i], y[i])
			idx = next.toInt()
		}
		if next != InvalidLocation && s.piece(idx) && s.black(idx) != s.isBlack {
			res = append(res, NewMove(loc, next))
		}
	}
	return res
}

func (s State) queenMoves(loc Location) (res []*Move) {
	// https://en.wikipedia.org/wiki/Queen_(chess)
	x := []int8{0, 1, 1, 1, -1, -1, 0, -1}
	y := []int8{1, 1, 0, -1, 1, 0, -1, -1}
	for i := 0; i < len(x); i++ {
		next := loc.offset(x[i], y[i])
		idx := next.toInt()
		for next != InvalidLocation && !s.piece(idx) {
			res = append(res, NewMove(loc, next))
			next = next.offset(x[i], y[i])
			idx = next.toInt()
		}
		if next != InvalidLocation && s.piece(idx) && s.black(idx) != s.isBlack {
			res = append(res, NewMove(loc, next))
		}
	}
	return res
}

func (s State) kingMoves(loc Location) (res []*Move) {
	// https://en.wikipedia.org/wiki/King_(chess)
	// TODO castling https://en.wikipedia.org/wiki/Castling
	x := []int8{0, 1, 1, 1, -1, -1, 0, -1}
	y := []int8{1, 1, 0, -1, 1, 0, -1, -1}
	for i := 0; i < len(x); i++ {
		if m := loc.offset(x[i], y[i]); m != InvalidLocation {
			idx := m.toInt()
			if s.piece(idx) && s.black(idx) == s.isBlack {
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
				m.castling = NewMove(rook, bishop)
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
				m.castling = NewMove(rook, queen) // rook move
				res = append(res, m)
			}
		}
	}

	return res
}

func (s State) black(idx uint8) bool {
	return 'a' < s.board[idx] && s.board[idx] < 'z'
}

func (s State) piece(idx uint8) bool {
	return s.board[idx] != '1'
}
