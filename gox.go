// This file contains general utilities that do not warrant their own file.

package gox

// Wrap returns its arguments as a slice.
//
// General use of Wrap is to wrap function calls, so the return values of the
// "wrapped" function will be available as a slice. Which then can be passed
// to variadic functions that have other parameters too.
//
// Most notable example is fmt.Printf(). This code doesn't compile:
//   // Compile-time error!
//   fmt.Printf("Year: %d, month: %d, day: %d", time.Now().Date())
//
// But with the help of this Wrap:
//   // This is OK!
//   fmt.Printf("Year: %d, month: %d, day: %d",
//       Wrap(time.Now().Date())...)
//
// For details, see https://stackoverflow.com/a/52654950/1705598
func Wrap(vs ...interface{}) []interface{} {
	return vs
}
