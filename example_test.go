package gox_test

import (
	"fmt"
	"strconv"

	. "github.com/icza/gox"
)

func Example() {
	b, one := NewBool(true), NewUint(1) // Pointers to non-zero values
	fmt.Printf("*b: %t, *one: %d\n", *b, *one)

	for _, age := range []int{10, 20} {
		state := If(age < 18).String("child", "adult")
		fmt.Printf("Age:           %d, state: %s\n", age, state)
	}

	for _, tempC := range []int{-5, 10} {
		fmt.Printf("Temperature: %d°C, state: %s\n",
			tempC, IfString(tempC < 0, "solid", "liquid"))
	}

	n, err := strconv.Atoi("3")
	Pie(err)
	fmt.Println(n)

	// Output:
	// *b: true, *one: 1
	// Age:           10, state: child
	// Age:           20, state: adult
	// Temperature: -5°C, state: solid
	// Temperature: 10°C, state: liquid
	// 3
}
