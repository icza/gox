package i18n

import "testing"

// Check if Dict.Get matches Translator
var _ Translator = Dict{}.Get

func TestDict(t *testing.T) {
	const (
		en = iota
		de
		hu
		gr
	)
	get := Dict{
		en: "I%d",
		hu: "én%d",
		de: Empty,
	}.Get

	cases := []struct {
		name   string
		locale int
		args   []interface{}
		exp    string
	}{
		{"en-without-args", en, nil, "I%d"},
		{"en-with-zero-args", en, []interface{}{}, "I%!d(MISSING)"},
		{"en-with-args", en, []interface{}{1}, "I1"},
		{"hu-without-args", hu, nil, "én%d"},
		{"hu-with-zero-args", hu, []interface{}{}, "én%!d(MISSING)"},
		{"hu-with-args", hu, []interface{}{1}, "én1"},
		{"de-without-args", de, nil, ""},
		{"de-with-zero-args", de, []interface{}{}, ""},
		{"de-with-args", de, []interface{}{1}, ""},
		{"gr-without-args", gr, nil, "I%d"},
		{"gr-with-zero-args", gr, []interface{}{}, "I%!d(MISSING)"},
		{"gr-with-args", gr, []interface{}{1}, "I1"},
	}

	for _, c := range cases {
		if got := get(c.locale, c.args...); got != c.exp {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.exp, got)
		}
	}
}
