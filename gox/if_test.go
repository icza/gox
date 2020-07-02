package gox

import (
	"reflect"
	"testing"
)

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
