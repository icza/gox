package stringsx

import (
	"strings"
	"unicode"
)

// Clean removes non-graphic characters from the given string.
// Removable characters are the ones for which unicode.IsGraphic() returns false.
//
// For details, see https://stackoverflow.com/a/58994297/1705598
func Clean(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}
