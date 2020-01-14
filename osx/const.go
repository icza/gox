// This file contains constants.

package osx

const (
	maxUint = ^uint(0)
)

const (
	// Arch32bit tells if the target architecture is 32-bit.
	Arch32bit = maxUint == 0xffffffff

	// Arch64bit tells if the target architecture is 64-bit.
	Arch64bit = uint64(maxUint) == 0xffffffffffffffff
)
