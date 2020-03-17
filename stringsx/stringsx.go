package stringsx

import (
	"strings"
	"unicode"
)

// Clean removes non-graphic characters from the given string.
// Removable characters are the ones for which unicode.IsGraphic() returns false.
func Clean(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, s)
}
