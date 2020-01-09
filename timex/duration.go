// This file contains duration-related utilities.

package timex

import (
	"strings"
	"time"
)

var roundDivs = []time.Duration{
	time.Duration(1), time.Duration(10), time.Duration(100), time.Duration(1000),
	time.Duration(10000), time.Duration(100000), time.Duration(1000000), time.Duration(10000000),
}

// Round rounds the given duration so that when it is printed (when its String()
// method is called), result will have the given fraction digits at most.
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

// ShortDuration formats the given duration into a short format.
// The short format eludes trailing 0 units in the string.
//
// For example:
//   -5h4m3s remains 5h4m3s
//   -5h4m0s becomes 5h4m
//   -5h0m3s remains 5h0m3s
//   -4m3s   remains 4m3s
//   -5h0m0s becomes 5h
//   -4m0s   becomes 4m
//   -3s     remains 3s
//
// For details, see https://stackoverflow.com/questions/41335155/time-duration-to-string-2h-instead-2h0m0s/41336257#41336257
func ShortDuration(d time.Duration) string {
	s := d.String()
	if strings.HasSuffix(s, "m0s") {
		s = s[:len(s)-2]
	}
	if strings.HasSuffix(s, "h0m") {
		s = s[:len(s)-2]
	}
	return s
}
