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

func TestLimitRunes(t *testing.T) {
	cases := []struct {
		name string
		s    string
		n    int
		exp  string
	}{
		{"empty", "", 0, ""},
		{"empty-2", "", 10, ""},
		{"normal", "abc", 0, ""},
		{"normal-2", "abc", 2, "ab"},
		{"normal-3", "abc", 3, "abc"},
		{"normal-3", "abc", 20, "abc"},
		{"unicode", "世界世界", 0, ""},
		{"unicode-2", "世界世界", 2, "世界"},
		{"unicode-3", "世界世界", 20, "世界世界"},
		{"invalid-utf-8", "\xff", 0, ""},
		{"invalid-utf-8-2", "\xff", 1, "\xff"},
		{"invalid-utf-8-3", "\xff", 2, "\xff"},
		{"invalid-utf-8#2", "a\xffb", 0, ""},
		{"invalid-utf-8#2-2", "a\xffb", 1, "a"},
		{"invalid-utf-8#2-3", "a\xffb", 2, "a\xff"},
		{"invalid-utf-8#2-4", "a\xffb", 3, "a\xffb"},
		{"invalid-utf-8#2-5", "a\xffb", 4, "a\xffb"},
		{"invalid-utf-8#3", "\xff\xffb", 0, ""},
		{"invalid-utf-8#3-2", "\xff\xffb", 1, "\xff"},
		{"invalid-utf-8#3-3", "\xff\xffb", 2, "\xff\xff"},
		{"invalid-utf-8#3-4", "\xff\xffb", 3, "\xff\xffb"},
		{"invalid-utf-8#3-5", "\xff\xffb", 4, "\xff\xffb"},
	}

	for _, c := range cases {
		if got := LimitRunes(c.s, c.n); got != c.exp {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.exp, got)
		}
	}
}
