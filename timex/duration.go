// This file contains duration-related utilities.

package timex

import "time"

var roundDivs = []time.Duration{
	time.Duration(1), time.Duration(10), time.Duration(100), time.Duration(1000),
	time.Duration(10000), time.Duration(100000), time.Duration(1000000), time.Duration(10000000),
}

// Round rounds the given duration so that when it is printed (when its String()
// method is called), result will have the given fraction digits.
//
// For details, see https://stackoverflow.com/questions/58414820/limiting-significant-digits-in-formatted-durations/58415564#58415564
func Round(d time.Duration, digits int) time.Duration {
	if digits < 0 {
		digits = 0
	}
	if digits > len(roundDivs)-1 {
		digits = len(roundDivs) - 1
	}

	switch {
	case d > time.Second:
		d = d.Round(time.Second / roundDivs[digits])
	case d > time.Millisecond:
		d = d.Round(time.Millisecond / roundDivs[digits])
	case d > time.Microsecond:
		d = d.Round(time.Microsecond / roundDivs[digits])
	}
	return d
}
