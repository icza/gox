# gox

[![GoDoc](https://pkg.go.dev/badge/github.com/icza/gox/gox)](https://pkg.go.dev/github.com/icza/gox/gox)

Package `gox` contains functions and types which could have been builtin, which
could have been part of Go itself.

Reasonable to "dot-import" the package, so identifiers will be directly available:

```golang
import (
	"fmt"
	"strconv"

	. "github.com/icza/gox/gox"
)

func main() {
	// Pass multiple return values to variadic functions:
	now := time.Date(2020, 3, 4, 0, 0, 0, 0, time.UTC)
	fmt.Printf("Year: %d, month: %d, day: %d\n",
		Wrap(now.Date())...)

	// Quick "handling" of error:
	n, err := strconv.Atoi("3")
	Pie(err)
	fmt.Println("Parsed:", n)

	// Output:
	// Year: 2020, month: 3, day: 4
	// Parsed: 3
}
```
