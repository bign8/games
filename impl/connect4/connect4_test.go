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
