package connect4

var (
	master []point
	buffer [4]point
)

func add(p point, c, r int8) point { return point{col: p.col + c, row: p.row + r} }
func valid(p point) bool           { return p.col >= 0 && p.col < 8 && p.row >= 0 && p.row < 7 }

// provides the 4 points to determin 4-in-a-row
func offsets(s point, dir int8) []point {
	buffer[0] = s
	switch dir {
	case 0: // vertical
		buffer[1] = add(s, 0, 1)
		buffer[2] = add(s, 0, 2)
		buffer[3] = add(s, 0, 3)
	case 1: // horizontal
		buffer[1] = add(s, 1, 0)
		buffer[2] = add(s, 2, 0)
		buffer[3] = add(s, 3, 0)
	case 2: // /
		buffer[1] = add(s, 1, 1)
		buffer[2] = add(s, 2, 2)
		buffer[3] = add(s, 3, 3)
	case 3: // \
		buffer[1] = add(s, 1, -1)
		buffer[2] = add(s, 2, -2)
		buffer[3] = add(s, 3, -3)
	}
	return buffer[:]
}

func addIfValid(col, row, dir int8) {
	four := offsets(point{col, row}, dir)
	for i := 1; i < 4; i++ {
		if !valid(four[i]) {
			return
		}
	}
	master = append(master, four...)
}

// Build master list of all possible 4-in-a-row's
func createList() {
	master = make([]point, 0, 93*4)
	for i := int8(0); i < 7; i++ {
		for j := int8(0); j < 6; j++ {
			for k := int8(0); k < 4; k++ {
				addIfValid(i, j, k)
			}
		}
	}
}

func init() {
	createList()
}
