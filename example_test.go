package gox_test

import (
	"fmt"
	"strconv"
	"time"

	. "github.com/icza/gox"
)

func Example() {
	// Pointers to non-zero values
	b, i, s := NewBool(true), NewUint(1), NewString("hi")
	fmt.Printf("b: %t, i: %d, s: %s\n", *b, *i, *s)

	// One way of mimicing the ternary operator:
	for _, age := range []int{10, 20} {
		state := If(age < 18).String("child", "adult")
		fmt.Printf("Age: %d, state: %s\n", age, state)
	}

	// And another one:
	for _, tempC := range []int{-5, 10} {
		fmt.Printf("Temperature: %d°C, state: %s\n",
			tempC, IfString(tempC < 0, "solid", "liquid"))
	}

	// Pass multiple return values to variadic functions:
	now := time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Year: %d, month: %d, day: %d\n",
		Wrap(now.Date())...)

	// Quick "handling" of error:
	n, err := strconv.Atoi("3")
	Pie(err)
	fmt.Println("Parsed:", n)

	// Output:
	// b: true, i: 1, s: hi
	// Age: 10, state: child
	// Age: 20, state: adult
	// Temperature: -5°C, state: solid
	// Temperature: 10°C, state: liquid
	// Year: 2020, month: 3, day: 4
	// Parsed: 3
}
