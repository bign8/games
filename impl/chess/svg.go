package chess

import "strings"

func (s State) SVG(bool) string {
	top := `<svg xmlns="http://www.w3.org/2000/svg" viewBox="-.05 -.05 8.1 8.1"><g style="font-size:1px">`
	pieces := make([]string, 0, 32)
	pieces = append(pieces, `<text x="7" y="7.75">♟</text>`)
	return top + strings.Join(pieces, "") + `</g></svg>`
}
