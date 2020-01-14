// This file contains constants.

package mathx

// Compile time constants to determine the min and max values of the int and uint types.
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)
