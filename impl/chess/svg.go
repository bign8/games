package chess

import (
	"fmt"
	"sort"
	"strings"
)

const (
	svgPieceTPL = `<text x="%d" y="%d">%c</text>`
	svgHoverTPL = `<rect x="%.2f" y="%d" height="1" width="%.2f" data-show="%s"/>`
	svgShownTPL = `<g data-slug="%s">
	<text x="%d" y="%d">%c</text>
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
			sort.Sort(ltr(group)) // sort so overlays appear the same each time
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
	return svgHead + `<style>text{text-anchor:middle;transform:translate(.5px,.5px);dominant-baseline:middle}</style>` + parts + svgTail
}

// Sort a list of moves from left to right based on start positiion.
// Iff they are the same, sort top down.
type ltr []*Move

func (a ltr) Len() int      { return len(a) }
func (a ltr) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ltr) Less(i, j int) bool {
	ir, ic := a[i].Start.rowCol()
	jr, jc := a[j].Start.rowCol()
	if ic == jc {
		return ir < jr
	}
	return ic < jc
}
