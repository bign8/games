package chess

import "fmt"

// Location defines a location on the chess board that allows fast lookups
type Location uint8

// InvalidLocation is returned if there is an error parsing a location
const InvalidLocation = 255

var colLookup = map[byte]uint8{
	'a': 0, 'b': 1, 'c': 2, 'd': 3, 'e': 4, 'f': 5, 'g': 6, 'h': 7,
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
}

var rowLookup = map[byte]uint8{
	'1': 7, '2': 6, '3': 5, '4': 4, '5': 3, '6': 2, '7': 1, '8': 0,
}

func locFromRowCol(row, col int8) Location {
	if 0 <= row && row < 8 && 0 <= col && col < 8 {
		return Location(row*8 + col)
	}
	return InvalidLocation
}

func (l Location) rowCol() (int8, int8) {
	return int8(l) / 8, int8(l) % 8
}

func (l Location) offset(down, right int8) Location {
	row, col := l.rowCol()
	return locFromRowCol(row+down, col+right)
}

func (l Location) toInt() uint8 {
	return uint8(l)
}

// String prints the human readable format of a Location
func (l Location) String() string {
	if l == InvalidLocation {
		return "-"
	}
	return fmt.Sprintf("%c%d", l%8+'A', 8-l/8)
}

// ParseLocation returns a new chess board location object
func ParseLocation(in string) Location {
	if len(in) != 2 {
		return InvalidLocation
	}
	var ok bool
	var row, col uint8
	if col, ok = colLookup[in[0]]; !ok {
		return InvalidLocation
	}
	if row, ok = rowLookup[in[1]]; !ok {
		return InvalidLocation
	}
	return Location(row*8 + col)
}
