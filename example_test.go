package gox_test

import (
	"fmt"

	. "github.com/icza/gox"
)

func Example() {
	s := struct {
		B      *bool
		U, Max *uint
	}{NewBool(true), NewUint(1), NewUint(MaxUint)}

	fmt.Println(*s.B, *s.U)
	// Output: true 1
}
