package fmtx

import (
	"fmt"
	"testing"
)

// ExampleFormatInt shows how to use the FormatInt() function.
func ExampleFormatInt() {
	fmt.Println("groupSize:                  3               4")
	fmt.Println("---------------------------------------------")
	for _, v := range []int64{12, 123, 1234, 123456789} {
		for sign := int64(1); sign >= -1; sign -= 2 {
			x := v * sign
			fmt.Printf("n: %10s", fmt.Sprint(x))
			for groupingSize := 3; groupingSize <= 4; groupingSize++ {
				fmt.Printf(" =%14s", FormatInt(x, groupingSize, ','))
			}
			fmt.Println()
		}
	}

	// Output:
	// groupSize:                  3               4
	// ---------------------------------------------
	// n:         12 =            12 =            12
	// n:        -12 =           -12 =           -12
	// n:        123 =           123 =           123
	// n:       -123 =          -123 =          -123
	// n:       1234 =         1,234 =          1234
	// n:      -1234 =        -1,234 =         -1234
	// n:  123456789 =   123,456,789 =   1,2345,6789
	// n: -123456789 =  -123,456,789 =  -1,2345,6789
}

func TestFormatInt(t *testing.T) {
	cases := []struct {
		n         int64
		groupSize int
		grouping  byte
		exp       string
	}{
		{1, 3, ',', "1"},
		{-1, 3, ',', "-1"},
		{1234, 3, ',', "1,234"},
		{-1234, 3, ',', "-1,234"},
		{1234, -1, ',', "1,234"},
		{-1234, 0, ',', "-1,234"},
		{1234, 2, '.', "12.34"},
		{-1234, 2, '.', "-12.34"},
		{12345, 4, ';', "1;2345"},
		{-12345, 4, ';', "-1;2345"},
	}

	for i, c := range cases {
		if got := FormatInt(c.n, c.groupSize, c.grouping); got != c.exp {
			t.Errorf("[%d] Expected: %v, got: %v", i, c.exp, got)
		}
	}
}
