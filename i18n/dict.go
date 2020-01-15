package i18n

import "fmt"

// Empty is a string which translates into the empty string "".
// Needed because empty strings in Dict denote missing translations.
const Empty = "\xff"

// Dict describes a dictionary, it holds translations of a phrase or sentence.
type Dict []string

// Get returns the translation for the given locale.
// If no translation exists for the given locale, translation for the default
// locale (which is 0) is returned.
//
// If arguments are provided, the translation is treated as a format string,
// and fmt.Sprintf() is called to generate the result.
func (d Dict) Get(locale int, a ...interface{}) string {
	var format string
	if locale < len(d) {
		format = d[locale]
	}

	if format == "" {
		format = d[0]
	}

	if format == Empty {
		return ""
	}
	if a == nil {
		return format
	}
	return fmt.Sprintf(format, a...)
}

// Translator describes the type of Dict.Get method.
type Translator func(locale int, a ...interface{}) string
