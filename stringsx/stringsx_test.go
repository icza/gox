package stringsx

import "testing"

func TestClean(t *testing.T) {
	cases := []struct {
		name string
		s    string
		exp  string
	}{
		{"empty", "", ""},
		{"normal", "abc", "abc"},
		{"contains", "a\x00bc", "abc"},
		{"contains-multi-byte-runes", "\x01世\x02界\x03", "世界"},
		{"all-removable", "\x01\x02\x03", ""},
		{"invalid-utf-8", "\xff", "\ufffd"},
		{"invalid-utf-8-2", "a\x00\xffb\x00", "a\ufffdb"},
	}

	for _, c := range cases {
		if got := Clean(c.s); got != c.exp {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.exp, got)
		}
	}
}
