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

Reasonable to "dot-import" the package, so identifiers
will be directly available:

	import (
		"fmt"
		"strconv"

		. "github.com/icza/gox"
	)

	func main() {
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
