package gox

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestIf(t *testing.T) {
	{
		i1, i2 := 1, 2
		exp, got := i1, If(true, i1, i2)
		if got != exp {
			t.Errorf("[int-1] Expected %d, got: %d", exp, got)
		}
		exp, got = i2, If(false, i1, i2)
		if got != exp {
			t.Errorf("[int-2] Expected %d, got: %d", exp, got)
		}
	}

	{
		s1, s2 := "first", "second"
		exp, got := s1, If(true, s1, s2)
		if got != exp {
			t.Errorf("[string-1] Expected %s, got: %s", exp, got)
		}
		exp, got = s2, If(false, s1, s2)
		if got != exp {
			t.Errorf("[string-2] Expected %s, got: %s", exp, got)
		}
	}
}

func TestIfFunc(t *testing.T) {
	count1, count2 := 0, 0
	f1, f2 := func() int { count1++; return 1 }, func() int { count2++; return 2 }

	exp, got := 1, IfFunc(true, f1, f2)
	if got != exp {
		t.Errorf("[int-1] Expected %d, got: %d", exp, got)
	}
	// Check deferred, on-demand calls:
	if expCount1 := 1; count1 != expCount1 {
		t.Errorf("[int-1] Expected count1 %d, got: %d", expCount1, count1)
	}
	if expCount2 := 0; count2 != expCount2 {
		t.Errorf("[int-1] Expected count2 %d, got: %d", expCount2, count2)
	}

	exp, got = 2, IfFunc(false, f1, f2)
	if got != exp {
		t.Errorf("[int-2] Expected %d, got: %d", exp, got)
	}
	// Check deferred, on-demand calls:
	if expCount1 := 1; count1 != expCount1 {
		t.Errorf("[int-1] Expected count1 %d, got: %d", expCount1, count1)
	}
	if expCount2 := 1; count2 != expCount2 {
		t.Errorf("[int-1] Expected count2 %d, got: %d", expCount2, count2)
	}
}

func TestPtr(t *testing.T) {
	s := "a"
	sp := Ptr(s)
	if *sp != s {
		t.Errorf("Ptr[string] failed")
	}

	i := 2
	ip := Ptr(i)
	if *ip != i {
		t.Errorf("Ptr[int] failed")
	}
}

func TestMust(t *testing.T) {
	i := 1
	if got := Must(i, nil); got != i {
		t.Errorf("Must[int] failed")
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("Expected panic")
			}
		}()
		Must(i, errors.New("test")) // Expecting panic
		t.Error("Not expected to reach this")
	}()
}

func manyResults() (i, j, k int, s string, f float64) {
	return 1, 2, 3, "four", 5.0
}

func TestFirst(t *testing.T) {
	exp, got := 1, First(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}

func TestSecond(t *testing.T) {
	exp, got := 2, Second(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}
func TestThird(t *testing.T) {
	exp, got := 3, Third(manyResults())
	if got != exp {
		t.Errorf("Expected %d, got: %d", exp, got)
	}
}

func TestCoalesce(t *testing.T) {
	p1, p2 := Ptr(1), Ptr(2)

	cases := []struct {
		name     string
		exp, got any
	}{
		{
			"strings",
			"1", Coalesce("", "1", "2"),
		},
		{
			"strings first",
			"1", Coalesce("1", "2", "3"),
		},
		{
			"strings last",
			"1", Coalesce("", "", "1"),
		},
		{
			"strings all zero",
			"", Coalesce("", "", ""),
		},
		{
			"strings no args",
			"", Coalesce[string](),
		},
		{
			"ints",
			1, Coalesce(0, 1, 2, 3),
		},
		{
			"ints first",
			1, Coalesce(1, 2, 3),
		},
		{
			"ints last",
			1, Coalesce(0, 0, 0, 0, 1),
		},
		{
			"ints all zero",
			0, Coalesce(0, 0, 0, 0),
		},
		{
			"ints no args",
			0, Coalesce[int](),
		},
		{
			"pointers",
			p1, Coalesce(nil, p1, p2),
		},
		{
			"pointers first",
			p1, Coalesce(p1, p2),
		},
		{
			"pointers last",
			p1, Coalesce(nil, nil, p1),
		},
		{
			"pointers all zero",
			(*int)(nil), Coalesce[*int](nil, nil, nil),
		},
		{
			"pointers no args",
			(*int)(nil), Coalesce[*int](),
		},
	}

	for _, c := range cases {
		if c.exp != c.got {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, c.got)
		}
	}
}

func TestDeref(t *testing.T) {
	cases := []struct {
		name     string
		exp, got any
	}{
		{
			"*int",
			1, Deref(Ptr(1)),
		},
		{
			"*int nil",
			0, Deref[int](nil),
		},
		{
			"*int default",
			2, Deref[int](nil, 2),
		},
		{
			"*int not needing default",
			1, Deref[int](Ptr(1), 2),
		},
		{
			"*string",
			"1", Deref(Ptr("1")),
		},
		{
			"*string nil",
			"", Deref[string](nil),
		},
		{
			"*string default",
			"2", Deref[string](nil, "2"),
		},
		{
			"*string not needing default default",
			"1", Deref[string](Ptr("1"), "2"),
		},
	}

	for _, c := range cases {
		if c.exp != c.got {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, c.got)
		}
	}
}

func TestPie(t *testing.T) {
	Pie(nil)

	var wasPanic bool
	func() {
		defer func() {
			if x := recover(); x != nil {
				wasPanic = true
			}
		}()
		Pie(errors.New("test error"))
	}()
	if !wasPanic {
		t.Errorf("Expected panic!")
	}
}

// ExampleWrap shows how to use the Wrap() function.
func ExampleWrap() {
	now := time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC)
	// Note that without Wrap it's a compile-time error.
	fmt.Printf("Year: %d, month: %d, day: %d\n",
		Wrap(now.Date())...)

	// Output:
	// Year: 2020, month: 3, day: 4
}

func TestWrap(t *testing.T) {
	cases := []struct {
		name string
		in   []interface{}
		out  []interface{}
	}{
		{"nil", nil, nil},
		{"empty", []interface{}{}, []interface{}{}},
		{"one", []interface{}{1}, []interface{}{1}},
		{"two", []interface{}{byte(1), "two"}, []interface{}{byte(1), "two"}},
	}

	for _, c := range cases {
		if got := Wrap(c.in...); !reflect.DeepEqual(got, c.out) {
			t.Errorf("[%s] Expected: %#v, got: %#v",
				c.name, c.out, got)
		}
	}
}
