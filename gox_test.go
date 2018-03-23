package gox

import (
	"errors"
	"reflect"
	"testing"
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
	type is = []interface{}
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
