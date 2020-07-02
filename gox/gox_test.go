package gox

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

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
