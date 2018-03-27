/*
Package gox is a minimalistic extension to Go. It contains constants, helpers
and utilities which could have been part of Go itself (could have been
built-in).

The package is minimalistic, and introduces no dependency to any package.
Most of the functions are eligible for inlining. And don't worry if you're not
using some of the functions, the compiler will exclude those from your binary.

An easy way to use this library is to "dot-import" the package so identifiers
will be directly available, see the package example.

*/
package gox

// Compile time constants to determine the min and max values of the int and uint types.
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// Untyped bool constants telling if the architecture is 32 or 64 bit.
const (
	// Arch32bit tells if the target architecture is 32-bit.
	Arch32bit = MaxUint == 0xffffffff

	// Arch64bit tells if the target architecture is 64-bit.
	Arch64bit = uint64(MaxUint) == 0xffffffffffffffff
)

// Pie is a "panic-if-error" utility: panics if the passed error is not nil.
// Should not be over-used, but may come handy to write code quickly.
func Pie(err error) {
	if err != nil {
		panic(err)
	}
}

// NewBool returns a pointer to the given bool value.
func NewBool(b bool) *bool { return &b }

// NewString returns a pointer to the given string value.
func NewString(s string) *string { return &s }

// NewInt returns a pointer to the given int value.
func NewInt(i int) *int { return &i }

// NewInt8 returns a pointer to the given int8 value.
func NewInt8(i int8) *int8 { return &i }

// NewInt16 returns a pointer to the given int16 value.
func NewInt16(i int16) *int16 { return &i }

// NewInt32 returns a pointer to the given int32 value.
func NewInt32(i int32) *int32 { return &i }

// NewInt64 returns a pointer to the given int64 value.
func NewInt64(i int64) *int64 { return &i }

// NewUint returns a pointer to the given uint value.
func NewUint(i uint) *uint { return &i }

// NewUint8 returns a pointer to the given uint8 value.
func NewUint8(i uint8) *uint8 { return &i }

// NewUint16 returns a pointer to the given uint16 value.
func NewUint16(i uint16) *uint16 { return &i }

// NewUint32 returns a pointer to the given uint32 value.
func NewUint32(i uint32) *uint32 { return &i }

// NewUint64 returns a pointer to the given uint64 value.
func NewUint64(i uint64) *uint64 { return &i }

// NewFloat32 returns a pointer to the given float32 value.
func NewFloat32(f float32) *float32 { return &f }

// NewFloat64 returns a pointer to the given float64 value.
func NewFloat64(f float64) *float64 { return &f }

// NewByte returns a pointer to the given byte value.
func NewByte(b byte) *byte { return &b }

// NewRune returns a pointer to the given rune value.
func NewRune(r rune) *rune { return &r }
