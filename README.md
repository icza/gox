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

    import . "github.com/icza/gox"

    func main() {

    }
