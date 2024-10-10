// This file contains general utilities that do not warrant their own file.

package gox

// If returns vtrue if cond is true, vfalse otherwise.
//
// Useful to avoid an if statement when initializing variables, for example:
//
//	min := If(i > 0, i, 0)
func If[T any](cond bool, vtrue, vfalse T) T {
	if cond {
		return vtrue
	}
	return vfalse
}

// IfFunc returns the return value of ftrue if cond is true, the return value of ffalse otherwise.
//
// In contrast to [If], this can be used to deferred, on-demand evaluation of values depending on the condition.
func IfFunc[T any](cond bool, ftrue func() T, ffalse func() T) T {
	if cond {
		return ftrue()
	}
	return ffalse()
}

// Ptr returns a pointer to the passed value.
//
// Useful when you have a value and need a pointer, e.g.:
//
//	func f() string { return "foo" }
//
//	foo := struct{
//	    Bar *string
//	}{
//	    Bar: Ptr(f()),
//	}
func Ptr[T any](v T) *T {
	return &v
}

// Must takes 2 arguments, the second being an error.
// If err is not nil, Must panics. Else the first argument is returned.
//
// Useful when inputs to some function are provided in the source code,
// and you are sure they are valid (if not, it's OK to panic).
// For example:
//
//	t := Must(time.Parse("2006-01-02", "2022-04-20"))
func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

// First returns the first argument.
// Useful when you want to use the first result of a function call that has more than one return values
// (e.g. in a composite literal or in a condition).
//
// For example:
//
//	func f() (i, j, k int, s string, f float64) { return }
//
//	p := image.Point{
//	    X: First(f()),
//	}
func First[T any](first T, _ ...any) T {
	return first
}

// Second returns the second argument.
// Useful when you want to use the second result of a function call that has more than one return values
// (e.g. in a composite literal or in a condition).
//
// For example:
//
//	func f() (i, j, k int, s string, f float64) { return }
//
//	p := image.Point{
//	    X: Second(f()),
//	}
func Second[T any](_ any, second T, _ ...any) T {
	return second
}

// Third returns the third argument.
// Useful when you want to use the third result of a function call that has more than one return values
// (e.g. in a composite literal or in a condition).
//
// For example:
//
//	func f() (i, j, k int, s string, f float64) { return }
//
//	p := image.Point{
//	    X: Third(f()),
//	}
func Third[T any](_, _ any, third T, _ ...any) T {
	return third
}

// Coalesce returns the first non-zero value from listed arguments.
// Returns the zero value of the type parameter if no arguments are given or all are the zero value.
// Useful when you want to initialize a variable to the first non-zero value from a list of fallback values.
//
// For example:
//
//	hostVal := Coalesce(hostName, os.Getenv("HOST"), "localhost")
//
// Note: the same functionality has been added in Go 1.22 as cmp.Or()
func Coalesce[T comparable](values ...T) (v T) {
	var zero T
	for _, v = range values {
		if v != zero {
			return
		}
	}
	return
}

// Deref "safely" dereferences a pointer, returns the pointed value.
// If the pointer is nil, the (first) def is returned.
// If def is not specified, the zero value of T is returned.
func Deref[T any](p *T, def ...T) (result T) {
	if p != nil {
		return *p
	}
	if len(def) > 0 {
		return def[0]
	}
	return
}

// Pie is a "panic-if-error" utility: panics if the passed error is not nil.
// Should not be over-used, but may come handy to write code quickly.
func Pie(err error) {
	if err != nil {
		panic(err)
	}
}

// Wrap returns its arguments as a slice.
//
// General use of Wrap is to wrap function calls, so the return values of the
// "wrapped" function will be available as a slice. Which then can be passed
// to variadic functions that have other parameters too.
//
// Most notable example is fmt.Printf(). This code doesn't compile:
//
//	// Compile-time error!
//	fmt.Printf("Year: %d, month: %d, day: %d", time.Now().Date())
//
// But with the help of this Wrap:
//
//	// This is OK!
//	fmt.Printf("Year: %d, month: %d, day: %d",
//	    Wrap(time.Now().Date())...)
//
// For details, see https://stackoverflow.com/a/52654950/1705598
func Wrap(vs ...interface{}) []interface{} {
	return vs
}
