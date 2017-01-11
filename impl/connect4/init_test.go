package connect4

import "testing"

func BenchmarkCreateList(b *testing.B) {
	for i := 0; i < b.N; i++ {
		master = nil
		createList()
	}
}

func TestCreateList(t *testing.T) {
	master = nil
	createList()
	if len(master)/4 != 69 {
		t.Errorf("Master invalid length: %d != 69 * 4", len(master))
	}
}
