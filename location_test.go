package chess

import "testing"

var loc Location

func TestLocationParseGood(t *testing.T) {
	loc = ParseLocation("e4")
	if loc != 36 {
		t.Errorf("Parsing Location is bad: %d != 36", loc)
	}
	loc = ParseLocation("z4")
	if loc != InvalidLocation {
		t.Errorf("Invalid location created valid location (letter): %d", loc)
	}
	loc = ParseLocation("e9")
	if loc != InvalidLocation {
		t.Errorf("Invalid location created valid location (numer): %d", loc)
	}
	loc = ParseLocation("-")
	if loc != InvalidLocation {
		t.Errorf("Invalid location created valid location (size): %d", loc)
	}
}

func BenchmarkLocationParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		loc = ParseLocation("e4")
	}
}
