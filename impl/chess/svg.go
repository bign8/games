package chess

import (
	"fmt"
	"strings"
)

func (s State) SVG(bool) string {
	pieces := make([]string, 0, 32)
	for i, chr := range s.toGrid() {
		if chr == ' ' {
			continue
		}
		str := fmt.Sprintf(`<text x="%d" y="%d.75">%c</text>`, i%8, i/8, chr)
		pieces = append(pieces, str)
	}
	parts := `<g style="font-size:1px">` + strings.Join(pieces, "") + "</g>"
	return `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1">` + parts + `</svg>`
}
