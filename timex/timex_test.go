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
