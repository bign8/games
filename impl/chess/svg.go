package chess

import (
	"fmt"
	"strings"
)

// SVG shows the game state given the state of the document
func (s State) SVG(active bool) string {
	pieces := make([]string, 0, 32)
	gridz := s.toGrid()
	for i, chr := range gridz {
		if chr == ' ' {
			continue
		}
		str := fmt.Sprintf(`<text x="%d" y="%d.75">%c</text>`, i%8, i/8, chr)
		pieces = append(pieces, str)
	}
	parts := `<g>` + strings.Join(pieces, "") + "</g>"

	// Draw active SVG move states
	if active {
		actions := s.Actions()
		// TODO: make a hide toggler for the pieces above that are being removed (needs JS to be updated)
		moves := make([]string, len(actions))
		for i, action := range actions {
			move := action.(*Move)
			chr := gridz[int(move.Start.toInt())]
			row, col := move.Stop.rowCol()
			moves[i] = fmt.Sprintf(`<g data-slug="%s"><text x="%d" y="%d.75">%c</text></g>`, action.Slug(), col, row, chr)
		}
		parts += `<g>` + strings.Join(moves, "") + "</p>"
	}
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1" style="font-size:1px">` + parts + `</svg>`
}
