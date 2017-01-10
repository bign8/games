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

func TestIsInARow(t *testing.T) {
	board := &c4{
		board: [7][]byte{
			[]byte{'a'}, []byte{'a'}, []byte{'a'}, []byte{'a'},
			[]byte{' '}, []byte{' '}, []byte{' '},
		},
	}
	x := isInARow(board)
	if x < 0 {
		t.Error("Did not find 4 in a row")
	}
}
