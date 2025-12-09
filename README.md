# gox

![Build Status](https://github.com/icza/gox/actions/workflows/go.yml/badge.svg)
[![GoDoc](https://pkg.go.dev/badge/github.com/icza/gox)](https://pkg.go.dev/github.com/icza/gox)
[![Go Report Card](https://goreportcard.com/badge/github.com/icza/gox)](https://goreportcard.com/report/github.com/icza/gox)

The `gox` module is a minimalistic, lightweigt extension to Go.
It contains constants, helpers and utilities which could have been part of Go itself.

## Module structure

- [`gox`](https://github.com/icza/gox/tree/main/gox): functions which could have been builtin, reasonable to "dot-import" this package
- [`fmtx`](https://github.com/icza/gox/tree/main/fmtx): formatting utilities,
complement to the standard `fmt` package.
- [`i18n`](https://github.com/icza/gox/tree/main/i18n): internationalization utilities.
- [`imagex/colorx`](https://github.com/icza/gox/tree/main/imagex/colorx): color utilities,
complement to the standard `image/color` package.
- [`logx/slogx`](https://github.com/icza/gox/tree/main/logx/slogx): structured logging utilities,
complement to the standard `log/slog` package.
- [`mathx`](https://github.com/icza/gox/tree/main/mathx): math utilities,
complement to the standard `math` package.
- [`mathx/randx`](https://github.com/icza/gox/tree/main/mathx/randx): random-related utilities,
complement to the standard `math/rand` package.
- [`netx/httpx`](https://github.com/icza/gox/tree/main/netx/httpx): HTTP utilities,
complement to the standard `net/http` package.
- [`osx`](https://github.com/icza/gox/tree/main/osx): operating system utilities,
complement to the standard `os` package.
- [`slicesx`](https://github.com/icza/gox/tree/main/slicesx): slice utilities,
complement to the standard `slices` package.
- [`stringsx`](https://github.com/icza/gox/tree/main/stringsx): string utilities,
complement to the standard `strings` package.
- [`timex`](https://github.com/icza/gox/tree/main/timex): time and duration related calculations and utilities,
complement to the standard `time` package.
