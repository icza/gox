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
		args   []any
		exp    string
	}{
		{"en-without-args", en, nil, "I%d"},
		{"en-with-zero-args", en, []any{}, "I%!d(MISSING)"},
		{"en-with-args", en, []any{1}, "I1"},
		{"hu-without-args", hu, nil, "én%d"},
		{"hu-with-zero-args", hu, []any{}, "én%!d(MISSING)"},
		{"hu-with-args", hu, []any{1}, "én1"},
		{"de-without-args", de, nil, ""},
		{"de-with-zero-args", de, []any{}, ""},
		{"de-with-args", de, []any{1}, ""},
		{"gr-without-args", gr, nil, "I%d"},
		{"gr-with-zero-args", gr, []any{}, "I%!d(MISSING)"},
		{"gr-with-args", gr, []any{1}, "I1"},
	}

	for _, c := range cases {
		if got := get(c.locale, c.args...); got != c.exp {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.exp, got)
		}
	}
}
