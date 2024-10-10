package gox_test

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"time"

	. "github.com/icza/gox/gox"
)

func Example() {
	// Mimicing the ternary operator:
	for _, age := range []int{10, 20} {
		state := If(age < 18, "child", "adult")
		fmt.Printf("[If] Age: %d, state: %s\n", age, state)
	}

	// Pointers to non-zero values"
	b, i, s := Ptr(true), Ptr[uint](1), Ptr("hi")
	fmt.Printf("[Ptr] b: %t, i: %d, s: %s\n", *b, *i, *s)

	// Safely dereference pointers:
	var nilPtr *string
	fmt.Printf("[Deref] s: %q, nilPtr: %q\n", Deref(s), Deref(nilPtr))

	// Pass multiple return values to variadic functions:
	now := time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC)
	fmt.Printf("[Wrap] Year: %d, month: %d, day: %d\n",
		Wrap(now.Date())...)

	// Using one of multiple return values:
	fmt.Println("[First-Third] Time coordinates:", image.Point{
		X: First(now.Date()),
		Y: Third(now.Date()),
	})

	// First non-zero of many values:
	hostName := ""
	fmt.Println("[Coalesce] Host name:", Coalesce(hostName, os.Getenv("HOST"), "localhost"))

	// Quick "handling" of error:
	n, err := strconv.Atoi("3")
	Pie(err)
	fmt.Println("[Pie] Parsed:", n)
	fmt.Println("[Must] Parsed:", Must(strconv.Atoi("4")))

	// Output:
	// [If] Age: 10, state: child
	// [If] Age: 20, state: adult
	// [Ptr] b: true, i: 1, s: hi
	// [Deref] s: "hi", nilPtr: ""
	// [Wrap] Year: 2020, month: 3, day: 4
	// [First-Third] Time coordinates: (2020,4)
	// [Coalesce] Host name: localhost
	// [Pie] Parsed: 3
	// [Must] Parsed: 4
}
