package timex

import (
	"fmt"
	"testing"
	"time"
)

// ExampleDiff shows how to use the Diff() function.
func ExampleDiff() {
	// Your birthday: let's say it's January 2nd, 1980, 3:30 AM
	birthday := time.Date(1980, 1, 2, 3, 30, 0, 0, time.UTC)
	now := time.Date(2020, 3, 9, 13, 57, 46, 0, time.UTC)

	year, month, day, hour, min, sec := Diff(birthday, now)

	fmt.Printf(
		"You are %d years, %d months, %d days, %d hours, %d mins and %d seconds old.",
		year, month, day, hour, min, sec,
	)

	// Output:
	// You are 40 years, 2 months, 7 days, 10 hours, 27 mins and 46 seconds old.
}

func TestDiff(t *testing.T) {
	cases := []struct {
		name                             string    // Name of the test case
		a, b                             time.Time // inputs
		year, month, day, hour, min, sec int       // expected output
	}{
		{
			"same",
			time.Date(1980, 1, 2, 3, 30, 0, 0, time.UTC),
			time.Date(1980, 1, 2, 3, 30, 0, 0, time.UTC),
			0, 0, 0, 0, 0, 0,
		},
		{
			"normal",
			time.Date(1980, 1, 2, 3, 30, 0, 0, time.UTC),
			time.Date(1985, 2, 4, 6, 34, 5, 0, time.UTC),
			5, 1, 2, 3, 4, 5,
		},
		{
			"normal-swapped",
			time.Date(1985, 2, 4, 6, 34, 5, 0, time.UTC),
			time.Date(1980, 1, 2, 3, 30, 0, 0, time.UTC),
			5, 1, 2, 3, 4, 5,
		},
		{
			"normal-diff-zone",
			time.Date(1980, 1, 2, 3, 30, 0, 0, time.UTC),
			time.Date(1985, 2, 4, 6, 34, 5, 0, time.FixedZone("", 3600)),
			5, 1, 2, 2, 4, 5,
		},
		{
			"normalize-negative-sec-min-hour-day-month",
			time.Date(1980, 1, 1, 10, 10, 10, 0, time.UTC),
			time.Date(1981, 1, 1, 0, 0, 0, 0, time.UTC),
			0, 11, 30, 13, 49, 50,
		},
	}

	for _, c := range cases {
		year, month, day, hour, min, sec := Diff(c.a, c.b)
		if year != c.year || month != c.month || day != c.day || hour != c.hour || min != c.min || sec != c.sec {
			t.Errorf("[%s] Expected: %d years, %d months, %d days, %d hours, %d mins and %d seconds, got: %d years, %d months, %d days, %d hours, %d mins and %d seconds",
				c.name,
				c.year, c.month, c.day, c.hour, c.min, c.sec,
				year, month, day, hour, min, sec,
			)
		}
	}
}

// ExampleWeekStart shows how to use the WeekStart() function.
func ExampleWeekStart() {
	inputs := []struct{ year, week int }{
		{2018, -1},
		{2018, 0},
		{2018, 1},
		{2018, 2},
		{2019, 1},
		{2019, 2},
		{2019, 53},
		{2019, 54},
	}
	for _, c := range inputs {
		fmt.Printf("Week (%d,%2d) starts on: %v\n",
			c.year, c.week, WeekStart(c.year, c.week))
	}

	// Output:
	// Week (2018,-1) starts on: 2017-12-18 00:00:00 +0000 UTC
	// Week (2018, 0) starts on: 2017-12-25 00:00:00 +0000 UTC
	// Week (2018, 1) starts on: 2018-01-01 00:00:00 +0000 UTC
	// Week (2018, 2) starts on: 2018-01-08 00:00:00 +0000 UTC
	// Week (2019, 1) starts on: 2018-12-31 00:00:00 +0000 UTC
	// Week (2019, 2) starts on: 2019-01-07 00:00:00 +0000 UTC
	// Week (2019,53) starts on: 2019-12-30 00:00:00 +0000 UTC
	// Week (2019,54) starts on: 2020-01-06 00:00:00 +0000 UTC
}

func TestParseMonth(t *testing.T) {
	cases := []struct {
		name   string     // Name of the test case
		inputs []string   // inputs
		exp    time.Month // expected output
		err    bool       // true if error is expected
	}{
		{"January", []string{"January", "Jan"}, time.January, false},
		{"February", []string{"February", "Feb"}, time.February, false},
		{"March", []string{"March", "Mar"}, time.March, false},
		{"April", []string{"April", "Apr"}, time.April, false},
		{"May", []string{"May"}, time.May, false},
		{"June", []string{"June", "Jun"}, time.June, false},
		{"July", []string{"July", "Jul"}, time.July, false},
		{"August", []string{"August", "Aug"}, time.August, false},
		{"September", []string{"September", "Sep"}, time.September, false},
		{"October", []string{"October", "Oct"}, time.October, false},
		{"November", []string{"November", "Nov"}, time.November, false},
		{"December", []string{"December", "Dec"}, time.December, false},
		{"invalid", []string{"invalid", "ii", "january"}, time.January, true},
	}

	for _, c := range cases {
		for _, s := range c.inputs {
			got, err := ParseMonth(s)
			if c.err != (err != nil) {
				t.Errorf("[%s(%s)] Expected error: %v, got: %v", c.name, s, c.err, err)
			} else {
				if got != c.exp {
					t.Errorf("[%s(%s)] Expected: %v, got: %v", c.name, s, c.exp, got)
				}
			}
		}
	}
}

func TestParseWeekday(t *testing.T) {
	cases := []struct {
		name   string       // Name of the test case
		inputs []string     // inputs
		exp    time.Weekday // expected output
		err    bool         // true if error is expected
	}{
		{"Monday", []string{"Monday", "Mon"}, time.Monday, false},
		{"Tuesday", []string{"Tuesday", "Tue"}, time.Tuesday, false},
		{"Wednesday", []string{"Wednesday", "Wed"}, time.Wednesday, false},
		{"Thursday", []string{"Thursday", "Thu"}, time.Thursday, false},
		{"Friday", []string{"Friday", "Fri"}, time.Friday, false},
		{"Saturday", []string{"Saturday", "Sat"}, time.Saturday, false},
		{"Sunday", []string{"Sunday", "Sun"}, time.Sunday, false},
		{"invalid", []string{"invalid", "ii", "january"}, time.Sunday, true},
	}

	for _, c := range cases {
		for _, s := range c.inputs {
			got, err := ParseWeekday(s)
			if c.err != (err != nil) {
				t.Errorf("[%s(%s)] Expected error: %v, got: %v", c.name, s, c.err, err)
			} else {
				if got != c.exp {
					t.Errorf("[%s(%s)] Expected: %v, got: %v", c.name, s, c.exp, got)
				}
			}
		}
	}
}

// ExampleNextWeekday shows how to use the NextWeekday() function.
func ExampleNextWeekday() {
	// Given the date today is Tuesday, Jan 26th, 2021
	now := time.Date(2021, time.January, 26, 11, 0, 0, 0, time.UTC)

	// Get the timestamp of next Monday at 12pm UTC
	t := NextWeekday(now, "Monday", 12, time.UTC)

	fmt.Println(t)

	// Output:
	// 2021-02-01 12:00:00 +0000 UTC
}

func Test_NextWeekday(t *testing.T) {
	t.Parallel()

	// Testing against Wednesday, Jan 13th, 2021 12:00 UTC
	now := time.Date(2021, time.January, 13, 12, 0, 0, 0, time.UTC)

	testCases := []struct {
		name     string
		weekday  string
		hour     int
		expected time.Time
		loc      *time.Location
	}{
		{"next Friday", "Friday", 12, time.Date(2021, time.January, 15, 12, 0, 0, 0, time.UTC), nil},
		{"next Saturday", "Saturday", 15, time.Date(2021, time.January, 16, 15, 0, 0, 0, time.UTC), nil},
		{"next Monday", "Monday", 20, time.Date(2021, time.January, 18, 20, 0, 0, 0, time.UTC), nil},
		{"next Tuesday", "Tuesday", 9, time.Date(2021, time.January, 19, 9, 0, 0, 0, time.UTC), nil},
		{"next Tuesday", "Tuesday", 9, time.Date(2021, time.January, 19, 9, 0, 0, 0, time.UTC), time.UTC},
		{"next Wednesday", "Wednesday", 9, time.Date(2021, time.January, 20, 9, 0, 0, 0, time.UTC), nil},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := NextWeekday(now, tt.weekday, tt.hour, tt.loc)
			if tt.expected != got {
				t.Errorf("[%s] Expected: %v, got: %v", tt.name, tt.expected, got)
			}
		})
	}

	// Verify the new date is in the new year
	testName := "new year next Monday"
	t.Run(testName, func(t *testing.T) {
		// Wednesday, December 30th, 2020 12:00 UTC
		now := time.Date(2020, time.December, 30, 12, 0, 0, 0, time.UTC)

		got := NextWeekday(now, "Monday", 9, time.UTC)
		expected := time.Date(2021, time.January, 4, 9, 0, 0, 0, time.UTC)

		if expected != got {
			t.Errorf("[%s] Expected: %v, got: %v", testName, expected, got)
		}
	})
}
