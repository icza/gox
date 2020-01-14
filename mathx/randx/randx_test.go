package randx

import (
	"math/rand"
	"testing"
)

func TestRandomWeights(t *testing.T) {
	// Save-restore randFloat64 func
	oldRandFloat64 := randFloat64
	defer func() {
		randFloat64 = oldRandFloat64
	}()

	cases := []struct {
		name    string
		randf   func() float64
		weights []float64
		exp     int
	}{
		{"no-weights", rand.Float64, nil, -1},
		{"one-weight", rand.Float64, []float64{1}, 0},
		{"not-1-sum", rand.Float64, []float64{0}, -1},
		{"normal-1", func() float64 { return 0.2 }, []float64{0.5, 0.5}, 0},
		{"normal-2", func() float64 { return 0.6 }, []float64{0.5, 0.5}, 1},
		{"normal-3", func() float64 { return 0.3 }, []float64{0.1, 0.1, 0.7, 0.1}, 2},
		{"normal-4", func() float64 { return 0.95 }, []float64{0.1, 0.1, 0.7, 0.1}, 3},
	}

	for _, c := range cases {
		randFloat64 = c.randf
		if got := RandomWeight(c.weights...); got != c.exp {
			t.Errorf("[%s] Expected %d, got: %d", c.name, c.exp, got)
		}
	}
}

func TestRandomIntWeights(t *testing.T) {
	// Save-restore randIntn func
	oldRandIntn := randIntn
	defer func() {
		randIntn = oldRandIntn
	}()

	cases := []struct {
		name    string
		randf   func(int) int
		weights []int
		exp     int
	}{
		{"no-weights", rand.Intn, nil, -1},
		{"one-weight", rand.Intn, []int{1}, 0},
		{"not-1-sum", rand.Intn, []int{0}, -1},
		{"normal-1", func(int) int { return 0 }, []int{1, 1}, 0},
		{"normal-2", func(int) int { return 5 }, []int{10, 10}, 0},
		{"normal-2", func(int) int { return 15 }, []int{10, 10}, 1},
		{"normal-3", func(int) int { return 3 }, []int{1, 1, 7, 1}, 2},
		{"normal-4", func(int) int { return 9 }, []int{1, 1, 7, 1}, 3},
	}

	for _, c := range cases {
		randIntn = c.randf
		if got := RandomIntWeight(c.weights...); got != c.exp {
			t.Errorf("[%s] Expected %d, got: %d", c.name, c.exp, got)
		}
	}
}
