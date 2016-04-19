package chess

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"
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
	}
	y := strings.Join(lines, "\n")
	if x != y {
		b.Errorf("String Missmatch\n%s", compare(x, y))
	}
}

func compare(a, b string) string {
	al := strings.Split(a, "\n")
	bl := strings.Split(b, "\n")

	max := 0
	for i := 0; i < len(al); i++ {
		if len(al[i]) > max {
			max = utf8.RuneCountInString(al[i])
		}
	}
	mask := fmt.Sprintf("%%-%ds", max+10)
	// panic(mask)
	log.Printf("Mash: %q", mask)

	out := make([]string, len(al))
	for i := 0; i < len(al); i++ {
		out[i] = fmt.Sprintf(mask, al[i]) + bl[i]
	}

	return strings.Join(out, "\n") + " " + strconv.Itoa(max) + " " + mask
}
