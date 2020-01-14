package mathx

import "math"

// AbsInt returns the absolute value of i.
func AbsInt(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

// Round returns x rounded to the given unit.
//
// For details, see https://stackoverflow.com/a/39544897/1705598
func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}
