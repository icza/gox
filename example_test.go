package gox_test

import (
	"fmt"
	"strconv"

	. "github.com/icza/gox"
)

func Example() {
	s := struct {
		B      *bool
		U, Max *uint
	}{NewBool(true), NewUint(1), NewUint(MaxUint)}

	fmt.Println("*s.B:", *s.B, "*s.U:", *s.U)

	for _, age := range []int{10, 20} {
		state := If(age < 18).String("child", "adult")
		fmt.Printf("%d-years-old: %s\n", age, state)
	}

	for _, tempC := range []int{-5, 10} {
		fmt.Printf("State of water at %d°C: %s\n",
			tempC, IfString(tempC < 0, "solid", "liquid"))
	}

	n, err := strconv.Atoi("3")
	Pie(err)
	fmt.Println(n)

	// Output:
	// *s.B: true *s.U: 1
	// 10-years-old: child
	// 20-years-old: adult
	// State of water at -5°C: solid
	// State of water at 10°C: liquid
	// 3
}
