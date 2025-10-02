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
		name      string
		n         int64
		groupSize int
		grouping  byte
		exp       string
	}{
		{"1group", 1, 3, ',', "1"},
		{"1group-neg", -1, 3, ',', "-1"},
		{"2groups", 1234, 3, ',', "1,234"},
		{"2groups-neg", -1234, 3, ',', "-1,234"},
		{"2groups-neg-group-size", 1234, -1, ',', "1,234"},
		{"2groups-neg-zero-group-size", -1234, 0, ',', "-1,234"},
		{"2groups-dot-separator", 1234, 2, '.', "12.34"},
		{"2groups-neg-dot-separator", -1234, 2, '.', "-12.34"},
		{"2groups-semi-separator", 12345, 4, ';', "1;2345"},
		{"2groups-neg-semi-separator", -12345, 4, ';', "-1;2345"},
	}

	for _, c := range cases {
		if got := FormatInt(c.n, c.groupSize, c.grouping); got != c.exp {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, got)
		}
	}
}

func TestFormatSize(t *testing.T) {
	cases := []struct {
		name           string
		size           int64
		unit           SizeUnit
		fractionDigits int
		exp            string
	}{
		{"bytes-0", 0, SizeUnitByte, 0, "0 bytes"},
		{"bytes-100", 100, SizeUnitByte, 0, "100 bytes"},
		{"bytes-999", 999, SizeUnitByte, 0, "999 bytes"},
		{"bytes-1200", 1200, SizeUnitByte, 0, "1200 bytes"},
		{"bytes-1200-frac", 1200, SizeUnitByte, 2, "1200 bytes"},

		{"kb-0", 0, SizeUnitKB, 0, "0 KB"},
		{"kb-0-frac", 0, SizeUnitKB, 1, "0.0 KB"},
		{"kb-100", 100, SizeUnitKB, 0, "0 KB"},
		{"kb-600", 600, SizeUnitKB, 0, "1 KB"},
		{"kb-600-frac", 600, SizeUnitKB, 1, "0.6 KB"},
		{"kb-0-frac2", 600, SizeUnitKB, 2, "0.59 KB"},
		{"kb-2048-frac2", 2048, SizeUnitKB, 2, "2.00 KB"},

		{"mb-0", 0, SizeUnitMB, 0, "0 MB"},
		{"mb-0-frac", 0, SizeUnitMB, 1, "0.0 MB"},
		{"mb-100kb", 100 << 10, SizeUnitMB, 0, "0 MB"},
		{"mb-600kb", 600 << 10, SizeUnitMB, 0, "1 MB"},
		{"mb-600kb-frac", 600 << 10, SizeUnitMB, 1, "0.6 MB"},
		{"mb-600kb-frac2", 600 << 10, SizeUnitMB, 2, "0.59 MB"},
		{"mb-2mb-frac2", 2048 << 10, SizeUnitMB, 2, "2.00 MB"},

		{"gb-0", 0, SizeUnitGB, 0, "0 GB"},
		{"gb-0-frac", 0, SizeUnitGB, 1, "0.0 GB"},
		{"gb-100mb", 100 << 20, SizeUnitGB, 0, "0 GB"},
		{"gb-600mb", 600 << 20, SizeUnitGB, 0, "1 GB"},
		{"gb-600mb-frac", 600 << 20, SizeUnitGB, 1, "0.6 GB"},
		{"gb-600mb-frac2", 600 << 20, SizeUnitGB, 2, "0.59 GB"},
		{"gb-2gb-frac2", 2048 << 20, SizeUnitGB, 2, "2.00 GB"},

		{"tb-0", 0, SizeUnitTB, 0, "0 TB"},
		{"tb-0-frac", 0, SizeUnitTB, 1, "0.0 TB"},
		{"tb-100mb", 100 << 30, SizeUnitTB, 0, "0 TB"},
		{"tb-600mb", 600 << 30, SizeUnitTB, 0, "1 TB"},
		{"tb-600mb-frac", 600 << 30, SizeUnitTB, 1, "0.6 TB"},
		{"tb-600mb-frac2", 600 << 30, SizeUnitTB, 2, "0.59 TB"},
		{"tb-2tb-frac2", 2048 << 30, SizeUnitTB, 2, "2.00 TB"},

		{"pb-0", 0, SizeUnitPB, 0, "0 PB"},
		{"pb-0-frac", 0, SizeUnitPB, 1, "0.0 PB"},
		{"pb-100gb", 100 << 40, SizeUnitPB, 0, "0 PB"},
		{"pb-600gb", 600 << 40, SizeUnitPB, 0, "1 PB"},
		{"pb-600gb-frac", 600 << 40, SizeUnitPB, 1, "0.6 PB"},
		{"pb-600gb-frac2", 600 << 40, SizeUnitPB, 2, "0.59 PB"},
		{"pb-2pmb-frac2", 2048 << 40, SizeUnitPB, 2, "2.00 PB"},

		{"eb-0", 0, SizeUnitEB, 0, "0 EB"},
		{"eb-0-frac", 0, SizeUnitEB, 1, "0.0 EB"},
		{"eb-100pb", 100 << 50, SizeUnitEB, 0, "0 EB"},
		{"eb-600pb", 600 << 50, SizeUnitEB, 0, "1 EB"},
		{"eb-600pb-frac", 600 << 50, SizeUnitEB, 1, "0.6 EB"},
		{"eb-600pb-frac2", 600 << 50, SizeUnitEB, 2, "0.59 EB"},
		{"eb-2eb-frac2", 2048 << 50, SizeUnitEB, 2, "2.00 EB"},

		{"auto-0", 0, SizeUnitAuto, 0, "0 bytes"},
		{"auto-100", 100, SizeUnitAuto, 0, "100 bytes"},
		{"auto-100-frac", 100, SizeUnitAuto, 1, "100 bytes"},
		{"auto-999", 999, SizeUnitAuto, 0, "999 bytes"},
		{"auto-1000", 1000, SizeUnitAuto, 0, "1 KB"},
		{"auto-1000-frac", 1000, SizeUnitAuto, 1, "1.0 KB"},
		{"auto-1000-frac2", 1000, SizeUnitAuto, 2, "0.98 KB"},
		{"auto-1024-frac2", 1024, SizeUnitAuto, 2, "1.00 KB"},
		{"auto-999kb-frac2", 999 << 10, SizeUnitAuto, 2, "999.00 KB"},
		{"auto-999kb500-frac2", 999<<10 + 500, SizeUnitAuto, 2, "999.49 KB"},
		{"auto-1000kb-frac2", 1000 << 10, SizeUnitAuto, 2, "0.98 MB"},
		{"auto-2mb-frac", 2 << 20, SizeUnitAuto, 1, "2.0 MB"},
		{"auto-2gb-frac", 2 << 30, SizeUnitAuto, 1, "2.0 GB"},
		{"auto-2tb-frac", 2 << 40, SizeUnitAuto, 1, "2.0 TB"},
		{"auto-2pb-frac", 2 << 50, SizeUnitAuto, 1, "2.0 PB"},
		{"auto-2eb-frac", 2 << 60, SizeUnitAuto, 1, "2.0 EB"},
	}

	for _, c := range cases {
		if got := FormatSize(c.size, c.unit, c.fractionDigits); got != c.exp {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, got)
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
		name   string
		format string
		args   []any
		exp    string
	}{
		{"empty", "", nil, ""},
		{"no-args", "Foo", nil, "Foo"},
		{"1arg", "%s", []any{"bar"}, "bar"},
		{"1arg-2", "Foo%s", []any{"bar"}, "Foobar"},
		{"1-extra", "Foo%s", []any{"bar", "extra"}, "Foobar"},
		{"no-verb-1-extra", "Foo", []any{"bar"}, "Foo"},
		{"explicit-arg-indices", "foo%s%[3]s", []any{"bar", "baz", "1"}, "foobar1"},
		{"explicit-arg-indices-star-1-extra", "foo%[2]*d=%[1]d+%d", []any{2, 3, 5, 9, "x"}, "foo  5=2+3"},
	}

	for _, c := range cases {
		if got := CondSprintf(c.format, c.args...); got != c.exp {
			t.Errorf("[%s] Expected: %v, got: %v", c.name, c.exp, got)
		}
	}
}
