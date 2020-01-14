package mathx

import (
	"fmt"
	"math"
	"testing"
)

func TestAbsInt(t *testing.T) {
	cases := []struct {
		i, exp int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{math.MinInt32, -math.MinInt32},
		{math.MaxInt32, math.MaxInt32},
	}

	for i, c := range cases {
		if got := AbsInt(c.i); got != c.exp {
			t.Errorf("[%d] Expected: %d, got: %d", i, c.exp, c.i)
		}
	}
}

func ExampleRound() {
	fmt.Printf("%.4f\n", Round(0.363636, 0.001)) // 0.364
	fmt.Printf("%.4f\n", Round(0.363636, 0.01))  // 0.36
	fmt.Printf("%.4f\n", Round(0.363636, 0.1))   // 0.4
	fmt.Printf("%.4f\n", Round(0.363636, 0.05))  // 0.35
	fmt.Printf("%.4f\n", Round(3.2, 1))          // 3
	fmt.Printf("%.4f\n", Round(32, 5))           // 30
	fmt.Printf("%.4f\n", Round(33, 5))           // 35
	fmt.Printf("%.4f\n", Round(32, 10))          // 30

	fmt.Printf("%.4f\n", Round(-0.363636, 0.001)) // -0.364
	fmt.Printf("%.4f\n", Round(-0.363636, 0.01))  // -0.36
	fmt.Printf("%.4f\n", Round(-0.363636, 0.1))   // -0.4
	fmt.Printf("%.4f\n", Round(-0.363636, 0.05))  // -0.35
	fmt.Printf("%.4f\n", Round(-3.2, 1))          // -3
	fmt.Printf("%.4f\n", Round(-32, 5))           // -30
	fmt.Printf("%.4f\n", Round(-33, 5))           // -35
	fmt.Printf("%.4f\n", Round(-32, 10))          // -30

	// Output:
	// 0.3640
	// 0.3600
	// 0.4000
	// 0.3500
	// 3.0000
	// 30.0000
	// 35.0000
	// 30.0000
	// -0.3640
	// -0.3600
	// -0.4000
	// -0.3500
	// -3.0000
	// -30.0000
	// -35.0000
	// -30.0000
}

func TestNear(t *testing.T) {
	inf, neginf, nan := math.Inf(1), math.Inf(-1), math.NaN()
	cases := []struct {
		a, b, eps float64
		exp       bool
	}{
		{1.0, 1.0, 1e-6, true},
		{1.0, 1.001, 1e-6, false},
		{1.0, 1.001, 1e-2, true},

		// Corner cases
		{inf, 1.001, 1e-2, false},
		{neginf, 1.001, 1e-2, false},
		{inf, inf, 1e-2, true},
		{neginf, neginf, 1e-2, true},
		{inf, neginf, 1e-2, false},

		{1.0, 1.1, inf, true},
		{1.0, inf, inf, false},
		{inf, inf, inf, true},
		{neginf, neginf, inf, true},

		{1.0, nan, 1e10, false},
		{1.0, nan, inf, false},
		{nan, nan, 1e10, false},
		{nan, nan, inf, false},

		{1.0, 1.0, nan, true},
		{1.0, 1.001, nan, false},
		{inf, inf, nan, true},
		{neginf, neginf, nan, true},
	}

	for i, c := range cases {
		if got := Near(c.a, c.b, c.eps); c.exp != got {
			t.Errorf("[i=%d] Expected: %v, got: %v", i, c.exp, got)
		}
	}
}
