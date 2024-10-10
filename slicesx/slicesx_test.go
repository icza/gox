package slicesx

import (
	"testing"
)

func TestIndex(t *testing.T) {
	cases := []struct {
		name     string
		exp, got any
	}{
		{
			"[]int",
			2, Index([]int{1, 2, 3}, 1),
		},
		{
			"[]int nil",
			0, Index[int](nil, 0),
		},
		{
			"[]int default",
			2, Index[int](nil, 0, 2),
		},
		{
			"[]int default negative index",
			3, Index([]int{1, 2}, -1, 3),
		},
		{
			"[]int not needing default",
			2, Index[int]([]int{1, 2}, 1, 3),
		},
		{
			"[]string",
			"2", Index([]string{"1", "2", "3"}, 1),
		},
		{
			"[]string nil",
			"", Index[string](nil, 0),
		},
		{
			"[]string default",
			"2", Index[string](nil, 0, "2"),
		},
		{
			"[]int default negative index",
			"3", Index([]string{"1", "2"}, -1, "3"),
		},
		{
			"[]string not needing default",
			"2", Index([]string{"1", "2"}, 1, "3"),
		},
	}

	for _, c := range cases {
		if c.exp != c.got {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, c.got)
		}
	}
}
