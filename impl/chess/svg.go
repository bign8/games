package chess

import (
	"fmt"
	"strings"
)

func (s State) SVG(bool) string {
	whites := make([]string, 0, 16)
	blacks := make([]string, 0, 16)
	for i, chr := range s.toGrid() {
		if chr == ' ' {
			continue
		}
		row := i % 8
		col := i / 8
		str := fmt.Sprintf(`<text x="%d" y="%d.75">%c</text>`, row, col, chr)
		if row%2 == 0 && col%2 == 0 || row%2 == 1 && col%2 == 1 {
			blacks = append(blacks, str)
		} else {
			whites = append(whites, str)
		}
	}
	black := `<g style="font-size:1px">` + strings.Join(blacks, "") + "</g>"
	white := `<g style="font-size:1px;fill:white">` + strings.Join(whites, "") + "</g>"
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1">` + black + white + `</svg>`
}
