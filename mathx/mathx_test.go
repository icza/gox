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
	fmt.Printf("%.4f\n", Round(0.363636, 0.05)) // 0.35
	fmt.Printf("%.4f\n", Round(3.232, 0.05))    // 3.25
	fmt.Printf("%.4f\n", Round(0.4888, 0.05))   // 0.5

	fmt.Printf("%.4f\n", Round(-0.363636, 0.05)) // -0.35
	fmt.Printf("%.4f\n", Round(-3.232, 0.05))    // -3.25
	fmt.Printf("%.4f\n", Round(-0.4888, 0.05))   // -0.5

	// Output:
	// 0.3500
	// 3.2500
	// 0.5000
	// -0.3500
	// -3.2500
	// -0.5000
}
