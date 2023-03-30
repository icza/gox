# gox

![Build Status](https://github.com/icza/gox/actions/workflows/go.yml/badge.svg)
[![Go Reference](https://pkg.go.dev/badge/github.com/icza/gox.svg)](https://pkg.go.dev/github.com/icza/gox)
[![Go Report Card](https://goreportcard.com/badge/github.com/icza/gox)](https://goreportcard.com/report/github.com/icza/gox)
[![codecov](https://codecov.io/gh/icza/gox/branch/master/graph/badge.svg)](https://codecov.io/gh/icza/gox)

_This module is in beta phase._

The `gox` module is a minimalistic, lightweigt extension to Go.
It contains constants, helpers and utilities which could have been part of Go itself.

## Module structure

- [`gox`](https://github.com/icza/gox/tree/master/gox): functions and types which could have been builtin, reasonable to "dot-import" this package
- [`fmtx`](https://github.com/icza/gox/tree/master/fmtx): formatting utilities,
complement to the standard `fmt` package.
- [`i18n`](https://github.com/icza/gox/tree/master/i18n): internationalization utilities.
- [`imagex/colorx`](https://github.com/icza/gox/tree/master/imagex/colorx): color utilities,
complement to the standard `image/color` package.
- [`mathx`](https://github.com/icza/gox/tree/master/mathx): math utilities,
complement to the standard `math` package.
- [`mathx/randx`](https://github.com/icza/gox/tree/master/mathx/randx): random-related utilities,
complement to the standard `math/rand` package.
- [`netx/httpx`](https://github.com/icza/gox/tree/master/netx/httpx): HTTP utilities,
complement to the standard `net/http` package.
- [`osx`](https://github.com/icza/gox/tree/master/osx): operating system utilities,
complement to the standard `os` package.
- [`stringsx`](https://github.com/icza/gox/tree/master/stringsx): string utilities,
complement to the standard `strings` package.
- [`timex`](https://github.com/icza/gox/tree/master/timex): time and duration related calculations and utilities,
complement to the standard `time` package.
