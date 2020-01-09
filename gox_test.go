package gox

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestMinMaxUint(t *testing.T) {
	var max uint
	max = ^max

	if MaxUint != max {
		t.Errorf("MaxUint is not max!")
	}
	if MinUint != max+1 {
		t.Errorf("MinUint is not min!")
	}
}

func TestMinMaxInt(t *testing.T) {
	maxuint := uint(0)
	maxuint = ^maxuint
	max := int(maxuint >> 1)

	if MaxInt != max {
		t.Errorf("MaxInt is not max!")
	}
	if MinInt != -max-1 {
		t.Errorf("MinInt is not min!")
	}
}

func TestArchConsts(t *testing.T) {
	arch32bit := reflect.TypeOf(0).Size() == 4

	if arch32bit != Arch32bit {
		t.Errorf("Arch32bit is incorrect!")
	}
	if arch32bit == Arch64bit {
		t.Errorf("Arch64bit is incorrect!")
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

func TestNewValues(t *testing.T) {
	type is []interface{}
	cases := []struct {
		name   string        // Name of the test case
		f      interface{}   // Function to call
		params []interface{} // Params to pass in subsequent calls
	}{
		{"NewBool", NewBool, is{false, true}},
		{"NewString", NewString, is{"", "a"}},
		{"NewInt", NewInt, is{0, 1}},
		{"NewInt8", NewInt8, is{int8(0), int8(1)}},
		{"NewInt16", NewInt16, is{int16(0), int16(1)}},
		{"NewInt32", NewInt32, is{int32(0), int32(1)}},
		{"NewInt64", NewInt64, is{int64(0), int64(1)}},
		{"NewUint", NewUint, is{uint(0), uint(1)}},
		{"NewUint8", NewUint8, is{uint8(0), uint8(1)}},
		{"NewUint16", NewUint16, is{uint16(0), uint16(1)}},
		{"NewUint32", NewUint32, is{uint32(0), uint32(1)}},
		{"NewUint64", NewUint64, is{uint64(0), uint64(1)}},
		{"NewFloat32", NewFloat32, is{float32(0), float32(1)}},
		{"NewFloat64", NewFloat64, is{float64(0), float64(1)}},
		{"NewByte", NewByte, is{byte(0), byte(1)}},
		{"NewRune", NewRune, is{rune(0), rune(1)}},
	}

	for _, c := range cases {
		fv := reflect.ValueOf(c.f)
		for _, param := range c.params {
			results := fv.Call([]reflect.Value{reflect.ValueOf(param)})
			if len(results) != 1 {
				t.Errorf("[%s(%v)] Expected 1 result(s), got: %d", c.name, param, len(results))
			}
			if got := results[0].Elem().Interface(); got != param {
				t.Errorf("[%s(%v)] Expected: %v, got: %v", c.name, param, param, got)
			}
		}
	}
}

func TestIfFuncs(t *testing.T) {
	cases := []struct {
		name string      // Name of the test case
		f    interface{} // Function to call
		a, b interface{} // a and b params to pass in
	}{
		{"IfIf", IfIf, "one", 2.0},
		{"IfBool", IfBool, false, true},
		{"IfString", IfString, "a", "b"},
		{"IfInt", IfInt, 1, 2},
		{"IfInt8", IfInt8, int8(1), int8(2)},
		{"IfInt16", IfInt16, int16(1), int16(2)},
		{"IfInt32", IfInt32, int32(1), int32(2)},
		{"IfInt64", IfInt64, int64(1), int64(2)},
		{"IfUint", IfUint, uint(1), uint(2)},
		{"IfUint8", IfUint8, uint8(1), uint8(2)},
		{"IfUint16", IfUint16, uint16(1), uint16(2)},
		{"IfUint32", IfUint32, uint32(1), uint32(2)},
		{"IfUint64", IfUint64, uint64(1), uint64(2)},
		{"IfFloat32", IfFloat32, float32(1), float32(2)},
		{"IfFloat64", IfFloat64, float64(1), float64(2)},
		{"IfByte", IfByte, byte(1), byte(2)},
		{"IfRune", IfRune, rune(1), rune(2)},
	}

	for _, c := range cases {
		for _, cond := range []bool{false, true} {
			var exp interface{}
			if cond {
				exp = c.a
			} else {
				exp = c.b
			}

			// First test the standalone function:
			fv := reflect.ValueOf(c.f)
			params := []reflect.Value{reflect.ValueOf(cond), reflect.ValueOf(c.a), reflect.ValueOf(c.b)}

			results := fv.Call(params)
			if len(results) != 1 {
				t.Errorf("[%s(%v, %v, %v)] Expected 1 result(s), got: %d",
					c.name, cond, c.a, c.b, len(results))
			}
			if got := results[0].Interface(); got != exp {
				t.Errorf("[%s(%v, %v, %v)] Expected: %v, got: %v",
					c.name, cond, c.a, c.b, exp, got)
			}

			// Next test the similar method of If:
			name := c.name[2:] // Trim leading "If"
			fv = reflect.ValueOf(If(cond)).MethodByName(name)
			params = params[1:] // cond is not a param

			results = fv.Call(params)
			if len(results) != 1 {
				t.Errorf("[If(%v).%s(%v, %v)] Expected 1 result(s), got: %d",
					cond, name, c.a, c.b, len(results))
			}
			if got := results[0].Interface(); got != exp {
				t.Errorf("[If(%v).%s(%v, %v)] Expected: %v, got: %v",
					cond, name, c.a, c.b, exp, got)
			}
		}
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
