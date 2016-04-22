package chess

import (
	"strings"
	"testing"
)

func BenchmarkNewString(b *testing.B) {
	game := New()
	for i := 0; i < b.N; i++ {
		game.String()
	}
}

func TestNewString(b *testing.T) {
	x := New().String()
	lines := []string{
		"╔═══╦═══╦═══╦═══╦═══╦═══╦═══╦═══╗",
		"║ ♜ ║ ♞ ║ ♝ ║ ♛ ║ ♚ ║ ♝ ║ ♞ ║ ♜ ║  8",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║ ♟ ║ ♟ ║ ♟ ║ ♟ ║ ♟ ║ ♟ ║ ♟ ║ ♟ ║  7",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║   ║   ║   ║   ║   ║   ║   ║   ║  6",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║   ║   ║   ║   ║   ║   ║   ║   ║  5",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║   ║   ║   ║   ║   ║   ║   ║   ║  4",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║   ║   ║   ║   ║   ║   ║   ║   ║  3",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║ ♙ ║ ♙ ║ ♙ ║ ♙ ║ ♙ ║ ♙ ║ ♙ ║ ♙ ║  2",
		"╠═══╬═══╬═══╬═══╬═══╬═══╬═══╬═══╣",
		"║ ♖ ║ ♘ ║ ♗ ║ ♕ ║ ♔ ║ ♗ ║ ♘ ║ ♖ ║  1",
		"╚═══╩═══╩═══╩═══╩═══╩═══╩═══╩═══╝",
		"  A   B   C   D   E   F   G   H",
		"",
		"White's Turn",
	}
	y := strings.Join(lines, "\n")
	if x != y {
		b.Errorf("String Missmatch\n%s", compare(x, y))
	}
}

func compare(a, b string) string {
	al := BlockifyLines(strings.Split(a, "\n"))
	bl := BlockifyLines(strings.Split(b, "\n"))
	padding := strings.Repeat(" ", 10)

	out := make([]string, len(al))
	for i := 0; i < len(al); i++ {
		out[i] = al[i] + padding + bl[i]
	}

	return strings.Join(out, "\n")
}
