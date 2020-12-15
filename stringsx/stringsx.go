package stringsx

import (
	"strings"
	"unicode"
	"unicode/utf8"
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
// If s has less runes than n, s is returned.
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
	// s has n or less runes
	return s
}

// SplitQuotes splits the given string by sep.
// If sep appears inside quotes, it is skipped.
// Quotes are not removed from the parts.
func SplitQuotes(s string, sep, quote rune) (parts []string) {
	sepLen := utf8.RuneLen(sep)

	start, inQuotes := 0, false

	for i, r := range s {
		if r == sep && !inQuotes {
			parts = append(parts, s[start:i])
			start = i + sepLen
			continue
		}
		if r == quote {
			inQuotes = !inQuotes
		}
	}

	return append(parts, s[start:])
}
