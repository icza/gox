package mathx

import (
	"fmt"
	"math"
	"testing"

	"github.com/icza/gox/gox"
	"github.com/icza/gox/osx"
)

func TestAbsInt(t *testing.T) {
	cases := []struct {
		name   string
		i, exp int
	}{
		{"zero", 0, 0},
		{"pos", 1, 1},
		{"neg", -1, 1},
		{"maxint32", math.MaxInt32, math.MaxInt32},
		// On 32-bit arch, -math.MinInt32 overflows and this test fails
		{"minint32", math.MinInt32, gox.If(osx.Arch32bit, math.MinInt32, -math.MinInt32)},
	}

	for _, c := range cases {
		if got := AbsInt(c.i); got != c.exp {
			t.Errorf("[%s] Expected: %d, got: %d", c.name, c.exp, c.i)
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
		name      string
		a, b, eps float64
		exp       bool
	}{
		{"normal-1", 1.0, 1.0, 1e-6, true},
		{"normal-1", 1.0, 1.001, 1e-6, false},
		{"normal-1", 1.0, 1.001, 1e-2, true},

		// Corner cases
		{"corner-case-inf", inf, 1.001, 1e-2, false},
		{"corner-case-neginf", neginf, 1.001, 1e-2, false},
		{"corner-case-inf-inf", inf, inf, 1e-2, true},
		{"corner-case-neginf-neginf", neginf, neginf, 1e-2, true},
		{"corner-case-inf-neginf", inf, neginf, 1e-2, false},

		{"corner-case-eps-inf", 1.0, 1.1, inf, true},
		{"corner-case-inf-eps-inf", 1.0, inf, inf, false},
		{"corner-case-all-inf", inf, inf, inf, true},
		{"corner-case-neginf-neginf-eps-inf", neginf, neginf, inf, true},

		{"corner-case-nan", 1.0, nan, 1e10, false},
		{"corner-case-nan-eps-inf", 1.0, nan, inf, false},
		{"corner-case-nan-nan", nan, nan, 1e10, false},
		{"corner-case-nan-nan-eps-inf", nan, nan, inf, false},

		{"corner-case-eps-nan", 1.0, 1.0, nan, true},
		{"corner-case-eps-nan-2", 1.0, 1.001, nan, false},
		{"corner-case-inf-inf-eps-nan", inf, inf, nan, true},
		{"corner-case-neginf-neginf-eps-nan", neginf, neginf, nan, true},
	}

	for _, c := range cases {
		if got := Near(c.a, c.b, c.eps); c.exp != got {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, got)
		}
	}
}
