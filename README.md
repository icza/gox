# gox

[![Build Status](https://travis-ci.org/icza/gox.svg?branch=master)](https://travis-ci.org/icza/gox)
[![GoDoc](https://godoc.org/github.com/icza/gox?status.svg)](https://godoc.org/github.com/icza/gox)
[![Go Report Card](https://goreportcard.com/badge/github.com/icza/gox)](https://goreportcard.com/report/github.com/icza/gox)
[![codecov](https://codecov.io/gh/icza/gox/branch/master/graph/badge.svg)](https://codecov.io/gh/icza/gox)

_This module is in experimental phase._

The `gox` module is a minimalistic, lightweigt extension to Go.
It contains constants, helpers and utilities which could have been part of Go itself.

## Module structure

- `gox`: General extensions and utilities to Go itself. Reasonable to "dot-import"
the package.
- [`gox/timex`](https://github.com/icza/gox/tree/master/timex): contains time and duration related calculations and utilities.
It means to be a complement to the standard `time` package.

## `gox` package

Reasonable to "dot-import" the package, so identifiers will be directly available:

	import (
		"fmt"
		"strconv"

		. "github.com/icza/gox"
	)

	func main() {
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
		// Parsed: 3
	}
