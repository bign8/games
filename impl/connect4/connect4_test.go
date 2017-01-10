package connect4

import "testing"

func BenchmarkIsInARow(b *testing.B) {
	board := &c4{
		board: [7][]byte{
			[]byte{'a'}, []byte{'a'}, []byte{'a'}, []byte{'a'},
			[]byte{' '}, []byte{' '}, []byte{' '},
		},
	}
	for i := 0; i < b.N; i++ {
		isInARow(board)
	}
}

func BenchmarkIsInARow2(b *testing.B) {
	board := &c4{
		board: [7][]byte{
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
		},
	}
	for i := 0; i < b.N; i++ {
		isInARow(board)
	}
}

func TestIsInARow(t *testing.T) {
	b := &c4{
		board: [7][]byte{
			[]byte{'a'}, []byte{'a'}, []byte{'a'}, []byte{'a'},
			[]byte{' '}, []byte{' '}, []byte{' '},
		},
	}
	if isInARow(b) < 0 {
		t.Error("Did not find 4 in a row")
	}
	b.board = [7][]byte{
		[]byte{'a'}, []byte{'a'}, []byte{'a'},
		[]byte{' '}, []byte{' '}, []byte{' '}, []byte{' '},
	}
	if isInARow(b) >= 0 {
		t.Error("Found 4 in a row")
	}
}

func BenchmarkString(b *testing.B) {
	board := &c4{
		board: [7][]byte{
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
		},
	}
	for i := 0; i < b.N; i++ {
		board.String()
	}
}

func TestString(t *testing.T) {
	board := &c4{
		board: [7][]byte{
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'Y', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'l', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g'},
			[]byte{'h', 'i', 'j', 'k', 'R', 'm', 'n'},
			[]byte{'a', 'b', 'c', 'd', 'e', 'f'},
		},
	}
	board.String()
}
