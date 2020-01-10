// This file contains general formatting utilities.

package fmtx

import (
	"strconv"
)

// FormatInt formats an integer with grouping decimals, in decimal radix.
// Grouping signs are inserted after every groupSize digits, starting from the right.
// A groupingSize less than 1 will default to 3.
// Only ASCII grouping decimal signs are supported which may be provided with
// grouping.
//
// For details, see https://stackoverflow.com/a/31046325/1705598
func FormatInt(n int64, groupSize int, grouping byte) string {
	if groupSize < 1 {
		groupSize = 3
	}

	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / groupSize

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == groupSize {
			j, k = j-1, 0
			out[j] = grouping
		}
	}
}
