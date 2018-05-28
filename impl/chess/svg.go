package chess

import (
	"fmt"
	"strings"
)

const (
	svgPieceTPL = `<text x="%d" y="%d.75">%c</text>`
	svgHoverTPL = `<rect x="%.2f" y="%d" height="1" width="%.2f" data-show="%s"/>`
	svgShownTPL = `<g data-slug="%s">
	<text x="%d" y="%d.75">%c</text>
	<line x1="%d.5" y1="%d.5" x2="%d.5" y2="%d.5" stroke-linecap="round" stroke="rgba(255, 0, 0, 0.1)" stroke-width="0.1"/>
</g>`
	svgHead = `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1" style="font-size:1px">`
	svgTail = `</svg>`
)

// SVG shows the game state given the state of the document
func (s State) SVG(active bool) string {
	pieces := make([]string, 0, 32)
	gridz := s.toGrid()
	for i, chr := range gridz {
		if chr == ' ' {
			continue
		}
		str := fmt.Sprintf(svgPieceTPL, i%8, i/8, chr)
		pieces = append(pieces, str)
	}
	parts := `<g>` + strings.Join(pieces, "") + "</g>"

	// Draw active SVG move states
	if active {
		actions := s.Actions()

		// group actions by destination
		groups := make(map[Location][]*Move, 10)
		for _, action := range actions {
			move := action.(*Move)

			group, ok := groups[move.Stop]
			if !ok {
				group = make([]*Move, 0, 1)
			}
			groups[move.Stop] = append(group, move)
		}

		// for each group, split the cell into the various move targets
		hovers := make([]string, 0, len(actions))
		for spot, group := range groups {
			fraction := 1 / float64(len(group))
			row, col := spot.rowCol()
			for i, move := range group {
				left := float64(col) + float64(i)*fraction
				hovers = append(hovers, fmt.Sprintf(svgHoverTPL, left, row, fraction, move.Slug()))
			}
		}

		// Create destination piece
		moves := make([]string, len(actions))
		for i, action := range actions {
			move := action.(*Move)
			chr := gridz[int(move.Start.toInt())]
			row, col := move.Stop.rowCol()
			sr, sc := move.Start.rowCol()
			moves[i] = fmt.Sprintf(svgShownTPL, action.Slug(), col, row, chr, sc, sr, col, row)
		}
		parts += `<g>` + strings.Join(moves, "") + strings.Join(hovers, "") + "</g>"
	}
	return svgHead + parts + svgTail
}
