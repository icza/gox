# gox

[![Build Status](https://travis-ci.org/icza/gox.svg?branch=master)](https://travis-ci.org/icza/gox)
[![GoDoc](https://godoc.org/github.com/icza/gox?status.svg)](https://godoc.org/github.com/icza/gox)
[![Go Report Card](https://goreportcard.com/badge/github.com/icza/gox)](https://goreportcard.com/report/github.com/icza/gox)
[![codecov](https://codecov.io/gh/icza/gox/branch/master/graph/badge.svg)](https://codecov.io/gh/icza/gox)

_This package is in EXPERIMENTAL phase._

Package `gox` is a minimalistic extension to Go. It contains constants, helpers
and utilities which could have been part of Go itself (could have been
built-in).

The package is minimalistic, and introduces no dependency to any package.
Most of the functions are eligible for inlining. And don't worry if you're not
using some of the functions, the compiler will exclude those from your binary.

An easy way to use this library is to "dot-import" the package so identifiers
will be directly available:

	import (
		"fmt"

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
