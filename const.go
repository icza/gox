// This file contains constants and general utilities.

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
