package gox_test

import (
	"fmt"
	"time"

	"github.com/icza/gox/gox"
	"github.com/icza/gox/slicesx"
)

// This example demonstrates how to use OpCache to cache the results
// of an existing function.
func ExampleOpCache() {
	type Point struct {
		X, Y    int
		Counter int // To track invocations
	}

	counter := 0
	// Existing GetPoint() function we want to add caching for:
	GetPoint := func(x, y int) (*Point, error) {
		counter++
		return &Point{X: x, Y: y, Counter: counter}, nil
	}

	var getPointCache = gox.NewOpCache[gox.Struct2[int, int], *Point](gox.OpCacheConfig{ResultExpiration: 100 * time.Millisecond})

	// Function to use which utilizes getPointCache (has identical signature to that of GetPoint):
	GetPointFast := func(x, y int) (*Point, error) {
		return getPointCache.Get(
			gox.Struct2Of(x, y), // Key constructed from all arguments
			func() (*Point, error) { return GetPoint(x, y) },
		)
	}

	p, err := GetPointFast(1, 2) // This will call GetPoint()
	fmt.Printf("%+v %v\n", p, err)
	p, err = GetPointFast(1, 2) // This will come from the cache
	fmt.Printf("%+v %v\n", p, err)

	time.Sleep(110 * time.Millisecond)
	p, err = GetPointFast(1, 2) // Cache expired, will call GetPoint() again
	fmt.Printf("%+v %v\n", p, err)

	// Output:
	// &{X:1 Y:2 Counter:1} <nil>
	// &{X:1 Y:2 Counter:1} <nil>
	// &{X:1 Y:2 Counter:2} <nil>
}

// This example demonstrates how to use OpCache to cache the results
// of an existing function that has multiple result types (besides the error).
func ExampleOpCache_multi_return() {
	type Point struct {
		X, Y    int
		Counter int // To track invocations
	}

	counter := 0
	// Existing GetPoint() function we want to add caching for:
	GetPoint := func(x, y int) (*Point, int, error) {
		counter++
		return &Point{X: x, Y: 2 * x, Counter: counter}, counter * 10, fmt.Errorf("test_error_%d", counter)
	}

	var getPointCache = gox.NewOpCache[gox.Struct2[int, int], gox.Struct2[*Point, int]](gox.OpCacheConfig{ResultExpiration: 100 * time.Millisecond})

	// Function to use which utilizes getPointCache (has identical signature to that of GetPoint):
	GetPointFast := func(x, y int) (*Point, int, error) {
		mr, err := getPointCache.Get(
			gox.Struct2Of(x, y), // Key constructed from all arguments
			func() (gox.Struct2[*Point, int], error) {
				p, n, err := GetPoint(x, y)
				return gox.Struct2Of(p, n), err // packing multiple results
			},
		)
		return mr.V1, mr.V2, err // Unpacking multiple results
	}

	p, n, err := GetPointFast(1, 2) // This will call GetPoint()
	fmt.Printf("%+v %d %v\n", p, n, err)
	p, n, err = GetPointFast(1, 2) // This will come from the cache
	fmt.Printf("%+v %d %v\n", p, n, err)

	time.Sleep(110 * time.Millisecond)
	p, n, err = GetPointFast(1, 2) // Cache expired, will call GetPoint() again
	fmt.Printf("%+v %d %v\n", p, n, err)

	// Output:
	// &{X:1 Y:2 Counter:1} 10 test_error_1
	// &{X:1 Y:2 Counter:1} 10 test_error_1
	// &{X:1 Y:2 Counter:2} 20 test_error_2
}

// This example demonstrates how to use OpCache.MultiGet().
func ExampleOpCache_MultiGet() {
	type CalcResult struct {
		Y       int
		Counter int // To track invocations
	}

	counter := 0
	// Existing Calc() function we want to add caching for:
	Calc := func(x int) (CalcResult, error) {
		counter++
		return CalcResult{Y: 2 * x, Counter: counter}, nil
	}
	// Existing MultiCalc() that can do the same for multiple inputs:
	MultiCalc := func(xs []int) (cs []CalcResult, errs []error) {
		for _, x := range xs {
			counter++
			cs = append(cs, CalcResult{Y: 2 * x, Counter: counter})
			errs = append(errs, nil)
		}
		return
	}

	var calcCache = gox.NewOpCache[int, CalcResult](gox.OpCacheConfig{
		ResultExpiration:      100 * time.Millisecond,
		ResultGraceExpiration: 50 * time.Millisecond,
	})

	// Function to use which utilizes calcCache (has identical signature to that of Calc):
	CalcFast := func(x int) (CalcResult, error) {
		return calcCache.Get(
			x, // Key constructed from all arguments
			func() (CalcResult, error) { return Calc(x) },
		)
	}

	// Function to use which utilizes calcCache (has identical signature to that of MultiCalc):
	MultiCalcFast := func(xs []int) ([]CalcResult, []error) {
		return calcCache.MultiGet(
			xs,
			func(keyIndices []int) ([]CalcResult, []error) {
				return MultiCalc(slicesx.SelectByIndices(xs, keyIndices))
			},
		)
	}

	c, err := CalcFast(1) // This will call Calc()
	fmt.Printf("%+v %v\n", c, err)

	cs, errs := MultiCalcFast([]int{1, 2, 3}) // First from cache, other 2 will be passed to MultiCalc()
	fmt.Printf("%+v %v\n", cs, errs)

	time.Sleep(110 * time.Millisecond)

	// First 2 from cache, third will be passed to MultiCalc()
	// Also background MultiCalc() will be called for first 2.
	cs, errs = MultiCalcFast([]int{1, 2, 4})
	fmt.Printf("%+v %v\n", cs, errs)

	time.Sleep(10 * time.Millisecond)

	// All from cache, first 2 with updated counter from the background refresh
	cs, errs = MultiCalcFast([]int{1, 2, 4})
	fmt.Printf("%+v %v\n", cs, errs)

	// Output:
	// {Y:2 Counter:1} <nil>
	// [{Y:2 Counter:1} {Y:4 Counter:2} {Y:6 Counter:3}] [<nil> <nil> <nil>]
	// [{Y:2 Counter:1} {Y:4 Counter:2} {Y:8 Counter:4}] [<nil> <nil> <nil>]
	// [{Y:2 Counter:5} {Y:4 Counter:6} {Y:8 Counter:4}] [<nil> <nil> <nil>]
}

// This example demonstrates how to use OpCache.MultiGet() when the operaiton has multiple input arguments.
func ExampleOpCache_MultiGet_multi_inputargs() {
	type CalcResult struct {
		Y       int
		Counter int // To track invocations
	}

	counter := 0
	// Existing Calc() function we want to add caching for:
	Calc := func(x, y int) (CalcResult, error) {
		counter++
		return CalcResult{Y: 2*x + y, Counter: counter}, nil
	}
	// Existing MultiCalc() that can do the same for multiple inputs:
	MultiCalc := func(xs, ys []int) (cs []CalcResult, errs []error) {
		for i, x := range xs {
			counter++
			cs = append(cs, CalcResult{Y: 2*x + ys[i], Counter: counter})
			errs = append(errs, nil)
		}
		return
	}

	var calcCache = gox.NewOpCache[gox.Struct2[int, int], CalcResult](gox.OpCacheConfig{
		ResultExpiration:      100 * time.Millisecond,
		ResultGraceExpiration: 50 * time.Millisecond,
	})

	// Function to use which utilizes calcCache (has identical signature to that of Calc):
	CalcFast := func(x, y int) (CalcResult, error) {
		return calcCache.Get(
			gox.Struct2Of(x, y), // Key constructed from all arguments
			func() (CalcResult, error) { return Calc(x, y) },
		)
	}

	// Function to use which utilizes calcCache (has identical signature to that of MultiCalc):
	MultiCalcFast := func(xs, ys []int) ([]CalcResult, []error) {
		keys := make([]gox.Struct2[int, int], len(xs))
		for i, x := range xs {
			keys[i] = gox.Struct2Of(x, ys[i]) // Key constructed from all arguments
		}
		return calcCache.MultiGet(
			keys,
			func(keyIndices []int) ([]CalcResult, []error) {
				return MultiCalc(slicesx.SelectByIndices(xs, keyIndices), slicesx.SelectByIndices(ys, keyIndices))
			},
		)
	}

	c, err := CalcFast(1, 100) // This will call Calc()
	fmt.Printf("%+v %v\n", c, err)

	cs, errs := MultiCalcFast([]int{1, 2, 3}, []int{100, 200, 300}) // First from cache, other 2 will be passed to MultiCalc()
	fmt.Printf("%+v %v\n", cs, errs)

	time.Sleep(110 * time.Millisecond)

	// First 2 from cache, third will be passed to MultiCalc()
	// Also background MultiCalc() will be called for first 2.
	cs, errs = MultiCalcFast([]int{1, 2, 4}, []int{100, 200, 400})
	fmt.Printf("%+v %v\n", cs, errs)

	time.Sleep(10 * time.Millisecond)

	// All from cache, first 2 with updated counter from the background refresh
	cs, errs = MultiCalcFast([]int{1, 2, 4}, []int{100, 200, 400})
	fmt.Printf("%+v %v\n", cs, errs)

	// Output:
	// {Y:102 Counter:1} <nil>
	// [{Y:102 Counter:1} {Y:204 Counter:2} {Y:306 Counter:3}] [<nil> <nil> <nil>]
	// [{Y:102 Counter:1} {Y:204 Counter:2} {Y:408 Counter:4}] [<nil> <nil> <nil>]
	// [{Y:102 Counter:5} {Y:204 Counter:6} {Y:408 Counter:4}] [<nil> <nil> <nil>]
}
