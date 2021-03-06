package randx

import "math/rand"

var randFloat64 = rand.Float64 // Mockable rand.Float64 function

// RandomWeight chooses an index randomly using the listed (non-negative)
// probabilities as weights. Weights must add up to 1.
//
// The default Rand of math/rand package is used for random data.
//
// Implementation guarantees to return an integer in the range of [0..len(weights)).
// If weights don't add up to 1, -1 may be returned.
func RandomWeight(weights ...float64) int {
	r := randFloat64()

	sum := 0.0
	for i, w := range weights {
		sum += w
		if r < sum {
			return i
		}
	}

	// We can only end up here if sum of weights is less than 1.
	return -1
}

var randIntn = rand.Intn // Mockable rand.Intn function

// RandomIntWeight chooses an index randomly using the listed non-negative
// relative weights.
//
// The default Rand of math/rand package is used for random data.
//
// Implementation guarantees to return an integer in the range of [0..len(weights)).
// If no weight is provided or they add up to 0, -1 is returned.
// Behavior for negative weights is undefined.
func RandomIntWeight(weights ...int) int {
	sum := 0
	for _, w := range weights {
		sum += w
	}

	// Return early if sum is not positive, Intn() would panic.
	if sum <= 0 {
		return -1
	}

	r, s := randIntn(sum), 0
	for i, w := range weights {
		s += w
		if r < s {
			return i
		}
	}

	// We could only end up here if weights add up to a non-positive number,
	// but that is checked earlier (and returned early).
	panic("unreachable")
}
