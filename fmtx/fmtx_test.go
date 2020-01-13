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

func TestFormatSize(t *testing.T) {
	cases := []struct {
		size           int64
		unit           SizeUnit
		fractionDigits int
		exp            string
	}{
		{0, SizeUnitByte, 0, "0 bytes"},
		{100, SizeUnitByte, 0, "100 bytes"},
		{999, SizeUnitByte, 0, "999 bytes"},
		{1200, SizeUnitByte, 0, "1200 bytes"},

		{0, SizeUnitKB, 0, "0 KB"},
		{0, SizeUnitKB, 1, "0.0 KB"},
		{100, SizeUnitKB, 0, "0 KB"},
		{600, SizeUnitKB, 0, "1 KB"},
		{600, SizeUnitKB, 1, "0.6 KB"},
		{600, SizeUnitKB, 2, "0.59 KB"},
		{2048, SizeUnitKB, 2, "2.00 KB"},

		{0, SizeUnitMB, 0, "0 MB"},
		{0, SizeUnitMB, 1, "0.0 MB"},
		{100 << 10, SizeUnitMB, 0, "0 MB"},
		{600 << 10, SizeUnitMB, 0, "1 MB"},
		{600 << 10, SizeUnitMB, 1, "0.6 MB"},
		{600 << 10, SizeUnitMB, 2, "0.59 MB"},
		{2048 << 10, SizeUnitMB, 2, "2.00 MB"},

		{0, SizeUnitGB, 0, "0 GB"},
		{0, SizeUnitGB, 1, "0.0 GB"},
		{100 << 20, SizeUnitGB, 0, "0 GB"},
		{600 << 20, SizeUnitGB, 0, "1 GB"},
		{600 << 20, SizeUnitGB, 1, "0.6 GB"},
		{600 << 20, SizeUnitGB, 2, "0.59 GB"},
		{2048 << 20, SizeUnitGB, 2, "2.00 GB"},

		{0, SizeUnitTB, 0, "0 TB"},
		{0, SizeUnitTB, 1, "0.0 TB"},
		{100 << 30, SizeUnitTB, 0, "0 TB"},
		{600 << 30, SizeUnitTB, 0, "1 TB"},
		{600 << 30, SizeUnitTB, 1, "0.6 TB"},
		{600 << 30, SizeUnitTB, 2, "0.59 TB"},
		{2048 << 30, SizeUnitTB, 2, "2.00 TB"},

		{0, SizeUnitPB, 0, "0 PB"},
		{0, SizeUnitPB, 1, "0.0 PB"},
		{100 << 40, SizeUnitPB, 0, "0 PB"},
		{600 << 40, SizeUnitPB, 0, "1 PB"},
		{600 << 40, SizeUnitPB, 1, "0.6 PB"},
		{600 << 40, SizeUnitPB, 2, "0.59 PB"},
		{2048 << 40, SizeUnitPB, 2, "2.00 PB"},

		{0, SizeUnitEB, 0, "0 EB"},
		{0, SizeUnitEB, 1, "0.0 EB"},
		{100 << 50, SizeUnitEB, 0, "0 EB"},
		{600 << 50, SizeUnitEB, 0, "1 EB"},
		{600 << 50, SizeUnitEB, 1, "0.6 EB"},
		{600 << 50, SizeUnitEB, 2, "0.59 EB"},
		{2048 << 50, SizeUnitEB, 2, "2.00 EB"},

		{0, SizeUnitAuto, 0, "0 bytes"},
		{100, SizeUnitAuto, 0, "100 bytes"},
		{100, SizeUnitAuto, 1, "100 bytes"},
		{999, SizeUnitAuto, 0, "999 bytes"},
		{1000, SizeUnitAuto, 0, "1 KB"},
		{1000, SizeUnitAuto, 1, "1.0 KB"},
		{1000, SizeUnitAuto, 2, "0.98 KB"},
		{1024, SizeUnitAuto, 2, "1.00 KB"},
		{999 << 10, SizeUnitAuto, 2, "999.00 KB"},
		{999<<10 + 500, SizeUnitAuto, 2, "999.49 KB"},
		{1000 << 10, SizeUnitAuto, 2, "0.98 MB"},
		{2 << 20, SizeUnitAuto, 1, "2.0 MB"},
		{2 << 30, SizeUnitAuto, 1, "2.0 GB"},
		{2 << 40, SizeUnitAuto, 1, "2.0 TB"},
		{2 << 50, SizeUnitAuto, 1, "2.0 PB"},
		{2 << 60, SizeUnitAuto, 1, "2.0 EB"},
	}

	for i, c := range cases {
		if got := FormatSize(c.size, c.unit, c.fractionDigits); got != c.exp {
			t.Errorf("[%d] Expected: %v, got: %v", i, c.exp, got)
		}
	}
}

// ExampleFormatInt shows how to use the CondSprintf() function.
func ExampleCondSprintf() {
	fmt.Println(CondSprintf("Foo%s", "bar", "baz"))
	fmt.Println(CondSprintf("%d + %d = %d", 1, 2, 3, "extra", 4))

	// Output:
	// Foobar
	// 1 + 2 = 3
}

func TestCondSprintf(t *testing.T) {
	cases := []struct {
		format string
		args   []interface{}
		exp    string
	}{
		{"", nil, ""},
		{"Foo", nil, "Foo"},
		{"%s", []interface{}{"bar"}, "bar"},
		{"Foo%s", []interface{}{"bar"}, "Foobar"},
		{"Foo%s", []interface{}{"bar", "extra"}, "Foobar"},
		{"Foo", []interface{}{"bar"}, "Foo"},
		{"foo%s%[3]s", []interface{}{"bar", "baz", "1"}, "foobar1"},
		{"foo%[2]*d=%[1]d+%d", []interface{}{2, 3, 5, 9, "x"}, "foo  5=2+3"},
	}

	for i, c := range cases {
		if got := CondSprintf(c.format, c.args...); got != c.exp {
			t.Errorf("[%d] Expected: %v, got: %v", i, c.exp, got)
		}
	}
}
