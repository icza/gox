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

// LimitRunes returns a slice of s that contains at most the given number of runes.
// If n is 0 or negative, the empty string is returned.
// If s has more runes than n, s is returned.
//
// Each byte of invalid UTF-8 sequences count as one when counting the limit, e.g.
//   LimitRunes("\xff\xffab", 3) // returns "\xff\xffa"
func LimitRunes(s string, n int) string {
	if n <= 0 || s == "" {
		return ""
	}

	for i := range s {
		n--
		if n == -1 {
			return s[:i]
		}
	}
	// s has n or more runes
	return s
}
