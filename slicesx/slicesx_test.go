package slicesx

import (
	"slices"
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
			2, Index(nil, 0, 2),
		},
		{
			"[]int default negative index",
			3, Index([]int{1, 2}, -1, 3),
		},
		{
			"[]int not needing default",
			2, Index([]int{1, 2}, 1, 3),
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
			"2", Index(nil, 0, "2"),
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

func TestSortByPropValueOrder(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	ageGetter := func(p Person) int { return p.Age }

	persons := []Person{
		{Name: "A", Age: 30},
		{Name: "B", Age: 43},
		{Name: "C", Age: 20},
		{Name: "D", Age: 11},
	}

	cases := []struct {
		name       string
		vals       []Person
		propGetter func(v Person) int
		propOrder  []int
		exp        []Person // Might be smaller than input
	}{
		{
			"empty input",
			nil,
			ageGetter,
			[]int{20, 30, 11, 43},
			nil,
		},
		{
			"normal case",
			persons,
			ageGetter,
			[]int{20, 30, 11, 43},
			[]Person{{Name: "C", Age: 20}, {Name: "A", Age: 30}, {Name: "D", Age: 11}, {Name: "B", Age: 43}},
		},
		{
			"missing props",
			persons,
			ageGetter,
			[]int{11, 43},
			[]Person{{Name: "D", Age: 11}, {Name: "B", Age: 43}},
		},
		{
			"missing props #2",
			persons,
			ageGetter,
			[]int{43, 30, 11},
			[]Person{{Name: "B", Age: 43}, {Name: "A", Age: 30}, {Name: "D", Age: 11}, {Name: "C", Age: 20}},
		},
		{
			"missing props #3",
			persons,
			ageGetter,
			nil,
			nil,
		},
	}

	for _, c := range cases {
		got := slices.Clone(c.vals)

		SortByPropValueOrder(got, c.propGetter, c.propOrder)

		if len(got) > len(c.exp) {
			got = got[:len(c.exp)]
		}

		if !slices.Equal(got, c.exp) {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, got)
		}
	}
}
