// This file contains helpers to provide a "ternary operator like" functionality.

package gox

// If is a helper type to form expressive ternary expressions being the
// concatenation of a type conversion and a method call such as:
//
//   i := If(cond).Int(a, b)
type If bool

// If returns a if c is true, b otherwise.
func (c If) If(a, b interface{}) interface{} {
	if c {
		return a
	}
	return b
}

// Bool returns a if c is true, b otherwise.
func (c If) Bool(a, b bool) bool {
	if c {
		return a
	}
	return b
}

// String returns a if c is true, b otherwise.
func (c If) String(a, b string) string {
	if c {
		return a
	}
	return b
}

// Int returns a if c is true, b otherwise.
func (c If) Int(a, b int) int {
	if c {
		return a
	}
	return b
}

// Int8 returns a if c is true, b otherwise.
func (c If) Int8(a, b int8) int8 {
	if c {
		return a
	}
	return b
}

// Int16 returns a if c is true, b otherwise.
func (c If) Int16(a, b int16) int16 {
	if c {
		return a
	}
	return b
}

// Int32 returns a if c is true, b otherwise.
func (c If) Int32(a, b int32) int32 {
	if c {
		return a
	}
	return b
}

// Int64 returns a if c is true, b otherwise.
func (c If) Int64(a, b int64) int64 {
	if c {
		return a
	}
	return b
}

// Uint returns a if c is true, b otherwise.
func (c If) Uint(a, b uint) uint {
	if c {
		return a
	}
	return b
}

// Uint8 returns a if c is true, b otherwise.
func (c If) Uint8(a, b uint8) uint8 {
	if c {
		return a
	}
	return b
}

// Uint16 returns a if c is true, b otherwise.
func (c If) Uint16(a, b uint16) uint16 {
	if c {
		return a
	}
	return b
}

// Uint32 returns a if c is true, b otherwise.
func (c If) Uint32(a, b uint32) uint32 {
	if c {
		return a
	}
	return b
}

// Uint64 returns a if c is true, b otherwise.
func (c If) Uint64(a, b uint64) uint64 {
	if c {
		return a
	}
	return b
}

// Float32 returns a if c is true, b otherwise.
func (c If) Float32(a, b float32) float32 {
	if c {
		return a
	}
	return b
}

// Float64 returns a if c is true, b otherwise.
func (c If) Float64(a, b float64) float64 {
	if c {
		return a
	}
	return b
}

// Byte returns a if c is true, b otherwise.
func (c If) Byte(a, b byte) byte {
	if c {
		return a
	}
	return b
}

// Rune returns a if c is true, b otherwise.
func (c If) Rune(a, b rune) rune {
	if c {
		return a
	}
	return b
}

// IfIf returns a if c is true, b otherwise.
func IfIf(c bool, a, b interface{}) interface{} {
	if c {
		return a
	}
	return b
}

// IfBool returns a if c is true, b otherwise.
func IfBool(c bool, a, b bool) bool {
	if c {
		return a
	}
	return b
}

// IfString returns a if c is true, b otherwise.
func IfString(c bool, a, b string) string {
	if c {
		return a
	}
	return b
}

// IfInt returns a if c is true, b otherwise.
func IfInt(c bool, a, b int) int {
	if c {
		return a
	}
	return b
}

// IfInt8 returns a if c is true, b otherwise.
func IfInt8(c bool, a, b int8) int8 {
	if c {
		return a
	}
	return b
}

// IfInt16 returns a if c is true, b otherwise.
func IfInt16(c bool, a, b int16) int16 {
	if c {
		return a
	}
	return b
}

// IfInt32 returns a if c is true, b otherwise.
func IfInt32(c bool, a, b int32) int32 {
	if c {
		return a
	}
	return b
}

// IfInt64 returns a if c is true, b otherwise.
func IfInt64(c bool, a, b int64) int64 {
	if c {
		return a
	}
	return b
}

// IfUint returns a if c is true, b otherwise.
func IfUint(c bool, a, b uint) uint {
	if c {
		return a
	}
	return b
}

// IfUint8 returns a if c is true, b otherwise.
func IfUint8(c bool, a, b uint8) uint8 {
	if c {
		return a
	}
	return b
}

// IfUint16 returns a if c is true, b otherwise.
func IfUint16(c bool, a, b uint16) uint16 {
	if c {
		return a
	}
	return b
}

// IfUint32 returns a if c is true, b otherwise.
func IfUint32(c bool, a, b uint32) uint32 {
	if c {
		return a
	}
	return b
}

// IfUint64 returns a if c is true, b otherwise.
func IfUint64(c bool, a, b uint64) uint64 {
	if c {
		return a
	}
	return b
}

// IfFloat32 returns a if c is true, b otherwise.
func IfFloat32(c bool, a, b float32) float32 {
	if c {
		return a
	}
	return b
}

// IfFloat64 returns a if c is true, b otherwise.
func IfFloat64(c bool, a, b float64) float64 {
	if c {
		return a
	}
	return b
}

// IfByte returns a if c is true, b otherwise.
func IfByte(c bool, a, b byte) byte {
	if c {
		return a
	}
	return b
}

// IfRune returns a if c is true, b otherwise.
func IfRune(c bool, a, b rune) rune {
	if c {
		return a
	}
	return b
}
