package connect4

var inARow [][]point

func (p point) add(dc, dr int8) point {
	return point{col: p.col + dc, row: p.row + dr}
}

func (p point) valid() bool {
	return p.col >= 0 && p.col < 8 && p.row >= 0 && p.row < 7
}

// provides the 4 points to determin 4-in-a-row
func offsets(s point, dir int) []point {
	res := []point{s}
	switch dir {
	case 0: // vertical
		res = append(res, s.add(0, 1), s.add(0, 2), s.add(0, 3))
	case 1: // horizontal
		res = append(res, s.add(1, 0), s.add(2, 0), s.add(3, 0))
	case 2: // /
		res = append(res, s.add(1, 1), s.add(2, 2), s.add(3, 3))
	case 3: // \
		res = append(res, s.add(1, -1), s.add(2, -2), s.add(3, -3))
	}
	return res
}

func init() {

	// Build all possible 4-in-a-row's
	for i := int8(0); i < 7; i++ {
		for j := int8(0); j < 6; j++ {
			inARow = append(inARow, offsets(point{i, j}, 0))
			inARow = append(inARow, offsets(point{i, j}, 1))
			inARow = append(inARow, offsets(point{i, j}, 2))
			inARow = append(inARow, offsets(point{i, j}, 3))
		}
	}

	// Remove invalid array sets
outer:
	for i := 0; i < len(inARow); i++ {
		for j := 0; j < len(inARow[i]); j++ {
			if !inARow[i][j].valid() {
				inARow = append(inARow[:i], inARow[i+1:]...)
				i--
				continue outer
			}
		}
	}
}
