package chess

import (
	"strings"
	"testing"
	"unicode/utf8"
)

func BenchmarkStateString(b *testing.B) {
	game := New(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = game.String()
	}
}

func TestNewString(b *testing.T) {
	x := New(2).String()
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

// BlockifyString adds enough spaces to the end of a string to appear "square"
func BlockifyString(str string) string {
	lines := strings.Split(str, "\n")
	lines = BlockifyLines(lines)
	return strings.Join(lines, "\n")
}

// BlockifyLines does the same as BlokifyString but with lists of lines
func BlockifyLines(lines []string) []string {
	cache := make([]int, len(lines))

	// Find max length line
	max := 0
	for i, line := range lines {
		cache[i] = utf8.RuneCountInString(line)
		if cache[i] > max {
			max = cache[i]
		}
	}

	// Pad each line to max
	for i := range lines {
		lines[i] += strings.Repeat(" ", max-cache[i])
	}
	return lines
}
